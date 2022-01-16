current_dir=$(shell pwd)

gen-front:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate -i /local/api/openapi.yaml -g typescript-axios -o /local/front/gen/openapi_cli -c /local/front/api.json

gen-backend:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate -i /local/api/openapi.yaml -g go-server -o /local/backend/gen/opencliapi --git-repo-id=fs/backend --git-user-id=mytord --additional-properties=sourceFolder=.,addResponseHeaders=true,router=chi

push-docker-images:
	docker-compose build && docker-compose push

start-dev:
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d

start-prod:
	docker-compose pull
	docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d
	docker-compose run fs-migrate up
