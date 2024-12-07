package main

import (
	"net/http"

	"github.com/baskaev/db/datab"
	"github.com/labstack/echo/v4"
)

func main() {
	// Инициализация базы данных
	if err := datab.InitDB(); err != nil {
		panic(err)
	}

	// Создаем экземпляр Echo
	e := echo.New()

	// Определяем маршруты
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	e.GET("/api/movies", func(c echo.Context) error {
		movies, err := datab.FetchMovies()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, movies)
	})

	// Запускаем сервер
	e.Logger.Fatal(e.Start(":8081"))
}

// package main

// import (
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// func main() {
// 	// Создаем экземпляр Echo
// 	e := echo.New()

// 	// Определяем маршрут и обработчик
// 	e.GET("/", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "Hello, Echo!")
// 	})
// 	// Определяем маршрут и обработчик
// 	e.POST("/", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "Hello, Echo!")
// 	})

// 	// Обработчик для маршрута /api
// 	e.GET("/api", func(c echo.Context) error {
// 		return c.JSON(http.StatusOK, map[string]string{
// 			"message": "API is working!",
// 		})
// 	})

// 	// Обработчик для маршрута /api
// 	e.POST("/api", func(c echo.Context) error {
// 		return c.JSON(http.StatusOK, map[string]string{
// 			"message": "API is working!",
// 		})
// 	})

// 	// Запускаем сервер на порту 8080
// 	e.Logger.Fatal(e.Start(":8081"))
// }
