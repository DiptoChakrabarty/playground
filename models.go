package main

import (
	"database/sql"

	"gorm.io/gorm"
)

// User has one `Account` (has one), many `Pets` (has many) and `Toys` (has many - polymorphic)
// He works in a Company (belongs to), he has a Manager (belongs to - single-table), and also managed a Team (has many - single-table)
// He speaks many languages (many to many) and has many friends (many to many - single-table)
// His pet also has one Toy (has one - polymorphic)
type User struct {
	Id        uint64 `gorm:"primaryKey" json:"id"`
	FirstName string `json:"first" binding:"required" gorm:"type:varchar(10)"`
	LastName  string `json:"last" binding:"required" gorm:"type:varchar(10)"`
	Email     string `json:"email" binding:"required" gorm:"type:varchar(10)"`
	//Created   string `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

type Account struct {
	gorm.Model
	UserID sql.NullInt64
	Number string
}

type Pet struct {
	gorm.Model
	UserID *uint
	Name   string
	Toy    Toy `gorm:"polymorphic:Owner;"`
}

type Toy struct {
	gorm.Model
	Name      string
	OwnerID   string
	OwnerType string
}

type Company struct {
	ID   int
	Name string
}

type Language struct {
	Code string `gorm:"primarykey"`
	Name string
}
