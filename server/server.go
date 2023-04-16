package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"goworkshop/handler"
	"goworkshop/repository"
)

type Server struct {
	app *fiber.App
}

func NewServer(productHandler handler.ProductHandlerInterface) *Server {
	app := fiber.New()

	app.Get("/products", func(c *fiber.Ctx) error {
		return c.JSON(productHandler.ListProducts())
	})

	app.Post("/products", func(c *fiber.Ctx) error {
		var product repository.Product
		if err := c.BodyParser(&product); err != nil {
			return err
		}

		if _, err := productHandler.CreateProduct(&product); err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(product)
	})

	return &Server{app: app}
}

func (server Server) Listen(port string) error {
	return server.app.Listen(port)
}

func (server Server) Test(req *http.Request) (*http.Response, error) {
	return server.app.Test(req)
}
