package main

import (
	"fmt"
	"github.com/LuxAeterna-git/my-parcer/pkg/repository"
	"github.com/LuxAeterna-git/my-parcer/pkg/service"
)

func main() {

	db := repository.NewPg()
	parser := service.NewService(db)

	parser.ParseGoods()
	all := db.FindAll()
	for _, e := range all {
		fmt.Println("ID:", e.Id, "\n", "Name:", e.Name, "\n", "Price:", e.Price, "\n", "URL:", e.Url, "\n", "ImageURL:", e.UrlImg)
		fmt.Println("\n", "==============================================", "\n")
	}

}
