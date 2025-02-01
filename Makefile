# Variables  
APP_NAME = cs-ticketing
DOCKER_IMAGE = $(APP_NAME):latest  
DOCKERFILE = Dockerfile  
CONTAINER_NAME = cs-ticketing
  
# Default target  
all: build  
  
# Build the Docker image  
build:  
	docker build -t $(DOCKER_IMAGE) -f $(DOCKERFILE) .  
  
# Run the Docker container  
run: 
	nodemon
  
# Stop and remove the Docker container  
stop:  
	docker-compose down  
  
# Clean up Docker images and containers  
clean:  
	docker rm -f $$(docker ps -aq) || true  
	docker rmi -f $(DOCKER_IMAGE) || true  
	docker rmi -f cs-ticketing-sod-app || true  

rebuild: stop clean run

# Run tests  
test:  
	go test $(TEST_FLAGS) ./...  

# Migration
db-migrate:
	go run cmd/migrate/main.go