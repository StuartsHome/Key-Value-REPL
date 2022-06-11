## How to run
```
go run main.go
```

You're able to provide an optional flag `name` to name the app; otherwise it defaults to
the name `web-app`, e.g.
```
go run main.go --name="web-app"
```



## Flow
```
backend-go> Welcome to web-app!
web-app> Type HELP for a list of commands.
web-app> START   
web-app> WRITE 3 3
web-app> WRITE 4 4
web-app> READALL    
Key: 3 Value: 3
Key: 4 Value: 4
web-app> COMMIT
web-app> READALL
web-app> ABORT
web-app> READALL
Key: 3 Value: 3
Key: 4 Value: 4
```