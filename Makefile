api-gen:
	docker compose exec payment oapi-codegen -generate "server" -package handler ./openapi/payment.yaml > ./internal/payment/handler/server.gen.go
	docker compose exec payment oapi-codegen -generate "types" -package handler ./openapi/payment.yaml > ./internal/payment/handler/types.gen.go
migrate-%:
	docker run -v ./database/migration:/migrations --network host migrate/migrate \
    	-path=/migrations/ -database "mysql://$(shell cat ./database/conf/${@:migrate-%=%}.json | jq -r .user):$(shell cat ./database/conf/${@:migrate-%=%}.json | jq -r .password)@tcp($(shell cat ./database/conf/${@:migrate-%=%}.json | jq -r .host):$(shell cat ./database/conf/${@:migrate-%=%}.json | jq -r .port))/$(shell cat ./database/conf/${@:migrate-%=%}.json | jq -r .database)" up
	docker run -v ./database/migration:/migrations --network host migrate/migrate \
    	-path=/migrations/ -database "mysql://$(shell cat ./database/conf/${@:migrate-%=%}.json | jq -r .user):$(shell cat ./database/conf/${@:migrate-%=%}.json | jq -r .password)@tcp($(shell cat ./database/conf/${@:migrate-%=%}.json | jq -r .host):$(shell cat ./database/conf/${@:migrate-%=%}.json | jq -r .port))/$(shell cat ./database/conf/${@:migrate-%=%}.json | jq -r .database)_test" up
seed-%:
	cat ./database/seed/service.sql | docker compose exec -T rdb mysql -u$(shell cat ./database/conf/${@:seed-%=%}.json | jq -r .user) -p$(shell cat ./database/conf/${@:seed-%=%}.json | jq -r .password) $(shell cat ./database/conf/${@:seed-%=%}.json | jq -r .database)

dbclear-%:
	docker compose exec -T rdb mysql -u$(shell cat ./database/conf/${@:dbclear-%=%}.json | jq -r .user) -p$(shell cat ./database/conf/${@:dbclear-%=%}.json | jq -r .password) -e "DROP DATABASE $(shell cat ./database/conf/${@:dbclear-%=%}.json | jq -r .database);"
	docker compose exec -T rdb mysql -u$(shell cat ./database/conf/${@:dbclear-%=%}.json | jq -r .user) -p$(shell cat ./database/conf/${@:dbclear-%=%}.json | jq -r .password) -e "CREATE DATABASE $(shell cat ./database/conf/${@:dbclear-%=%}.json | jq -r .database);"
	docker compose exec -T rdb mysql -u$(shell cat ./database/conf/${@:dbclear-%=%}.json | jq -r .user) -p$(shell cat ./database/conf/${@:dbclear-%=%}.json | jq -r .password) -e "DROP DATABASE $(shell cat ./database/conf/${@:dbclear-%=%}.json | jq -r .database)_test;"
	docker compose exec -T rdb mysql -u$(shell cat ./database/conf/${@:dbclear-%=%}.json | jq -r .user) -p$(shell cat ./database/conf/${@:dbclear-%=%}.json | jq -r .password) -e "CREATE DATABASE $(shell cat ./database/conf/${@:dbclear-%=%}.json | jq -r .database)_test;"