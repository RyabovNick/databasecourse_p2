OS=linux
ARCH=amd64

run:
	docker-compose up -d

# для использования --profile раскомментируйте profiles в docker-compose
db-run:
	docker-compose --profile db up -d

app-run:
	docker-compose --profile app up

app-run-build:
	docker-compose --profile app up --build

stop:
	docker-compose down