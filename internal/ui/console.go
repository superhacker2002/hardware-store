package ui

import (
	"fmt"
	"github.com/superhacker2002/shop/internal/entity"
	"strings"
)

func OutputData(orderNums string, orders []entity.ShelfProduct) {
	fmt.Println("Страница сборки заказов", orderNums)
	if len(orders) == 0 {
		fmt.Println("Такого(их) заказа(ов) нет в базе данных")
		return
	}

	var shelf string
	for _, order := range orders {
		if order.ShelfName != shelf {
			fmt.Println("===Стеллаж", order.ShelfName)
		}
		fmt.Printf("%s (id=%d)\n"+
			"заказ %d, %d шт\n",
			order.ProductName, order.ProductID, order.OrderID, order.Quantity)
		if len(order.AdditionalShelves) != 0 {
			fmt.Printf("доп стеллаж: %s\n", strings.Join(order.AdditionalShelves, ","))
		}
		fmt.Print("\n")
		shelf = order.ShelfName
	}
}
