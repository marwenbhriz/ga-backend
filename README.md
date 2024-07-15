module cmd/go-api-sample-todo

go 1.22.3

run mysql db 

docker run --name tasks-db -e MYSQL_ROOT_PASSWORD=root -d -p 3307:3306 mysql:latest 

docker exec -it tasks-db bash 


mysql -u root -proot 

CREATE DATABASE tasks;