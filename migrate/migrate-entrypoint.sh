#!/bin/sh

DSN="mysql://$MYSQL_USER:$MYSQL_PASSWORD@tcp($MYSQL_HOST:3306)/$MYSQL_DATABASE?query"

migrate -path=/migrations -database=${DSN} up
