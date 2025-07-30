export MYSQL_URL='mysql://root:babi123@tcp(localhost:3306)/mysqlgo'

migrate-create:
	@ migrate create -ext sql -dir scripts/migrations -seq $(name)
migrate-up:
	@ migrate -database ${MYSQL_URL} -path scripts/migrations up
migrate-down:
	@ migrate -database ${MYSQL_URL} -path scripts/migrations down
migrate-force:
	@ migrate -database ${MYSQL_URL} -path scripts/migrations force 1 
