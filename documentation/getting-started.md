# Getting Started Guide

The following procedure is intended to help you quickly customize this starter kit for your unique application.

## Verify Working App
1. Clone the starter kit.
2. Delete the local `.git` directory.
3. Ensure you can successfully run the app for local development and test coverage. See the [Development](../README.md#development) section of the [README](../README.md) for further instructions.

### Ports
4. Choose a unique service port for the HTTP server to avoid collisions with other service port bindings on the host machine. Replace `9202` with your unique service port in the following files:
  - `/config/config.toml`
  - `/docker-compose.yml`
  - `/Dockerfile`
  - `/test/Dockerfile`
5. Choose a unique service port for the PostgreSQL server. In the `/docker-compose.yml` file, replace the port mapping `25432:5432` with `{custom-postgres-port}:5432`.
6. Choose a unique service port for the Redis server. In the `/docker-compose.yml` file, replace the port mapping `26379:6379` with `{custom-redis-port}:6379`. Alternatively, just remove the `redis` service and the `api.depends_on.redis` configuration.

### App Name
7. Choose a name for the app. This will be referred to as `{appname}` in the rest of this guide.
8. Update the go module file and all import paths:
  - Find all occurrences of `github.com/BetterWorks/gosk-api` and replace with `github.com/BetterWorks/{appname}`
  - Delete `/go.sum`
  - Run `go mod tidy && go mod vendor`
9. Update the `/package.json` `name` field to `{appname}`.
10. Update the `/justfile` `project` field to `{appname}`.

At this point, check that the app is still in a working state by rebuilding and rerunning the docker service containers:
```sh
$ docker compose down
$ docker compose build --no-cache
$ docker compose run --rm --service-ports api
```
:exclamation: Please familiarize yourself with the [Application Design & Architecture](architecture.md) before proceeding with [Code Changes](#code-changes) and [Migrations](#migrations).

## Adding Resources
The starter kit contains a single generic resource called `Resource`. You can use this type as a reference when creating the resources for the app, and then delete this generic code later.

### Code Changes
By example, let's say we're adding a new resource called `TShirt`.

The following is a list of locations where code needs to be modified to add the new resource:
- add `/internal/types/tshirt.go` containing the various type definitions
- modify `/internal/types/domain.go` to add:
  - `TShirt string` to the `DomainRegistry` struct
  - `TShirt: "tshirt"` to the `DomainType` variable
- add `/internal/repo/tshirt.go` repository code
- add `/internal/domain/tshirt.go` service code
- modify `/internal/domain/domain.go` to add:
  - `TShirtService types.Service` to the `Services` struct
- add `/internal/httpapi/routes/tshirt.go` with `TShirtRouter` route definitions
- modify `/internal/httpapi/router.go` to add:
  - `TShirtController *controllers.Controller` to the `controllerRegistry` struct
  - `TShirtController: ...` to the `controllerRegistry` in the `registerControllers` function
  - `routes.TShirtRouter(app, c.TShirtController, ns)` to the `registerRoutes` method
- modify `/internal/resolver/singletons.go` to add:
  - a `RepositoryTShirt` method mirroring the `RepositoryResource` method
  - an entry for `svcTShirt` in the `Domain` method
- modify `/internal/resolver/resolver.go` to add:
  - `RepoTShirt types.Repository` to the `Config` struct
  - `repoTShirt types.Repository` to the `Resolver` struct
  - `repoTShirt: c.RepoTShirt,` to the `Resolver` instantiation in the `NewResolver` function
  - an entry for `r.RepositoryTShirt()` in the `Initialize` method

:warning: It will likely take some time to get working code in all those locations. Once complete, the `Resource` code can be removed, and the app code developed from there.

:exclamation: Please note that these patterns are just a beginning. Codebases with higher complexity require more than what is demonstrated here with a simple, generic use case.

### Migrations
The existing migrations work for the `Resource` type. It may be advantageous to simply add new migrations for the new resources, and then circle back with a clean `initial-schema` migration once the app is in a stable state.

To create a new set of initial migrations:
  - delete the existing set of `/database/migrations/*.sql` files
  - use the commands described in the [Migrations](../README.md#migrations) section of the [README](../README.md).

## Last Steps
Run `git init` and go build cool stuff!
