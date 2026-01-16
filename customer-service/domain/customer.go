package domain

import "time"

type Customer struct {
	ID        int64  `gorm:"column:id;primary_key;default:nextval('customer_id_seq')auto_increment"`
	FirstName string `gorm:"column:first_name;type:varchar(255);not null"`
	LastName  string `gorm:"column:last_name;type:varchar(255);not null"`
	Email     string `gorm:"column:email;not null"`
	Address   Address
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:now()"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null;default:now()"`
}
