INSERT INTO shelves (name) VALUES ('A'), ('Б'), ('В'), ('Ж'), ('З');



INSERT INTO products (name, shelf_id, main_shelf) VALUES
                                          ('Ноутбук', 1, true),
                                          ('Телевизор', 1, true),
                                          ('Телефон', 2, true),
                                          ('Системный блок', 4, true),
                                          ('Часы', 4, true),
                                          ('Микрофон', 4, true),
                                          ('Телефон', 3, false),
                                          ('Телефон', 5, false),
                                          ('Часы', 1, false);

INSERT INTO orders (product_id, quantity, order_number) VALUES
                                                            (1, 2, 10),
                                                            (3, 1, 10),
                                                            (6, 1, 10),
                                                            (2, 3, 11),
                                                            (1, 3, 14),
                                                            (4, 4, 14),
                                                            (5, 1, 15);

SELECT s2.name
FROM shelves s2
         JOIN products p2 on s2.id = p.shelf_id
WHERE p2.main_shelf = false AND p2.name = p.name


