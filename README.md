# Discord-Hooks-Go
Wrapper to easily create and send Discord webhooks and embeds

```go
embed := hooks.NewEmbed()
embed.Title = "My Embed Title"
embed.SetAuthor("An Author", "https://google.com", "https://i.imgur.com/vSxLPRg.jpg")
embed.SetFooter("A Footer", "https://i.imgur.com/vSxLPRg.jpg")

isSent := embed.Send(
  "WEBHOOK_URL",
  "WEBHOOK_CONTENT",
  "WEBHOOK_NAME",
  "WEBHOOK_ICON_URL")
if isSent {
  fmt.Println("Sent Embed!")
}
```