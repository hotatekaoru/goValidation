package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"goValidation/validate"
)

const defaultPort = "8080"

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/", showIndexScreen)
	router.POST("/", onClickSubmitButton)
	http.ListenAndServe(":"+port(), router)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}

	return port
}

func showIndexScreen(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func onClickSubmitButton(c *gin.Context) {
	form, errs := validate.ValidateForm(c)

	if errs != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"errorList": errs,
			"form":      form,
		})
		return
	}

	c.HTML(http.StatusOK, "result.html", gin.H{
		"form": form,
	})
}