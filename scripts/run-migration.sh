#!/bin/bash

export POSTGRES_USER=asadbek
export POSTGRES_PASSWORD=3066586
export POSTGRES_DATABASE=orderdb

migrate -database "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@localhost:5432/$POSTGRES_DATABASE?sslmode=disable" -path="../migrations" up
