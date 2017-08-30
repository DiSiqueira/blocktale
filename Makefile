COMPOSE ?= docker-compose
COMPILEDAEMON ?= CompileDaemon
CNT_NAME_APP ?= blocktale
GLIDE ?= glide

BUILD_DIR ?= build/

build:
	@printf "==> Building images\n"
	@$(COMPOSE) build

run:
	@printf "==> Starting the application\n"
	@$(COMPOSE) up -d

deps:
	@printf "==> Installing all deps\n"
	@$(GLIDE) install

start: deps build run
