#!/bin/bash


# clean up
# remove ~/.cache/go-build
#go clean -cache -modcache

# Dependencies:
#go get -u github.com/go-sql-driver/mysql
#go get -u github.com/go-sql-driver/mysql
#go get -u github.com/mattn/go-sqlite3
#go get -u github.com/gorilla/mux
#go get -u github.com/gorilla/rpc
#go get -u github.com/gorilla/rpc/json
#go get -u github.com/rs/cors
#go get -u github.com/gorilla/websocket

go run rpc_server_main.go
