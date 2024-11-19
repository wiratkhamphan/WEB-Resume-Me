package getproduct

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wiratkhamphan/WEBResumeMe/config/database"
	models "github.com/wiratkhamphan/WEBResumeMe/models/Product"
)

// Get_product_list retrieves a list of products from the database
func Getporductlist(c *fiber.Ctx) error {
	// Connect to the database
	db, err := database.Connect()
	if err != nil {
		log.Println("Database connection error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to connect to the database",
		})
	}
	defer db.Close()

	// Define a slice to hold the products
	var products []models.Product

	// Execute the query
	rows, err := db.Query("SELECT id, name, price, qty, detail FROM product")
	if err != nil {
		log.Println("Query execution error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch products",
		})
	}
	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Qty, &product.Detail); err != nil {
			log.Println("Row scan error:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to parse product data",
			})
		}
		products = append(products, product)
	}

	// Check for errors from the iteration
	if err := rows.Err(); err != nil {
		log.Println("Rows iteration error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error during rows iteration",
		})
	}

	// Return the product list as JSON
	return c.JSON(products)
}
