package postgresql

import (
	"database/sql"
	"github.com/lib/pq"
	"github.com/superhacker2002/shop/internal/entity"
)

type PostgresRepository struct {
	db *sql.DB
}

func New(db *sql.DB) PostgresRepository {
	return PostgresRepository{db: db}
}

func (p PostgresRepository) ShelfProducts(orders []int) ([]entity.ShelfProduct, error) {
	//query := `SELECT s.name AS "Стеллаж",
	//				p.name AS "Товар",
	//				p.id AS "id",
	//				o.id AS "заказ",
	//				o.quantity AS "шт"
	//FROM shelves s
	//JOIN products p ON s.id = p.shelf_id
	//JOIN orders o ON p.id = o.product_id
	//WHERE o.id = ANY($1::int[])
	//ORDER BY s.name ASC;`

	query := `SELECT s.name, p.name, p.id, o.id, o.quantity
	FROM shelves s
	JOIN products p ON s.id = p.shelf_id
	JOIN orders o ON p.id = o.product_id
	WHERE order_number = ANY($1)`

	rows, err := p.db.Query(query, pq.Array(orders))
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

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return shelfProducts, nil
}
