docker-compose build
docker-compose run --service-ports db firebase login
docker-compose run --service-ports db firebase init
docker-compose up

Run VSCode and attach to the running golang container using Remote - Containers

