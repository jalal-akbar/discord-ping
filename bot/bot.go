package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/jalal-akbar/discord-ping/config"
)

var (
	BotID string
	goBot *discordgo.Session
)

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
	}
	user, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}
	BotID = user.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("bot is running")

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}
	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}
}
