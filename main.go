package main

import (
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
)

var config Config

func main() {
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatalln(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln(err)
		return
	}

	discord, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatalln(err)
		return
	}

	discord.AddHandler(messageHandler)
	discord.AddHandler(messageReactionAddHandler)
	discord.AddHandler(messageReactionRemoveHandler)

	err = discord.Open()
	if err != nil {
		log.Fatalln(err)
		return
	}

	err = discord.UpdateGameStatus(0, "Freero.am")
	if err != nil {
		log.Fatalln(err)
		return
	}

	<-make(chan struct{})
	return
}
