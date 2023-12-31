package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"zdim/spotbot/pkg/api"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	GuildID  = flag.String("guild", "", "Test Guild ID")
)

func init() {
	flag.Parse()
}

func main() {
	fmt.Println("starting bot...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	botToken := os.Getenv("TOKEN")
	appId := os.Getenv("APP_ID")

	session, err := discordgo.New("Bot " + botToken)
	if err != nil {
		fmt.Println("error starting bot:", err)
		return
	}

	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("bot started")
	})

	session.AddHandler(func(s *discordgo.Session, ixn *discordgo.InteractionCreate) {
		switch ixn.Type {
		case discordgo.InteractionApplicationCommand:
			if ixn.ApplicationCommandData().Name == "album" {
				// make the request to spotify in here
				album := api.GetAlbum("guid_goes_here")
				s.InteractionRespond(ixn.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content:    "Response message",
						Embeds: []*discordgo.MessageEmbed{
							{
								URL:         album.URL,
								Title:       album.Title,
								Description: fmt.Sprintf("%v - %v", album.Artist, album.Year),
								Image:       &discordgo.MessageEmbedImage{URL: album.Image},
							},
						},
					},
				})
			}

		}
	})

	_, err = session.ApplicationCommandCreate(appId, *GuildID, &discordgo.ApplicationCommand{
		Name:        "album",
		Description: "Post a compact album preview for a provided link or uri",
	})

	if err != nil {
		log.Fatalf("Failed to create slash command: %v", err)
	}

	err = session.Open()
	if err != nil {
		log.Fatalf("Failed to open the session: %v", err)
	}
	defer session.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Shutting down")
}
