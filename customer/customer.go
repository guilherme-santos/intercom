package customer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math"
)

type (
	Customer struct {
		UserID  int64 `json:"user_id"`
		Name    string
		Latlong `json:",inline"`
	}

	Latlong struct {
		Latitude, Longitude float64
	}
)

var DublinOffice = Latlong{
	Latitude:  53.3381985,
	Longitude: -6.2592576,
}

func NewCustomer(data []byte) (*Customer, error) {
	// All fields are string cannot unmarshal directly to Customer
	var (
		customerData map[string]interface{}
		err          error
	)

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	err = decoder.Decode(&customerData)
	if err != nil {
		return nil, err
	}

	if customerData["user_id"] != nil {
		if userID, ok := customerData["user_id"].(json.Number); ok {
			customerData["user_id"], err = userID.Int64()
			if err != nil {
				return nil, err
			}
		}
	}
	if customerData["latitude"] != nil {
		if latitudeStr, ok := customerData["latitude"].(string); ok {
			customerData["latitude"] = json.Number(latitudeStr)
		}

		if latitude, ok := customerData["latitude"].(json.Number); ok {
			customerData["latitude"], err = latitude.Float64()
			if err != nil {
				return nil, err
			}
		}
	}
	if customerData["longitude"] != nil {
		if longitudeStr, ok := customerData["longitude"].(string); ok {
			customerData["longitude"] = json.Number(longitudeStr)
		}

		if longitude, ok := customerData["longitude"].(json.Number); ok {
			customerData["longitude"], err = longitude.Float64()
			if err != nil {
				return nil, err
			}
		}
	}

	return NewCustomerByObject(customerData)
}

func NewCustomerByObject(data map[string]interface{}) (*Customer, error) {
	if data["user_id"] == nil {
		return nil, errors.New("Field 'user_id' is missing")
	}
	if data["name"] == nil {
		return nil, errors.New("Field 'name' is missing")
	}
	if data["latitude"] == nil {
		return nil, errors.New("Field 'latitude' is missing")
	}
	if data["longitude"] == nil {
		return nil, errors.New("Field 'longitude' is missing")
	}

	userID, ok := data["user_id"].(int64)
	if !ok {
		return nil, errors.New("Field 'user_id' should be an integer")
	}

	name, ok := data["name"].(string)
	if !ok {
		return nil, errors.New("Field 'name' should be a string")
	}

	latitude, ok := data["latitude"].(float64)
	if !ok {
		return nil, errors.New("Field 'latitude' should be a float")
	}

	longitude, ok := data["longitude"].(float64)
	if !ok {
		return nil, errors.New("Field 'longitude' should be a float")
	}

	return &Customer{
		UserID: userID,
		Name:   name,
		Latlong: Latlong{
			Latitude:  latitude,
			Longitude: longitude,
		},
	}, nil
}

func (customer *Customer) String() string {
	return fmt.Sprintf("id[%02d] name[%s]", customer.UserID, customer.Name)
}

func (customer *Customer) ShouldInvite() bool {
	dlat := customer.Latitude - DublinOffice.Latitude
	dlong := customer.Longitude - DublinOffice.Longitude
	calc := math.Pow(math.Sin(dlat/2), 2) + math.Cos(customer.Latitude)*math.Cos(DublinOffice.Latitude)*math.Pow(math.Sin(dlong/2), 2)
	calc = 2 * math.Asin(math.Sqrt(calc))
	distance := 6371 * calc * 0.0174533

	if distance < 100 {
		return true
	}

	return false
}
