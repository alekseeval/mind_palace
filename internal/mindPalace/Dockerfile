# syntax=docker/dockerfile:1

# BUILD stage
FROM golang:1.19.0-alpine3.15 as build-image

WORKDIR /etc/mp_app/

COPY . ./
RUN go mod download
RUN go build -o /bin/mp_app/server /etc/mp_app/cmd/mpapp

# DEPLOY stage
FROM alpine:3.15

WORKDIR /etc/mp_app

COPY --from=build-image /bin/mp_app/server /bin/mp_app/server
COPY --from=build-image /etc/mp_app/internal/mindPalace/config.yaml /etc/mp_app/config.yaml

CMD [ "/bin/mp_app/server" ]