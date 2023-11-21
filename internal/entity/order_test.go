package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfItGetAnErrorIfIdsBlank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "id is required")

}

func TestIfItGetsAnErrorIfPriceIsBlank(t *testing.T) {
	order := Order{Id: "123"}
	assert.Error(t, order.Validate(), "price must be greater than zero")
}

func TestIfItGetsAnErrorIfTaxIsBlank(t *testing.T) {
	order := Order{Id: "123", Price: 10.0}
	assert.Error(t, order.Validate(), "invalid tax")
}
