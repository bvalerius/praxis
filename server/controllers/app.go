package controllers

import (
	"net/http"

	"github.com/convox/praxis/api"
)

func AppCreate(w http.ResponseWriter, r *http.Request, c *api.Context) error {
	name := c.Form("name")

	c.LogParams("name")

	app, err := Provider.AppCreate(name)
	if err != nil {
		return err
	}

	return c.RenderJSON(app)
}

func AppDelete(w http.ResponseWriter, r *http.Request, c *api.Context) error {
	name := c.Var("name")

	if err := Provider.AppDelete(name); err != nil {
		return err
	}

	return nil
}

func AppGet(w http.ResponseWriter, r *http.Request, c *api.Context) error {
	name := c.Var("name")

	app, err := Provider.AppGet(name)
	if err != nil {
		return err
	}

	return c.RenderJSON(app)
}

func AppList(w http.ResponseWriter, r *http.Request, c *api.Context) error {
	apps, err := Provider.AppList()
	if err != nil {
		return err
	}

	return c.RenderJSON(apps)
}