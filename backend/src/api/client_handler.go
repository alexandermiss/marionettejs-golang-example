package api

import (
	"encoding/json"
	_ "fmt"
	"github.com/labstack/echo/v4"
    "linguanski/src/models"
	"net/http"
)

func (s *Server) GetAllClient(ctx echo.Context) error {
	context := ctx.Request().Context()
	clients, err := s.clientManager.GetAllClient(context)
	if err != nil {
		s.logger.WithContext(context).Error(err)
		return ctx.JSON(http.StatusInternalServerError, internalErrorMessage)
	}

	return ctx.JSON(http.StatusOK, clients)
}

func (s *Server) GetClientById(ctx echo.Context) error {
	context := ctx.Request().Context()
    id := ctx.Param("id")
	client, err := s.clientManager.FindByID(context, id) 
	if err != nil {
		s.logger.WithContext(context).Error(err)
		return ctx.JSON(http.StatusInternalServerError, internalErrorMessage)
	}

	return ctx.JSON(http.StatusOK, client)
}

func (s *Server) DeleteClientById(ctx echo.Context) error {
	context := ctx.Request().Context()
    id := ctx.Param("id")
    rows, err := s.clientManager.DeleteByID(context, id) 
	if err != nil {
		s.logger.WithContext(context).Error(err)
		return ctx.JSON(http.StatusInternalServerError, internalErrorMessage)
	}

	return ctx.JSON(http.StatusOK, rows)
}

func (s *Server) CreateClient(ctx echo.Context) error {
	context := ctx.Request().Context()
	var clientRequest = new(models.ClientRequest)

	body := ctx.Request().Body
	if body == nil {
		s.logger.Warn("Body is nil")
		return ctx.JSON(http.StatusBadRequest, NewErrorMessage("Body is nil"))
	}

	decoder := json.NewDecoder(body)
	err := decoder.Decode(&clientRequest)
	if err != nil {
		s.logger.WithContext(context).Error(err)
		return ctx.JSON(http.StatusBadRequest, NewErrorMessage("invalid input: cannot decode"))
	}

    client := models.Client{
        Name: clientRequest.Name,
        ProviderID: clientRequest.ProviderID,
    }

	err = s.clientManager.Save(context, &client)
	if err != nil {
		s.logger.WithContext(context).Error(err)
		return ctx.JSON(http.StatusInternalServerError, internalErrorMessage)
	}

	return ctx.JSON(http.StatusOK, client)
}

func (s *Server) UpdateClient(ctx echo.Context) error {
	id := ctx.Param("id")
	context := ctx.Request().Context()

	var clientRequest = new(models.Client)

	body := ctx.Request().Body
	if body == nil {
		s.logger.Warn("Body is nil")
		return ctx.JSON(http.StatusBadRequest, NewErrorMessage("Body is nil"))
	}

	decoder := json.NewDecoder(body)
	err := decoder.Decode(&clientRequest)
	if err != nil {
		s.logger.WithContext(context).Error(err)
		return ctx.JSON(http.StatusBadRequest, NewErrorMessage("invalid input: cannot decode"))
	}

    client := models.Client{
        Name: clientRequest.Name,
        ProviderID: clientRequest.ProviderID,
    }

    rows, err := s.clientManager.UpdateClient(context, client, id)
	if err != nil {
		s.logger.WithContext(context).Error(err)
		return ctx.JSON(http.StatusInternalServerError, internalErrorMessage)
	}

	return ctx.JSON(http.StatusOK, rows)
}

