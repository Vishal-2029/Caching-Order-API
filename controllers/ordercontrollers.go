package controllers

import (
	db "Exercise/OrderAPI/config"
	"Exercise/OrderAPI/models"
	"fmt"

	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

func PlaceOrder(c *fiber.Ctx) error {
	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	if order.Items == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "No items selected!",
		})
	}

	order.Status = "Pending"
	db.DB.Create(&order)
	return c.Status(201).JSON(order)
}

func ApproveOrder(c *fiber.Ctx) error {
	var order models.Order
	CustomerId := c.Params("CustomerId")

	if err := db.DB.First(&order, CustomerId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Order not found",
		})
	}

	order.Status = "Approved"
	db.DB.Save(&order)
	return c.JSON(order)

}

func CancelOrder(c *fiber.Ctx) error {
	var order models.Order
	CustomerId := c.Params("CustomerId")

	if err := db.DB.First(&order, CustomerId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Order not found",
		})
	}

	order.Status = "Cancel"
	db.DB.Save(&order)
	return c.JSON(order)
}

var billCache = cache.New(10*time.Minute, 15*time.Second)

func GetBill(c *fiber.Ctx) error {
	var order models.Order
	CustomerId := c.Params("CustomerId")

	if cachedBill, found := billCache.Get(CustomerId); found {
		fmt.Println("Cache hit")
		return c.JSON(cachedBill)
	}

	if err := db.DB.First(&order, CustomerId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Order not found",
		})
	}

	if order.Status != "Approved" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Order is not approved!",
		})
	}

	bill := fiber.Map{
		"ID":         order.ID,
		"items":      order.Items,
		"totalprice": order.TotalPrice,
		"status":     order.Status,
		"message":    "Bill generated successfully",
	}

	billCache.Set(CustomerId, bill, cache.DefaultExpiration)

	return c.JSON(bill)
}
