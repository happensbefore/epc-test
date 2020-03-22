package product

import (
	"errors"
	"net/http"

	"epc-test/models"
)

type CalculateProductDTO struct {
	Product    models.Product
	Conditions []models.Condition
}

func (req *CalculateProductDTO) Check() (error, int) {
	if req.Product.Components == nil || len(req.Product.Components) == 0 {
		return errors.New("Must be at least one product component presented"), http.StatusBadRequest
	}

	var haveMain bool

	for _, c := range req.Product.Components {
		if c.IsMain {
			haveMain = true
		}
	}

	if !haveMain {
		return errors.New("Must be at least one main component in product"), http.StatusBadRequest
	}

	if req.Conditions == nil || len(req.Conditions) == 0 {
		return errors.New("Must be presented at least one condition"), http.StatusBadRequest
	}

	return nil, 0
}
