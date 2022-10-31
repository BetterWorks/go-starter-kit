# syntax=docker/dockerfile:1

FROM golang:1.19
ENV MIGRATE_VERSION=v4.15.2
WORKDIR /app
COPY go.* package.json ./

RUN \
    # just
    curl --proto '=https' --tlsv1.2 -sSf https://just.systems/install.sh | bash -s -- --to /usr/local/bin && \
    chmod +x /usr/local/bin/just && \
    # migrate
    curl -L https://github.com/golang-migrate/migrate/releases/download/${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/bin/migrate && \
    chmod +x /usr/bin/migrate && \
    # air
    curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh && \
    chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

COPY . .
RUN go mod vendor
# RUN just lint

EXPOSE 9202
EXPOSE 9204

CMD ["just"]
