package models

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status int         `json:"-"`
	Data   interface{} `json:"data"`
	Err    string      `json:"err"`
}

func (r *Response) Error() string {
	return r.Err
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	res, ok := err.(*Response)
	if ok {
		return ctx.Status(res.Status).JSON(res)
	}
	return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
}
