package main

import (
	"encoding/json"
	"fmt"
	"os"
	"pizzaria/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var pizzas []models.Pizza

func main() {
	loadPizzas()
	r := gin.Default()
	r.GET("pizzas", getPizzas)
	r.GET("/pizzas/:id", getPizzasByID)
	r.POST("/pizzas", postPizzas)
	r.Run()
}

func getPizzas(c *gin.Context) {
	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}

func postPizzas(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	newPizza.ID = len(pizzas) + 1
	pizzas = append(pizzas, newPizza)
	if err := savePizza(); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(201, newPizza)
}

func getPizzasByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	for _, p := range pizzas {
		if p.ID == id {
			c.JSON(200, p)
			return
		}
	}
	c.JSON(404, gin.H{"message": "Pizza not found"})
}

func loadPizzas() {
	file, err := os.Open("data/pizza.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&pizzas); err != nil {
		fmt.Println("Error decoding JSON:", err)
	}
}

func savePizza() error {
	if err := os.MkdirAll("data", os.ModePerm); err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}
	file, err := os.Create("data/pizza.json")
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(pizzas); err != nil {
		return fmt.Errorf("error encoding JSON: %w", err)
	}
	return nil
}
