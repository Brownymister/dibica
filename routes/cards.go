package routes

import (
	"dibica/db"
	"fmt"
	"image/color"
	"time"

	"github.com/Brownymister/imgtext"
	"github.com/google/uuid"
)

type CardData struct {
	Id         string `json:"id"`         // card id
	Name       string `json:"name"`       // name of the recever of the card
	Message    string `json:"message"`    // message on the backsite
	Template   string `json:"template"`   // card template
	CardLink   string `json:"cardlink"`   // image link
	CreateDate string `json:"createdate"` // create date
}

type Card interface {
	GenerateAndSaveCard()
	GetCardLink() string
	StoreCard()
}

func NewCard(name string, message string, tpl string) Card {

	var card Card

	card = &CardData{Name: name, Template: tpl, Id: uuid.New().String(), Message: message, CreateDate: GetDate()}

	return card
}

func GetCardById(id string) db.CardData {
	card := db.GetCardById(id)
	return card
}

func GetDate() string {
	dt := time.Now()
	return dt.Format("01-02-2006")
}

func (card *CardData) GenerateAndSaveCard() {

	img := imgtext.NewImage("./templates/" + card.Template + ".png")

	err := img.AddTextToImage(card.Name, 250, 250, 60, "./fonts/DancingScript-VariableFont_wght.ttf", color.Black)
	if err != nil {
		fmt.Print(err)
	}

	err = img.Save(fmt.Sprintf("./createdCards/%s.png", card.Id))
	if err != nil {
		fmt.Print(err)
	}

}

func (card *CardData) StoreCard() {
	card.CardLink = fmt.Sprintf("/cardsimg/%s.png", card.Id)
	fmt.Print(card)

	db.InsertIntoCards(db.CardData{
		Id:         card.Id,
		Name:       card.Name,
		Message:    card.Message,
		Template:   card.Template,
		CardLink:   card.CardLink,
		CreateDate: card.CreateDate,
	})

}

func (card *CardData) GetCardLink() string {
	return card.CardLink
}
