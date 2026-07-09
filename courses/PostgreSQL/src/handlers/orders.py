from dataclasses import dataclass
from datetime import datetime
from decimal import Decimal

from prompt_toolkit import prompt
from prompt_toolkit.completion import WordCompleter
from prompt_toolkit.shortcuts import choice
from psycopg.rows import dict_row
from rich.panel import Panel
from rich.table import Table

from commands import command, CATEGORY_ORDERS
from console import console, render_error
from db import get_conn
from validators import (
    NonEmptyValidator,
    PriceValidator,
    QuantityValidator,
    YesNoValidator,
)

UNPUBLISHED_STATUS = "unpublished"


@dataclass
class Order:
    id: int
    status: str
    total_amount: Decimal
    created_at: datetime
    warehouse: str
    warehouse_id: int


@dataclass
class OrderItem:
    product_id: int
    sku: str
    name: str
    price: Decimal
    quantity: int


def _warehouse_label(city: str, label: str | None) -> str:
    if label:
        return f"{city} ({label})"
    return city


def _fetch_order(_id: str) -> Order | None:
    conn = get_conn()
    with conn.cursor(row_factory=dict_row) as cur:
        cur.execute(
            """
            SELECT o.id, o.status, o.total_amount, o.created_at,
                   o.warehouse_id, w.city, w.label
            FROM sales.orders o
            JOIN catalog.warehouses w ON w.id = o.warehouse_id
            WHERE o.id = %s
            """,
            (_id,),
        )
        row = cur.fetchone()

    if row is None:
        return None

    return Order(
        id=row["id"],
        status=row["status"],
        total_amount=row["total_amount"],
        created_at=row["created_at"],
        warehouse=_warehouse_label(row["city"], row["label"]),
        warehouse_id=row["warehouse_id"],
    )


def _fetch_order_items(order_id: str) -> list[OrderItem]:
    conn = get_conn()
    with conn.cursor(row_factory=dict_row) as cur:
        cur.execute(
            """
            SELECT oi.product_id, p.sku, p.name, oi.price, oi.quantity
            FROM sales.order_items oi
            JOIN catalog.products p ON p.id = oi.product_id
            WHERE oi.order_id = %s
            ORDER BY oi.product_id
            """,
            (order_id,),
        )
        rows = cur.fetchall()

    return [
        OrderItem(
            product_id=row["product_id"],
            sku=row["sku"],
            name=row["name"],
            price=row["price"],
            quantity=row["quantity"],
        )
        for row in rows
    ]


def _recalculate_total(conn, order_id: str) -> None:
    conn.execute(
        """
        UPDATE sales.orders
        SET total_amount = (
            SELECT COALESCE(SUM(price * quantity), 0)
            FROM sales.order_items
            WHERE order_id = %s
        )
        WHERE id = %s
        """,
        (order_id, order_id),
    )


def _require_unpublished(order: Order) -> bool:
    if order.status != UNPUBLISHED_STATUS:
        render_error(
            f"Заказ #{order.id} нельзя изменить: статус «{order.status}» "
            "(доступно только для unpublished)"
        )
        return False
    return True


def _get_warehouses() -> list[dict]:
    conn = get_conn()
    with conn.cursor(row_factory=dict_row) as cur:
        cur.execute("SELECT id, city, label FROM catalog.warehouses ORDER BY id")
        return cur.fetchall()


def _prompt_warehouse_id(default_id: int | None = None) -> str:
    warehouses = _get_warehouses()
    if not warehouses:
        raise ValueError("no_warehouses")

    options = [
        (str(w["id"]), _warehouse_label(w["city"], w["label"])) for w in warehouses
    ]
    default = str(default_id) if default_id is not None else None
    return choice(message="Склад:", options=options, default=default)


def _get_available_products(order_id: str) -> list[dict]:
    """Вернуть товары, которых ещё нет в заказе (1 запрос к БД)."""
    conn = get_conn()
    with conn.cursor(row_factory=dict_row) as cur:
        cur.execute(
            """
            SELECT p.id, p.sku, p.name, p.price
            FROM catalog.products p
            WHERE NOT EXISTS (
                SELECT 1
                FROM sales.order_items oi
                WHERE oi.order_id = %s AND oi.product_id = p.id
            )
            ORDER BY p.id
            """,
            (order_id,),
        )
        return cur.fetchall()


def _product_display(product: dict) -> str:
    return f"{product['sku']} — {product['name']}"


def _prompt_product(available_products: list[dict]) -> dict | None:
    if not available_products:
        return None

    labels = [_product_display(p) for p in available_products]
    label_to_product = {_product_display(p): p for p in available_products}
    completer = WordCompleter(labels, ignore_case=True, sentence=True)

    label = prompt(
        "Товар (SKU — название): ",
        completer=completer,
        validator=NonEmptyValidator(message="Выберите товар из списка"),
    ).strip()

    return label_to_product.get(label)


def _prompt_order_item(order_id: str) -> OrderItem | None:
    items = _fetch_order_items(order_id)
    if not items:
        render_error("В заказе нет позиций")
        return None

    options = [(str(item.product_id), f"{item.sku} — {item.name}") for item in items]
    product_id = choice(message="Позиция заказа:", options=options)
    return next(item for item in items if str(item.product_id) == product_id)


def _render_order(order: Order, items: list[OrderItem]) -> None:
    table = Table(show_header=False, box=None, padding=(0, 2))

    table.add_column("Поле", style="bold cyan", width=18)
    table.add_column("Значение", style="white")

    table.add_row("ID", str(order.id))
    table.add_row("Статус", order.status)
    table.add_row("Сумма", f"{order.total_amount:.2f}")
    table.add_row("Создан", order.created_at.strftime("%Y-%m-%d %H:%M:%S"))
    table.add_row("Склад", order.warehouse)

    panel = Panel(
        table,
        expand=False,
        title=f"[bold green]Заказ #{order.id}[/bold green]",
        border_style="green",
    )
    console.print(panel)

    if items:
        items_table = Table(
            title="Позиции заказа", show_header=True, header_style="bold cyan"
        )
        items_table.add_column("SKU", style="cyan", min_width=12)
        items_table.add_column("Название", style="green", min_width=20)
        items_table.add_column("Цена", style="yellow", justify="right", min_width=10)
        items_table.add_column("Кол-во", style="magenta", justify="right", min_width=8)
        items_table.add_column("Сумма", style="white", justify="right", min_width=10)

        for item in items:
            line_total = item.price * item.quantity
            items_table.add_row(
                item.sku,
                item.name,
                f"{item.price:.2f}",
                str(item.quantity),
                f"{line_total:.2f}",
            )

        console.print(items_table)
    else:
        console.print("[dim]Позиции заказа отсутствуют[/dim]")


def _add_order_item_interactive(order_id: str) -> bool:
    conn = get_conn()
    available = _get_available_products(order_id)

    if not available:
        render_error("Нет доступных товаров для добавления в заказ")
        return False

    product = _prompt_product(available)
    if product is None:
        render_error("Товар не найден. Выберите значение из списка (Tab)")
        return False

    price_str = prompt(
        "Цена: ",
        default=str(product["price"]),
        validator=PriceValidator(),
    ).strip()
    quantity_str = prompt("Количество: ", validator=QuantityValidator()).strip()

    conn.execute(
        """
        INSERT INTO sales.order_items (order_id, product_id, price, quantity)
        VALUES (%s, %s, %s, %s)
        """,
        (order_id, product["id"], Decimal(price_str), int(quantity_str)),
    )
    _recalculate_total(conn, order_id)
    console.print(f"[green]Товар «{product['name']}» добавлен в заказ[/green]")
    return True


@command("list orders", "список всех заказов", CATEGORY_ORDERS)
def list_orders() -> None:
    conn = get_conn()
    table = Table(title="Заказы", show_header=True, header_style="bold cyan")

    table.add_column("ID", style="dim", width=6, justify="right")
    table.add_column("Статус", style="cyan", min_width=14)
    table.add_column("Сумма", style="yellow", min_width=12, justify="right")
    table.add_column("Создан", style="green", min_width=20)
    table.add_column("Склад", style="magenta", min_width=25)

    with conn.cursor(row_factory=dict_row) as cur:
        cur.execute("""
            SELECT o.id, o.status, o.total_amount, o.created_at,
                   w.city, w.label
            FROM sales.orders o
            JOIN catalog.warehouses w ON w.id = o.warehouse_id
            ORDER BY o.id
            """)
        orders = cur.fetchall()

    for order in orders:
        table.add_row(
            str(order["id"]),
            order["status"],
            f"{order['total_amount']:.2f}",
            order["created_at"].strftime("%Y-%m-%d %H:%M:%S"),
            _warehouse_label(order["city"], order["label"]),
        )

    console.print(table)


@command("show order", "информация о заказе", CATEGORY_ORDERS)
def show_order(_id: str) -> None:
    order = _fetch_order(_id)
    if order is None:
        render_error(f"Заказ с ID {_id} не найден")
        return

    items = _fetch_order_items(_id)
    _render_order(order, items)


@command("add order", "добавить заказ (интерактивно)", CATEGORY_ORDERS)
def add_order() -> None:
    try:
        warehouse_id = _prompt_warehouse_id()
    except ValueError:
        render_error("Сначала добавьте хотя бы один склад")
        return

    conn = get_conn()
    with conn.cursor(row_factory=dict_row) as cur:
        cur.execute(
            """
            INSERT INTO sales.orders (warehouse_id)
            VALUES (%s)
            RETURNING id
            """,
            (warehouse_id,),
        )
        row = cur.fetchone()

    if row is None:
        render_error("Не удалось создать заказ")
        return

    order_id = str(row["id"])

    while True:
        answer = prompt(
            "Добавить товар в заказ? (y/n, д/н): ", validator=YesNoValidator()
        )
        if not YesNoValidator.is_yes(answer):
            break
        _add_order_item_interactive(order_id)

    _recalculate_total(conn, order_id)
    console.print(f"[green]Заказ #{order_id} создан[/green]")


@command("edit order", "редактировать заказ", CATEGORY_ORDERS)
def edit_order(_id: str) -> None:
    order = _fetch_order(_id)
    if order is None:
        render_error(f"Заказ с ID {_id} не найден")
        return

    if not _require_unpublished(order):
        return

    try:
        warehouse_id = _prompt_warehouse_id(default_id=order.warehouse_id)
    except ValueError:
        render_error("Нет доступных складов")
        return

    conn = get_conn()
    conn.execute(
        "UPDATE sales.orders SET warehouse_id = %s WHERE id = %s",
        (warehouse_id, _id),
    )
    console.print(f"[green]Заказ #{_id} обновлён[/green]")


@command("delete order", "удалить заказ", CATEGORY_ORDERS)
def delete_order(_id: str) -> None:
    order = _fetch_order(_id)
    if order is None:
        render_error(f"Заказ с ID {_id} не найден")
        return

    if not _require_unpublished(order):
        return

    items = _fetch_order_items(_id)
    _render_order(order, items)

    answer = prompt("Вы уверены? (y/n, д/н): ", validator=YesNoValidator())
    if not YesNoValidator.is_yes(answer):
        return

    conn = get_conn()
    conn.execute("DELETE FROM sales.orders WHERE id = %s", (_id,))
    console.print(f"[green]Заказ #{_id} удалён[/green]")


@command("publish order", "опубликовать заказ", CATEGORY_ORDERS)
def publish_order(_id: str) -> None:
    order = _fetch_order(_id)
    if order is None:
        render_error(f"Заказ с ID {_id} не найден")
        return

    if order.status != UNPUBLISHED_STATUS:
        render_error(
            f"Заказ #{_id} нельзя опубликовать: текущий статус «{order.status}»"
        )
        return

    conn = get_conn()
    with conn.cursor() as cur:
        cur.execute(
            """
            UPDATE sales.orders SET status = 'new'
            WHERE id = %s AND status = %s
            """,
            (_id, UNPUBLISHED_STATUS),
        )
        if cur.rowcount == 0:
            render_error(f"Не удалось опубликовать заказ #{_id}")
            return

    console.print(f"[green]Заказ #{_id} опубликован (статус: new)[/green]")


@command("add order_item", "добавить позицию в заказ", CATEGORY_ORDERS)
def add_order_item(order_id: str) -> None:
    order = _fetch_order(order_id)
    if order is None:
        render_error(f"Заказ с ID {order_id} не найден")
        return

    if not _require_unpublished(order):
        return

    _add_order_item_interactive(order_id)


@command("edit order_item", "редактировать позицию заказа", CATEGORY_ORDERS)
def edit_order_item(order_id: str) -> None:
    order = _fetch_order(order_id)
    if order is None:
        render_error(f"Заказ с ID {order_id} не найден")
        return

    if not _require_unpublished(order):
        return

    item = _prompt_order_item(order_id)
    if item is None:
        return

    price_str = prompt(
        "Цена: ", default=str(item.price), validator=PriceValidator()
    ).strip()
    quantity_str = prompt(
        "Количество: ", default=str(item.quantity), validator=QuantityValidator()
    ).strip()

    conn = get_conn()
    conn.execute(
        """
        UPDATE sales.order_items
        SET price = %s, quantity = %s
        WHERE order_id = %s AND product_id = %s
        """,
        (Decimal(price_str), int(quantity_str), order_id, item.product_id),
    )
    _recalculate_total(conn, order_id)
    console.print(f"[green]Позиция «{item.name}» обновлена[/green]")


@command("delete order_item", "удалить позицию из заказа", CATEGORY_ORDERS)
def delete_order_item(order_id: str) -> None:
    order = _fetch_order(order_id)
    if order is None:
        render_error(f"Заказ с ID {order_id} не найден")
        return

    if not _require_unpublished(order):
        return

    item = _prompt_order_item(order_id)
    if item is None:
        return

    answer = prompt("Вы уверены? (y/n, д/н): ", validator=YesNoValidator())
    if not YesNoValidator.is_yes(answer):
        return

    conn = get_conn()
    conn.execute(
        """
        DELETE FROM sales.order_items
        WHERE order_id = %s AND product_id = %s
        """,
        (order_id, item.product_id),
    )
    _recalculate_total(conn, order_id)
    console.print(f"[green]Позиция «{item.name}» удалена из заказа[/green]")
