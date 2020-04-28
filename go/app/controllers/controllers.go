package controllers

import (
	"math/rand"
	"net/http"
	"time"

	"../templates"
	"../utils"
)

func HomeController(res http.ResponseWriter, req *http.Request) {

	images := []string{
		"https://media1.giphy.com/media/yFQ0ywscgobJK/giphy.webp",
		"https://media0.giphy.com/media/WYEWpk4lRPDq0/200w.webp",
		"https://media0.giphy.com/media/zWuSfeDJkqj0A/200w.webp",
		"https://media2.giphy.com/media/mJqQXx8vK9kD6/200w.webp",
		"https://media2.giphy.com/media/12mWfQYoxRqslq/200.webp",
		"https://media3.giphy.com/media/Vj5ZgHbXa3kWY/200w.webp",
		"https://media0.giphy.com/media/40Fpxgn6Yq640/giphy.webp",
		"https://media3.giphy.com/media/jpPTyP6YghtiU/200.webp",
		"https://media2.giphy.com/media/3oriO0OEd9QIDdllqo/200.webp",
		"https://media0.giphy.com/media/OmK8lulOMQ9XO/200w.webp",
	}

	data := make(map[string]interface{})

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s) // initialize local pseudorandom generator
	data["url"] = images[r.Intn(len(images))]
	controllerTemplate := templates.HOME
	if req.Method == "GET" {
		utils.CustomTemplateExecute(res, req, controllerTemplate, data)
	}
}
