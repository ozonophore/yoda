#!/bin/sh
migrate -path changelog -database "postgresql://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable" -verbose up
