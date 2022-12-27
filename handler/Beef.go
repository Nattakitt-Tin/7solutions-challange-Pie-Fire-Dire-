package handler

import (
	"PieFireDire/service"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetBeefSummary(c echo.Context) error {
	b := service.NewBeef()
	text, err := b.GetText()
	if err != nil {
		c.String(http.StatusBadRequest, "error get text from url")
	}
	result := strings.Split(b.CleanText(text), " ")
	beefMap := b.CountWord(result)
	return c.JSON(http.StatusOK, beefMap)
}
