![Github Actions](https://github.com/sf9v/nero/workflows/test/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/sf9v/nero/badge.svg?branch=master)](https://coveralls.io/github/sf9v/nero?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/sf9v/nero)](https://goreportcard.com/report/github.com/sf9v/nero)

# Nero

An experimental repository generator for Go.

## Motivation

Throughout our codebase, we heavily use the *[repository pattern](https://martinfowler.com/eaaCatalog/repository.html)* and we often [write our queries manually](https://golang.org/pkg/database/sql/#example_DB_QueryContext). It becomes tedious when you have more and more tables to maintain. One small change and you end-up changing a lot of things. So, we decided to experiment on creating this project to automatically generate our *[data-access layer](https://en.wikipedia.org/wiki/Data_access_layer)* code.

## Inspiration

Before we thought of this project, we tried many different ORMs and found [ent](https://entgo.io/). *Ent* is an ORM that uses *code generation* to make fluent and convenient data-access API. It is so convinient that we started using it on some parts of our system. 

However, due to our heavy use of *repository pattern*, it is still very difficult to for us to migrate all of our code. That's when we started to ponder about the idea of generating our data access layer code.

## Goals

- Decouple database specific logic by using `Repository` interface 
- Integrate with existing codebase
- Minimal API

## Assumption

The only assumption is every *collection/table* will always have a surrogate key i.e. *id column*.
