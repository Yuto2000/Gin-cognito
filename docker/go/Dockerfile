FROM golang:1.18-alpine

WORKDIR /work

RUN apk update && \
    apk add git && \
    apk add curl && \
    apk add build-base

# RUN go get "github.com/aws/aws-sdk-go/aws" && \
#     go get "github.com/aws/aws-sdk-go/aws/session" && \
#     go get "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"

# RUN go install gorm.io/gorm@latest && \
#     go install gorm.io/driver/mysql@latest && \
#     go install github.com/gin-gonic/gin@latest && \
#     go install github.com/rakyll/gotest@latest && \
#     go install github.com/jmoiron/sqlx@latest && \
#     go install github.com/golang-migrate/migrate@latest && \
#     go install github.com/go-sql-driver/mysql@latest