#!/bin/bash

#Preventing Credential leak in shell history
export HISTCONTROL=ignorespace

#Set variables
#Remember to set appropriate values where "xxx" is used
ENVIRONMENT='xxx' 

# the following commands are prefixed with whitespace
 SECRET_ARN_SUFFIX='xxx'

 SECRET_ARN_SUFFIX=$(printf '{
     "token": "%s",
     }' "$BOT_TOKEN")

TAGS=$(printf '[
    {
        "Key": "environment",
        "Value": "%s"
    },
]' "$ENVIRONMENT") 

#Create Secret
aws ssm put-parameter \
--name "/ralphbot/token/arn-suffix" \
--description "The ARN suffix for Ralphbot's Secrets Manager secret. This is due to CDK not \"knowing\" what the full ARN is for a Secrets Manager secret by design." \
--value "$SECRET_ARN_SUFFIX" \
--tags "$TAGS"