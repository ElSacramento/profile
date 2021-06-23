package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pariz/gountries"
	"github.com/sirupsen/logrus"

	"github.com/profile/configuration"
	"github.com/profile/dto"
	"github.com/profile/storage"
)

type Server struct {
	cfg     configuration.API
	db      storage.Storage
	logger  *logrus.Entry
	country *gountries.Query
}

func NewServer(cfg configuration.API, db storage.Storage, logger *logrus.Entry) *Server {
	countryConverter := gountries.New()

	return &Server{
		cfg:     cfg,
		db:      db,
		logger:  logger,
		country: countryConverter,
	}
}

func (s *Server) CreateUser(ctx echo.Context) error {
	request := dto.CreateUser{}
	if err := ctx.Bind(&request); err != nil {
		s.logger.Error(err)
		return err
	}

	if request.Email == "" || request.Password == "" {
		s.logger.Error("empty email or password")
		return ctx.JSON(http.StatusUnprocessableEntity, dto.ErrorResponse{Message: "empty email or password"})
	}

	if request.Country != "" {
		country, err := s.country.FindCountryByName(request.Country)
		if err != nil {
			errMsg := fmt.Sprintf("wrong country name: %v", err.Error())
			s.logger.Error(errMsg)
			return ctx.JSON(http.StatusUnprocessableEntity, dto.ErrorResponse{Message: "wrong country name"})
		}

		request.Country = country.Name.Common
	}

	dbUser := request.ToDatabase()
	created, err := s.db.Create(dbUser)
	if err != nil {
		s.logger.Errorf("failed to create user: %s", err.Error())
		return ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "failed to create user"})
	}

	return ctx.JSON(http.StatusOK, dto.UserFromDatabase(created))
}

func (s *Server) UpdateUser(ctx echo.Context, id uint64) error {
	request := dto.UpdateUser{}
	if err := ctx.Bind(&request); err != nil {
		s.logger.Error(err)
		return err
	}

	if request.Email == "" || request.Password == "" {
		s.logger.Error("empty email or password")
		return ctx.JSON(http.StatusUnprocessableEntity, dto.ErrorResponse{Message: "empty email or password"})
	}

	if request.Country != "" {
		country, err := s.country.FindCountryByName(request.Country)
		if err != nil {
			errMsg := fmt.Sprintf("wrong country name: %v", err.Error())
			s.logger.Error(errMsg)
			return ctx.JSON(http.StatusUnprocessableEntity, dto.ErrorResponse{Message: "wrong country name"})
		}

		request.Country = country.Name.Common
	}

	dbUser := request.ToDatabase(id)
	updated, err := s.db.Update(dbUser)
	if err == storage.ErrNotFound {
		s.logger.Errorf("failed to update user: %v", storage.ErrNotFound)
		return ctx.JSON(http.StatusNotFound, struct{}{})
	}
	if err != nil {
		errMsg := fmt.Sprintf("failed to update user: %v", err.Error())
		s.logger.Error(errMsg)
		return ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "failed to update user"})
	}

	return ctx.JSON(http.StatusOK, dto.UserFromDatabase(updated))
}

func (s *Server) DeleteUser(ctx echo.Context, id uint64) error {
	ok, err := s.db.Delete(id)
	if err != nil {
		s.logger.Errorf("failed to delete user: %v", err)
		return ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "failed to delete user"})
	}
	if !ok {
		errMsg := "could not delete user"
		s.logger.Error(errMsg)
		return ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: errMsg})
	}
	return ctx.JSON(http.StatusNoContent, struct{}{})
}

func (s *Server) User(ctx echo.Context, id uint64) error {
	dbUser, err := s.db.Get(id)
	if err == storage.ErrNotFound {
		s.logger.Errorf("failed to get user: %v", storage.ErrNotFound)
		return ctx.JSON(http.StatusNotFound, struct{}{})
	}
	if err != nil {
		s.logger.Errorf("failed to get user: %v", err)
		return ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "failed to get user"})
	}

	return ctx.JSON(http.StatusOK, dto.UserFromDatabase(dbUser))
}

func (s *Server) Users(ctx echo.Context, filter dto.Filter) error {
	dbFilter := filter.ToDatabase()
	users, err := s.db.List(dbFilter)
	if err != nil {
		s.logger.Errorf("failed to list users: %v", err)
		return ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "failed to list users"})
	}

	response := make([]dto.User, 0, len(users))
	for _, obj := range users {
		response = append(response, dto.UserFromDatabase(obj))
	}

	return ctx.JSON(http.StatusOK, response)
}
