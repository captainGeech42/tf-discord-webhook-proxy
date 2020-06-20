# Terraform -> Discord Webhook Proxy

[![Go Report Card](https://goreportcard.com/badge/github.com/captainGeech42/tf-discord-webhook-proxy)](https://goreportcard.com/report/github.com/captainGeech42/tf-discord-webhook-proxy) [![Build Status](https://travis-ci.com/captainGeech42/tf-discord-webhook-proxy.svg?branch=master)](https://travis-ci.com/captainGeech42/tf-discord-webhook-proxy)

Have you ever wanted to use the webhook notification system on Terraform to notify a Discord channel of your infrastructure state changes, only to realize that they can't natively talk to each other? Well not anymore, because here's the tool you've been searching for!

_Terraform webhook comes in, Discord webhook goes out, profit._

By default, an embedded rich message will be sent to Discord but this can be disabled in the config file.

![Rich Message](https://i.imgur.com/Q9uNRqV.png)

## Usage

1. Download: `go get github.com/captainGeech42/tf-discord-webhook-proxy`
2. Copy the `config.ex.json` file to your current directory as `config.json`
3. Update the `WebhookURL` field with a Discord webhook URL ([Discord docs on webhooks](https://support.discord.com/hc/en-us/articles/228383668))
4. Run it: `tf-discord-webhook-proxy`

## Docker Image

This tool is also available via a Docker image. When running via the Docker image, you can either use this image as a base image to `COPY` your `config.json` into `/app`, or set the following environment variables instead:

* `TF_PROXY_ENV=YES` (without this, a `config.json` will be looked for)
* `TF_PROXY_WEBHOOK_URL="https://discordapp.com/api/webhooks/xxxxxxxx/yyyyyyyyyyyyy"`
* `TF_PROXY_RICH_MESSAGES=YES` (optional, disable rich messages by setting to `NO`)

The Docker image will always have the proxy running on port 8080 in the container, you can choose to forward this outside of the container to whatever port you need.

Example execution of container:

```
docker run --rm -it -p8080:8080 \
           -e TF_PROXY_ENV=YES \
           -e TF_PROXY_WEBHOOK_URL="https://discordapp.com/api/webhooks/xxxxxxxx/yyyyyyyyyyyyy" \
           zzzanderw/tf-discord-webhook-proxy:latest
```