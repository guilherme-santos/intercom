package customer_test

import (
	"testing"

	"github.com/guilherme-santos/intercom/customer"
	"github.com/stretchr/testify/assert"
)

func TestNewCustomer_InvalidJson(t *testing.T) {
	data := `invalid json`
	_, err := customer.NewCustomer([]byte(data))
	assert.Error(t, err)
}

func TestNewCustomer_InvalidUserID(t *testing.T) {
	data := `{"user_id": "invalid", "name": "Christina McArdle", "latitude": "52.986375", "longitude": "-6.043701"}`
	_, err := customer.NewCustomer([]byte(data))
	assert.Error(t, err)
}

func TestNewCustomer_InvalidLatitude(t *testing.T) {
	data := `{"user_id": "12", "name": "Christina McArdle", "latitude": "invalid", "longitude": "-6.043701"}`
	_, err := customer.NewCustomer([]byte(data))
	assert.Error(t, err)
}

func TestNewCustomer_InvalidLongitude(t *testing.T) {
	data := `{"user_id": "12", "name": "Christina McArdle", "latitude": "52.986375", "longitude": "invalid"}`
	_, err := customer.NewCustomer([]byte(data))
	assert.Error(t, err)
}

func TestNewCustomer_ValidJson(t *testing.T) {
	data := `{"user_id": 12, "name": "Christina McArdle", "latitude": "52.986375", "longitude": "-6.043701"}`
	obj, err := customer.NewCustomer([]byte(data))
	assert.NoError(t, err)
	assert.IsType(t, &customer.Customer{}, obj)
}

func TestNewCustomerByObject_MissingUserID(t *testing.T) {
	data := map[string]interface{}{}
	_, err := customer.NewCustomerByObject(data)
	assert.Error(t, err)
}

func TestNewCustomerByObject_UserIDAsStringInvalid(t *testing.T) {
	data := map[string]interface{}{"user_id": "invalid"}
	_, err := customer.NewCustomerByObject(data)
	assert.Error(t, err)
}

func TestNewCustomerByObject_UserIDAsStringWithNumber(t *testing.T) {
	data := map[string]interface{}{"user_id": "12"}
	_, err := customer.NewCustomerByObject(data)
	assert.Error(t, err)
}

func TestNewCustomerByObject_UserIDAsInteger(t *testing.T) {
	data := map[string]interface{}{"user_id": int64(12), "name": "Christina McArdle", "latitude": 52.986375, "longitude": -6.043701}
	_, err := customer.NewCustomerByObject(data)
	assert.NoError(t, err)
}

func TestNewCustomerByObject_MissingName(t *testing.T) {
	data := map[string]interface{}{"user_id": int64(12)}
	_, err := customer.NewCustomerByObject(data)
	assert.Error(t, err)
}

func TestNewCustomerByObject_NameAsInteger(t *testing.T) {
	data := map[string]interface{}{"user_id": int64(12), "name": 123}
	_, err := customer.NewCustomerByObject(data)
	assert.Error(t, err)
}

func TestNewCustomerByObject_MissingLatitude(t *testing.T) {
	data := map[string]interface{}{"user_id": int64(12), "name": "Christina McArdle"}
	_, err := customer.NewCustomerByObject(data)
	assert.Error(t, err)
}

func TestNewCustomerByObject_LatitudeAsStringInvalid(t *testing.T) {
	data := map[string]interface{}{"user_id": int64(12), "name": "Christina McArdle", "latitude": "invalid", "longitude": "-6.043701"}
	_, err := customer.NewCustomerByObject(data)
	assert.Error(t, err)
}

func TestNewCustomerByObject_LatitudeAsStringWithNumber(t *testing.T) {
	data := map[string]interface{}{"user_id": int64(12), "name": "Christina McArdle", "latitude": "52.986375", "longitude": "-6.043701"}
	_, err := customer.NewCustomerByObject(data)
	assert.Error(t, err)
}

func TestNewCustomerByObject_LatitudeAsFloat(t *testing.T) {
	data := map[string]interface{}{"user_id": int64(12), "name": "Christina McArdle", "latitude": 52.986375, "longitude": -6.043701}
	_, err := customer.NewCustomerByObject(data)
	assert.NoError(t, err)
}

func TestNewCustomerByObject_MissingLongitude(t *testing.T) {
	data := map[string]interface{}{"user_id": int64(12), "name": "Christina McArdle", "latitude": "52.986375"}
	_, err := customer.NewCustomerByObject(data)
	assert.Error(t, err)
}

func TestNewCustomerByObject_LongituteAsStringInvalid(t *testing.T) {
	data := map[string]interface{}{"user_id": int64(12), "name": "Christina McArdle", "latitude": 52.986375, "longitude": "invalid"}
	_, err := customer.NewCustomerByObject(data)
	assert.Error(t, err)
}

func TestNewCustomerByObject_LongituteAsStringWithNumber(t *testing.T) {
	data := map[string]interface{}{"user_id": int64(12), "name": "Christina McArdle", "latitude": 52.986375, "longitude": "-6.043701"}
	_, err := customer.NewCustomerByObject(data)
	assert.Error(t, err)
}

func TestNewCustomerByObject_LongituteAsFloat(t *testing.T) {
	data := map[string]interface{}{"user_id": int64(12), "name": "Christina McArdle", "latitude": 52.986375, "longitude": -6.043701}
	_, err := customer.NewCustomerByObject(data)
	assert.NoError(t, err)
}

func TestCustomerShouldInvite_LessThan100Km(t *testing.T) {
	customer := &customer.Customer{
		Latlong: customer.Latlong{53.1302756, -6.2397222},
	}
	assert.True(t, customer.ShouldInvite())
}

func TestCustomerShouldInvite_MoreThan100Km(t *testing.T) {
	customer := &customer.Customer{
		Latlong: customer.Latlong{52.986375, -6.043701},
	}
	assert.True(t, customer.ShouldInvite())
}
