// A simple example that uses the modules from the gsbot package and go-steam to log on
// to the Steam network.
//
// Use the right options for your account settings:
// Normal login: username + password
//
// Email code:   username + password
//               username + password + authcode
//
// Mobile code:  username + password + twofactorcode
//               username + loginkey
//
//     gsbot [username] [-p password] [-a authcode] [-t twofactorcode] [-l loginkey]

package main

import (
	"fmt"
	"os"

	"github.com/Philipp15b/go-steam/v2"
	"github.com/Philipp15b/go-steam/v2/gsbot"
	"github.com/Philipp15b/go-steam/v2/protocol/steamlang"
)

const usage string = "usage: gsbot [username] [-p password] [-a authcode] [-t twofactorcode] [-l loginkey]"

func main() {
	if len(os.Args) < 3 || len(os.Args) % 2 != 0 {
		fmt.Println(usage)
		return
	}

	details := &steam.LogOnDetails{
		Username: os.Args[1],
		ShouldRememberPassword: true,
	}

	for i := 2; i < len(os.Args) - 1; i += 2 {
		switch os.Args[i] {
		case "-p":
			details.Password = os.Args[i+1]
		case "-a":
			details.AuthCode = os.Args[i+1]
		case "-t":
			details.TwoFactorCode = os.Args[i+1]
		case "-l":
			details.LoginKey = os.Args[i+1]
		default:
			fmt.Println(usage)
			return
		}
	}

	bot := gsbot.Default()
	client := bot.Client
	auth := gsbot.NewAuth(bot, details, "sentry.bin")
	debug, err := gsbot.NewDebug(bot, "debug")
	if err != nil {
		panic(err)
	}
	client.RegisterPacketHandler(debug)
	serverList := gsbot.NewServerList(bot, "serverlist.json")
	serverList.Connect()

	for event := range client.Events() {
		auth.HandleEvent(event)
		debug.HandleEvent(event)
		serverList.HandleEvent(event)

		switch e := event.(type) {
		case error:
			fmt.Printf("Error: %v", e)
		case *steam.LoggedOnEvent:
			client.Social.SetPersonaState(steamlang.EPersonaState_Online)
		}
	}
}
