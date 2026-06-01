/* --------------------------------------------------------------
Project:    websocket test
Purpose:
Author:     Ho-Jung Kim (godmode2k@hotmail.com)
Date:       Since May 4, 2026
Filename:   rpc_client_main.go

Last modified:  May 4, 2026
License:

*
* Copyright (C) 2026 Ho-Jung Kim (godmode2k@hotmail.com)
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*
-----------------------------------------------------------------
Note:
-----------------------------------------------------------------
Reference:
 - https://pkg.go.dev/database/sql

Dependencies:
$ go get -u github.com/go-sql-driver/mysql


1. Build:
	$ go build rpc_client_main.go
    or
	$ go run rpc_client_main.go
-------------------------------------------------------------- */
package main



//! Header
// ---------------------------------------------------------------

import (
    "fmt"
    "log"

    //"net/rpc"
    "net/http"

    // HTTP JSON-RPC
    "bytes"
    gorilla_json "github.com/gorilla/rpc/json"

    "ws_apis_golang/types"
)



//! Definition
// --------------------------------------------------------------------

//var SERVER_ADDRESS = "127.0.0.1"
//var SERVER_PORT = "8544"
//var SERVER = SERVER_ADDRESS + ":" + SERVER_PORT
//var URL = "http://" + SERVER_ADDRESS + ":" + SERVER_PORT
//var HTTP_RPC_SERVER_HOST_PORT = ":4900" // Internal
var HTTP_JSONRPC_SERVER_HOST_PORT = ":5000" // External



//! Implementation
// --------------------------------------------------------------------

func json_rpc_request(api string, args interface{}, url string) types.RPC_response_st {
    var result types.RPC_response_st
    message, err := gorilla_json.EncodeClientRequest( api, args )

    log.Println("json_rpc_request(): ", args.(*types.RPC_params_st) )

    if err != nil {
        log.Fatalf("json_rpc_request(): %s", err)
    }
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
    if err != nil {
        log.Fatalf("json_rpc_request(): %s", err)
    }
    req.Header.Set("Content-Type", "application/json")
    client_jsonrpc := new(http.Client)
    resp, err := client_jsonrpc.Do(req)
    if err != nil {
        log.Fatalf("json_rpc_request(): http.Client.Do(): Error: URL = %s, %s", url, err)
    }
    defer resp.Body.Close()

    log.Println( "response = ", resp.Body )
    err = gorilla_json.DecodeClientResponse(resp.Body, &result)
    if err != nil {
        log.Fatalf("json_rpc_request(): DecodeClientResponse(): Error: %s", err)
    }

    return result
}

func main() {
    var result types.RPC_response_st
    url := "http://localhost" + HTTP_JSONRPC_SERVER_HOST_PORT + "/rpc"


    api := "RPCServerAPIs.JSONRPC_room_create"
    fmt.Println( api )
    //args := []*types.RPC_params_st{
    args :=
        &types.RPC_params_st{
            Req: map[string]interface{} {
                "host": "0.0.0.0", "port": "5001", "n_max_players": 50,
                "datetime": "2026-01-01 10:00:00","prep_time_s": 30,
            },
        //},
        }
    //}
    result = json_rpc_request( api, args, url )
    fmt.Println( "result = \n", result )
    fmt.Println( "\n\n" )


    api = "RPCServerAPIs.JSONRPC_room_close"
    fmt.Println( api )
    //args = []*types.RPC_params_st{
    args = 
        &types.RPC_params_st{
            Req: map[string]interface{} {
                "room_id": "",
                "port": "",
            },
        //},
        }
    //}
    result = json_rpc_request( api, args, url )
    fmt.Println( "result = \n", result )
    fmt.Println( "\n\n" )


    api = "RPCServerAPIs.JSONRPC_room_list"
    fmt.Println( api )
    //args = []*types.RPC_params_st{
    args = 
        &types.RPC_params_st{
            Req: map[string]interface{} {},
        //},
        }
    //}
    result = json_rpc_request( api, args, url )
    fmt.Println( "result = \n", result )
    fmt.Println( "\n\n" )
}



