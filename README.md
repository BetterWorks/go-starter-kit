# go-starter-kit
Go Starter Kit for Betterworks Server Applications

## Documentation
- [Architecture](./documentation/architecture.md)
- [Getting Started](./documentation/getting-started.md)
- [AWS Lambda](./documentation/working-with-lambda.md)

## Installation
Clone the repository
```sh
$ git clone git@github.com:BetterWorks/go-starter-kit.git
$ cd go-starter-kit
```

## Development
**Prerequisites**
- *[Docker Desktop](https://www.docker.com/products/docker-desktop)*
- *[Go 1.20+](https://golang.org/doc/install)*
- *[AWS SAM](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)*

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

**Create new migration**
```sh
$ docker compose run --rm api just migrate-create {name}
```

### Server
**Run the server in development mode**
```sh
$ docker compose run --rm --service-ports api
```

### Testing
**Run the integration test suite with code coverage**
```sh
$ docker compose run --rm api just coverage
```

## Building
**Compile server binary**
```sh
$ go build -mod vendor -o out/bin/domain ./cmd/httpserver
```

## Contributing
1. Create feature branch (`git switch -c new-feature`)
1. Commit changes using [conventional changelog standards](https://www.conventionalcommits.org) (`git commit -m 'feat(scope): adds new feature'`)
1. Push to the branch (`git push origin new-feature`)
1. Ensure linting and tests are passing
1. Create new pull request

## License
UNLICENSED
