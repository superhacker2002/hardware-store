CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    shelf_id INTEGER NOT NULL
);

CREATE TABLE shelves (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL
);


