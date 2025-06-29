package ec2test

import (
	"context"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type Server struct {
	pool           *pgxpool.Pool
	echo           *echo.Echo
	httpServerPort string
	apiKey         string
}

func NewServer(httpServerPort string, pool *pgxpool.Pool, apiKey string) *Server {
	return &Server{
		pool:           pool,
		echo:           echo.New(),
		httpServerPort: httpServerPort,
		apiKey:         apiKey,
	}
}
func (s *Server) SomeMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		if auth == "" || auth != s.apiKey {
			return echo.ErrUnauthorized
		}
		return next(c)
	}

}
func (s *Server) Run() error {
	s.echo.Use(s.SomeMiddleware)

	s.echo.GET("/ping", s.Pong)
	return s.echo.Start(s.httpServerPort)
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return s.echo.Shutdown(ctx)
}

type Response struct {
	Pong string `json:"pong"`
}

func (s *Server) Pong(ctx echo.Context) error {

	var resp Response
	resp.Pong = "pong"
	return ctx.JSON(http.StatusOK, resp)
}
