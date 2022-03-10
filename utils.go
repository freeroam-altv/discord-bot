package main

func getReactionRoles(channelID, emoji string) string {
	reactionsRole, foundReaction := config.ReactionRoles[channelID]
	if !foundReaction {
		return ""
	}

	role, foundRole := reactionsRole[emoji]
	if !foundRole {
		return ""
	}

	return role
}
