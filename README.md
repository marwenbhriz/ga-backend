module cmd/go-api-sample-todo

go 1.22.3

run mysql db 

docker run --name books-db -e MYSQL_ROOT_PASSWORD=root -d -p 3307:3306 mysql:latest 

docker exec -it books-db bash 


mysql -u root -proot 

CREATE DATABASE books;