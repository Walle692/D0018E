CREATE SCHEMA IF NOT EXISTS myschema;

DROP TABLE IF EXISTS myschema.users;
DROP TABLE IF EXISTS myschema.products;
DROP TABLE IF EXISTS myschema.order;
DROP TABLE IF EXISTS myschema.orderitem;
DROP TABLE IF EXISTS myschema.basket;
DROP TABLE IF EXISTS myschema.basketitem;
DROP TABLE IF EXISTS myschema.rating;

CREATE TABLE IF NOT EXISTS myschema.users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    role VARCHAR NOT NULL,
    address VARCHAR,
    country VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS myschema.products (
    product_id SERIAL PRIMARY KEY,
    product_name VARCHAR NOT NULL,
    manufacturer VARCHAR NOT NULL,
    seller_user_id INT NOT NULL,
    description VARCHAR,
    screen_size FLOAT,
    picture_url VARCHAR,
    price FLOAT NOT NULL,
    stock INT NOT NULL,
    active BOOLEAN DEFAULT TRUE,
    FOREIGN KEY (seller_user_id) REFERENCES myschema.users(user_id)
);

CREATE TABLE IF NOT EXISTS myschema.order (
    order_id SERIAL PRIMARY KEY,
    order_user_id INT NOT NULL,
    orderdate TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    totalprice FLOAT NOT NULL,
    FOREIGN KEY (order_user_id) REFERENCES myschema.users(user_id)
);

CREATE TABLE IF NOT EXISTS myschema.orderitem (
    order_item_id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    price FLOAT NOT NULL,
    FOREIGN KEY (order_id) REFERENCES myschema.order(order_id),
    FOREIGN KEY (product_id) REFERENCES myschema.products(product_id)
);

CREATE TABLE IF NOT EXISTS myschema.basket (
    basket_id SERIAL PRIMARY KEY,
    basket_user_id INT NOT NULL,
    FOREIGN KEY (basket_user_id) REFERENCES myschema.users(user_id)
);

CREATE TABLE IF NOT EXISTS myschema.basketitem (
    basket_item_id SERIAL PRIMARY KEY,
    basket_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    FOREIGN KEY (basket_id) REFERENCES myschema.basket(basket_id),
    FOREIGN KEY (product_id) REFERENCES myschema.products(product_id)
);

CREATE TABLE IF NOT EXITS myschema.rating (
    comment_id SERIAL PRIMARY KEY,
    customer_user_id INT NOT NULL,
    product_id INT NOT NULL,
    comment VARCHAR DEFAULT '',
    rating INT CHECK (rating >= 1 AND rating <= 3),
)

