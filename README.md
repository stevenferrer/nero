![Github Actions](https://github.com/sf9v/nero/workflows/ci/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/sf9v/nero/badge.svg?branch=master)](https://coveralls.io/github/sf9v/nero?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/sf9v/nero)](https://goreportcard.com/report/github.com/sf9v/nero)

# Nero

An experimental repository generator in Go.

## Goals

- By using the `Repository` interface, database specific logic could be implemented separately e.g. PostgreSQL, MySQL/MariaDB, SQLite, MSSQL, Oracle, MongoDB etc.
- Integrate with existing codebase by implementing the `Schemaer` interface

## Assumption

Every `collection/table` shall have a surrogate key i.e. `id` column.
