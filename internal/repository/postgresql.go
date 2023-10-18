package repository

import (
	"database/sql"
	"fmt"
)

type PostgresRepository struct {
	db *sql.DB
}

type ShelfProduct struct {
	ShelfName   string
	ProductName string
	ProductID   int
	OrderID     int
	Quantity    int
}

func New(db *sql.DB) PostgresRepository {
	return PostgresRepository{db: db}
}

func (p PostgresRepository) ShelfProducts() ([]ShelfProduct, error) {
	query := `SELECT
	s.name AS "Стеллаж",
		p.name AS "Товар",
		p.id AS "id",
		o.id AS "заказ",
		o.quantity AS "шт"
	FROM
	shelves s
	JOIN
	products p ON s.id = p.shelf_id
	JOIN
	orders o ON p.id = o.product_id
	WHERE
	o.id IN (10, 11, 14, 15)
	ORDER BY
	s.name ASC;`

	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shelfProducts []ShelfProduct

	for rows.Next() {
		var sp ShelfProduct
		err = rows.Scan(&sp.ShelfName, &sp.ProductName, &sp.ProductID, &sp.OrderID, &sp.Quantity)
		if err != nil {
			return nil, err
		}
		shelfProducts = append(shelfProducts, sp)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	fmt.Println(shelfProducts)

	return shelfProducts, nil
}
