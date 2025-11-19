.PHONY: bash bash-debug

bash:
	docker compose up user-service -d --build

bash-debug:
	docker compose --profile debug up user-service-debug -d --build
