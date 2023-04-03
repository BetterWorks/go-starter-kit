# gosk-api
Go Starter Kit for Betterworks Server Applications

## Documentation
- [Architecture](./documentation/architecture.md)
- [Getting Started](./documentation/getting-started.md)

## Installation
Clone the repository
```sh
$ git clone git@github.com:BetterWorks/gosk-api.git
$ cd gosk-api
```

## Development
**Prerequisites**
- *[Docker Desktop](https://www.docker.com/products/docker-desktop)*
- *[Go 1.20+](https://golang.org/doc/install)*

**Show all commands**
```sh
$ docker compose run --rm api just
```

### Migrations
**Run all up migrations**
```sh
$ docker compose run --rm api just migrate
```

**Run up migrations {n} steps**
```sh
$ docker compose run --rm api just migrate-up svcdb {n}
```

**Run down migrations {n} steps**
```sh
$ docker compose run --rm api just migrate-down svcdb {n}
```

### Server
**Run the server in development mode**
```sh
$ docker compose run --rm --service-ports api
```

## Building
**Compile server binary**
```sh
$ go build -mod vendor -o out/bin/domain ./cmd/server
```

## Contributing
1. Create feature branch (`git switch -c new-feature`)
1. Commit changes using [conventional changelog standards](https://www.conventionalcommits.org) (`git commit -m 'feat(scope): adds new feature'`)
1. Push to the branch (`git push origin new-feature`)
1. Ensure linting and tests are passing
1. Create new pull request

## License
Copyright (c) 2022 Jason Sites
