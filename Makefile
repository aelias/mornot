MUTANTORNOT_IMAGE_NAME=mutantornot
DNASTATS_IMAGE_NAME=dnastats

# Build and run containers
all: docker docker-test docker-run

# Build both containers
docker:
		docker build -t $(MUTANTORNOT_IMAGE_NAME) -f mutantornot/Dockerfile .
		docker build -t $(DNASTATS_IMAGE_NAME) -f dnastats/Dockerfile .

# Running tests
docker-test:
		# docker run $(MUTANTORNOT_IMAGE_NAME) bash -c "go test ./... -v"
		# docker run $(DNASTATS_IMAGE_NAME) bash -c "go test ./... -v"

# Run both containers in deatached mode
docker-run:
		docker run -d -p8081:8081 $(MUTANTORNOT_IMAGE_NAME)
		docker run -d -p8082:8082 $(DNASTATS_IMAGE_NAME)		