# homeserver

docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=dbpass -e MYSQL_DATABASE=homeserver -d mysql:latest
docker run -p 3306:3306 --name mysql_test -e MYSQL_ROOT_PASSWORD=dbpass -e MYSQL_DATABASE=test_homeserver -d mysql:latest

migrate create -ext sql -dir ./internal/db/migrations/ -seq create_users_table
migrate -database mysql://root:dbpass@(172.17.0.2:3306)/homeserver -path ./internal/db/migrations/ up

"root:dbpass@(172.17.0.2:3306)/test_homeserver?parseTime=true&multiStatements=true"

sudo mysql -u root -p
sudo mysql -D homeserver -u web -p