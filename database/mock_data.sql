INSERT INTO shelves (name) VALUES ('A'), ('Б'), ('Ж');



INSERT INTO products (name, shelf_id) VALUES
                                          ('Ноутбук', 1),
                                          ('Телевизор', 1),
                                          ('Телефон', 2),
                                          ('Системный блок', 3),
                                          ('Часы', 3),
                                          ('Микрофон', 3);

INSERT INTO orders (product_id, quantity, order_number) VALUES
                                                            (1, 2, 10),
                                                            (3, 1, 10),
                                                            (6, 1, 10),
                                                            (2, 3, 11),
                                                            (1, 3, 14),
                                                            (4, 4, 14),
                                                            (5, 1, 15);

