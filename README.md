# RPC, WebSocket Examples


Summary
----------
> RPC, WebSocket Examples </br>
>
> WORK IN-PROGRESS </br>
> This is a test version. so, USE THIS AT YOUR OWN RISK.


Environment
----------
> build all and tested on GNU/Linux

    GNU/Linux: Ubuntu 24.04_x64 LTS
    Go: go1.26.2 linux/amd64
    HTML, JavaScript: Apache2, websocket, axios


Run
----------
```sh
-----------------------------------------------------
Server
-----------------------------------------------------
(golang)
$ cd ws_apis_golang
$ go run rpc_server_main.go

 
-----------------------------------------------------
Clients
-----------------------------------------------------
(golang)
$ cd ws_apis_golang
$ go run rpc_client_main.go


(cURL)
$ cd ws_apis_golang
$ bash run_rpc_client_curl_test.sh

or

// room: create
$ curl
    -H "Content-Type: application/json"
    -d '{ "method":"RPCServerAPIs.JSONRPC_room_create",
        "params":[{"req": {"host": "0.0.0.0", "port": "5001", "n_max_players": 50,
        "datetime": "2026-01-01 10:00:00","prep_time_s": 30}}],
        "id":0}'
    http://127.0.0.1:5000/rpc

// room: close
$ curl
    -H "Content-Type: application/json"
    -d '{ "method":"RPCServerAPIs.JSONRPC_room_close",
        "params":[{"req": {"room_id": "", "port": ""}}],
        "id":0}'
    http://127.0.0.1:5000/rpc

// room: list
$ curl
    -H "Content-Type: application/json"
    -d '{ "method":"RPCServerAPIs.JSONRPC_room_list",
        "params":[{"req": {}}],
        "id":0}'
    http://127.0.0.1:5000/rpc


(Web)
(project home)/web/lobby.html
$ sudo ln -s (project home)/web/ /var/www/html/

http://127.0.0.1:8080/web/lobby.html


-----------------------------------------------------
CORS error
-----------------------------------------------------
(golang)
(rpc_server_main.go)
...
cors := rs_cors.New( rs_cors.Options {
    AllowedOrigins: []string {
        ...,
        "http://host:port",  // <-- ADD here...
    },
...


(golang)
(ws_server/ws_server.go)
...
t.ws_upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        origin := r.Header.Get("Origin")
        allowed_origins := []string {
            ...,
            "http://host:port",  // <-- ADD here...
        }


(apache2)
$ sudo vim /etc/apache2/conf-available/cors.conf
<IfModule mod_headers.c>
    Header always set Access-Control-Allow-Origin "*"
    Header always set Access-Control-Allow-Methods "POST, GET, OPTIONS, DELETE, PUT"
    Header always set Access-Control-Allow-Headers "x-requested-with, Content-Type, Authorization"
</IfModule>
```



