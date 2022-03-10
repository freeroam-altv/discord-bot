package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	switch m.ChannelID {
	case config.PublicRelationsMessages.ListenChannel:
		publicRelationsMessageHandler(s, m)
		break
	case config.ManageBotChannel:
		// TODO: manage bot channel
		break
	}
}

func messageReactionAddHandler(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	role := getReactionRoles(m.ChannelID, m.Emoji.MessageFormat())
	if role == "" {
		return
	}

	err := s.GuildMemberRoleAdd(m.GuildID, m.UserID, role)
	if err != nil {
		log.Println(err)
	}
}

func messageReactionRemoveHandler(s *discordgo.Session, m *discordgo.MessageReactionRemove) {
	role := getReactionRoles(m.ChannelID, m.Emoji.MessageFormat())
	if role == "" {
		return
	}

	err := s.GuildMemberRoleRemove(m.GuildID, m.UserID, role)
	if err != nil {
		log.Println(err)
	}
}

func publicRelationsMessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// TODO: add embed.

	_, err := s.ChannelMessageSend(config.PublicRelationsMessages.AnnouncementsChannel, m.Content)
	if err != nil {
		log.Println(err)
		return
	}
}
