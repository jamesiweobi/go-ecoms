package domain

import "time"

type Address struct {
	ID          int64     `gorm:"column:id;primary_key;default:nextval('customer_id_seq')"`
	CustomerID  int64     `gorm:"column:customer_id;not null"`
	Street      string    `gorm:"column:street;not null"`
	HouseNumber string    `gorm:"column:house_number;not null"`
	City        string    `gorm:"column:city;not null"`
	ZipCode     string    `gorm:"column:zip_code;not null"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp;not null;default:now()"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp;not null;default:now()"`
}
