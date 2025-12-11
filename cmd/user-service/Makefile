.PHONY: bash bash-debug

bash-user-service:
	docker compose up user-service -d --build

bash-user-service-debug:
	docker compose --profile debug up user-service-debug -d --build
