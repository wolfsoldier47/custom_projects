# URL shortner

This is a custom URL shortner which you can use to generate SHORT URLS

## how to run

docker-compose up -d

## structure

User ---> GO,Go-fibre ---> REDIS


API ----> DB,Helpers,Routes,Dockerfile
DATA
DB ----> DockerFile


DockerComposefile

## API
http://localhost:3000/api/v1
/:url

## Test

wget --method=POST --header="Content-Type: application/json" --body-data='{"url": "https://www.youtube.com/watch?v=pL_l_YnzPAE"}' http://localhost:3000/api/v1
