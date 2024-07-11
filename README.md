# Service Status Checker

Service Status Checker is a Go library that allows you to monitor the status of a service and send messages to a Discord channel when the status changes.

## Usage

Here's an example of how to use the library to monitor a service and send messages to a Discord channel:

```go
package main

import (
	"log"

	"github.com/tapmah/service-status-checker/checker"
	"github.com/tapmah/service-status-checker/configs"
	"github.com/tapmah/service-status-checker/sendservices"
	"github.com/spf13/viper"
)

func main() {
	if err := configs.InitConfig(); err != nil {
		log.Fatalf("Failed to init config: %v", err.Error())
	}

	channelID := viper.GetString("discord.channelID")
	token := viper.GetString("discord.token")
	serviceName := viper.GetString("checker.serviceName")

	bot := sendservices.NewDiscordBot(channelID, token)
	services := sendservices.NewService(bot)
	checker := checker.NewChecker(serviceName, services)
	checker.Start()
}

Before running the example, you need to configure the config.yml file. Heres an example of what the configuration file might look like:

discord:
  channelID: "1234567890"
  token: "your-bot-token"
checker:
  serviceName: "your-service-name"

In this configuration file, channelID is the ID of the channel where messages will be sent, token is the token of the bot that will send the messages, and serviceName is the name of the service that will be monitored.

To run the example, execute the following commands:

go get -u github.com/tapmah/service-status-checker
go run main.go