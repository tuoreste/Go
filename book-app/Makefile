GO := go

SRC_DIR := sources

# MAIN_FILE := main.go

run:
	$(GO) run . $(wildcard $(SRC_DIR)/*.go)

build:
	$(GO) build -o app . $(wildcard $(SRC_DIR)/*.go)

clean:
	rm -f app