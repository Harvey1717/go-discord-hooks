# Discord-Hooks-Go
Wrapper to easily create and send Discord webhooks and embeds

```go
embed := hooks.NewEmbed()
embed.Title = "Click Me!"
embed.TitleURL = "https://twitter.com"
embed.Description = "Hey! This is a decription"
embed.SetColour("42f5b9")
embed.SetTimestamp()
embed.AddField("Field Title One", "Field Value One", false)
embed.AddField("Field Title Two", "Field Value Two", false)
embed.SetAuthor("An Author", "https://google.com", "https://i.imgur.com/vSxLPRg.jpg")
embed.SetFooter("A Footer", "https://i.imgur.com/vSxLPRg.jpg")

isSent := embed.Send(
  "https://discordapp.com/api/webhooks/XXX/XXX",
  "This is some content",
  "DMC",
  "https://i.imgur.com/vSxLPRg.jpg")

if isSent {
  fmt.Println("Sent Embed!")
}
```