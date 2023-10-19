package controller

import (
	"github.com/superhacker2002/shop/internal/entity"
	"github.com/superhacker2002/shop/internal/ui"
	"strings"
)

type repository interface {
	ShelfProducts(orders string) ([]entity.ShelfProduct, error)
}

type Controller struct {
	r repository
}

func New(r repository) Controller {
	return Controller{r: r}
}

func (c Controller) ShelfProducts(orderNums []string) error {
	ordersStr := strings.Join(orderNums, " ")
	ordersInfo, err := c.r.ShelfProducts(ordersStr)
	if err != nil {
		return err
	}
	ui.OutputData(ordersInfo)
	return nil
}
