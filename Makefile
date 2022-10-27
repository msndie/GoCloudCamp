all: up

up:
	docker-compose up --build -d;

client:
	docker build -t grpc-client -f ./client.Dockerfile .;
	docker run --network="host" grpc-client;

down:
	docker-compose down;

clean:
	docker system prune -a;

.PHONY: all up client down clean