package sendservices

import (
	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	ChannelID string
	Token     string
}

func NewDiscordBot(channelID string, token string) *Bot {
	return &Bot{
		ChannelID: channelID,
		Token:     token,
	}
}

func (bot *Bot) SendMessage(message string) error {
	dg, err := discordgo.New("Bot " + bot.Token)
	if err != nil {
		return err
	}

	_, err = dg.ChannelMessageSend(bot.ChannelID, message)
	if err != nil {
		return err
	}
	return nil
}
