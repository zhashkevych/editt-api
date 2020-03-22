if [ "$(docker ps -q -f name=editt-api)" ]; then
    if [ ! "$(docker ps -aq -f status=exited -f name=editt-api)" ]; then
        docker stop editt-api
    fi
    docker run -e HOST --rm -d --publish 8000:8000 --network editt --name editt-api editt:0.1
fi