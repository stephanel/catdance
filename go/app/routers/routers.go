package routers

import (
	"../controllers"
	"../urls"

	"github.com/gorilla/pat"
)

// registers all routes for the application.
func GetRouter() *pat.Router {
	// url paths imported from urls package
	urlPatterns := urls.ReturnURLS()
	app := pat.New()
	app.Get(urlPatterns.HOME_PATH, controllers.HomeController)

	common := pat.New()
	common.Get(urlPatterns.HOME_PATH, controllers.HomeController)

	return common
}
