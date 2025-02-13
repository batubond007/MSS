### Setup

Change to ./build directory, then

``` docker-compose up```

### Mock Data

To drop mock data

``` go run .\.build\build.go Drop ```


To insert mock user data

``` go run .\.build\build.go InsertUser ```


To insert mock message data

``` go run .\.build\build.go InsertMessage ```

### API'S

Register request is used to register to messaging service.

``
curl --location 'http://localhost:8080/register' \
--header 'Content-Type: application/json' \
--data '{
    "phone": "+905553333333",
    "password": "password"
}'
``


Unregister request is used to unregister to messaging service.

``
curl --location 'http://localhost:8080/unregister' \
--header 'Content-Type: application/json' \
--data '{
    "phone": "+905553333333",
    "password": "password"
}'
``


Sent List request returns the sent messages for user.

``
curl --location 'http://localhost:8080/sent-list/?page_size=4' \
--header 'Content-Type: application/json' \
--data '{
    "phone": "+905551111111",
    "password": "password"
}'
``


### The Clean Architecture

![image](https://github.com/batubond007/MSS/assets/46222914/fec4c554-ba92-4cbb-bee8-8a346d4b675d)

### Domain Driven Design
![image](https://github.com/batubond007/MSS/assets/46222914/f13c6d0d-7f0c-422d-b1c1-8824dd07a672)


