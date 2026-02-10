BIN := bin/awesome-backend
SWAG_MAIN := ./cmd/awesome-backend/main.go
SWAG_OUT  := ./docs
SWAG_DIRS := .,./internal/server

.PHONY: swaga build
swaga:
	@echo ">> generating swagger"
	GOFLAGS=-mod=mod go run github.com/swaggo/swag/cmd/swag@latest \
		init -g $(SWAG_MAIN) -o $(SWAG_OUT) -d $(SWAG_DIRS) \
		--parseDependency --parseInternal

build: swaga
	@echo ">> building binary"
	mkdir -p bin
	CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o $(BIN) ./cmd/awesome-backend
