include .env
export

service-run:
	@go run main.go

service-deploy:
	docker compose up -d application

service-undeploy:
	docker compose down application

mig-up:
	migrate -path migrations -database ${abiba} up 1

mig-down:
	migrate -path migrations -database ${abiba} down 1

mig-v:
	migrate -path migrations -database ${abiba} version
