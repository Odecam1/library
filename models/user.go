package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement;type:int unsigned"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Loans    []Loan `json:"loans" gorm:"foreignKey:UserID"`
}
