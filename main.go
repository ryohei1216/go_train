package main

import (
	"fmt"
	"main/app/controllers"
	"main/app/models"
)


func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()

}