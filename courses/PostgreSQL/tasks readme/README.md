# Отчёт по заданию

**Автор:** Шитов Григорий

REPL-приложение для управления каталогом товаров, складами и заказами. Реализовано на Python с PostgreSQL.

---

## Задание 1. CRUD каталога

### Что сделано

Реализовано интерактивное REPL-приложение с CRUD-операциями над тремя сущностями каталога:

| Сущность | Модуль | Команды |
|----------|--------|---------|
| Категории товаров | `src/handlers/product_categories.py` | `list/show/add/edit/delete product_category` |
| Товары | `src/handlers/products.py` | `list/show/add/edit/delete product` |
| Склады | `src/handlers/warehouses.py` | `list/show/add/edit/delete warehouse` |

### Схема БД (схема `catalog`)

- **product_categories** — `id`, `name` (UNIQUE)
- **products** — `id`, `sku` (UNIQUE, до 30 символов), `name`, `price`, `category_id` → FK на `product_categories`
- **warehouses** — `id`, `city`, `address`, `label` (nullable), `is_central`

### Особенности реализации

- Выбор категории товара через `choice()` (по названию, без показа ID пользователю)
- Валидация SKU (`SkuValidator`), цены (`PriceValidator`)
- Логика центрального склада: первый склад — центральный; нельзя снять флаг с единственного центрального; нельзя удалить центральный склад, пока есть другие

### ER-диаграмма (draw.io, Crow's Foot)

Исходник: [img/er-drawio.drawio](../img/er-drawio.drawio)

![ER-диаграмма draw.io](../../img/er-drawio.png)

---

## Задание 2. Миграции, ER-диаграммы, заказы

### Что сделано

1. **ER-диаграмма в draw.io** — целевая схема (5 таблиц, связи Crow's Foot)
2. **Первая миграция Alembic** (`7903e1779201`) — пересоздание `catalog` с FK, создание схемы `sales`
3. **ER-диаграмма из DBeaver** — по фактической БД после миграции (схемы `catalog`, `sales`)
4. **Сравнение диаграмм** — структура совпадает: 5 таблиц, 4 внешних ключа, составной PK у `order_items`
5. **CRUD заказов** и **CRUD позиций заказа** в `src/handlers/orders.py`

### Миграция

Файлы: [alembic/sql/7903e1779201/up.sql](../alembic/sql/7903e1779201/up.sql), [down.sql](../alembic/sql/7903e1779201/down.sql)

Создано:

| Схема | Таблица | Назначение |
|-------|---------|------------|
| `catalog` | `product_categories`, `products`, `warehouses` | каталог (пересозданы с FK) |
| `sales` | `orders` | заказы: `status`, `total_amount`, `created_at`, `warehouse_id` |
| `sales` | `order_items` | позиции: `order_id`, `product_id`, `price`, `quantity` |

Ограничения:

- `products.category_id` → `product_categories(id)`
- `orders.warehouse_id` → `warehouses(id)`
- `order_items` — составной PK `(order_id, product_id)`, FK на `orders` и `products`
- `orders.status` — CHECK (`unpublished`, `new`, `processing`, `pending`, `packing`, `shipped`), default `unpublished`

### ER-диаграммы

**draw.io (ручная, целевая схема):**

Исходник: [img/er-drawio.drawio](../img/er-drawio.drawio)

![ER-диаграмма draw.io](../../img/er-drawio.png)

**DBeaver (фактическая БД после миграции):**

![ER-диаграмма DBeaver](../../img/er-dbeaver.png)

### Сравнение ER-диаграмм

| Проверка | draw.io | DBeaver |
|----------|---------|---------|
| Таблиц | 5 | 5 |
| Схемы `catalog` / `sales` | 3 + 2 | 3 + 2 |
| FK products → product_categories | да | да |
| FK orders → warehouses | да | да |
| FK order_items → orders, products | да | да |
| Составной PK order_items | да | да |

### Заказы — команды

| Команда | Описание |
|---------|----------|
| `list orders` | Список заказов |
| `show order {id}` | Заказ и позиции |
| `add order` | Создание заказа с интерактивным добавлением товаров |
| `edit order {id}` | Изменение склада (только `unpublished`) |
| `delete order {id}` | Удаление (только `unpublished`) |
| `publish order {id}` | Публикация: `unpublished` → `new` |
| `add order_item {order_id}` | Добавить позицию |
| `edit order_item {order_id}` | Изменить цену/количество позиции |
| `delete order_item {order_id}` | Удалить позицию |

### Бизнес-правила заказов

- `status` и `created_at` задаются БД, пользователь их не вводит
- `status` нельзя менять в `edit order` — только через `publish order`
- Редактирование и удаление заказа и позиций доступно только при статусе `unpublished`
- После `publish order` заказ нельзя редактировать и удалять
- `total_amount` = сумма `price × quantity` по всем позициям; пересчитывается автоматически
- Один товар в заказе — один раз (составной PK)
- Выбор склада — `choice()`, товара — autocomplete по SKU, позиции — `choice()`
