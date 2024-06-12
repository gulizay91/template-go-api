.PHONY: run_docker_restapi stop_docker_restapi docker_latest_image update_swagger

run_docker_restapi:
	docker rm -f template-go-api || true && \
    		docker rmi -f \$(docker images -q template-go-api) || true && \
				docker build -t template-go-api ./app && \
					docker run --rm --name template-go-api -p 7001:7001 -e SERVICE__ENVIRONMENT=dev -d template-go-api

stop_docker_restapi:
	docker stop template-go-api

docker_latest_image:
	docker build -t template-go-api ./app && \
		docker tag template-go-api guliz91/template-go-api:latest && \
			docker push guliz91/template-go-api:latest

update_swagger:
	swag init --parseDependency -g app/cmd/main.go -o app/docs