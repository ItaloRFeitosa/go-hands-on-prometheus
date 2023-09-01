package internal

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type LinkHandler struct {
	linkDAO LinkDAO
	codec   Codec
}

func NewLinkHandler(linkDAO LinkDAO, codec Codec) *LinkHandler {
	return &LinkHandler{linkDAO, codec}
}

func (h *LinkHandler) Shorten(c *fiber.Ctx) error {
	link := new(Link)

	if err := c.BodyParser(link); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(link); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	if err := h.linkDAO.Save(c.UserContext(), link); err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	link.Slug = h.codec.Encode(link.ID)

	return c.Status(http.StatusCreated).JSON(link)
}

func (h *LinkHandler) Redirect(c *fiber.Ctx) error {
	slug := c.Params("slug")
	link, err := h.linkDAO.Get(c.UserContext(), h.codec.Decode(slug))
	if err != nil {
		err = fmt.Errorf("error on redirect with slug { %s }: %w", slug, err)
		return fiber.NewError(http.StatusNotFound, err.Error())
	}

	return c.Redirect(link.URL)
}

func RedirectTest(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendString(`<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Redirect Success</title>
	<style>
	  body {
		margin: 0;
		padding: 0;
		display: flex;
		justify-content: center;
		align-items: center;
		min-height: 100vh;
		background: linear-gradient(to bottom right, #3498db, #e74c3c);
		font-family: Arial, sans-serif;
	  }
	
	  .message-container {
		text-align: center;
		color: white;
	  }
	
	  .message {
		font-size: 3rem;
		font-weight: bold;
		text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5);
	  }
	</style>
	</head>
	<body>
	  <div class="message-container">
		<p class="message">Successful Redirect!!!</p>
	  </div>
	</body>
	</html>
	`)
}
