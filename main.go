package main

import (
	"fmt"
	"log"
	"os"

	tb "gopkg.in/tucnak/telebot.v2"
)

const (
	port      = os.Getenv("PORT")
	publicURL = os.Getenv("PUBLIC_URL")
	token     = os.Getenv("TOKEN")
)

func main() {

	fmt.Printf("Bot Started!")
	fmt.Printf(publicURL)
	fmt.Printf(token)

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
	} else {
		fmt.Printf("Success!")
	}

	b.Handle("/start", func(m *tb.Message) {
		b.Send(m.Sender, Information())
	})

	b.Handle("/info", func(m *tb.Message) {
		b.Send(m.Sender, Information())

	})

	b.Handle("/loc", func(m *tb.Message) {
		if m.Payload != "" {
			b.Send(m.Sender, GetInformationByLoc(m.Payload))

		} else {
			b.Send(m.Sender, AvailableLocation())
		}
	})

	b.Handle("/all", func(m *tb.Message) {
		b.Send(m.Sender, OverviewStatistic())
	})

	b.Handle(tb.OnLocation, func(m *tb.Message) {
		b.Send(m.Sender, GetInformationByCoordinate(m.Location.Lat, m.Location.Lng))
	})

	b.Start()
}
