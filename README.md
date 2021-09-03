Simple crud using go and graphql with clean architecture.

#### How to run

```
$ cd go-graphql
$ cp .env.example .env
$ cd go-graphql/src
$ go mod tidy

Run migration and seeder
$ go run db/migration_cli.go migrate
$ go run db/migration_cli.go seed

Run this app
$ go run cmd/main.go
```

Next, open `localhost:9090/grahpql`
