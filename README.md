# Nero ![Github Actions](https://github.com/sf9v/nero/workflows/test/badge.svg) [![Coverage Status](https://coveralls.io/repos/github/sf9v/nero/badge.svg?branch=master)](https://coveralls.io/github/sf9v/nero?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/sf9v/nero)](https://goreportcard.com/report/github.com/sf9v/nero)

Nero is a library for generating the repository layer code. Please see this [blog post](https://sf9v.github.io/posts/generating-the-repository-layer-in-go/) for more details.

## Installation

```console
$ go get github.com/sf9v/nero
```

## Example

See the [integration test](./test/integration) for other examples.

```go
func main() {
    dsn := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
    db, err := sql.Open("postgres", dsn)
    ...

    repo := repository.NewPostgresRepository(db).Debug()
    // create 
    productID, err := repo.Create(
        context.Background(), 
        repository.NewCreator().Name("Product 1"),
    )
	...	

	// query
	product, err := repo.QueryOne(
        context.Background(), 
        repository.NewQueryer().
            Where(repository.IDEq(product1ID)),
    )
	...

	// update
	now := time.Now()
	_, err = repo.Update(
        context.Background(), 
        repository.NewUpdater().
            Name("Updated Product 1").
            UpdatedAt(&now).
            Where(repository.IDEq(product1ID)),
    )
	...

	// delete 
	_, err = repo.Delete(
        context.Background(), 
        repository.NewDeleter().
            Where(repository.IDEq(product1ID)),
    )
    ...
}
```

## Motivation

We heavily use the *[repository pattern](https://threedots.tech/post/repository-pattern-in-go/)* in our codebases and we often [write our queries manually](https://golang.org/pkg/database/sql/#example_DB_QueryContext). It becomes tedious, boring and repetitive as we have more and more tables/models to maintain. One small change and we end-up changing a lot of things in different places. 

So, we decided to experiment on creating this library to auto-generate our repository layer code.

## Goals

- Decouple database specific logic by using `Repository` interface 
- Easy integration with existing codebase
- Minimal and easy to use API

## Supported back-ends

Currently, we have official support for [PostgreSQL](postgresql.org), ther back-ends will be supported soon.

| Back-end | Library | 
|---------| ------- |
| [PostgreSQL](https://postgresql.org) | [lib/pq](http://github.com/lib/pq) |
| [SQLite](https://sqlite.org) (soon) | |
| MySQL (soon) | |

## Custom back-ends

You can support custom back-ends by implementing the [_Templater_](./templater.go) interface. This interface is specifically created to support customisability.

To implement a custom back-end, you can refer to [postgres template](./template/postgres.go). 

## Inspired by

This library is inspired the the following projects:

* [*ent*](https://github.com/facebook/ent). An entity framework for Go. Simple, yet powerful ORM for modeling and querying data.
* [SQLBoiler](https://github.com/volatiletech/sqlboiler) is a tool to generate a Go ORM tailored to your database schema.

## Contributing

If you have any suggestions and ideas on how to improve this library, please feel free to [open an issue](https://github.com/sf9v/nero/issues) or making a [pull request](https://github.com/sf9v/nero/pulls).