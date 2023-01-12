run:
	 docker-compose build
	 docker-compose up

migrate:
	migrate -path ./migrations -database 'mysql://user:Password@123@tcp(localhost:3305)/golang?multiStatements=true' up


mod:
	go mod tidy
	go mod vendor


