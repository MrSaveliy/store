-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS users
(
   id          SERIAL PRIMARY KEY,
   email VARCHAR(255),
   password_hash VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS profiles
(
   id          SERIAL PRIMARY KEY,
   name        VARCHAR(255),
   surname     VARCHAR(255),
   patronymic  VARCHAR(255),
   password_hash VARCHAR(255),
   adress      VARCHAR(255),
   number      VARCHAR(255),
   date        DATE,
   user_id INT REFERENCES users(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS categories
(
   id          SERIAL PRIMARY KEY,
   name        VARCHAR(255),
   description VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS products
(
   id           SERIAL PRIMARY KEY,
   name         VARCHAR(255),
   description  VARCHAR(255),
   price        INT,
   quantity     INT,
   category_id  INT REFERENCES categories(id) ON DELETE SET NULL,
   img          VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS orders
(
   id           SERIAL PRIMARY KEY,
   user_id      INT,
   date         DATE,
   status       VARCHAR(255),
   total_price  INT,
   adress       VARCHAR(255),
   FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS order_items
(
   id          SERIAL PRIMARY KEY,
   order_id    INT REFERENCES orders(id) ON DELETE SET NULL,
   product_id  INT REFERENCES products(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS bascket
(
   id          SERIAL PRIMARY KEY,
   user_id    INT REFERENCES orders(id) ON DELETE SET NULL,
   product_id  INT REFERENCES products(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS favorites
(
   id          SERIAL PRIMARY KEY,
   user_id    INT REFERENCES orders(id) ON DELETE SET NULL,
   product_id  INT REFERENCES products(id) ON DELETE SET NULL
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE basket;
DROP TABLE favorites;
DROP TABLE order_items;
DROP TABLE orders;
DROP TABLE profiles;
DROP TABLE users;
DROP TABLE products;
DROP TABLE categories;

-- +goose StatementEnd
