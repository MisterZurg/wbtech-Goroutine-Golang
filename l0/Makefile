# ——————————————————————————————————————————————————————
# ———— Docker CumPose Version
# ——————————————————————————————————————————————————————
.PHONY: up down
up:
	docker compose -f ./docker-compose.yml rm && \
	docker compose -f ./docker-compose.yml build --no-cache && \
	docker compose -f ./docker-compose.yml up

down:
	docker-compose -f ./docker-compose.yml down