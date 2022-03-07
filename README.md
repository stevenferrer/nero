[![GoDoc Reference](https://pkg.go.dev/badge/github.com/stevenferrer/nero)](https://pkg.go.dev/github.com/stevenferrer/nero)
![Github Actions](https://github.com/stevenferrer/nero/workflows/test/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/stevenferrer/nero/badge.svg?branch=main)](https://coveralls.io/github/stevenferrer/nero?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/stevenferrer/nero)](https://goreportcard.com/report/github.com/stevenferrer/nero)

# Nero

A library for generating the repository pattern.

## Motivation

Writing _[repositories](https://www.martinfowler.com/eaaCatalog/repository.html)_ [manually](https://golang.org/pkg/database/sql/#example_DB_QueryContext) can be tedious, so why not just generate them?

## Goals

Below are the goals of this project.

- Decouple implementation from the `Repository` interface
- Easy integration with existing codebase
- Minimal and easy to use API

## Installation

```console
$ go get github.com/stevenferrer/nero
```

## Example

See the [official example](https://github.com/stevenferrer/nero-example) and [demo test](./test/demo-test/playerrepo) for a more complete example.

```go
import (
    "database/sql"

    // import the generated package
    "github.com/stevenferrer/nero-example/productrepo"
)

func main() {
    dsn := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
    db, err := sql.Open("postgres", dsn)
    ...

    ctx := context.Background()


    // initialize the repository (and optionally enable debug)
    productRepo := productrepo.NewPostgresRepository(db).Debug()

    // create
    creator := productrepo.NewCreator().Name("Product 1")
    productID, err := productRepo.Create(ctx, creator)
    ...

    // query
    queryer := productrepo.NewQueryer().Where(productrepo.IDEq(product1ID))
    product, err := productRepo.QueryOne(ctx, queryer)
    ...

    // update
    now := time.Now()
    updater := productrepo.NewUpdater().Name("Updated Product 1").
        UpdatedAt(&now).Where(productrepo.IDEq(product1ID))
    _, err = productRepo.Update(ctx, updater)
    ...

    // delete
    deleter := productrepo.NewDeleter().Where(productrepo.IDEq(product1ID))
    _, err = productRepo.Delete(ctx, deleter)
    ...
}
```

## Supported back-ends

Below is the list of supported back-ends.

| Back-end             | Library                                                       |
| -------------------- | ------------------------------------------------------------- |
| PostgreSQL           | [lib/pq](http://github.com/lib/pq)                            |
| SQLite               | [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3)       |
| MySQL/MariaDB (soon) | [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) |

If your your back-end is not yet supported, you can write your own [custom template](#custom-templates).

## Custom templates

For writing custom templates, the official templates can serve as a reference. You can write your own template with custom back-ends (Oracle, MSSQL, BoltDB, MongoDB, etc.) by implementing the [_Template_](./template.go) interface.

See official [postgres template](./postgres_template.go) for reference.

## Limitations

Currently, the supported operations are basic CRUD and aggregate (i.e. count, sum). For more complex requirements, we suggest that you just write your repositories manually for now.

We're still planning how to support compound predicates (and, or), joins and relationships. Feel free to share any ideas!

## Standing on the shoulders of giants

This project wouldn't be possible without these amazing open-source projects:

- [Masterminds/squirrel](https://github.com/Masterminds/squirrel) - Fluent SQL generation in golang
- [lib/pq](https://github.com/lib/pq) - Pure Go Postgres driver for database/sql
- [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) - sqlite3 driver conforming to the built-in database/sql interface
- [pkg/errors](https://github.com/pkg/errors) - Simple error handling primitives
- [hashicorp/multierror](https://github.com/hashicorp/go-multierror) - A Go (golang) package for representing a list of errors as a single error.
- [jinzhu/inflection](https://github.com/jinzhu/inflection) - Pluralizes and singularizes English nouns
- [stretchr/testify](https://github.com/stretchr/testify) - A toolkit with common assertions and mocks that plays nicely with the standard library

Also, the following have a huge influence on this project and deserves most of the credits:

- [ent](https://github.com/facebook/ent) - An entity framework for Go. Simple, yet powerful ORM for modeling and querying data.
- [SQLBoiler](https://github.com/volatiletech/sqlboiler) - Generate a Go ORM tailored to your database schema.

## What's in the name?

The project name was inspired by an [anti-bird](https://blackclover.fandom.com/wiki/Anti-bird) in an anime called [Black Clover](https://blackclover.fandom.com/wiki/Black_Clover_Wiki). The anti-bird, which the [Black Bulls](https://blackclover.fandom.com/wiki/Black_Bull) squad calls _Nero_ is apparently a human named [Secre Swallowtail](https://blackclover.fandom.com/wiki/Secre_Swallowtail). It's a really cool anime with lots of magic!

## Contributing

Any suggestions and ideas are very much welcome, feel free to [open an issue](https://github.com/stevenferrer/nero/issues) or [make a pull request](https://github.com/stevenferrer/nero/pulls)!

## License

[Apache-2.0](LICENSE)
