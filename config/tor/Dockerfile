FROM alpine:latest

RUN apk add --no-cache tor

# Copy torrc config
COPY config/tor/torrc /etc/tor/torrc

# Copy hidden service keys
COPY config/tor/vanity/* /var/lib/tor/hidden_service/

# Set correct ownership and permissions
RUN chown -R tor:tor /var/lib/tor && \
    chmod -R 700 /var/lib/tor

# Run tor as the correct user
USER tor

ENTRYPOINT [ "tor", "-f", "/etc/tor/torrc" ]
