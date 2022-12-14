package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CardCreateData struct {
	Name     string `json:"name"`
	Template string `json:"template"`
	Message  string `json:"message"`
}

func Cards(r *gin.Engine) {

	cards := r.Group("/cards")
	{
		cards.POST("/create", func(c *gin.Context) {

			body := CardCreateData{}
			if err := c.BindJSON(&body); err != nil {
				fmt.Print(err)
			}

			if body.Name == "" || body.Template == "" || body.Message == "" {
				c.JSON(406, map[string]string{"status": "Name and template needed"})
			}

			card := NewCard(body.Name, body.Message, body.Template)

			card.GenerateAndSaveCard()
			card.StoreCard()

			c.JSON(http.StatusOK, map[string]string{"link": card.GetCardLink()})
		})

		cards.GET("/message", func(c *gin.Context) {
			id := c.Query("id")
			if id != "" {
				c.JSON(http.StatusOK, map[string]string{"message": GetCardById(id).Message})
			} else {
				c.JSON(404, map[string]string{"status": "id needed"})
			}
		})

		cards.GET("/createDate", func(c *gin.Context) {
			id := c.Query("id")
			if id != "" {
				c.JSON(http.StatusOK, map[string]string{"date": GetCardById(id).CreateDate})
			} else {
				c.JSON(404, map[string]string{"status": "id needed"})
			}
		})
	}
}
