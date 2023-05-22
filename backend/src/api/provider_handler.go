package api

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
    "linguanski/src/models"
	"net/http"
)

func (s *Server) CreateProvider(ctx echo.Context) error {
	context := ctx.Request().Context()
	var userRequest = new(models.ProviderRequest)

	body := ctx.Request().Body
	if body == nil {
		s.logger.Warn("Body is nil")
		return ctx.JSON(http.StatusBadRequest, NewErrorMessage("Body is nil"))
	}

	decoder := json.NewDecoder(body)
	err := decoder.Decode(&userRequest)
	if err != nil {
		s.logger.WithContext(context).Error(err)
		return ctx.JSON(http.StatusBadRequest, NewErrorMessage("invalid input: cannot decode"))
	}

	user := models.Provider{
		ID:           0,
		Name:         userRequest.Name,
	}
	err = s.providerManager.Save(context, &user)
	if err != nil {
		s.logger.WithContext(context).Error(err)
		return ctx.JSON(http.StatusInternalServerError, internalErrorMessage)
	}

	providerManager := models.ProviderResponse(user)
	return ctx.JSON(http.StatusOK, providerManager)
}

func (s *Server) GetAllProvider(ctx echo.Context) error {
	context := ctx.Request().Context()
	providers, err := s.providerManager.GetAll(context)
	if err != nil {
		s.logger.WithContext(context).Error(err)
		return ctx.JSON(http.StatusInternalServerError, internalErrorMessage)
	}

	return ctx.JSON(http.StatusOK, providers)
}

func (s *Server) GetProviderById(ctx echo.Context) error {
	context := ctx.Request().Context()
    id := ctx.Param("id")
	user, err := s.providerManager.FindByID(context, id) 
	if err != nil {
		s.logger.WithContext(context).Error(err)
		return ctx.JSON(http.StatusInternalServerError, internalErrorMessage)
	}

    userResponse := models.ProviderResponse(user)
	return ctx.JSON(http.StatusOK, userResponse)
}

func (s *Server) DeleteByProviderId(ctx echo.Context) error {
	context := ctx.Request().Context()
    id := ctx.Param("id")
    rows, err := s.providerManager.DeleteByID(context, id) 
	if err != nil {
		s.logger.WithContext(context).Error(err)
		return ctx.JSON(http.StatusInternalServerError, internalErrorMessage)
	}

	return ctx.JSON(http.StatusOK, rows)
}

func (s *Server) UpdateProvider(ctx echo.Context) error {
    id := ctx.Param("id")
	context := ctx.Request().Context()

	var userRequest = new(models.ProviderRequest)

	body := ctx.Request().Body
	if body == nil {
		s.logger.Warn("Body is nil")
		return ctx.JSON(http.StatusBadRequest, NewErrorMessage("Body is nil"))
	}

	decoder := json.NewDecoder(body)
	err := decoder.Decode(&userRequest)
	if err != nil {
		s.logger.WithContext(context).Error(err)
		return ctx.JSON(http.StatusBadRequest, NewErrorMessage("invalid input: cannot decode"))
	}

	user := models.Provider{
		Name: userRequest.Name,
	}

    rows, err := s.providerManager.Update(context, user, id) 
	if err != nil {
		s.logger.WithContext(context).Error(err)
		return ctx.JSON(http.StatusInternalServerError, internalErrorMessage)
	}

	return ctx.JSON(http.StatusOK, rows)
}

func (s *Server) IndexTemplate(ctx echo.Context) error {
	context := ctx.Request().Context()

    s.logger.WithContext(context).Info("Show html template")
    return ctx.Render(http.StatusOK, "index.html", map[string]interface{}{
        "name": "Hello World!",
    })
}

