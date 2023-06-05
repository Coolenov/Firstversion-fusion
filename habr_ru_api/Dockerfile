FROM golang:alpine

WORKDIR /app

RUN go get -u github.com/gin-gonic/gin
RUN go get -u gorm.io/gorm
RUN go get github.com/joho/godotenv
RUN go get -u gorm.io/driver/mysql


COPY . .
COPY ../lib .
COPY ../initialize .


CMD ["go run", "cmd/api.habr.go"]