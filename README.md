# Nero ![Github Actions](https://github.com/sf9v/nero/workflows/test/badge.svg) [![Coverage Status](https://coveralls.io/repos/github/sf9v/nero/badge.svg?branch=main)](https://coveralls.io/github/sf9v/nero?branch=main) [![Go Report Card](https://goreportcard.com/badge/github.com/sf9v/nero)](https://goreportcard.com/report/github.com/sf9v/nero)

Nero is a library for generating the repository layer code. Please see this [blog post](https://sf9v.github.io/posts/generating-the-repository-layer-in-go/) for more details.

## Installation

```console
$ go get github.com/sf9v/nero
```

## Example

See the [integration test](./test/integration) for other examples.

```go
import (
    "database/sql"

    _ "github.com/lib/pq"

    // import the generated package
    "github.com/sf9v/nero-example/repository"
)

func main() {
    dsn := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
    db, err := sql.Open("postgres", dsn)
    ...

    // initialize a new postgres repository
    repo := repository.NewPostgresRepository(db).Debug()

    // create a product
    productID, err := repo.Create(
        context.Background(), 
        repository.NewCreator().Name("Product 1"),
    )
    ...	

    // query the product
    product, err := repo.QueryOne(
        context.Background(), 
        repository.NewQueryer().
            Where(repository.IDEq(product1ID)),
    )
    ...

    // update the product
    now := time.Now()
    _, err = repo.Update(
        context.Background(), 
        repository.NewUpdater().
            Name("Updated Product 1").
            UpdatedAt(&now).
            Where(repository.IDEq(product1ID),
        ),
    )
    ...
    
    // delete the product
    _, err = repo.Delete(
        context.Background(), 
        repository.NewDeleter().
            Where(repository.IDEq(product1ID),
        ),
    )
    ...
}
```

## Motivation

We heavily use the *[repository pattern](https://threedots.tech/post/repository-pattern-in-go/)* in our codebases and we often [write our queries manually](https://golang.org/pkg/database/sql/#example_DB_QueryContext). It becomes tedious, boring and repetitive as we have more and more tables/models to maintain. So, we decided to experiment on creating this library to auto-generate our repository layer code.

## Goals

- Decouple database specific logic by using `Repository` interface 
- Easy integration with existing codebase
- Minimal and easy to use API

## Supported back-ends

Currently, we have official support for [PostgreSQL](postgresql.org). Other back-ends shall be supported soon. Meanwhile, you can implement [custom back-ends](#custom-back-ends).

| Back-end | Library | 
|---------| ------- |
| [PostgreSQL](https://postgresql.org) | [lib/pq](http://github.com/lib/pq) |
| [SQLite](https://sqlite.org) (soon) | |
| MySQL/MariaDB (soon) | |
| MongoDB ??? | |

## Custom back-ends

Implementing your own custom back-end is very easy. In fact, you don't have to use the official back-ends. You can implement custom back-ends (ClickHouse, Cassandra, CouchDB , Badger, H2, etc.) by implementing the [_Templater_](./templater.go) interface. This interface is specifically created to support extensibility and customisability.

You can refer to the official [postgres template](./template/postgres.go) and this [example schema](./example/user.go#L46).

## Inspired by

This library is inspired by these amazing projects:

* [ent](https://github.com/facebook/ent). An entity framework for Go. Simple, yet powerful ORM for modeling and querying data.
* [SQLBoiler](https://github.com/volatiletech/sqlboiler) is a tool to generate a Go ORM tailored to your database schema.

## Contributing

Any suggestions and ideas on how to improve this library are very much welcome! Please feel free to [open an issue](https://github.com/sf9v/nero/issues) or by [making a pull request](https://github.com/sf9v/nero/pulls).