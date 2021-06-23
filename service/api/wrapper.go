package api

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/profile/configuration"
	"github.com/profile/dto"
	"github.com/profile/storage"
)

type ServerWrapper struct {
	srv    *Server
	logger *logrus.Entry
}

func NewServerWrapper(cfg configuration.API, db storage.Storage) *ServerWrapper {
	logger := logrus.New().WithField("layer", "http")

	wrapper := &ServerWrapper{
		srv:    NewServer(cfg, db, logger),
		logger: logger,
	}
	return wrapper
}

func (s *ServerWrapper) RegisterHandlers(e *echo.Echo) {
	e.POST("/users", s.CreateUser)
	e.GET("/users/:id", s.User)
	e.POST("/users/:id", s.UpdateUser)
	e.DELETE("/users/:id", s.DeleteUser)
	e.GET("/users", s.Users)
}

func (s *ServerWrapper) CreateUser(ctx echo.Context) error {
	return s.srv.CreateUser(ctx)
}

func (s *ServerWrapper) UpdateUser(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		s.logger.Errorf("wrong user id: %s", err.Error())
		return ctx.JSON(http.StatusBadRequest, "wrong user id")
	}
	return s.srv.UpdateUser(ctx, uint64(id))
}

func (s *ServerWrapper) User(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		s.logger.Errorf("wrong user id: %s", err.Error())
		return ctx.JSON(http.StatusBadRequest, "wrong user id")
	}
	return s.srv.User(ctx, uint64(id))
}

func (s *ServerWrapper) DeleteUser(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		s.logger.Errorf("wrong user id: %s", err.Error())
		return ctx.JSON(http.StatusBadRequest, "wrong user id")
	}
	return s.srv.DeleteUser(ctx, uint64(id))
}

func (s *ServerWrapper) Users(ctx echo.Context) error {
	params := ctx.QueryParams()
	filter := dto.Filter{}

	if val, ok := params["country"]; ok {
		if len(val) > 0 && val[0] != "" {
			filter.Country = val[0]
		}
	}

	return s.srv.Users(ctx, filter)
}

func (s *ServerWrapper) Stop(ctx context.Context) error {
	return nil
}
