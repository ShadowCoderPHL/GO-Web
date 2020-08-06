# GO-Web
Go web try

#### #go run http-server.go

Basic authentication

#### #go run http-server-basic-authentication.go

Test with curl

#curl --user admin:admin [http://localhost:8080](http://localhost:8080/)/

#### #go run http-server-basic-routing.go

Should show Hello world

#curl -X GET -i http://localhost:8080 

Should show Login Page

#curl -X GET -i http://localhost:8080/login

 Should show Logout Page

#curl -X GET -i http://localhost:8080/logout

#### #go run http-session.go 

Go first to http://localhost:8080/home , without going first to http://localhost:8080/login you will receive "unauthorized message", to test logout session just go to http://localhost:8080/logout

----------------------------------------------------------------------------

## Rest API

#### #go run http-rest-get.go

With Curl http-rest-get.go

`#curl -X GET http://localhost:8080/employees`

`#curl -X GET http://localhost:8080/employee/1`

#### #go run http-rest-post.go

With Curl http-rest-post.go

`#curl -H "Content-Type: application/json" -X POST -d '{"Id":"3",  "firstName":"Aron", "lastName":"Villanueva" }' http://localhost:8080/employee/add`

#### #go run http-rest-put.go

With Curl http-rest-put.go, this also insert if id dont exist yet

`#curl -H "Content-Type: application/json" -X PUT -d '{"Id":"3",  "firstName":"Aron", "lastName":"Villanueva" }' http://localhost:8080/employee/update`

#### #go run http-rest-delete.go

With Curl http-rest-delete.go

`#curl -H "Content-Type: application/json" -X DELETE -d '{"Id":"3",  "firstName":"Aron", "lastName":"Villanueva" }' http://localhost:8080/employee/delete`