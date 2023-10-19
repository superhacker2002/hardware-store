package ui

import (
	"fmt"
	"github.com/superhacker2002/shop/internal/entity"
)

func OutputData(orders []entity.ShelfProduct) {
	fmt.Println(orders)
}
