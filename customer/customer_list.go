package customer

import (
	"fmt"
	"sort"
	"strings"
)

type Customers struct {
	Invited []Customer
}

func NewCustomerList() *Customers {
	return &Customers{
		Invited: make([]Customer, 0),
	}
}

func (customers *Customers) Add(customer *Customer) {
	customers.Invited = append(customers.Invited, *customer)
}

func (customers *Customers) Print() string {
	if len(customers.Invited) == 0 {
		return "Any customer will be invited..."
	}

	sort.Sort(customers)

	invites := make([]string, len(customers.Invited))
	for k, customer := range customers.Invited {
		invites[k] = "- " + customer.String()
	}

	return fmt.Sprintf("Prepare your party we're going to invite %d customers:\n%s", len(invites), strings.Join(invites, "\n"))
}

func (c *Customers) Len() int {
	return len(c.Invited)
}
func (c *Customers) Swap(i, j int) {
	c.Invited[i], c.Invited[j] = c.Invited[j], c.Invited[i]
}
func (c *Customers) Less(i, j int) bool {
	return c.Invited[i].UserID < c.Invited[j].UserID
}
