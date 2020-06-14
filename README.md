# Discord-Hooks-Go
Wrapper to easily create and send Discord webhooks and embeds

```go
myEmbed := hooks.NewEmbed()
myEmbed.Title = "My Embed Title"
myEmbed.SetAuthor("An Author", "https://google.com", "https://i.imgur.com/vSxLPRg.jpg")
myEmbed.SetFooter("A Footer", "https://i.imgur.com/vSxLPRg.jpg")
isSent := myEmbed.Send(
  "https://discordapp.com/api/webhooks/693106893681590306/htOr6ptNmKx1ishvcoXZP0pv4PS78ol2YHZHh0sSH37x9-BpvNBXcYBVxY5ZvM8ibbfI",
  "@everyone",
  "Webhook",
  "")
if isSent {
  fmt.Println("Sent Embed!")
}
```