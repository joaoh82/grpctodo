**gRPC Client**

Client created main for testing the gRPC Server

**Build it**

```
env GO111MODULE=on go build client.go
```

**Run it**

List tasks
```
./client list
```

Add task
```
./client add "buy eggs" false
```

Help
```
help:
  client add <title> <done>
  client list
```