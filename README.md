# Service Status Checker

Service Status Checker is a Go library that allows you to monitor the status of a service and send messages to a Discord channel when the status changes.

## Usage

Here's an example of how to use the library to monitor a service and send messages to a Discord channel:

```go
package main

import (
	"github.com/tapmah/service-status-checker/pkg/checker"
	sendservices "github.com/tapmah/service-status-checker/pkg/send_services"
)

const (
	channelID   = "<YOUR CHANNEL ID>"
	token       = "<YOUR DISCORD BOT TOCKEN>"
	serviceName = "<SERVICE NAME>"
)

func main() {

	bot := sendservices.NewDiscordBot(channelID, token)
	services := sendservices.NewService(bot)
	checker := checker.NewChecker(serviceName, services)
	checker.Start()
}
```

To run the example, execute the following commands:

`go get -u github.com/tapmah/service-status-checker`
`go run main.go`
