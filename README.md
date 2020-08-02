![Github Actions](https://github.com/sf9v/nero/workflows/ci/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/sf9v/nero/badge.svg?branch=master)](https://coveralls.io/github/sf9v/nero?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/sf9v/nero)](https://goreportcard.com/report/github.com/sf9v/nero)

# Nero

An experimental repository generator in Go.

## Goals

- Decouple database specific logic by using `Repository` interface 
- Integrate with existing codebase
- Minimal

## Assumption

Every `collection/table` shall always have a surrogate key i.e. `id` column.
