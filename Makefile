CURRENT_DIR=$(shell pwd)

APP=$(shell basename ${CURRENT_DIR})
APP_CMD_DIR=${CURRENT_DIR}/cmd

TAG=latest
ENV_TAG=latest

migration-up:
	migrate -path ./migrations/postgresql -database 'postgres://testuser:0000@localhost:5432/voronej?sslmode=disable' up

migration-down:
	migrate -path ./migrations/postgresql -database 'postgres://testuser:0000@localhost:5432/voronej?sslmode=disable' down

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

run:
	go run cmd/main.go