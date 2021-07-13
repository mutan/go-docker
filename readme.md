# Запускаем контейнер с компилятором Go и сразу билдим
docker container run --rm -it -w /go-app -v ~/docker-for-go/cmd/go-app:/go-app --name test-go golang:1.14-buster go build