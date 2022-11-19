all: up logs

up:
	docker-compose up --build -d;

client:
	docker build -t grpc-client -f ./client.Dockerfile .;
	docker run --network="host" grpc-client;

down:
	docker-compose down;

logs:
	docker-compose logs;

clean:
	docker system prune -a;

.PHONY: all up client down clean