PHONY: clean test

test:
	docker compose down --remove-orphans
	docker compose run test

clean:
	docker compose down --remove-orphans
	docker system prune -f