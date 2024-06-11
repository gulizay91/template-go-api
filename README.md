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
docker run -p 8061:8061 -e SERVICE__ENVIRONMENT=dev --name template-go-api template-go-api
```