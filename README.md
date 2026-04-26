#  Booking-web-service

Film-web-service - это простой REST API сервис для тестирования запросов.

### Способ установки


Запустить docker-compose
```
docker-compose up -d
```

Импортировать JSON файлы в БД Mflix, предворительно скопировав содержиое data
```
mongoimport --uri=mongodb://admin:admin@localhost:27017 --db mflix --collection users --authenticationDatabase=admin --file users.json
mongoimport --uri=mongodb://admin:admin@localhost:27017 --db mflix --collection movies --authenticationDatabase=admin --file movies.json
mongoimport --uri=mongodb://admin:admin@localhost:27017 --db mflix --collection sessions --authenticationDatabase=admin --file sessions.json
mongoimport --uri=mongodb://admin:admin@localhost:27017 --db mflix --collection theaters --authenticationDatabase=admin --file theaters.json
mongoimport --uri=mongodb://admin:admin@localhost:27017 --db mflix --collection comments --authenticationDatabase=admin --file comments.json
```
