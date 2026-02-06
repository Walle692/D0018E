CREATE SCHEMA IF NOT EXISTS myschema;

DROP TABLE IF EXISTS myschema.users;
DROP TABLE IF EXISTS myschema.products;
DROP TABLE IF EXISTS myschema.order;
DROP TABLE IF EXISTS myschema.orderitem;
DROP TABLE IF EXISTS myschema.basket;
DROP TABLE IF EXISTS myschema.basketitem;

CREATE TABLE IF NOT EXISTS myschema.users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(50) NOT NULL,
    role VARCHAR(10) NOT NULL,
    address VARCHAR(50),
    country VARCHAR(37),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS myschema.products (
    product_id SERIAL PRIMARY KEY,
    product_name VARCHAR(100) NOT NULL,
    manufacturer VARCHAR(50) NOT NULL,
    seller_user_id INT NOT NULL,
    description VARCHAR(400),
    screen_size FLOAT NOT NULL,
    picture_url VARCHAR(400),
    sku VARCHAR(50),
    price FLOAT NOT NULL,
    stock INT NOT NULL,
    FOREIGN KEY (seller_user_id) REFERENCES myschema.users(user_id)
)

CREATE TABLE IF NOT EXISTS myschema.order (
    order_id SERIAL PRIMARY KEY,
    order_user_id INT NOT NULL,
    orderdate TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    totalprice FLOAT NOT NULL,
    FOREIGN KEY (customer_user_id) REFERENCES myschema.users(user_id)
)

CREATE TABLE IF NOT EXISTS myschema.orderitem (
    order_item_id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    price FLOAT NOT NULL,
    FOREIGN KEY (order_id) REFERENCES myschema.order(order_id),
    FOREIGN KEY (product_id) REFERENCES myschema.products(product_id)
)

CREATE TABLE IF NOT EXISTS myschema.basket (
    basket_id SERIAL PRIMARY KEY,
    basket_user_id INT NOT NULL,
    FOREIGN KEY (basket_user_id) REFERENCES myschema.users(user_id)
)

CREATE TABLE IF NOT EXISTS myschema.basketitem (
    basket_item_id SERIAL PRIMARY KEY,
    basket_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    price FLOAT NOT NULL,
    FOREIGN KEY (basket_id) REFERENCES myschema.basket(basket_id),
    FOREIGN KEY (product_id) REFERENCES myschema.products(product_id)
)


/* Legacy incase of nuke
CREATE TABLE IF NOT EXISTS myschema.users (
    username VARCHAR(50) UNIQUE NOT NULL PRIMARY KEY,
    password VARCHAR(50) NOT NULL,
    role VARCHAR(10) NOT NULL,
    token VARCHAR(256),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS myschema.token (
    username VARCHAR(50) REFERENCES myschema.users(username) PRIMARY KEY,
    token    VARCHAR(255) NOT NULL
);*/
