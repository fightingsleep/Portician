build:
	go mod download
	go mod verify
	go build -o bin/portician github.com/fightingsleep/portician/cmd/portician

build_image:
	docker build -t fightingsleep/portician -f ./deployments/Dockerfile .

push_image:
	docker push fightingsleep/portician:latest

run_image:
	docker compose -f ./deployments/docker-compose.yml up

run_image_background:
	docker compose -f ./deployments/docker-compose.yml up -d