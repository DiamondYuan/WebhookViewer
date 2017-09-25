package main

import (
	"github.com/GeorgeYuen/webhookViewer/webhook"
	"github.com/go-macaron/macaron"
)

func main() {
	m := macaron.Classic()
	m.Combo("/*").
	Get(webhook.GetWebhook).
	Post(webhook.SetWebhook).
	Put(webhook.SetWebhook)
	m.Run()
}
