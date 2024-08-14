# Docker
up:
	docker compose up -d

down:
	docker compose down

build:
	docker compose build

exec:
	docker exec -it grpc_protoc bash