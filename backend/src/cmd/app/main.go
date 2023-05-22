package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"linguanski/src/api"
	"linguanski/src/config"
	"linguanski/src/repositories"
	"linguanski/src/service"
	"linguanski/src/service/db"
	"os"
    "io"
    "html/template"
)


type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}


func main() {

	cfg, err := config.LoadConfigFromEnv()
	if err != nil {
		println(err.Error())
		os.Exit(-1)
	}

    dbClient:= db.NewDBClientFromConfig(cfg)
    connection := dbClient.GetConnection()

	e := echo.New()

    e.Static("/assets", "dist/assets")

    renderer := &TemplateRenderer{
        templates: template.Must(template.ParseGlob("dist/*.html")),
    }

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: cfg.Logger.Out,
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	providerRepository := repositories.NewProviderRepository(connection)
	providerManager := service.NewProviderManager(providerRepository)

	clientRepository := repositories.NewClientRepository(connection)
	clientManager := service.NewClientManager(clientRepository)

	apiServer := api.NewServer(providerManager, clientManager, cfg.Logger)

    e.Renderer = renderer

	e.Router().Add("GET", "/", apiServer.IndexTemplate)

	r := e.Group("/api/v1")

	r.Add("GET", "/provider", apiServer.GetAllProvider)
	r.Add("POST", "/provider", apiServer.CreateProvider)
    r.Add("GET", "/provider/:id", apiServer.GetProviderById)
    r.Add("PUT", "/provider/:id", apiServer.UpdateProvider)
    r.Add("DELETE", "/provider/:id", apiServer.DeleteByProviderId)

	r.Add("GET", "/client", apiServer.GetAllClient)
	r.Add("POST", "/client", apiServer.CreateClient)
    r.Add("GET", "/client/:id", apiServer.GetClientById)
    r.Add("PUT", "/client/:id", apiServer.UpdateClient)
    r.Add("DELETE", "/client/:id", apiServer.DeleteClientById)

	err = e.Start(fmt.Sprintf(":%v", cfg.ServerPort))
	if err != nil {
		cfg.Logger.Error("Cannot start the server")
		os.Exit(-1)
	}

}
