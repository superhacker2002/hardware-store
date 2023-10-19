package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/superhacker2002/shop/internal/entity"
)

type PostgresRepository struct {
	db *sql.DB
}

func New(db *sql.DB) PostgresRepository {
	return PostgresRepository{db: db}
}

func (p PostgresRepository) ShelfProducts(orders string) ([]entity.ShelfProduct, error) {
	query := fmt.Sprintf(`SELECT
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
	o.id IN (%s)
	ORDER BY
	s.name ASC;`, orders)

	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shelfProducts []entity.ShelfProduct

	for rows.Next() {
		var sp entity.ShelfProduct
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
