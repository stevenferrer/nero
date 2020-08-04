![Github Actions](https://github.com/sf9v/nero/workflows/ci/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/sf9v/nero/badge.svg?branch=master)](https://coveralls.io/github/sf9v/nero?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/sf9v/nero)](https://goreportcard.com/report/github.com/sf9v/nero)

# Nero

An experimental repository generator in Go.

## Motivation

Throughout our codebase, we heavily use the `Repository` pattern and we often [write our queries manually](https://golang.org/pkg/database/sql/#example_DB_QueryContext). It becomes tedious when you have more and more tables to maintain. So, we decided to create `nero` to automate the task of writing the repository implementations.

## Inspiration

In search for alternative [data-access layer](https://en.wikipedia.org/wiki/Data_access_layer), we discovered [ent](https://entgo.io/), an ORM that uses *code generation* to make fluent data-access API. We fell in-love with the idea and that's when we started to ponder about generating a *data-access layer* code that would *fit our needs*.

## Goals

- Decouple database specific logic by using `Repository` interface 
- Integrate with existing codebase
- Minimal API

## Assumption

Every `collection/table` shall always have a surrogate key i.e. `id` column.
