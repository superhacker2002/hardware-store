package controller

import (
	"github.com/superhacker2002/shop/internal/entity"
	"github.com/superhacker2002/shop/internal/ui"
	"strings"
)

type repository interface {
	ShelfProducts(orders []string) ([]entity.ShelfProduct, error)
}

type Controller struct {
	r repository
}

func New(r repository) Controller {
	return Controller{r: r}
}

func (c Controller) ShelfProducts(orders string) error {
	ordersIDs := strings.Split(orders, ",")
	ordersInfo, err := c.r.ShelfProducts(ordersIDs)
	if err != nil {
		return err
	}
	ui.OutputData(orders, ordersInfo)
	return nil
}
