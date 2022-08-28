package routes

import (
	"fmt"
	"image/color"

	"github.com/Brownymister/imgtext"
	"github.com/google/uuid"
)

type CardData struct {
	Id       string `json:"id"`       // card id
	Name     string `json:"name"`     // name of the recever of the card
	Message  string `json:"message"`  // message on the backsite
	Template string `json:"template"` // card template
	CardLink string `json:"cardlink"` // image link
}

type Card interface {
	GenerateAndSaveCard()
	GetCardLink() string
	StoreCard()
}

var CardsArray []CardData

func NewCard(name string, message string, tpl string) Card {

	var card Card

	card = &CardData{Name: name, Template: tpl, Id: uuid.New().String(), Message: message}

	return card
}

func GetMessageById(id string) CardData {
	for _, ele := range CardsArray {
		if ele.Id == id {
			return ele
		}
	}
	return CardData{}
}

func (card *CardData) GenerateAndSaveCard() {
	fmt.Print("create 1")

	img := imgtext.NewImage("./templates/" + card.Template + ".png")

	err := img.AddTextToImage(card.Name, 250, 250, 60, "./fonts/DancingScript-VariableFont_wght.ttf", color.Black)
	if err != nil {
		fmt.Print(err)
	}

	err = img.Save(fmt.Sprintf("./createdCards/%s.png", card.Id))
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print("create 2")
}

func (card *CardData) StoreCard() {
	card.CardLink = fmt.Sprintf("/cardsimg/%s.png", card.Id)

	CardsArray = append(CardsArray, CardData{
		Id:       card.Id,
		Name:     card.Name,
		Message:  card.Message,
		Template: card.Template,
		CardLink: card.CardLink,
	})
}

func (card *CardData) GetCardLink() string {
	return card.CardLink
}
