-include .env
export

.PHONY: mock migrate rollback drop force migration

mock:
	mockery --all --keeptree

migrate:
	migrate -source file://postgres/migrations -database "$(DB_URL)" up

rollback:
	migrate -source file://postgres/migrations -database "$(DB_URL)" down

drop:
	migrate -source file://postgres/migrations -database "$(DB_URL)" drop

force:
	migrate -source file://postgres/migrations -database "$(DB_URL)" force -- -1

migration:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir postgres/migrations -seq $$name

run:
	go run ./cmd/graphqlserver 

generate:
	go generate ./...