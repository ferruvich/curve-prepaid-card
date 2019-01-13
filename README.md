# Curve Challenge
Curve Ldt. Development Challenge: a small project for payments via prepaid card.

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
A Postman documentation has been created in order to perform example calls. It is available [here](https://github.com/ferruvich/curve-prepaid-card/tree/master/api/postman-collection)