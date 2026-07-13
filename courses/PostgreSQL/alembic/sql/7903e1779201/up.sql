CREATE SCHEMA IF NOT EXISTS catalog;
CREATE SCHEMA IF NOT EXISTS sales;

DROP TABLE IF EXISTS sales.order_items;
DROP TABLE IF EXISTS sales.orders;
DROP TABLE IF EXISTS catalog.products;
DROP TABLE IF EXISTS catalog.product_categories;
DROP TABLE IF EXISTS catalog.warehouses;

CREATE TABLE catalog.product_categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL UNIQUE
);

CREATE TABLE catalog.warehouses (
    id SERIAL PRIMARY KEY,
    city VARCHAR NOT NULL,
    address VARCHAR NOT NULL,
    label VARCHAR,
    is_central BOOLEAN NOT NULL
);

CREATE TABLE catalog.products (
    id SERIAL PRIMARY KEY,
    sku VARCHAR(30) NOT NULL UNIQUE,
    name VARCHAR NOT NULL,
    price NUMERIC(20, 4) NOT NULL,
    category_id INTEGER NOT NULL
        REFERENCES catalog.product_categories(id)
);

CREATE TABLE sales.orders (
    id SERIAL PRIMARY KEY,
    status VARCHAR NOT NULL DEFAULT 'unpublished'
        CHECK (status IN (
            'unpublished', 'new', 'processing',
            'pending', 'packing', 'shipped'
        )),
    total_amount NUMERIC(20, 4) NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    warehouse_id INTEGER NOT NULL
        REFERENCES catalog.warehouses(id)
);

CREATE TABLE sales.order_items (
    order_id INTEGER NOT NULL
        REFERENCES sales.orders(id) ON DELETE CASCADE,
    product_id INTEGER NOT NULL
        REFERENCES catalog.products(id),
    price NUMERIC(20, 4) NOT NULL,
    quantity INTEGER NOT NULL CHECK (quantity > 0),
    PRIMARY KEY (order_id, product_id)
);
