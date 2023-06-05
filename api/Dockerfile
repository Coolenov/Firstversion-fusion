FROM golang:alpine

WORKDIR /app
COPY go.mod .
RUN go mod download

RUN go get -u github.com/gin-gonic/gin \
    && go get -u gorm.io/gorm \
    && go get github.com/joho/godotenv \
    && go get -u gorm.io/driver/mysql

COPY . .


RUN go build -o main cmd/apiserver.go

CMD ["./main"]