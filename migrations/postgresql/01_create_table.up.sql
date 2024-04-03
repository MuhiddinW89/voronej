CREATE TABLE product (
    product_id SERIAL PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    category VARCHAR(50),
    price DECIMAL(10, 2) NOT NULL
);

CREATE TABLE shelf (
    shelf_id SERIAL PRIMARY KEY,
    shelf_name VARCHAR(50) NOT NULL
);

CREATE TABLE productShelf (
    productshelf_id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    shelf_id INT NOT NULL,
    is_main BOOLEAN NOT NULL DEFAULT FALSE,
    FOREIGN KEY (product_id) REFERENCES product(product_id),
    FOREIGN KEY (shelf_id) REFERENCES shelf(shelf_id)
);

CREATE TABLE orders (
    order_id SERIAL PRIMARY KEY,
    order_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    customer_name VARCHAR(255),
    total_price DECIMAL(10, 2) NOT NULL
);

CREATE TABLE orderDetail (
    order_detail_id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(order_id),
    FOREIGN KEY (product_id) REFERENCES product(product_id)
);
