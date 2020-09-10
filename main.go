package main

import (
	"log"

	"github.com/adibaulia/lacak-pesanan-bot/api"
	"github.com/adibaulia/lacak-pesanan-bot/config"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {

	webhook := &tb.Webhook{
		Listen:   ":" + config.PORT,
		Endpoint: &tb.WebhookEndpoint{PublicURL: config.PublicURL},
	}

	pref := tb.Settings{
		Token:  config.TELETOKEN,
		Poller: webhook,
	}

	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	api.Route(b)
	b.Start()
}
