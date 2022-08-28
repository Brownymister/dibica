package main

import (
	"image/color"
	"path"
	"path/filepath"

	"github.com/Brownymister/imgtext"
	"github.com/gin-gonic/gin"
)

func main() {

	// init a new Image instance
	img := imgtext.NewImage("./templates/bd1.png")
	// add Text to Image
	img.AddTextToImage("Ute", 250, 250, 60, "./fonts/DancingScript-VariableFont_wght.ttf", color.Black)
	// save Image on drive
	img.Save("out.png")

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

	r.Static("/templates", "./templates/")

	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}
