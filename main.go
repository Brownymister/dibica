package main

import (
	"dibica/routes"
	"os"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("./ui/dist/ui/index.html")
		} else {
			c.File("./ui/dist/ui/" + path.Join(dir, file))
		}
	})

	routes.Cards(r)

	r.Static("/templates", "./templates/")
	r.Static("/cardsimg", "./createdCards/")
	r.Static("/img", "./img/")

	err := r.Run(":" + os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
}
