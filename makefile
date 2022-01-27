createDB:
	psql -h localhost -U postgres -w -c "create database test;"

dropDB:
	psql -h localhost -U postgres -w -c "drop database test;"

migrateUP:
	migrate -path db/migrations -database "postgres://postgres:0712@localhost:5432/test?sslmode=disable" -verbose up

migrateDOWN:
	migrate -path db/migrations -database "postgres://postgres:0712@localhost:5432/test?sslmode=disable" -verbose down
	
.PHOBY: createDB dropDB migrateUP migrateDOWN