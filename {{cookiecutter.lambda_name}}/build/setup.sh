#!/usr/bin/env bash
set -x
set -e

# Run DynamoDB docker container
DYNAMODB_CONTAINER_NAME=$1

# Check if it's running
if [ ! "$(docker ps -q -f name=$1)" ]; then
    if [ "$(docker ps -aq -f status=exited -f name=$1)" ]; then
        # cleanup
        docker rm $1
    fi
    # Run it with local data dir
    DATADIR="$(pwd)/db"
    rm -rf $DATADIR; mkdir $DATADIR
    echo "Created temporary data directory: $DATADIR"
    docker run -d --name $DYNAMODB_CONTAINER_NAME -p 8000:8000 -v $DATADIR:/db/ amazon/dynamodb-local -jar DynamoDBLocal.jar -dbPath /db/ -sharedDb
    sleep 2

    # Create table
    aws dynamodb --endpoint-url http://localhost:8000 \
        create-table \
        --table-name flights \
        --attribute-definitions AttributeName=id,AttributeType=S \
        --key-schema AttributeName=id,KeyType=HASH \
        --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1
fi
