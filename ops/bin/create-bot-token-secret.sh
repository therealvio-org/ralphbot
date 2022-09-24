#!/bin/bash

#Preventing Credential leak in shell history
export HISTCONTROL=ignorespace

#Set variables
#Remember to set appropriate values where "xxx" is used
ENVIRONMENT='xxx' 

# the following commands are prefixed with whitespace
 BOT_TOKEN='xxx' #the Discord Bot Token from your Developer portal

 BOT_TOKEN_SECRET=$(printf '{
     "token": "%s",
     }' "$BOT_TOKEN")

TAGS=$(printf '[
    {
        "Key": "environment",
        "Value": "%s"
    },
]' "$ENVIRONMENT") 

#Create Secret
aws secretsmanager create-secret \
--name "ralphbot/token" \
--description 'Ralphbot Discord API Token' \
--secret-string "$BOT_TOKEN_SECRET" \
--tags "$TAGS"