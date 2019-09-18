package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

const (
	botStatus string = "Bathtime :)"
	commandPrefix = "!vesi"
	memeCommands = true
	botTokenFileName = "token.txt"
)

var (
	vesiID string
	botToken string
)

func main() {
	botToken = readKey(botTokenFileName)

	vesi, err := discordgo.New("Bot " + botToken)
	if err != nil {
		fmt.Printf("Error creating bot due to %v", err)
	}

	vesiUser, err := vesi.User("@me")
	if err != nil {
		fmt.Printf("Error retrieving bot account due to %v", err)
	}

	vesiID = vesiUser.ID

	vesi.AddHandler(commandHandler)
	vesi.AddHandler(func(vesi *discordgo.Session, ready *discordgo.Ready) {
		err = vesi.UpdateStatus(0, botStatus)
		if err != nil {
			fmt.Printf("Error updating bot status %v", err)
		}

		servers := vesi.State.Guilds
		fmt.Printf("VesiBot has started on %d servers", len(servers))
	})

	err = vesi.Open()
	if err != nil {
		fmt.Printf("Error opening connection to discord due to %v", err)
	}


	defer vesi.Close()

	<-make(chan struct{})
}

func commandHandler(vesi *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	channelID := message.ChannelID
	var commandString string
	var paramString string

	if user.ID == vesiID || user.Bot {
		return
	}

	content := message.Content

	if len(content) > 4 {
		messagePrefix := content[0:5]
		if strings.Compare(commandPrefix, messagePrefix) == 0 {

			if len(content) > 9 {
				commandString = content[6:10]
			}
			if len(content) > 10 {
				paramString = content[11:]
			}

			fmt.Printf("\nCommand invoked on %v by %v - %v\n", message.GuildID, message.Author, content)

			switch commandString {
			case "help":
				helpText := HelpCommand()
				sendMessage(vesi, channelID, helpText)

				break

			case "sorc":
				helpText := commandSource()
				sendMessage(vesi, channelID, helpText)

				break

			case "find":
				result := findCommand(paramString)
				sendMessage(vesi, channelID, result)

				break

			case "item":
				result := itemCommand(paramString)
				sendMessage(vesi, channelID, result)

				break

			case "zone":
				result := zoneCommand(paramString)
				sendMessage(vesi, channelID, result)

				break

			case "dung":
				result := dungeonCommand(paramString)
				sendMessage(vesi, channelID, result)

				break

			case "bisl":
				result := bislCommand(paramString)
				sendMessage(vesi, channelID, result)

				break

			case "qust":
				result := questCommand(paramString)
				sendMessage(vesi, channelID, result)

				break

			case "strs":
				if memeCommands {
					result := streetsCommand()
					sendMessage(vesi, channelID, result)
					break
				} else {
					sendMessage(vesi, channelID, "Meme commands not enabled by Bot Shell :(")
				}

			}


		}
	}
}

func sendMessage(vesi *discordgo.Session, channelID, message string) {
	_, err := vesi.ChannelMessageSend(channelID, message)
	if err != nil {
		fmt.Printf("Error sending message due to %v", err)
	}
}
