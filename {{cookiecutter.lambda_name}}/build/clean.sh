#!/usr/bin/env bash

# Stop and remove DynamoDB container
DYNAMODB_CONTAINER_NAME=$1
DYNAMOBD_DOCKERID=$(docker ps -aqf "name=$DYNAMODB_CONTAINER_NAME")
docker stop $DYNAMOBD_DOCKERID && docker rm $DYNAMOBD_DOCKERID || true
sleep 2