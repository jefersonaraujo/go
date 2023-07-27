package usecase

import "github.com/jefersonaraujo/go/internal/entity"

type OrderInput struct {
	ID     string
	Prince float64
	Tax    float64
}
type OrderOutput struct {
	ID         string
	Prince     float64
	Tax        float64
	FinalPrice float64
}

// SOLID - "D" - Dependency Inversion Principle
type CalculateFinalPrince struct {
	OrderRepository entity.OrderRepositoryInterface
}

func (c *CalculateFinalPrince) Execute(input OrderInput) (*OrderOutput, error) {
	order, err := entity.NewOrder(input.ID, input.Prince, input.Tax)
	if err != nil {
		return nil, err
	}
	err = order.CalculateFinalPrince()
	if err != nil {
		return nil, err
	}

	err = c.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}

	return &OrderOutput{
		ID:         order.ID,
		Prince:     order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
