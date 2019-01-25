MUTANTORNOT_IMAGE_NAME=mutantornot
DNASTATS_IMAGE_NAME=dnastats
# You have to export your docker user name
DOCKER_USER=${docker_user_env}

# Build and run containers
all: docker docker-test

# Build both containers
docker-build:
		docker build -t $(MUTANTORNOT_IMAGE_NAME) -f mutantornot/Dockerfile .
		docker build -t $(DNASTATS_IMAGE_NAME) -f dnastats/Dockerfile .

# Running tests
docker-test:
		docker run $(MUTANTORNOT_IMAGE_NAME) bash -c "go test ./... -v"
		docker run $(DNASTATS_IMAGE_NAME) bash -c "go test ./... -v"

# Run both containers in deatached mode
docker-run:
		docker run -d -p8081:8081 $(MUTANTORNOT_IMAGE_NAME)
		docker run -d -p8082:8082 $(DNASTATS_IMAGE_NAME)		

# Tag and push images to docker registry
docker-push: docker-build
ifeq ($(DOCKER_USER),)
	@echo "You have to export your docker user (export docker_user_env=xxx)"
else		
		docker tag $(MUTANTORNOT_IMAGE_NAME) $(DOCKER_USER)/$(MUTANTORNOT_IMAGE_NAME)
		docker tag $(DNASTATS_IMAGE_NAME) $(DOCKER_USER)/$(DNASTATS_IMAGE_NAME)
		docker push $(DOCKER_USER)/$(MUTANTORNOT_IMAGE_NAME)
		docker push $(DOCKER_USER)/$(DNASTATS_IMAGE_NAME)
endif		
