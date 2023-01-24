package echo

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"testing"
)

func TestController(t *testing.T) {

	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		fmt.Println("拦截器1")
		return next
	})
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		fmt.Println("拦截器2")
		return next
	})
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		fmt.Println("拦截器3")
		return next
	})
	g := e.Group("/user", func(next echo.HandlerFunc) echo.HandlerFunc {
		fmt.Println("控制器 拦截器 1")
		return next
	})
	g.GET("/info", func(c echo.Context) error {
		return c.String(http.StatusOK, "this is get request!")
	}, func(next echo.HandlerFunc) echo.HandlerFunc {
		fmt.Println("get 拦截器1")
		return next
	}, func(next echo.HandlerFunc) echo.HandlerFunc {
		fmt.Println("get 拦截器2")
		return next
	}, func(next echo.HandlerFunc) echo.HandlerFunc {
		fmt.Println("get 拦截器3")
		return next
	})
	err := e.Start(":8086")
	if err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
