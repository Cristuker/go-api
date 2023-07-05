# Go API

API de produtos criada em Go usando arquitetura hexagonal.

Técnologias e métodologias usadas

* Arquitetura Hexagonal
* Testes
* Injeção de dependências


## Adapters

* Web server
* CLI

## Como configurar

```bash
// Para criar um container do projeto
$ docker compose up -d

// Para entrar no container do projeto
$ docker exec -it  appproduct bash

//Dentro do container para link o sqlite
$ sqlite3 db.sqlite

```

## Comandos

Para rodar o projeto via CLI

```go
$ go run main.go cli
```

Para rodar o projeto via Webserver
```go
$ go run main.go http
```

## SQL

```sql
1. create table products(id string, name string, price float, status string);

2. sqlite3 db.sqlite
```

<p align="center">Made whit ❤️ by <strong><a href="https://bento.me/cmsdev" target="blank" >Cristian</></p></strong>