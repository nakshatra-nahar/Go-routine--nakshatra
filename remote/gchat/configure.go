// SPDX-License-Identifier: Apache-2.0

package gchat

import (
	"github.com/rs/zerolog/log"

	"github.com/target/businessBot/models"
	"github.com/target/businessBot/utils"
)

func Configure(bot *models.Bot) {
	// emptyMap for substitute function
	// (it will only replace from env vars)
	emptyMap := map[string]string{}

	// google_chat_project_id
	projectID, err := utils.Substitute(bot.GoogleChatProjectID, emptyMap)
	if err != nil {
		log.Error().Msgf("could not set 'google_chat_project_id': %s", err.Error())
	}

	bot.GoogleChatProjectID = projectID

	// google_chat_credentials
	credentials, err := utils.Substitute(bot.GoogleChatCredentials, emptyMap)
	if err != nil {
		log.Error().Msgf("could not set 'google_chat_credentials': %s", err.Error())
	}

	bot.GoogleChatCredentials = credentials

	// google_chat_subscription_id
	subscriptionID, err := utils.Substitute(bot.GoogleChatSubscriptionID, emptyMap)
	if err != nil {
		log.Error().Msgf("could not set 'google_chat_subscription_id': %s", err.Error())
	}

	bot.GoogleChatSubscriptionID = subscriptionID
}
