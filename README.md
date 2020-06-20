# Terraform -> Discord Webhook Proxy

[![Go Report Card](https://goreportcard.com/badge/github.com/captainGeech42/tf-discord-webhook-proxy)](https://goreportcard.com/report/github.com/captainGeech42/tf-discord-webhook-proxy) [![Build](https://github.com/captainGeech42/tf-discord-webhook-proxy/workflows/Build/badge.svg)](https://github.com/captainGeech42/tf-discord-webhook-proxy/actions?query=workflow%3ABuild) [![Docker Hub Publish](https://github.com/captainGeech42/tf-discord-webhook-proxy/workflows/Docker%20Hub%20Publish/badge.svg)](https://github.com/captainGeech42/tf-discord-webhook-proxy/actions?query=workflow%3A%22Docker+Hub+Publish%22) [![Docker Hub Image](https://img.shields.io/docker/v/zzzanderw/tf-discord-webhook-proxy?color=blue)](https://hub.docker.com/repository/docker/zzzanderw/tf-discord-webhook-proxy/general)

Have you ever wanted to use the webhook notification system on Terraform Cloud to notify a Discord channel of your infrastructure state changes, only to realize that they can't natively talk to each other? Well not anymore, because here's the tool you've been searching for!

_Terraform webhook comes in, Discord webhook goes out, profit._

By default, an embedded rich message will be sent to Discord but this can be disabled in the config file.

Sample embed message:

![Rich Message](https://i.imgur.com/hkjoS4Z.png)

The color of the message (on the left side) will be either green (success), red (error), blue (plan needs confirmation), or yellow (other).

## Usage

1. Download: `go get github.com/captainGeech42/tf-discord-webhook-proxy`
2. Copy the `config.ex.json` file to your current directory as `config.json`
3. Update the `WebhookURL` field with a Discord webhook URL ([Discord docs on webhooks](https://support.discord.com/hc/en-us/articles/228383668))
4. Run it: `tf-discord-webhook-proxy`

The proxy will be available at `http://host:8080/webhook`. Add that URL to a new Webhook Notification in the Notifications settings in your Terraform Cloud workspace.

## Docker Image

This tool is also available via a Docker image on Docker Hub ([`zzzanderw/tf-discord-webhook-proxy`](https://hub.docker.com/repository/docker/zzzanderw/tf-discord-webhook-proxy)). When running via the Docker image, you can either use this image as a base image to `COPY` your `config.json` into `/app`, or set the following environment variables instead:

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