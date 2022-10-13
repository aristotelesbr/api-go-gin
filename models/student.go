package models

type Student struct {
	Name      string `json:"name"`
	CPF       string `json:"cpf"`
	DocNumber string `json:"doc_number"`
}

var Students []Student
