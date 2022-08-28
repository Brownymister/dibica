package routes

import (
	"fmt"
	"net/http"
	"reflect"

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

			if reflect.DeepEqual(body, CardCreateData{}) {
				c.JSON(500, map[string]string{"status": "Name and template needed"})
			}

			card := NewCard(body.Name, body.Message, body.Template)

			card.GenerateAndSaveCard()
			card.StoreCard()

			c.JSON(http.StatusOK, map[string]string{"link": card.GetCardLink()})
		})
	}
}
