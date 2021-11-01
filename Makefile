PHONY: help
help: ## Show this help.
	@echo Welcome to Altair Technical Test. Choice one of the following options to start:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY: app-up
app-up: ## Command to run the application
	docker-compose -f docker-compose.yml up --build

.PHONY: app-down
app-down: ## Command to shut down the application
	docker-compose -f docker-compose.yml down --remove-orphans