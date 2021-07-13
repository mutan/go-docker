### Запускаем контейнер с компилятором Go и сразу билдим
docker container run --rm -it -w /go/src/app -v "$PWD":/go/src/app --name test-go golang:1.16.6-buster go build -v

docker-compose -f docker/docker-compose.yml run golang go run hello.go