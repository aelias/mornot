MUTANTORNOT_IMAGE_NAME=mutantornot
DNASTATS_IMAGE_NAME=dnastats
MNGINX_IMAGE_NAME=mnginx
# You have to export your docker user name
DOCKER_USER=${docker_user_env}

# Build run and test solution
all: docker-build docker-run docker-test docker-stop

# Build both containers
docker-build:
		docker build -t $(MUTANTORNOT_IMAGE_NAME) -f mutantornot/Dockerfile .
		docker build -t $(DNASTATS_IMAGE_NAME) -f dnastats/Dockerfile .
		docker build -t $(MNGINX_IMAGE_NAME) -f mnginx/Dockerfile .
		docker-compose build

docker-run:
	docker-compose up -d
	@echo "Pause for wait docker is running"
	sleep 10
	@echo "Continuing"

docker-stop:
	docker-compose stop

# Running tests
docker-test:
		docker-compose exec mutant bash -c "go test ./... -v"
		docker-compose exec dna bash -c "go test ./... -v"

# Tag and push images to docker registry
docker-push: docker-build
ifeq ($(DOCKER_USER),)
	@echo "You have to export your docker user (export docker_user_env=xxx)"
else		
		docker tag $(MUTANTORNOT_IMAGE_NAME) $(DOCKER_USER)/$(MUTANTORNOT_IMAGE_NAME)
		docker tag $(DNASTATS_IMAGE_NAME) $(DOCKER_USER)/$(DNASTATS_IMAGE_NAME)
		docker push $(DOCKER_USER)/$(MUTANTORNOT_IMAGE_NAME)
		docker push $(DOCKER_USER)/$(DNASTATS_IMAGE_NAME)
		docker tag $(MNGINX_IMAGE_NAME) $(DOCKER_USER)/$(MNGINX_IMAGE_NAME)
		docker push $(DOCKER_USER)/$(MNGINX_IMAGE_NAME)
endif		
