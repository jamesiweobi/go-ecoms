package response

import "time"

type CreateCustomerResponse struct {
	ID         int64           `json:"id"`
	CustomerID int64           `json:"customer_id"`
	FirstName  string          `json:"first_name"`
	LastName   string          `json:"last_name"`
	Email      string          `json:"email"`
	Address    CustomerAddress `json:"address"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
}

type CustomerAddress struct {
	ID          int64  `json:"id"`
	Street      string `json:"street"`
	HouseNumber string `json:"house_number"`
	City        string `json:"city"`
	ZipCode     string `json:"zip_code"`
}
