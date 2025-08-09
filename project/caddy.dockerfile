FROM caddy:builder-alpine AS builder

COPY Caddyfile /etc/caddy/Caddyfile

FROM caddy:alpine

COPY --from=builder /etc/caddy/Caddyfile /etc/caddy/Caddyfile