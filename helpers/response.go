package helpers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// Response нь амжилттай хариу буцаана (HTTP 200).
//
// Response(c) -> {"message": "Амжилттай"}
// Response(c, "Амжилттай бүртгэгдлээ") -> {"message": "Амжилттай бүртгэгдлээ"}
// Response(c, anyData) -> anyData-г JSON болгон буцаана
func Response(c *fiber.Ctx, data ...interface{}) error {
	if len(data) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Амжилттай"})
	}

	switch val := data[0].(type) {
	case string:
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": val})
	default:
		return c.Status(fiber.StatusOK).JSON(val)
	}
}

// ResponseCreated нь амжилттай үүсгэсэн хариу буцаана (HTTP 201).
func ResponseCreated(c *fiber.Ctx, data ...interface{}) error {
	if len(data) == 0 {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Амжилттай үүсгэлээ"})
	}

	switch val := data[0].(type) {
	case string:
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": val})
	default:
		return c.Status(fiber.StatusCreated).JSON(val)
	}
}

// ResponseBadRequest нь буруу хүсэлтийн хариу буцаана (HTTP 400).
func ResponseBadRequest(c *fiber.Ctx, message string) error {
	zap.L().Error("Bad Request", zap.String("error", message))
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": message})
}

// ResponseUnauthorized нь зөвшөөрөлгүй хариу буцаана (HTTP 401).
func ResponseUnauthorized(c *fiber.Ctx) error {
	msg := "Токен хугацаа дууссан эсвэл буруу байна"
	zap.L().Error("Unauthorized", zap.String("error", msg))
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": msg})
}

// ResponseForbidden нь хандах эрхгүй хариу буцаана (HTTP 403).
func ResponseForbidden(c *fiber.Ctx) error {
	msg := "Хандах эрхгүй"
	zap.L().Error("Forbidden", zap.String("error", msg))
	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": msg})
}

// ResponseNotFound нь олдоогүй хариу буцаана (HTTP 404).
func ResponseNotFound(c *fiber.Ctx, message string) error {
	zap.L().Error("Not Found", zap.String("error", message))
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": message})
}

// ResponseErr нь серверийн алдааны хариу буцаана (HTTP 500).
func ResponseErr(c *fiber.Ctx, message string) error {
	zap.L().Error("Internal Server Error", zap.String("error", message))
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": message})
}
