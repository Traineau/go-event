DROP TABLE IF EXISTS FoodOrderLine;

DROP TABLE IF EXISTS FoodOrder;

CREATE TABLE IF NOT EXISTS FoodOrder (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    reference TEXT NOT NULL,
    client TEXT NOT NULL,
    date DATETIME,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS FoodOrderLine (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    meal TEXT NOT NULL,
    price BIGINT NOT NULL,
    id_food_order BIGINT UNSIGNED,
    PRIMARY KEY(id),
    FOREIGN KEY (id_food_order) REFERENCES FoodOrder(id) ON DELETE CASCADE
);

INSERT INTO FoodOrder (id, reference, client, date) VALUES
	(1, 'ER45T', 'Edwin', '2020-07-22 15:41:11'),
	(2, 'FG21S', 'Jason', '2020-07-22 13:21:11'),
	(3, '21TZE', 'Thomas', '2020-07-21 12:01:11');

INSERT INTO FoodOrderLine (id, meal, price, id_food_order) VALUES
	(1, 'Steak', '1400', 1),
	(2, 'Paté', '600', 1),
	(3, 'Glace', '400', 1),
	(4, 'Patates', '800', 2),
	(5, 'Menu B2', '2100', 3);

