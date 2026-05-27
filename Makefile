db_login:
	psql ${POSTGRES_URL}

db_create_migration:
	migrate create -ext sql -dir migrations -seq ${name}

db_migrate:
	migrate -database ${POSTGRES_URL} -path migrations up