package hooks

// Package hooks allows easy creation of Discord embeds
// and sending of webhooks

import (
	"time"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// WebhookData is the data sent to a Discord Webhook
type WebhookData struct {
	Content   string  `json:"content,omitempty"`
	Username  string  `json:"username,omitempty"`
	AvatarURL string  `json:"avatar_url,omitempty"`
	Embeds    []Embed `json:"embeds"`
}

// Embed is a Discord embed
type Embed struct {
	Title       string  `json:"title,omitempty"`
	TitleURL    string  `json:"url,omitempty"`
	Description string  `json:"description,omitempty"`
	Colour      string  `json:"color,omitempty"`
	Author      Author  `json:"author"`
	Footer      Footer  `json:"footer"`
	Fields      []Field `json:"fields,omitempty"`
	Timestamp   string    `json:"timestamp,omitempty"`
}

// Field is a Discord embed field
type Field struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline,omitempty"`
}

// Author is Discord embed author data
type Author struct {
	Text    string `json:"name,omitempty"`
	URL     string `json:"url,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

// Footer is Discord embed footer data
type Footer struct {
	Text    string `json:"text,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

// NewEmbed creates and returns an Embed
func NewEmbed() Embed {
	embed := Embed{Footer: Footer{}}
	return embed
}

// AddField creates a Field and adds it to an Embed
func (e *Embed) AddField(name, value string, inline bool) {
	field := Field{
		Name:   name,
		Value:  value,
		Inline: inline,
	}
	e.Fields = append(e.Fields, field)
}

// SetAuthor sets the author of an Embed
func (e *Embed) SetAuthor(text, url, iconURL string) {
	e.Author = Author{Text: text, URL: url, IconURL: iconURL}
}

// SetFooter sets the footer of an Embed
func (e *Embed) SetFooter(text, iconURL string) {
	e.Footer = Footer{Text: text, IconURL: iconURL}
}

// SetTimestamp sets the timestamp of an Embed
func (e *Embed) SetTimestamp() {
	e.Timestamp = time.Now().Format(time.RFC3339)
}

func (e Embed) getWebhookData(content, username, avatarURL string) []byte {
	webhookData := WebhookData{Content: content, Username: username, AvatarURL: avatarURL, Embeds: []Embed{e}}
	data, err := json.Marshal(webhookData)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

// Send sends the webhook data
func (e Embed) Send(webhook, content, username, avatarURL string) bool {
	d := e.getWebhookData(content, username, avatarURL)
	resp, err := http.Post(webhook, "application/json", bytes.NewBuffer(d))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 204 {
		return true
	}
	return false
}
