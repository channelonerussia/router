// Package router фасад над https://github.com/gin-gonic/gin.
// Сразу предустановлен пакет graceful, gin-contrib/cors,
package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/graceful"
	"github.com/gin-gonic/gin"
)

// Router дефолтный Gin роутер с Graceful middleware.
type Router = *graceful.Graceful

// CorsConfig cors.Config
type CorsConfig = cors.Config

// Context оболочка над gin.Context.
type Context = gin.Context

// H оболочка над gin.H.
type H = gin.H

// HandlerFunc оболочка над gin.HandlerFunc.
type HandlerFunc = gin.HandlerFunc

// New создает инстанс роутера.
func New() (Router, error) {
	// Создаем дефолтный роутер
	router, GFErr := graceful.Default()

	// Если не получилось
	if GFErr != nil {
		return nil, GFErr
	}

	return router, nil
}

// Cors возвращает мидлвару с настройками корсов для роутера
func Cors(config CorsConfig) HandlerFunc {
	return cors.New(config)
}
