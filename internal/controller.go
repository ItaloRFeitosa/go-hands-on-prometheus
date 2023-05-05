package internal

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type LinkController struct {
	linkRepo LinkRepository
	codec    Codec
}

func NewLinkController(linkRepo LinkRepository, codec Codec) *LinkController {
	return &LinkController{linkRepo, codec}
}

func (ctl *LinkController) Shorten(c *fiber.Ctx) error {
	link := new(Link)

	if err := c.BodyParser(&link); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(link); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	if err := ctl.linkRepo.Save(c.UserContext(), link); err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	link.Slug = ctl.codec.Encode(link.ID)

	return c.Status(201).JSON(link)
}

func (ctl *LinkController) Redirect(c *fiber.Ctx) error {
	link, err := ctl.linkRepo.Get(c.UserContext(), ctl.codec.Decode(c.Params("slug")))
	if err != nil {
		return fiber.NewError(http.StatusNotFound, err.Error())
	}

	return c.Redirect(link.URL)
}
