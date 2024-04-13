# COMMON #
#################################################################################################################
# Пароль универсальный
PASSWORD=qwerty

# Порт на хосте
HOST_PORT=5436

# Указание что это не названия файлов
.PHONY: docker-run docker-stop docker-status docker-logs migrate-up migrate-down git-push

# FOR DOCKER #
#################################################################################################################
# Имя контейнера Docker
CONTAINER_NAME=serviceBanner-db

# Запуск контейнера Docker c удалением при остановке
docker-run:
	docker run --name=$(CONTAINER_NAME) -e POSTGRES_PASSWORD=$(PASSWORD) -p $(HOST_PORT):5432 -d --rm postgres

# Остановка контейнера Docker
docker-stop:
	docker stop $(CONTAINER_NAME)

# Вывод статуса всех контейнеров Docker
docker-status:
	docker ps -a

# Вывод логов контейнера Docker
docker-logs:
	docker logs $(CONTAINER_NAME)

# FOR MIGRATE #
#################################################################################################################
# Migrate up
migrate-up:
	migrate -path ./schema -database 'postgres://postgres:$(PASSWORD)@localhost:$(HOST_PORT)/postgres?sslmode=disable' up

# Migrate down
migrate-down:
	migrate -path ./schema -database 'postgres://postgres:$(PASSWORD)@localhost:$(HOST_PORT)/postgres?sslmode=disable' down

# FOR GIT #
#################################################################################################################
# Commit+push в свою ветку
git-push:
	git add .
	git commit -m "$(m)"
	git push origin develop