bin: server.go session.go session_mgr.go web.go cli.go client.go
	govendor build -o bin server.go session.go session_mgr.go web.go cli.go client.go
