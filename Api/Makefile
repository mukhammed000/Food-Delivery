CURRENT_DIR := $(shell pwd)

DBURL := "postgres://postgres:1234@localhost:5432/delivery?sslmode=disable"

run:
	go run cmd/main.go

migrate-create:
	migrate create -ext sql -dir migrations create_users_table

mig-up:
	migrate -path migrations -database ${DBURL} -verbose up

mig-down:
	migrate -path migrations -database ${DBURL} -verbose down

proto-gen: 
	./scripts/gen-proto.sh ${CURRENT_DIR}

SWAGGER := $(HOME)/go/bin/swag
SWAGGER_DOCS := docs
SWAGGER_INIT := $(SWAGGER) init -g ./api/router.go -o $(SWAGGER_DOCS)

swag-gen:
	@echo "Running: $(SWAGGER_INIT)"
	$(SWAGGER_INIT)
