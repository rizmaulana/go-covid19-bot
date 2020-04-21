package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ivandzf/go-covid19-bot/client"
	"github.com/ivandzf/go-covid19-bot/service"
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	port          = os.Getenv("PORT")
	publicURL     = os.Getenv("PUBLIC_URL")
	token         = os.Getenv("TOKEN")
	hereMapsURL   = "" //Put HERE Maps API
	hereMapsToken = os.Getenv("HEREMAPS_API_KEY")
	covidKalseURL = "" //Put url data source API
)

func main() {

	fmt.Println("Bot Started!")
	fmt.Println(publicURL)
	fmt.Println(token)

	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}

	config := tb.Settings{
		Token:  token,
		Poller: webhook,
	}

	b, err := tb.NewBot(config)
	if err != nil {
		log.Fatal(err)
	}

	// init client
	clientSvc := client.NewClient(covidKalseURL, hereMapsURL, hereMapsToken)

	// init service
	covidService := service.NewCovidService(clientSvc)

	fmt.Println("Success!")

	b.Handle("/start", func(m *tb.Message) {
		b.Send(m.Sender, covidService.GetInformation())
	})

	b.Handle("/info", func(m *tb.Message) {
		b.Send(m.Sender, covidService.GetInformation())

	})

	b.Handle("/loc", func(m *tb.Message) {
		if m.Payload != "" {
			b.Send(m.Sender, covidService.GetInformationByLocation(m.Payload))

		} else {
			b.Send(m.Sender, covidService.GetAvailableLocation())
		}
	})

	b.Handle("/all", func(m *tb.Message) {
		b.Send(m.Sender, covidService.GetOverviewStatistic())
	})

	b.Handle(tb.OnLocation, func(m *tb.Message) {
		b.Send(m.Sender, covidService.GetInformationByCoordinate(m.Location.Lat, m.Location.Lng))
	})

	b.Start()
}
