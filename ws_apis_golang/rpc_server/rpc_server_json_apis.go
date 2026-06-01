/* --------------------------------------------------------------
Project:    websocket test
Purpose:
Author:     Ho-Jung Kim (godmode2k@hotmail.com)
Date:       Since May 4, 2026
Filename:   rpc_server_json_apis.go

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
 - https://golang.org/pkg/net/rpc/
 - https://pkg.go.dev/database/sql
 - https://pkg.go.dev/github.com/mattn/go-sqlite3

Dependencies:
 - $ go get -u github.com/go-sql-driver/mysql
 - $ go get github.com/mattn/go-sqlite3
-------------------------------------------------------------- */
package rpc_server



//! Header
// ---------------------------------------------------------------

import (
    //"fmt"
    "log"
    //"encoding/json"
    //"math/big"

    //"runtime"
    //"regexp"

    "net/http"

    // UUID
    "github.com/google/uuid"

    // HTTP JSON-RPC Server
    //"github.com/gorilla/mux"
    //"github.com/gorilla/rpc"
    //"github.com/gorilla/rpc/json"

    // $ go get -u github.com/go-sql-driver/mysql
    //"database/sql"
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/mattn/go-sqlite3"

    "ws_apis_golang/types"
    "ws_apis_golang/ws_server"
)



//! Definition
// --------------------------------------------------------------------
//var g_ws_server_list = new( []ws_server.WSServer )
var g_ws_server_list = []*ws_server.WSServer{}

func ws_server_list_init() {
    g_ws_server_list = []*ws_server.WSServer{}
}

func ws_server_list_add(server *ws_server.WSServer) {
    g_ws_server_list = append( g_ws_server_list, server )
}

func ws_server_list_add_with_create() *ws_server.WSServer {
    server := &ws_server.WSServer{}
    g_ws_server_list = append( g_ws_server_list, server )
    return server
}

func ws_server_list() []*ws_server.WSServer {
    return g_ws_server_list
}

func ws_server_list_len() int {
    return len(g_ws_server_list)
}

func ws_server_list_get(room_id string) *ws_server.WSServer {
    if ws_server_list_len() > 0 {
        for _, v := range g_ws_server_list {
            if v == nil { continue }
            if v.Get_room_id() == room_id {
                return v
            }
        }
    }

    return nil
}

func ws_server_list_checks_dup(port string) bool {
    ret := false

    if ws_server_list_len() > 0 {
        for _, v := range g_ws_server_list {
            if v == nil { continue }
            if v.Get_host_port()[1] == port {
                ret = true
                break
            }
        }
    }

    return ret
}

func ws_server_list_del(room_id string) {
    idx := int(0)

    if ws_server_list_len() > 0 {
        for _, v := range g_ws_server_list {
            if v == nil { continue }
            if v.Get_room_id() == room_id {
                g_ws_server_list[idx] = nil
                g_ws_server_list = append( g_ws_server_list[:idx], g_ws_server_list[idx+1:]... )
            }
        }
    }
}



//! Implementation
// --------------------------------------------------------------------


// HTTP JSON-RPC Server: for frontend

func (t *RPCServerAPIs) JSONRPC_test(
    request *http.Request,
    rpc_params *types.RPC_params_st,
    //response *string,
    response *types.RPC_response_st,
) error {
    //*response = fmt.Sprintf( "%d", rpc_params.Req )
    //result, _ := json.Marshal( _result )
    //*response = string(result)
    return nil
}

/*
func (t *RPCServerAPIs) JSONRPC_get_(
    request *http.Request,
    rpc_params *types.RPC_params_st,
    response *string,
) error {
    log.Println( "JSONRPC_get_()" )

    var _result []types.Test_st
    //var result_str string
    //OFFSET := uint(0)

    //err := t.Db_memory_select_txns_all_mixed( OFFSET, &result_str )
    var err []string = nil
    //err := t.Db_memory_select_txns_all_mixed( OFFSET, &_result )
    if err != nil {
        log.Fatal( "JSONRPC_get_(): Error: ", err )
    }

    //*response = result_str


    result, err_marshal := json.Marshal( _result )
    if err_marshal != nil {
        panic( err_marshal.Error() )
    }


    //var json_map_arr = make( map[string]interface{} )
    //json_map_arr["txid"] = _result
    //
    //result, err_marshal := json.Marshal( json_map_arr )
    //if err_marshal != nil {
    //    panic( err_marshal.Error() )
    //}


    *response = string(result)
    return nil
}
*/


func (t *RPCServerAPIs) JSONRPC_room_create(
    r *http.Request,
    rpc_params *types.RPC_params_st,
    //response *string,
    response *types.RPC_response_st,
) error {
    log.Println( "JSONRPC_room_create()" )

    _result := []map[string]interface{} {}

    /*
    {
        "host": "0.0.0.0",
        "port": "5001",
        "timer": {"start_datetime": "2026-01-01 10:10:10", "prep_time_s": "30"},
        "n_max_player": "50",
    }
    */

    log.Println( "JSONRPC_room_create(): params = ", rpc_params )

    // types.RPC_params_st: Req interface{}
    // type assertion required
    params, ret := rpc_params.Req.(map[string]interface{})
    if !ret {
        log.Println( "JSONRPC_room_create(): params error..." )
        return nil
    }

    // test
    //room_id := "room_test"
    //host := "0.0.0.0"
    //port := "5001"
    //n_max_players := 50
    //timer := map[string]interface{} {"datetime": "2026-01-01 10:00:00", "prep_time_s": "30" }

    room_id := uuid.New().String()
    host := ""
    port := ""
    n_max_players := int(0)
    timer := map[string]interface{} {}

    if v, has := params["host"]; has { host = v.(string) } else { return nil }
    if v, has := params["port"]; has { port = v.(string) } else { return nil }
    if v, has := params["n_max_players"]; has { n_max_players = int(v.(float64)) } else { return nil }
    if v, has := params["datetime"]; has { timer["datetime"] = v.(string) } else { return nil }
    if v, has := params["prep_time_s"]; has { timer["prep_time_s"] = int(v.(float64)) } else { return nil }

    // checks port duplicated
    if !ws_server_list_checks_dup(port) {
        p := ws_server_list_add_with_create()
        log.Println( p )
        log.Println( g_ws_server_list )

        go p.WSServer_init(
            room_id,
            host, port,
            n_max_players,
            timer,
        )
    }

    //*response = fmt.Sprintf( "%d", rpc_params.Req )
    //result, _ := json.Marshal( _result )
    //*response = string(result)
    response.Result = _result
    return nil
}

func (t *RPCServerAPIs) JSONRPC_room_close(
    r *http.Request,
    rpc_params *types.RPC_params_st,
    //response *string,
    response *types.RPC_response_st,
) error {
    log.Println( "JSONRPC_room_close()" )

    _result := []map[string]interface{} {}

    log.Println( "JSONRPC_room_close(): params = ", rpc_params )
    params, ret := rpc_params.Req.(map[string]interface{})
    if !ret {
        log.Println( "JSONRPC_room_close(): params error..." )
        return nil
    }

    room_id := ""
    port := ""

    if v, has := params["room_id"]; has { room_id = v.(string) } else { return nil }
    if v, has := params["port"]; has { port = v.(string) } else { return nil }

    if ws_server_list_len() > 0 {
        for _, p := range ws_server_list() {
            if p == nil { continue }
            if p.Get_room_id() == room_id && p.Get_host_port()[1] == port {
                log.Println( "JSONRPC_room_close(): found..." )
                p.WSServer_release()
                break
            }
        }

        ws_server_list_del( room_id )
    }


    //*response = fmt.Sprintf( "%d", rpc_params.Req )
    //result, _ := json.Marshal( _result )
    //*response = string(result)
    response.Result = _result
    return nil
}

func (t *RPCServerAPIs) JSONRPC_room_list(
    r *http.Request,
    rpc_params *types.RPC_params_st,
    //response *string,
    response *types.RPC_response_st,
) error {
    log.Println( "JSONRPC_room_list()" )

    _result := []map[string]interface{} {}

    for _, p := range ws_server_list() {
        res := map[string]interface{} {
            "room_id": p.Get_room_id(),
            "host": p.Get_host_port()[0],
            "port": p.Get_host_port()[1],
            "n_max_players": p.Get_n_max_players(),
            "n_players": p.Get_n_players(),
            "timer": map[string]interface{} {
                "datetime": p.Get_timer().Initial_start_datetime,
                "prep_time_s": p.Get_timer().Initial_start_prep_time_s,
            },
        }
        _result = append( _result, res )
        log.Println( "JSONRPC_room_list(): info = ", res )
    }


    //*response = fmt.Sprintf( "%d", rpc_params.Req )
    //result, _ := json.Marshal( _result )
    //*response = string(result)
    response.Result = _result
    return nil
}




