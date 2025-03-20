package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Loans    []Loan `json:"loans" gorm:"foreignKey:UserID"`
}
