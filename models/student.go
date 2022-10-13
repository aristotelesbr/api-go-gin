package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name      string `json:"name"`
	CPF       string `json:"cpf"`
	DocNumber string `json:"doc_number"`
}

var Students []Student
