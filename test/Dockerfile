# syntax=docker/dockerfile:1

FROM golang:1.21

ARG BW_GH_TOKEN
RUN git config --global url."https://${BW_GH_TOKEN}@github.com".insteadOf "https://github.com"

ENV GOPRIVATE="github.com/BetterWorks/*"
ENV MIGRATE_VERSION=v4.15.2
ENV SWAG_VERSION=1.16.2

WORKDIR /app
COPY go.* package.json ./

# Set the mode for the application to run in HTTP mode when deployed with Docker
ENV APP_MODE=http

RUN \
    # just
    curl --proto '=https' --tlsv1.2 -sSf https://just.systems/install.sh | bash -s -- --to /usr/local/bin \
    && chmod +x /usr/local/bin/just \
    # air
    && curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh \
    && sh install.sh \
    && mv ./bin/air /bin/air \
    && rm -rf ./bin install.sh \
    && mkdir -p out/tmp \
    # gotestsum
    && curl -sSL https://github.com/gotestyourself/gotestsum/releases/download/v1.9.0/gotestsum_1.9.0_linux_amd64.tar.gz | tar -xz -C /usr/bin gotestsum \
    # migrate
    && curl -L https://github.com/golang-migrate/migrate/releases/download/${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz | tar xvz \
    && mv migrate /usr/bin/migrate \
    && chmod +x /usr/bin/migrate \
    && rm LICENSE README.md \
    # swag
    && curl -L https://github.com/swaggo/swag/releases/download/v${SWAG_VERSION}/swag_${SWAG_VERSION}_Linux_x86_64.tar.gz | tar xvz \
    && mv swag /usr/bin/swag \
    && chmod +x /usr/bin/swag

COPY . .
RUN go mod download

EXPOSE 9000

CMD ["just"]
