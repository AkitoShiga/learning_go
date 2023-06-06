package main

import (
	"fmt"
	"todo/controllers"
	"todo/models"
)

func main() {
	fmt.Println(models.Db)
	controllers.StartMainServer()

}
