package bodyparser

import (
	"encoding/json"
	"github.com/gofiber/fiber/v3"
)

func ParseBody(ctx fiber.Ctx, dest any) error {
	body := ctx.Body()
	err := json.Unmarshal(body, dest)
	if err != nil {
		return err
	}
	return nil
}
