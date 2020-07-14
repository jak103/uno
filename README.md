Read me
## To run (in docker)
docker-compose build
docker-compose run --service-ports db firebase login
docker-compose run --service-ports db firebase init
docker-compose up

Run VSCode and attach to the running golang container using Remote - Containers

## To run (no docker)

`cd server/ && go run .`

In seperate terminal

`cd server/web && npm run-script build-watch`

This will start the back end server, and the front end will hot-reload when editing the frontend

[action workflow tutorial](https://docs.github.com/en/actions/creating-actions/creating-a-docker-container-action)
