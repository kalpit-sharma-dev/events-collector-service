## Instructions to Run the service in docker

1. git clone git@github.com:kalpit-sharma-dev/kalpit.cool2006-gmail.com.git (clone the repo)
2. sudo docker-compose up (start the container)
3. open terminal connect to database inside docker
4. docker ps -a
5. docker exec -it containerid bash
6. psql -h localhost -p 5432 -U admin -d eventscollector
7. after connecting to the DB run the sql script to create the table in DB src/kalpit.cool2006-gmail.com/sql/create_tables.sql
8. Follow documentation in openapi.yml to test the API's