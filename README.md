takonews-batch
===========

[![Build Status](https://travis-ci.org/takonews/takonews-batch.png?branch=master)](https://travis-ci.org/takonews/takonews-batch)
[![codecov](https://codecov.io/gh/takonews/takonews-batch/branch/master/graph/badge.svg)](https://codecov.io/gh/takonews/takonews-batch)

## Development

```
go get -u github.com/takonews/takonews-batch
mysql -u root -p
> create database if not exists takonews_development;
godep restore
go run main.go
```

## Deployment

```
go get -u github.com/takonews/takonews-batch
mysql -u root -p
> create database if not exists takonews_production;
godep restore
go run main.go
```

## Local testing

```
go test -v $(go list ./... | grep -v vendor)
```

## LICENSE

MIT.
