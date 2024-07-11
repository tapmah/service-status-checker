package checker

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	sendservices "github.com/tapmah/service-status-checker/pkg/send_services"
)

type Checker struct {
	ServiceName string
	Services    *sendservices.Service
}

func NewChecker(serviceName string, services *sendservices.Service) *Checker {
	return &Checker{
		ServiceName: serviceName,
		Services:    services,
	}
}

func (c *Checker) Start() error {
	previousState := getServiceState(c.ServiceName)

	for {
		currentState := getServiceState(c.ServiceName)

		if previousState != currentState {
			timestamp := time.Now().Format("2006-01-02 15:04:05")
			message := fmt.Sprintf("[%s] wol.service state changed: %s -> %s\n", timestamp, previousState, currentState)
			c.Services.DiscordBot.SendMessage(message)
		}

		previousState = currentState
		time.Sleep(1 * time.Second)
	}
}
func getServiceState(serviceName string) string {
	cmd := exec.Command("systemctl", "show", "--property=ActiveState", serviceName)
	output, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(string(output))
}
