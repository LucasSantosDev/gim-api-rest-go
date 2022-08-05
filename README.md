##### To run container

```console
sudo docker-compose exec -d

sudo docker-compose exec postgres bash

> hostname -i // TO GET IP OF CONTAINER
```

---

##### To execute project GO in other path

```console
go mod init gim-api-rest-go

or

go mod init github.com/LucasSantosDev/gim-api-rest-go
```

---

##### To build and execute

```console
go run main.go
```

---

##### To test list students

```console
curl -H "Content-Type: application/json" -X GET http://localhost:8080/v1/students/
```

---

##### To test show specific student

```console
curl -H "Content-Type: application/json" -X GET http://localhost:8080/v1/students/1
```

---

##### To test create student

```console
curl -H "Content-Type: application/json" -X POST -d '{"name":"Lucas","cpf":"123.123.123-12","rg":"12.123.123-x"}' http://localhost:8080/v1/students/
```

---

##### To test delete student

```console
curl -H "Content-Type: application/json" -X DELETE http://localhost:8080/v1/students/2
```

---

##### To test update student

```console
curl -H "Content-Type: application/json" -X PATCH -d '{"name":"Teste atualizado","cpf":"xxx.xxx.xxx-xx"}' http://localhost:8080/v1/students/3
```

---

##### To test filter student by CPF

```console
curl -H "Content-Type: application/json" -X GET http://localhost:8080/v1/students/filter/123.123.123-12
```

---
