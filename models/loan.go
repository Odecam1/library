package models

import "time"

type Loan struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	UserID     uint      `json:"user_id"`
	ResourceID uint      `json:"resource_id"`
	LoanDate   time.Time `json:"loan_date"`
	ReturnDate time.Time `json:"return_date"`
	Status     string    `json:"status"`
}
