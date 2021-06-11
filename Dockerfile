# First stage
FROM golang:rc-alpine3.12 AS build

ARG WEBHOOK_URL
ARG TITLE
ARG MESSAGE

RUN apk add gcc  g++

ENV GO11MODULE=on

WORKDIR /app

COPY . .

RUN go build

# Second stage
FROM alpine

ARG WEBHOOK_URL
ARG TITLE
ARG MESSAGE

ENV WEBHOOK_URL $WEBHOOK_URL
ENV TITLE $TITLE
ENV MESSAGE $MESSAGE

COPY --from=build /app/msteams-webhook-message-action /bin/msteams-webhook-message-action

ENTRYPOINT ["/bin/msteams-webhook-message-action"]
