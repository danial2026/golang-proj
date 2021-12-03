package app

import (
	controllers "github.com/danial2026/golang-proj/controllers"
)

func mapUrls() {
	prefix := "/api/v1"

	router.GET(prefix+"/contact-me", controllers.UsersController.StoreNewUser)
}
