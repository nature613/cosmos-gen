# cosmos-gen

[![CircleCI](https://circleci.com/gh/shanev/cosmos-gen.svg?style=svg)](https://circleci.com/gh/shanev/cosmos-gen)

cosmos-gen generates boilerplate for building modules for [Cosmos SDK](https://github.com/cosmos/cosmos-sdk).

## Install

```sh
go get -u github.com/shanev/cosmos-gen
```

## Usage

```sh
cosmos-gen claim
```

This generates the following boilerplate files:

```
└── x
    ├── claim
    │   ├── codec.go
    │   ├── genesis.go
    │   ├── handler.go
    │   ├── keeper.go
    │   ├── module.go
    │   ├── params.go
    │   ├── querier.go
    │   └── types.go
```
