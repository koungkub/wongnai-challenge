# Wongnai Challenge

This repository is build for [Carrers WeChallenge Program](https://careers.wongnai.com/development/wechallenge1)

## Docs

### Diagram

directiry `docs/diagram` collection of diagram for explain how api worked !!

### API Specification

run api specification using docker

```sh
cd docs/api_spec
docker run -d --rm --name slate -p 4567:4567 -v $(pwd)/build:/srv/slate/build -v $(pwd)/source:/srv/slate/source slate
```

and open this browser [here](http://localhost:4567)

## Usage

```sh
docker-compose -f docker-compose.yml up -d
```

and wait for 30 seconds for migrate data in databases before loadtest

## Testing

```sh
go test -v -cover ./...
```
