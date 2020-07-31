![Github Actions](https://github.com/sf9v/nero/workflows/test/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/sf9v/nero/badge.svg?branch=master)](https://coveralls.io/github/sf9v/nero?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/sf9v/nero)](https://goreportcard.com/report/github.com/sf9v/nero)

# Nero

An experimental repository generator in Go.

## Assumptions

- a collection should always have an *identity* column i.e. *id*

## Goals

- Decouple database specific logic (CRUD) by using the `Repository` interface i.e. Postgres, MariaDB, MongoDB etc. 
  Actual repository then could be implemented separately.
- Use with pre-existing structs in your code by implementing the `Schemaer` interface
- Generate only the client that you need i.e. I only use Postgres, so include only Postgres implementation

## Priorty features

- Basic CRUD operations
- Support for transactions
- Sqlite3 & Postgresql