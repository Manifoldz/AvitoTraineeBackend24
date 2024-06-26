# COMMON #
#################################################################################################################
# Пароль универсальный
PASSWORD=qwerty

# Порт на хосте
HOST_PORT=5436

# FOR DOCKER #
#################################################################################################################
# Имя контейнера Docker
CONTAINER_NAME=serviceBanner-db

# Запуск контейнера Docker
docker-run:
	docker run --name=$(CONTAINER_NAME) -e POSTGRES_PASSWORD=$(PASSWORD) -p $(HOST_PORT):5432 -d postgres

# Остановка контейнера Docker
docker-stop:
	docker stop $(CONTAINER_NAME)

# Продолжить работу контейнера
docker-start:
	docker start $(CONTAINER_NAME)

# Удалить контейнер
docker-remove: docker-stop
	docker rm $(CONTAINER_NAME)

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
	migrate -path ./schema -database \
	'postgres://postgres:$(PASSWORD)@localhost:$(HOST_PORT)/postgres?sslmode=disable' up

# Migrate down
migrate-down:
	migrate -path ./schema -database \
	'postgres://postgres:$(PASSWORD)@localhost:$(HOST_PORT)/postgres?sslmode=disable' down
# FOR GIT #
#################################################################################################################
# Commit+push в свою ветку
git-push:
	git add .
	git commit -m "$(m)"
	git push origin develop

# FOR CURL #
#################################################################################################################
# Создать баннеры
post-createBanner:
	curl -i -X POST http://localhost:8000/banner/ \
	-H "Content-Type: application/json" \
	-H "token: admin" \
	-d '{"feature_id": 1, "tag_ids": [1,2,3], \
	"content": {"title": "another_title", "text": "another_text", "url": "another_url"}, "is_active": true}'
	curl -i -X POST http://localhost:8000/banner/ \
	-H "Content-Type: application/json" \
	-H "token: admin" \
	-d '{"feature_id": 2, "tag_ids": [1,5,4], \
	"content": {"title": "some_title", "text": "some_text", "url": "some_url"}, "is_active": false}'
	curl -i -X POST http://localhost:8000/banner/ \
	-H "Content-Type: application/json" \
	-H "token: admin" \
	-d '{"feature_id": 2, "tag_ids": [6,7,8], \
	"content": {"title": "title", "text": "text", "url": "url"}, "is_active": false}'
	curl -i -X POST http://localhost:8000/banner/ \
	-H "Content-Type: application/json" \
	-H "token: user" \
	-d '{"feature_id": 4, "tag_ids": [6,7,8], \
	"content": {"title": "title", "text": "text", "url": "url"}, "is_active": false}'

# Запросить все баннеры с фильтрацией
get-getBannerFiltered:
	curl -i -X GET http://localhost:8000/banner/ \
	-H "Content-Type: application/json" \
	-H "token: admin" \
	-d '{"tag_id": 1, "limit": 10, "offset": 0}'
	curl -i -X GET http://localhost:8000/banner/ \
	-H "Content-Type: application/json" \
	-H "token: admin" \
	-d '{"feature_id": 2, "limit": 10, "offset": 0}'
	curl -i -X GET http://localhost:8000/banner/ \
	-H "Content-Type: application/json" \
	-H "token: admin" \
	-d '{}'
	curl -i -X GET http://localhost:8000/banner/ \
	-H "Content-Type: application/json" \
	-H "token: user" \
	-d '{"feature_id": 2, "limit": 10, "offset": 0}'

# FOR INSTALL #
#################################################################################################################
# Установка docker и migrate. Установка внутрь докера postgres.
install:
	brew install --cask docker
	brew install golang-migrate
	docker pull postgres

# Запуск приложения
go-start:
	go run cmd/main.go

