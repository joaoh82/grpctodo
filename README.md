**Shell Todo Application**

**Requirements:**

* gRPC API

**Methods:**

AddTask

ListTasks

* RESTful API 

**Enpoints:**

  GET /listtasks

  Response: (Example from local test)
 ```
"{\"Tasks\":[{\"title\":\"get eggs\"},{\"title\":\"pay rent\"},{\"title\":\"buy chicken\"},{\"title\":\"teste\"},{\"title\":\"teste\"},{},{},{\"title\":\"test task3\"},{\"title\":\"buy-tv\"},{\"title\":\"buy eggs\"}]}"
 ```

  POST /addtasks
  
  Body:
  ```
  {
	"title": "test task3",
	"done": false
}   
```

* Some type of persistance
* RESTful API should be just a proxy to the the gRPC API.

**Dependencies:**
* Go - https://golang.org/doc/install
* PostgreSQL DB
* Protoc Setup - In order to perform code generation, you will need to install protoc on your computer. - 
On Mac 
```
brew install protobuf
```
```
go get -u google.golang.org/grpc
```
Then you need to install the protobuf dependency for Go by going to your command line again the typing
```
go get -u github.com/golang/protobuf/protoc-gen-go
```

Command to generate go file based on *proto file
```
protoc pg/todo.proto --go_out=plugins=grpc:.
```

**TODO:**

* Create DB/Repository Package and abstract DB action out of API layer
* Write Tests for gRPC, RESTful and Repository
* Implement proper logging
* Implement proper error handling
* Authentication via SSL or Token