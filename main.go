package main

import (
	"github.com/aristotelesbr/api-go-gin/models"
	"github.com/aristotelesbr/api-go-gin/routes"
)

func main() {
	models.Students = []models.Student{
		{Name: "Ari", CPF: "1234", DocNumber: "321"},
		{Name: "Ana", CPF: "999", DocNumber: "888"},
	}
	routes.HandleRequest()
}
