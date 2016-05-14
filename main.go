package main

import (
	"bufio"
	"os"

	"github.com/NeowayLabs/logger"
	"github.com/guilherme-santos/intercom/customer"
)

func main() {
	file, err := os.Open("./customers.txt")
	if err != nil {
		logger.Fatal("Error opening file: %s", err.Error())
	}
	defer file.Close()

	customers := customer.NewCustomerList()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		customer, err := customer.NewCustomer(scanner.Bytes())
		if err != nil {
			logger.Fatal("Error creating customer: %s", err.Error())
		}

		if !customer.ShouldInvite() {
			logger.Debug("Ignoring " + customer.String())
			continue
		}

		customers.Add(customer)
	}

	if err := scanner.Err(); err != nil {
		logger.Fatal("Error reading line: %s", err.Error())
	}

	logger.Info(customers.Print())
}
