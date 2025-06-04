package models

import (
	"gorm.io/gorm"
)

type PostgresConfiguration struct {
    Host     string
    Port     string
    Username string
    Password string
    DBName   string
    SSLMode  string
}

type DatabaseConfiguration struct {
    GormDB *gorm.DB
}

type User struct {
	ID                           uint64    `gorm:"primaryKey" json:"id"`
	Username                     string    `gorm:"column:username" json:"username"`
	Email                        string    `gorm:"column:email" json:"email"`
	PhoneNumber                  string    `gorm:"column:\"phonenumber\"" json:"phonenumber"`
	Age                          int       `gorm:"column:age" json:"age"`
	Password                     string    `gorm:"column:password" json:"password"`
}
