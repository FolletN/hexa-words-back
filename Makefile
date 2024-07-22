dev: clean
	DB_NAME="postgres" \
	DB_USER="user" \
	DB_PASSWORD="password" \
	DB_HOST="localhost" \
	DB_PORT="5432" \
	SERVER_PORT="8080" \
	docker-compose up -d

# run:
# 	DATABASE_DATABASE="postgres" \
#     DATABASE_USER="user" \
#     DATABASE_PASSWORD="password" \
#     DATABASE_ADDRESS="localhost" \
#     DATABASE_PORT="5432" \
# 	go run cmd/main.go

clean:
	docker rm -vf hexacrosswords-api && docker rmi -f hexacrosswords-back-api

log:
	docker logs hexacrosswords-api -f
