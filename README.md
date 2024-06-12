# template-go-api

## Swagger
## Generate Swagger Doc
```sh
# /template-go-api>
swag init --parseDependency -g app/cmd/main.go -o app/docs
```

## Docker
### build docker image
```sh
# /template-go-api>
docker build -t template-go-api ./app
docker run -p 7001:7001 -e SERVICE__ENVIRONMENT=dev --name template-go-api template-go-api
```
or you can use phony target
### makefile
```sh
# /template-go-api>
make run_docker_restapi
make stop_docker_restapi
make docker_latest_image
make update_swagger
```