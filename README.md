![Github Actions](https://github.com/sf9v/nero/workflows/test/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/sf9v/nero/badge.svg?branch=master)](https://coveralls.io/github/sf9v/nero?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/sf9v/nero)](https://goreportcard.com/report/github.com/sf9v/nero)

# Nero

A repository generator for Go.

## Motivation

Throughout our codebase, we heavily use the *[repository pattern](https://martinfowler.com/eaaCatalog/repository.html)* and we often [write our queries manually](https://golang.org/pkg/database/sql/#example_DB_QueryContext). It becomes tedious when you have more and more tables to maintain. One small change and you end-up changing a lot of things. So, we decided to experiment on creating this project to automatically generate our *[data-access layer](https://en.wikipedia.org/wiki/Data_access_layer)* code.

## Inspired by

* [*ent*](https://entgo.io) - An entity framework for Go. Simple, yet powerful ORM for modeling and querying data.

## Goals

- Decouple database specific logic by using `Repository` interface 
- Easy integration with existing codebase
- Minimal and easy to use API

## Assumption

The only assumption is every *collection/table* will always have a surrogate key i.e. *id column*.
