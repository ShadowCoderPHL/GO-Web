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