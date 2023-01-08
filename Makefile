server:
	DB_PW=${dbpass} ./homeserver -env=production

server_build:
	env GOOS=linux GOARCH=arm GOARM=5 go build -o homeserver ./cmd/api/

server_run:
	go run ./cmd/api/

test_db:
	docker run -p 3306:3306 --name mysql_test -e MYSQL_ROOT_PASSWORD=${dbpass} -e MYSQL_DATABASE=test_homeserver -d mysql:latest

clean_test_db:
	migrate -database mysql://root:${dbpass}@(172.17.0.2:3306)/test_homeserver -path ./internal/db/migrations/ down
	
graphql_gen:
	go run github.com/99designs/gqlgen generate

graphql_playground:
	go run server.go
