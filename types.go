package main

type Config struct {
	Token                   string `json:"token"`
	PublicRelationsMessages struct {
		ListenChannel        string `json:"listen_channel"`
		AnnouncementsChannel string `json:"announcements_channel"`
	} `json:"public_relations_messages"`
	ManageBotChannel string                       `json:"manage_bot_channel"`
	ReactionRoles    map[string]map[string]string `json:"reaction_roles"`
}
