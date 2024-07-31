#!/bin/bash

source app.env

migrate -path database/migrations -database $DB_URL -verbose up