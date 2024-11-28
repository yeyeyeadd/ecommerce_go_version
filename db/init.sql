DROP DATABASE IF EXISTS ecommerce_c2c;
CREATE DATABASE ecommerce_c2c CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

use ecommerce_c2c;

DROP TABLE IF EXISTS users;
CREATE TABLE users (
                       id INT AUTO_INCREMENT PRIMARY KEY,
                       username VARCHAR(50) UNIQUE NOT NULL,
                       email VARCHAR(100) UNIQUE NOT NULL,
                       password VARCHAR(255) NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS products;
CREATE TABLE products (
                          id INT AUTO_INCREMENT PRIMARY KEY,
                          name VARCHAR(100) NOT NULL,
                          description TEXT,
                          price DECIMAL(10, 2) NOT NULL,
                          stock INT NOT NULL DEFAULT 0,
                          seller_id INT NOT NULL,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          FOREIGN KEY (seller_id) REFERENCES users(id)
);

DROP TABLE IF EXISTS orders;
CREATE TABLE orders (
                        id INT AUTO_INCREMENT PRIMARY KEY,
                        buyer_id INT NOT NULL,
                        total_price DECIMAL(10, 2) NOT NULL,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        FOREIGN KEY (buyer_id) REFERENCES users(id)
);

DROP TABLE IF EXISTS order_items;
CREATE TABLE order_items (
                             id INT AUTO_INCREMENT PRIMARY KEY,
                             order_id INT NOT NULL,
                             product_id INT NOT NULL,
                             quantity INT NOT NULL,
                             price DECIMAL(10, 2) NOT NULL,
                             FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
                             FOREIGN KEY (product_id) REFERENCES products(id)
);

DROP TABLE IF EXISTS reviews;
CREATE TABLE reviews (
                         id INT AUTO_INCREMENT PRIMARY KEY,
                         user_id INT NOT NULL,
                         product_id INT NOT NULL,
                         order_id INT NOT NULL,
                         rating INT CHECK (rating BETWEEN 1 AND 5),
                         comment TEXT,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         FOREIGN KEY (user_id) REFERENCES users(id),
                         FOREIGN KEY (product_id) REFERENCES products(id),
                         FOREIGN KEY (order_id) REFERENCES orders(id)
);



