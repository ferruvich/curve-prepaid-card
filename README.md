# Curve Challenge
Curve Ldt. Development Challenge: a small project for payments via prepaid card.

## Project Description
This project simulates a payment via prepaid cards.

The **user**, after being created, can create one or more **cards** and **deposit** some money on them.

When the user decides to pay something with his card, a **merchant** (that has already been created) can make an **authorization request** to block some amount of money. If the card contains enough money, the request is accepted and the amount is blocked.

After that, the merchant can **capture** some or all of the amount requested before, creating a **transaction** and receiving the money.

The merchant can **revert** some of the authorization request amount, making him unable to capture this later.

After capturing it, the merchant can **refund** the user, creating a transaction sending the money back to him.

If the user wants to know how he's spending his money, he can request the **transaction list** of his prepaid card. These transactions are of two types: **payment** if they're result of a capture on an authorization request, **refund** if they're result of a refund of some amount already captured.

## What is missing
- Authentication for both user and merchant;
- Since authentication is missing, anyone can make any operation knowing the card IDs or the authorization request ones;
- Better logs;
- Better HTTP responses, both for HTTP codes and JSON Bodies.

## Technologies used
- [Go 1.11](https://golang.org/) with [go mod](https://github.com/golang/go/wiki/Modules) to manage dependencies;
- [gin-gonic/gin](https://github.com/gin-gonic/gin) as HTTP web framework;
- [golang/mock](https://github.com/golang/mock) to mock interfaces, for Unit testing;
- [PostgreSQL](https://www.postgresql.org/) as DBMS;
- [docker](https://www.docker.com/) and [docker-compose](https://docs.docker.com/compose/) to deploy the project;

## How to run the project

### What do you need
- Go 1.11, since this project uses `go mod`;
- Docker and docker-compose;

### How to do it
- Install [docker](https://www.docker.com/) and [docker-compose](https://docs.docker.com/compose/)
- Make sure that your `docker.service` is active;
- The application will be served in port 8080 and the db uses port 5432. Please make sure that these port are available;
- Make sure your are in the project's root folder;
- Run `docker-compose up`

## API Documentation

### API Blueprint
[Here](https://github.com/ferruvich/curve-prepaid-card/tree/master/api/api-blueprint) you can find the API Blueprint documentation.
It can be served using [aglio](https://www.npmjs.com/package/aglio) with the following command on project's root folder:
```sh
    aglio -i api/api-blueprint/curve-prepaid-card-api.apib -s
```
you can find the documentation served on port 3000.
If you do not want to use `aglio`, you can find this api [here](https://curveprepaidcard.docs.apiary.io/#)

### How to perform API Calls 
A Postman documentation has been created in order to perform example calls. It is available [here](https://github.com/ferruvich/curve-prepaid-card/tree/master/api/postman-collection). 
To use it, download it, open Postman, click on **import** and follow the instructions.