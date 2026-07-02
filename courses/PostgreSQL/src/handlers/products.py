from dataclasses import dataclass
from decimal import Decimal

from prompt_toolkit import prompt
from prompt_toolkit.shortcuts import choice
from psycopg.rows import dict_row
from rich.panel import Panel
from rich.table import Table

from commands import command, CATEGORY_PRODUCTS
from console import console, render_error
from db import get_conn
from handlers.product_categories import ProductCategory
from validators import NonEmptyValidator, PriceValidator, SkuValidator, YesNoValidator


@dataclass
class Product:
    id: int
    sku: str
    name: str
    price: Decimal
    category: str


def _get_categories() -> list[ProductCategory]:
    conn = get_conn()
    with conn.cursor(row_factory=dict_row) as cur:
        cur.execute("SELECT id, name FROM catalog.product_categories ORDER BY id")
        rows = cur.fetchall()
    return [ProductCategory(id=row["id"], name=row["name"]) for row in rows]


def _prompt_category_id(
    categories: list[ProductCategory], default_name: str | None = None
) -> str:
    options = [(str(category.id), category.name) for category in categories]
    default = None
    if default_name is not None:
        for category in categories:
            if category.name == default_name:
                default = str(category.id)
                break

    return choice(message="Категория:", options=options, default=default)


def _fetch_product(_id: str) -> Product | None:
    conn = get_conn()
    with conn.cursor(row_factory=dict_row) as cur:
        cur.execute(
            """
            SELECT p.id, p.sku, p.name, p.price, c.name AS category
            FROM catalog.products p
            JOIN catalog.product_categories c ON c.id = p.category_id
            WHERE p.id = %s
            """,
            (_id,),
        )
        row = cur.fetchone()

    if row is None:
        return None

    return Product(
        id=row["id"],
        sku=row["sku"],
        name=row["name"],
        price=row["price"],
        category=row["category"],
    )


def _render_product(product: Product) -> None:
    table = Table(show_header=False, box=None, padding=(0, 2))

    table.add_column("Поле", style="bold cyan", width=15)
    table.add_column("Значение", style="white")

    table.add_row("ID", str(product.id))
    table.add_row("SKU", product.sku)
    table.add_row("Название", product.name)
    table.add_row("Цена", f"{product.price:.2f}")
    table.add_row("Категория", product.category)

    panel = Panel(
        table,
        expand=False,
        title=f"[bold green]Товар #{product.id}[/bold green]",
        border_style="green",
    )

    console.print(panel)


@command("list products", "список всех товаров", CATEGORY_PRODUCTS)
def list_products() -> None:
    conn = get_conn()
    table = Table(title="Товары", show_header=True, header_style="bold cyan")

    table.add_column("ID", style="dim", width=6, justify="right")
    table.add_column("SKU", style="cyan", min_width=12)
    table.add_column("Название", style="green", min_width=20)
    table.add_column("Цена", style="yellow", min_width=12, justify="right")
    table.add_column("Категория", style="magenta", min_width=15)

    with conn.cursor(row_factory=dict_row) as cur:
        cur.execute("""
            SELECT p.id, p.sku, p.name, p.price, c.name AS category
            FROM catalog.products p
            JOIN catalog.product_categories c ON c.id = p.category_id
            ORDER BY p.id
            """)
        products = cur.fetchall()

    for product in products:
        table.add_row(
            str(product["id"]),
            product["sku"],
            product["name"],
            f"{product['price']:.2f}",
            product["category"],
        )

    console.print(table)


@command("show product", "информация о товаре", CATEGORY_PRODUCTS)
def show_product(_id: str) -> None:
    product = _fetch_product(_id)

    if product is None:
        render_error(f"Товар с ID {_id} не найден")
        return

    _render_product(product)


@command("add product", "добавить товар (интерактивно)", CATEGORY_PRODUCTS)
def add_product() -> None:
    categories = _get_categories()
    if not categories:
        render_error("Сначала добавьте хотя бы одну категорию товаров")
        return

    conn = get_conn()
    sku = prompt("SKU: ", validator=SkuValidator()).strip()
    name = prompt("Название: ", validator=NonEmptyValidator()).strip()
    price_str = prompt("Цена: ", validator=PriceValidator()).strip()
    category_id = _prompt_category_id(categories)

    try:
        conn.execute(
            """
            INSERT INTO catalog.products (sku, name, price, category_id)
            VALUES (%s, %s, %s, %s)
            """,
            (sku, name, Decimal(price_str), category_id),
        )
    except Exception as e:
        if "unique" in str(e).lower() and "sku" in str(e).lower():
            render_error(f"Товар с SKU «{sku}» уже существует")
            return
        raise

    console.print(f"[green]Товар «{name}» добавлен[/green]")


@command("edit product", "редактировать товар", CATEGORY_PRODUCTS)
def edit_product(_id: str) -> None:
    conn = get_conn()
    with conn.cursor(row_factory=dict_row) as cur:
        cur.execute(
            """
            SELECT p.id, p.sku, p.name, p.price, p.category_id, c.name AS category
            FROM catalog.products p
            JOIN catalog.product_categories c ON c.id = p.category_id
            WHERE p.id = %s
            """,
            (_id,),
        )
        row = cur.fetchone()

    if row is None:
        render_error(f"Товар с ID {_id} не найден")
        return

    categories = _get_categories()
    if not categories:
        render_error("Нет доступных категорий товаров")
        return

    sku = prompt("SKU: ", default=row["sku"], validator=SkuValidator()).strip()
    name = prompt(
        "Название: ", default=row["name"], validator=NonEmptyValidator()
    ).strip()
    price_str = prompt(
        "Цена: ", default=str(row["price"]), validator=PriceValidator()
    ).strip()
    category_id = _prompt_category_id(categories, default_name=row["category"])

    try:
        conn.execute(
            """
            UPDATE catalog.products
            SET sku = %s, name = %s, price = %s, category_id = %s
            WHERE id = %s
            """,
            (sku, name, Decimal(price_str), category_id, _id),
        )
    except Exception as e:
        if "unique" in str(e).lower() and "sku" in str(e).lower():
            render_error(f"Товар с SKU «{sku}» уже существует")
            return
        raise

    console.print(f"[green]Товар «{name}» обновлён[/green]")


@command("delete product", "удалить товар", CATEGORY_PRODUCTS)
def delete_product(_id: str) -> None:
    product = _fetch_product(_id)

    if product is None:
        render_error(f"Товар с ID {_id} не найден")
        return

    _render_product(product)

    answer = prompt("Вы уверены? (y/n, д/н): ", validator=YesNoValidator())

    if not YesNoValidator.is_yes(answer):
        return

    conn = get_conn()
    conn.execute("DELETE FROM catalog.products WHERE id = %s", (_id,))
    console.print(f"[green]Товар «{product.name}» удалён[/green]")
