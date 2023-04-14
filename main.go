package main

import (
	"catching-pokemons/controller"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/pokemon/:id", controller.GetPokemon)

	err := router.Run(":8080")
	if err != nil {
		fmt.Print("Error found")
	}
}
