init:
	docker-compose -p gin-cognito up -d --build
up:
	docker-compose -p gin-cognito up -d
serve:
	docker-compose -p gin-cognito exec gin-cognito-app go run main.go
sql:
	docker-compose -p gin-cognito exec gin-cognito-app go run sql.go
gin:
	docker-compose -p gin-cognito exec gin-cognito-app go run gin.go
app:
	docker-compose -p gin-cognito exec gin-cognito-app sh
db:
	docker-compose -p gin-cognito exec gin-cognito-db sh
test:
	docker-compose -p gin-cognito exec gin-cognito-app go clean -testcache
	docker-compose -p gin-cognito exec gin-cognito-app gotest -v ./tests/...
stop:
	docker-compose -p gin-cognito stop
down:
	docker-compose -p gin-cognito down
destroy:
	docker-compose -p gin-cognito down --rmi all --volumes
ps:
	docker-compose -p gin-cognito ps
migrate-create:
	migrate create -ext sql -dir ./database/migrations/ ${table}
migrate:
	migrate -database='mysql://test1234:test1234@tcp(127.0.0.1:3306)/test' -path=./database/migrations/ -verbose up
migrate-rollback:
	migrate -database='mysql://test1234:test1234@tcp(127.0.0.1:3306)/test' -path=./database/migrations/ -verbose down
