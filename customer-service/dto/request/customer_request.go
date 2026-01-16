package request

type CreateCustomerRequest struct {
	FirstName string          `json:"first_name" validate:"required"`
	LastName  string          `json:"last_name" validate:"required"`
	Email     string          `json:"email" validate:"required,email"`
	Address   CustomerAddress `json:"address" validate:"required"`
}

type CustomerAddress struct {
	Street      string `json:"street" validate:"required"`
	HouseNumber string `json:"house_number" validate:"required"`
	City        string `json:"city"`
	ZipCode     string `json:"zip_code" validate:"required"`
}
