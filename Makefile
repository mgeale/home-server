.PHONY: test_db
test_db:
	docker run -p 3306:3306 --name mysql_test -e MYSQL_ROOT_PASSWORD=dbpass -e MYSQL_DATABASE=test_homeserver -d mysql:latest

.PHONY: clean_test_db
clean_test_db:
	migrate -database mysql://root:dbpass@(172.17.0.2:3306)/test_homeserver -path ./internal/db/migrations/ down

.PHONY: server
server:
	go run ./cmd/api/
	
.PHONY: graphql_gen
graphql_gen:
	go run github.com/99designs/gqlgen generate

.PHONY: graphql_server
graphql_server:
	go run server.go
