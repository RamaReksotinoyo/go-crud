package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Order struct {
	Name     string  `form:"name" json:"name"`
	Qty      int     `form:"qty" json:"qty"`
	Prize    float64 `form:"prize" json:"prize"`
	Category string  `form:"category" json:"category"`
}

var orders []Order

func createOrder(c echo.Context) error {
	var order Order

	if err := c.Bind(&order); err != nil {
		return err
	}

	orders = append(orders, order)
	printOrders(orders)

	return c.NoContent(http.StatusCreated)
}

func printOrders(orders []Order) {
	for i, order := range orders {
		fmt.Printf("%d. Name: %s, Qty: %d, Prize: %.2f, Category: %s\n", i+1, order.Name, order.Qty, order.Prize, order.Category)
	}
	fmt.Printf("Total orders: %d\n\n", len(orders))
}

func showOrder(c echo.Context) error {
	orderID, err := strconv.Atoi(c.Param("id"))

	if len(orders) < orderID {
		return c.NoContent(http.StatusNotFound)
	}

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, orders[orderID])
}

func updateOrder(c echo.Context) error {
	orderID, err := strconv.Atoi(c.Param("id"))

	if len(orders) < orderID {
		return c.NoContent(http.StatusNotFound)
	}

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	var order Order
	if err = c.Bind(&order); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	orders[orderID] = order
	return c.NoContent(http.StatusOK)
}

func deleteOrder(c echo.Context) error {
	orderID, err := strconv.Atoi(c.Param("id"))

	if len(orders) < orderID {
		return c.NoContent(http.StatusNotFound)
	}

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	orders = append(orders[:orderID], orders[orderID+1:]...)

	return c.NoContent(http.StatusOK)
}

func main() {
	orders = make([]Order, 0)
	e := echo.New()

	e.POST("/orders", createOrder)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, orders)
	})

	e.GET("/orders/:id", showOrder)
	e.PUT("/orders/:id", updateOrder)
	e.DELETE("/orders/:id", deleteOrder)

	e.Logger.Fatal(e.Start(":1323"))
}

