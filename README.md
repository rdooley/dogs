# dogs
An Api for dogs


Made on go1.12


Install deps with `go mod vendor`


Run api server with `go run main.go`

Uses uri param and json bindings to parse parameters, dogs data stored in dogs.json


A create
```
dogs git:master ❯ curl -X POST \
  http://localhost:8080/dogs \
  -H 'content-type: application/json' \
  -d '{ "Name": "example", "Owner": "human", "Details":"a wonderful example" }'
{"ID":2}%
```

Get a specific dog
```
dogs git:master ❯ curl -X GET \
  http://localhost:8080/dogs/2
{"ID":2,"Name":"example","Owner":"human","Details":"a wonderful example"}%
```

Get all dogs
```
dogs git:master ❯ curl -X GET \
  http://localhost:8080/dogs
[{"ID":1,"Name":"manute","Owner":"rees","Details":"some deets"},{"ID":2,"Name":"example","Owner":"human","Details":"a wonderful example"}]%
```

Delete a dog
```
dogs git:master ❯ curl -X DELETE \
  http://localhost:8080/dogs/2
Dog 2 deleted%
dogs git:master ❯ curl -X DELETE \
  http://localhost:8080/dogs/2
Dog 2 not found% 
```

Update a dog
```
dogs git:master ❯ curl -X PUT \
  http://localhost:8080/dogs/2 \
  -H 'content-type: application/json' \
  -d '{ "Name": "updated example" }'
{"ID":2,"Name":"updated example","Owner":"human","Details":"a wonderful example"}%
dogs git:master ❯ curl -X GET \
  http://localhost:8080/dogs/2
{"ID":2,"Name":"updated example","Owner":"human","Details":"a wonderful example"}% 
```
