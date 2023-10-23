package controller

import (
	"github.com/superhacker2002/shop/internal/entity"
	"github.com/superhacker2002/shop/internal/ui"
	"strconv"
	"strings"
)

type repository interface {
	ShelfProducts(orders []int) ([]entity.ShelfProduct, error)
}

type Controller struct {
	r repository
}

func New(r repository) Controller {
	return Controller{r: r}
}

func (c Controller) ShelfProducts(orders string) error {
	ordersIDs := strings.Split(orders, ",")
	var orderNums []int
	for _, order := range ordersIDs {
		orderNum, err := strconv.Atoi(order)
		if err != nil {
			return err
		}
		orderNums = append(orderNums, orderNum)
	}
	ordersInfo, err := c.r.ShelfProducts(orderNums)
	if err != nil {
		return err
	}
	ui.OutputData(orders, ordersInfo)
	return nil
}
