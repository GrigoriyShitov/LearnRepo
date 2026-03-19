


all: infra pause apps

infra:
	@echo "Starting infra..."
	docker-compose -f docker-compose-infra.yaml up --build -d
pause:
	@echo "waiting for infra launch"
	timeout /t 5
apps:
	@echo "Starting apps"
	docker-compose -f docker-compose.yaml up --build -d
