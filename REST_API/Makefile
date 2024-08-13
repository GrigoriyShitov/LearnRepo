

all:
	sudo docker build -t go-rest-api .
	sudo docker run -d --network=host \
  -e "DB_DBNAME=restapi" \
  -e "DB_PORT=5432" \
  -e "DB_USER=grish" \
  -e "DB_PASS=1129" \
  -e "DB_HOST=127.0.0.1" -p 8080:8080 go-rest-api

build:
	sudo docker build -t go-rest-api .

run:
	sudo docker run -d --network=host \
  -e "DB_DBNAME=restapi" \
  -e "DB_PORT=5432" \
  -e "DB_USER=grish" \
  -e "DB_PASS=1129" \
  -e "DB_HOST=127.0.0.1" -p 8080:8080 go-rest-api

stop:
	sudo docker kill $(sudo docker ps -a -q)