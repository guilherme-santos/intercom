package customer_test

import (
	"testing"

	"github.com/guilherme-santos/intercom/customer"
	"github.com/stretchr/testify/assert"
)

func TestAddNewCustomers(t *testing.T) {
	customers := customer.NewCustomerList()
	assert.Len(t, customers.Invited, 0)

	customers.Add(&customer.Customer{})
	assert.Len(t, customers.Invited, 1)

	customers.Add(&customer.Customer{})
	assert.Len(t, customers.Invited, 2)
}
