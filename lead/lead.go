package lead

import (
	"github.com/RakibSiddiquee/go-fiber-crm/database"
	"github.com/gofiber/fiber/v2"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) error {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	return c.JSON(leads)
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	return c.JSON(lead)
}

func NewLead(c *fiber.Ctx) error {
	db := database.DBConn
	lead := new(Lead)
	err := c.BodyParser(lead)
	if err != nil {
		c.Status(503).Send([]byte(err.Error()))
	}
	db.Create(&lead)
	return c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send([]byte("No lead found with ID"))
	}

	db.Delete(&lead)
	return c.Send([]byte("Lead successfully deleted"))
}
