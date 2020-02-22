# go-chi-mysql-test-server

## Usage
```
docker-compose up -d

curl -i http://localhost:80/list

curl -i -X POST -H "Content-Type: application/json" -d '{"key":"asdf", "value": "some value"}' http://localhost:80/add

docker-compose down
````