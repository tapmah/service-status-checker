package sendservices

type Service struct {
	DiscordBot
}

type DiscordBot interface {
	SendMessage(message string) error
}

func NewService(bot *Bot) *Service {
	return &Service{
		DiscordBot: NewDiscordBot(
			bot.ChannelID,
			bot.Token,
		),
	}
}
