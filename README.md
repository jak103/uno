![Go](https://github.com/jak103/uno/workflows/Go/badge.svg?branch=master)

Read me

## To run 

`cd server/ && go run .`

In seperate terminal

`cd server/web && npm run-script build-watch`

This will start the back end server, and the front end will hot-reload when editing the frontend

## Running the frontend unit tests
Attach a shell to the container running the frontend (node) then run 
```shell
npm run test:unit
```
Coverage can be calculated by running
```shell
npm run test:coverage
```
