package controller

import (
	"github.com/superhacker2002/shop/internal/entity"
	"github.com/superhacker2002/shop/internal/ui"
	"strconv"
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

func (c Controller) ShelfProducts(orders []string) error {
	var orderNums []int
	for _, order := range orders {
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
	ui.OutputData(ordersInfo)
	return nil
}
