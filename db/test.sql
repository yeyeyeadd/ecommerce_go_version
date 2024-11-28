INSERT INTO users (username, email, password) VALUES
                                                  ('user1', 'user1@example.com', 'hashed_password1'),
                                                  ('user2', 'user2@example.com', 'hashed_password2'),
                                                  ('user3', 'user3@example.com', 'hashed_password3'),
                                                  ('user4', 'user4@example.com', 'hashed_password4'),
                                                  ('user5', 'user5@example.com', 'hashed_password5'),
                                                  ('user6', 'user6@example.com', 'hashed_password6'),
                                                  ('user7', 'user7@example.com', 'hashed_password7'),
                                                  ('user8', 'user8@example.com', 'hashed_password8'),
                                                  ('user9', 'user9@example.com', 'hashed_password9'),
                                                  ('user10', 'user10@example.com', 'hashed_password10');

INSERT INTO products (name, description, price, stock, seller_id) VALUES
                                                                      ('Smartphone', 'Latest model smartphone', 599.99, 50, 1),
                                                                      ('Laptop', 'High-performance laptop', 999.99, 30, 2),
                                                                      ('T-shirt', 'Cotton T-shirt', 19.99, 100, 3),
                                                                      ('Blender', 'Multi-purpose blender', 49.99, 20, 4),
                                                                      ('Novel', 'Bestselling novel', 9.99, 200, 5),
                                                                      ('Toy Car', 'Remote control car', 29.99, 40, 6),
                                                                      ('Football', 'Professional football', 24.99, 60, 7),
                                                                      ('Shampoo', 'Hair care shampoo', 12.99, 80, 8),
                                                                      ('Vitamin C', 'Health supplement', 14.99, 90, 9),
                                                                      ('Car Vacuum', 'Portable car vacuum cleaner', 39.99, 25, 10);


INSERT INTO orders (buyer_id, total_price, status) VALUES
                                                       (1, 79.98, 'completed'),
                                                       (2, 49.99, 'completed'),
                                                       (3, 19.99, 'pending'),
                                                       (4, 99.99, 'completed'),
                                                       (5, 129.98, 'shipped'),
                                                       (6, 9.99, 'completed'),
                                                       (7, 39.99, 'pending'),
                                                       (8, 24.99, 'completed'),
                                                       (9, 12.99, 'completed'),
                                                       (10, 59.98, 'shipped');


INSERT INTO order_items (order_id, product_id, quantity, price) VALUES
                                                                    (1, 4, 1, 599.99),
                                                                    (1, 2, 1, 19.99),
                                                                    (2, 3, 2, 19.99),
                                                                    (3, 4, 1, 49.99),
                                                                    (4, 5, 1, 9.99),
                                                                    (5, 6, 2, 29.99),



INSERT INTO reviews (order_id, rating, comment) VALUES
                                                    (1, 5, 'Amazing product! Highly recommend.'),
                                                    (2, 4, 'Good value for money.'),
                                                    (3, 3, 'Average quality, expected better.'),
                                                    (4, 5, 'Excellent quality and fast delivery.'),
                                                    (5, 4, 'Product was as described.'),
                                                    (6, 2, 'Not satisfied with the packaging.');
