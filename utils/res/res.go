package res

import (
	"github.com/gofiber/fiber/v2"
)

type Code int

const (
	SuccessCode       Code = iota
	ValidateErrorCode Code = 7
)

type Response struct {
	Code   Code   `json:"code"` // 业务状态码 0 为成功
	Msg    string `json:"msg"`
	Data   any    `json:"data"`
	Reason string `json:"reason"`
}

func response(c *fiber.Ctx, code Code, data any, msg string, reason string) error {
	res := Response{
		Code:   code,
		Msg:    msg,
		Data:   data,
		Reason: reason,
	}
	c.Status(fiber.StatusOK)
	return c.JSON(res)
}

func Ok(c *fiber.Ctx, data any, msg string, reason string) error {
	return response(c, SuccessCode, data, msg, reason)
}

func OkWithData(c *fiber.Ctx, data any) error {
	return Ok(c, data, "成功", "")
}

func OkWithMsgAndData(c *fiber.Ctx, msg string, data any) error {
	return Ok(c, data, msg, "")
}
func OkWithMsg(c *fiber.Ctx, msg string) error {
	return Ok(c, map[string]any{}, msg, "")
}

func Fail(c *fiber.Ctx, code Code, data any, msg string, reason string) error {
	return Ok(c, map[string]any{}, msg, reason)
}

func FailWithMsg(c *fiber.Ctx, msg string) error {
	return Fail(c, ValidateErrorCode, map[string]any{}, msg, "")
}

func FailWithMsgAndReason(c *fiber.Ctx, msg string, reason string) error {
	return Fail(c, ValidateErrorCode, map[string]any{}, msg, reason)
}

func FailWithCode(c *fiber.Ctx, code Code) error {
	return Fail(c, code, map[string]any{}, "", "")
}
