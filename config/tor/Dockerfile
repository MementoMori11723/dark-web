FROM alpine:latest

RUN apk add --no-cache tor

COPY config/tor/torrc /etc/tor/torrc

ENTRYPOINT [ "tor", "-f", "/etc/tor/torrc" ]
