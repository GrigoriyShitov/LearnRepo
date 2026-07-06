from dataclasses import dataclass

from prompt_toolkit import prompt
from prompt_toolkit.completion import WordCompleter
from psycopg.rows import class_row
from rich.panel import Panel
from rich.table import Table

from console import console, render_error
from db import get_conn
from validators import ChoiceValidator, NonEmptyValidator, YesNoValidator
from commands import command, CATEGORY_WAREHOUSES

cities = [
    "Москва",
    "Санкт-Петербург",
    "Новосибирск",
    "Екатеринбург",
    "Казань",
    "Нижний Новгород",
    "Челябинск",
    "Самара",
    "Омск",
    "Ростов-на-Дону",
    "Уфа",
    "Красноярск",
    "Воронеж",
    "Пермь",
    "Волгоград",
]

city_completer = WordCompleter(cities, ignore_case=True, sentence=True)
city_validator = ChoiceValidator(
    cities, message="Город должен быть из списка. Используйте Tab для автодополнения."
)

yes_no_choices = ["y", "n", "д", "нет", "да", "yes", "no"]
yes_no_completer = WordCompleter(yes_no_choices, ignore_case=True)


@dataclass
class Warehouse:
    id: int
    city: str
    address: str
    label: str | None
    is_central: bool


def _render_warehouse(warehouse: Warehouse) -> None:
    table = Table(show_header=False, box=None, padding=(0, 2))

    table.add_column("Поле", style="bold cyan", width=15)
    table.add_column("Значение", style="white")

    table.add_row("ID", str(warehouse.id))
    table.add_row("Город", warehouse.city)
    table.add_row("Адрес", warehouse.address)
    table.add_row("Метка", warehouse.label or "")
    table.add_row("Центральный", "да" if warehouse.is_central else "нет")

    panel = Panel(
        table,
        expand=False,
        title=f"[bold green]Склад #{warehouse.id}[/bold green]",
        border_style="green",
    )

    console.print(panel)


def _warehouse_count(conn) -> int:
    with conn.cursor() as cur:
        cur.execute("SELECT COUNT(*) FROM catalog.warehouses")
        return cur.fetchone()[0]


def _prompt_make_central() -> bool:
    answer = prompt(
        "Сделать этот склад центральным? (y/n, д/н): ",
        validator=YesNoValidator(),
        completer=yes_no_completer,
    )
    return YesNoValidator.is_yes(answer)


def _unset_other_central(conn) -> None:
    conn.execute(
        "UPDATE catalog.warehouses SET is_central = FALSE WHERE is_central = TRUE"
    )


def _warehouse_label_text(city: str, label: str | None) -> str:
    if label:
        return f"Склад в городе {city} ({label})"
    return f"Склад в городе {city}"


@command("list warehouses", "список всех складов", CATEGORY_WAREHOUSES)
def list_warehouses() -> None:
    conn = get_conn()
    table = Table(title="Склады", show_header=True, header_style="bold cyan")

    table.add_column("ID", style="dim", width=6, justify="right")
    table.add_column("Город", style="green", min_width=20)
    table.add_column("Адрес", style="yellow", min_width=30)
    table.add_column("Метка", style="magenta", min_width=15)
    table.add_column("Центральный", style="cyan", min_width=12)

    with conn.cursor(row_factory=class_row(Warehouse)) as cur:
        cur.execute(
            "SELECT id, city, address, label, is_central "
            "FROM catalog.warehouses ORDER BY id"
        )
        warehouses: list[Warehouse] = cur.fetchall()

    for warehouse in warehouses:
        table.add_row(
            str(warehouse.id),
            warehouse.city,
            warehouse.address,
            warehouse.label or "",
            "да" if warehouse.is_central else "нет",
        )
    console.print(table)


@command("show warehouse", "информация о складе", CATEGORY_WAREHOUSES)
def show_warehouse(_id: str) -> None:
    conn = get_conn()
    with conn.cursor(row_factory=class_row(Warehouse)) as cur:
        cur.execute(
            "SELECT id, city, address, label, is_central "
            "FROM catalog.warehouses WHERE id = %s",
            (_id,),
        )
        warehouse: Warehouse | None = cur.fetchone()

    if warehouse is None:
        render_error(f"Склад с ID {_id} не найден")
        return

    _render_warehouse(warehouse)


@command("add warehouse", "добавить склад (интерактивно)", CATEGORY_WAREHOUSES)
def add_warehouse() -> None:
    conn = get_conn()
    city = prompt("Город: ", validator=city_validator, completer=city_completer).strip()
    address = prompt("Адрес: ", validator=NonEmptyValidator()).strip()
    label = prompt("Метка (необязательно): ").strip() or None

    is_first = _warehouse_count(conn) == 0
    if is_first:
        is_central = True
        console.print("[dim]Первый склад автоматически назначен центральным[/dim]")
    else:
        is_central = _prompt_make_central()

    if is_central:
        _unset_other_central(conn)

    conn.execute(
        """
        INSERT INTO catalog.warehouses (city, address, label, is_central)
        VALUES (%s, %s, %s, %s)
        """,
        (city, address, label, is_central),
    )

    console.print(f"[green]{_warehouse_label_text(city, label)} добавлен[/green]")


@command("edit warehouse", "редактировать склад", CATEGORY_WAREHOUSES)
def edit_warehouse(_id: str) -> None:
    conn = get_conn()
    with conn.cursor(row_factory=class_row(Warehouse)) as cur:
        cur.execute(
            "SELECT id, city, address, label, is_central "
            "FROM catalog.warehouses WHERE id = %s",
            (_id,),
        )
        warehouse: Warehouse | None = cur.fetchone()

    if warehouse is None:
        render_error(f"Склад с ID {_id} не найден")
        return

    city = prompt(
        "Город: ",
        default=warehouse.city,
        validator=city_validator,
        completer=city_completer,
    ).strip()
    address = prompt(
        "Адрес: ", default=warehouse.address, validator=NonEmptyValidator()
    ).strip()
    label = (
        prompt("Метка (необязательно): ", default=warehouse.label or "").strip() or None
    )

    if warehouse.is_central:
        is_central = True
    else:
        is_central = _prompt_make_central()

    if is_central and not warehouse.is_central:
        _unset_other_central(conn)

    conn.execute(
        """
        UPDATE catalog.warehouses
        SET city = %s, address = %s, label = %s, is_central = %s
        WHERE id = %s
        """,
        (city, address, label, is_central, _id),
    )

    console.print(f"[green]{_warehouse_label_text(city, label)} обновлен[/green]")


@command("delete warehouse", "удалить склад", CATEGORY_WAREHOUSES)
def delete_warehouse(_id: str) -> None:
    conn = get_conn()
    with conn.cursor(row_factory=class_row(Warehouse)) as cur:
        cur.execute(
            "SELECT id, city, address, label, is_central "
            "FROM catalog.warehouses WHERE id = %s",
            (_id,),
        )
        warehouse: Warehouse | None = cur.fetchone()

    if warehouse is None:
        render_error(f"Склад с ID {_id} не найден")
        return

    total = _warehouse_count(conn)
    if warehouse.is_central and total > 1:
        render_error(
            "Нельзя удалить центральный склад: "
            "сначала назначьте другой склад центральным"
        )
        return

    _render_warehouse(warehouse)

    answer = prompt("Вы уверены? (y/n, д/н): ", validator=YesNoValidator())

    if not YesNoValidator.is_yes(answer):
        return

    conn.execute("DELETE FROM catalog.warehouses WHERE id = %s", (_id,))

    console.print(
        f"[green]{_warehouse_label_text(warehouse.city, warehouse.label)} удален[/green]"
    )
