/* --------------------------------------------------------------
Project:    websocket test
Purpose:
Author:     Ho-Jung Kim (godmode2k@hotmail.com)
Date:       Since May 4, 2026
Filename:   rpc_server_main.go

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
$ go get -u github.com/go-sql-driver/mysql
$ go get -u github.com/mattn/go-sqlite3
$ go get -u github.com/gorilla/mux
$ go get -u github.com/gorilla/rpc
$ go get -u github.com/gorilla/rpc/json


1. Build:
	$ go build rpc_server_main.go
    or
	$ go run rpc_server_main.go
-------------------------------------------------------------- */
package main



//! Header
// ---------------------------------------------------------------

import (
    //"fmt"
    "log"
    //"time"

    // sync.WaitGroup
    // sync.Mutex
    "sync"
    // context.Context: context.WithCancel(context.Background())
    //"context"

    // test
    //"math/rand"

    // HTTP RPC
    //"net"
    "net/http"
    //"net/rpc"

    // HTTP JSON-RPC
    gorilla_mux "github.com/gorilla/mux"
    gorilla_rpc "github.com/gorilla/rpc"
    gorilla_json "github.com/gorilla/rpc/json"
    //gorilla_handlers "github.com/gorilla/handlers"
    rs_cors "github.com/rs/cors"

    // $ go get -u github.com/go-sql-driver/mysql
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/mattn/go-sqlite3"

    //"ws_apis_golang/types"
    "ws_apis_golang/rpc_server"

    //"reflect"
)



//! Definition
// --------------------------------------------------------------------

var SERVER_ADDRESS = "127.0.0.1"
var SERVER_PORT = "5000"
var SERVER = SERVER_ADDRESS + ":" + SERVER_PORT
var URL = "http://" + SERVER_ADDRESS + ":" + SERVER_PORT
var DB_SERVER_ADDRESS = "127.0.0.1:3306"
var DB_NAME = "db_test"
var DB_LOGIN_USERNAME = "root"
var DB_LOGIN_PASSWORD = "mysql"
var gDB *sql.DB
var DB_MEMORY_NAME = "memdb_test"
var gDBMemory *sql.DB


// HTTP RPC Server
//var HTTP_RPC_SERVER_HOST_PORT = ":4999" // Internal
var HTTP_JSONRPC_SERVER_HOST_PORT = ":5000" // External
//var g_rpc_server_apis = new( rpc_server.RPCServerAPIs)
var g_rpc_server_apis = &rpc_server.RPCServerAPIs{}
var UPDATES_INTERVAL = int(30) // 30 seconds
var gWG sync.WaitGroup
var gChan = make( chan uint8, 1 )
var _E_CHAN__DONE = uint8(0)
var _E_CHAN__CANCEL = uint8(1)


//var gMutex sync.Mutex
// mutex.Lock()
// defer mutex.Unlock()




//! Implementation
// --------------------------------------------------------------------

func __init_db() bool {
    log.Println( "main: __init_db(): initialize..." )

    /*
    db, err := sql.Open( "mysql",
        DB_LOGIN_USERNAME + ":" + DB_LOGIN_PASSWORD + "@tcp(" + DB_SERVER_ADDRESS + ")/" + DB_NAME )

    if err != nil {
        panic( err.Error() )
    }

    gDB = db



    // In-memory DB (SQLite)
    //db_sqlite, err := sql.Open("sqlite3", "./localdb_sqlite.db")
    //db_sqlite, err := sql.Open("sqlite3", "file::memory:?mode=memory&cache=shared")
    //db_sqlite_filename := randomString(16) // func creates random string
    db_sqlite_filename := "localdb_sqlite3.db"
    db_sqlite, err := sql.Open( "sqlite3", fmt.Sprintf("file:%s?mode=memory&cache=shared", db_sqlite_filename) )
    if err != nil {
        //log.Fatal(err)
        panic( err.Error() )
    }


    // memory_txid
    query := "CREATE TABLE " + "<table-name>" + " (idx integer not null primary key autoincrement,"
    query += " txid_json text);"
    _, err = db_sqlite.Exec( query )
    if err != nil {
        //log.Printf( "Error: %q: %s\n", err, query )
        fmt.Printf( "Error: %q: %s\n", err, query )
        panic( err.Error() )
    }

    gDBMemory = db_sqlite

    g_localdb.Db = gDB
    g_localdb.DbMemory = gDBMemory
    */

    g_rpc_server_apis.Db = gDB
    g_rpc_server_apis.DbMemory = gDBMemory

    return true
}

/*
func __test_db() {
    var TABLE_NAME = "txid"
    //var TABLE_NAME_ERC1155 = "txid_erc1155"

    var query_str = fmt.Sprintf(
        "INSERT INTO %s VALUES (0," +
        "'%s', '%s', '%s', '%s', '%s'" +
        ")",
        TABLE_NAME,
        "", "", "", "", "" )

    fmt.Println( "query = ", query_str )

    //result, err := gDB.Query( "INSERT INTO test VALUES ( 2, 'TEST' )" )
    result, err := gDB.Query( query_str )

    if err != nil {
        panic( err.Error() )
    }

    fmt.Println( "result = ", result )
}
*/

/*
func db_insert_txns(_type uint8, _data *types.Fetch_transactions_st) {
    var TABLE_NAME = ""
    var query_str = ""

    if _type == 0 {
        TABLE_NAME = "txid"
        query_str = fmt.Sprintf(
            "INSERT INTO %s VALUES (0," +
            "'%s', '%s', '%s', '%s', '%s'" +
            ")",

            TABLE_NAME,

            "", "", "", "", ""
            )
    } else {
    }

    fmt.Println( "query = ", query_str )

    //result, err := gDB.Query( "INSERT INTO test VALUES ( 2, 'TEST' )" )
    result, err := gDB.Query( query_str )

    if err != nil {
        panic( err.Error() )
    }

    fmt.Println( "result = ", result )
}
*/



// --------------------------------------------------------------------



//func run_worker_cache(ctx context.Context, ch chan int) {
//    for {
//        select {
//        case <-ctx.Done():
//            fmt.Println( "run_worker_cache()", "context: Done" )
//            close( ch )
//            gWG.Done()
//            break
//        case <-ch:
//            fmt.Println( "run_worker_cache()", "chan: ", <-ch )
//            // ch: 0, 1, 2, ...
//        }
//    }
//
//    fmt.Println( "run_worker_cache()", "finished..." )
//}

// Goroutine
func run_worker_cache() {
    //if g_localdb == nil {
    if g_rpc_server_apis == nil {
        panic( "DB object == NULL" )
    }

    log.Println( "run_worker_cache()", "Starting caching..." )


    // SEE: var g_localdb = new( rpc_server.LocalDB )
    //g_localdb.Db_memory_update_txns_all_mixed()

    /*
    for {
        log.Println( "run_worker_cache()", "Updating aaa..." )
        // aaa()

        log.Println( "run_worker_cache()", "Updating bbb..." )
        // bbb()

        log.Println( "run_worker_cache()", "Updating ccc..." )
        // ccc()

        log.Println()

        //time.Sleep( time.Second * time.Duration(UPDATES_INTERVAL) )
        time.Sleep( time.Millisecond * 1000 * time.Duration(UPDATES_INTERVAL) )
    } // for ()
    */



    /*
    var is_done = false
    for {
        select {
        case <-gChan:
            switch <-gChan {
            case _E_CHAN__CANCEL:
                log.Println( "CHAN: CANCEL:" )
                is_done = true
                break
            case _E_CHAN__DONE:
                log.Println( "CHAN: DONE:" )
                is_done = true
                break
            }

            if  is_done == true {
                break
            }
        default:
            log.Println( "CHAN: Waiting..." )
        }

        if  is_done == true {
            break
        }

    }
    */

    log.Println( "run_worker_cache()", "finished..." )

    gWG.Done()
}

// Goroutine
/*
func run_http_rpc_server() {
    //if g_localdb == nil {
    if g_rpc_server_apis == nil {
        panic( "DB object == NULL" )
    }

    log.Println( "Starting HTTP RPC Server...: ", HTTP_RPC_SERVER_HOST_PORT )


    // SEE: var g_localdb = new( rpc_server.LocalDB )

    //rpc.Register( g_localdb )
    rpc.Register( g_rpc_server_apis ) # do not share 'g_rpc_server_apis'
    rpc.HandleHTTP()

    l, e := net.Listen( "tcp", HTTP_RPC_SERVER_HOST_PORT )
    if e != nil {
        log.Fatal( "listen error:", e )
    }

    //go http.Serve( l, nil )
    http.Serve( l, nil )
}
*/

// Goroutine
func run_http_jsonrpc_server() {
    //if g_localdb == nil {
    if g_rpc_server_apis == nil {
        panic( "DB object == NULL" )
    }

    log.Println( "Starting HTTP JSON-RPC Server...: ", HTTP_JSONRPC_SERVER_HOST_PORT )


    // SEE: var g_localdb = new( rpc_server.LocalDB )

    /*
    rpc.Register( g_localdb )
    rpc.HandleHTTP()

    l, e := net.Listen( "tcp", ":1234" )
    if e != nil {
        log.Fatal( "listen error:", e )
    }

    //go http.Serve( l, nil )
    http.Serve( l, nil )
    */

    _rpc := gorilla_rpc.NewServer()
    _rpc.RegisterCodec( gorilla_json.NewCodec(), "application/json" )
    _rpc.RegisterCodec( gorilla_json.NewCodec(), "application/json;charset==UTF-8" )
    //_rpc.RegisterService( g_localdb, "" )
    _rpc.RegisterService( g_rpc_server_apis , "" )
    _router := gorilla_mux.NewRouter()
    _router.Handle( "/rpc",  _rpc )

    //http.Header.Set( "Connection", "close" )

    // Without CORS
    //http.ListenAndServe( HTTP_JSONRPC_SERVER_HOST_PORT,  _router )


    {
        // CORS
        // SEE: https://pkg.go.dev/github.com/gorilla/handlers#CORS
        // Default
        //
        //http.ListenAndServe( HTTP_JSONRPC_SERVER_HOST_PORT, gorilla_handlers.CORS()(_router) )


        // Canonical Host
        // CanonicalHost is HTTP middleware that re-directs requests to the canonical domain.
        // It accepts a domain and a status code (e.g. 301 or 302) and re-directs clients to this domain.
        // Note: If the provided domain is considered invalid by url.Parse or 
        // otherwise returns an empty scheme or host, clients are not re-directed.
        //
        //_canonical_host := gorilla_handlers.CanonicalHost( "http://www.example.org", 302 )
        //http.ListenAndServe( HTTP_JSONRPC_SERVER_HOST_PORT, _canonical_host(_router) )


        /*
        // Source: https://stackoverflow.com/questions/40985920/making-golang-gorilla-cors-handler-work
        //headersOk := gorilla_handlers.AllowedHeaders([]string{"X-Requested-With"})
        //originsOk := gorilla_handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
        //
        // (works)
        headersOk := gorilla_handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
        originsOk := gorilla_handlers.AllowedOrigins([]string{"*"})
        methodsOk := gorilla_handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
        //
        http.ListenAndServe( HTTP_JSONRPC_SERVER_HOST_PORT,
                            gorilla_handlers.CORS(originsOk, headersOk, methodsOk)(_router) )
        */


        ///*
        // (works)
        // RS CORS
        // Source: https://stackoverflow.com/questions/40985920/making-golang-gorilla-cors-handler-work
        cors := rs_cors.New( rs_cors.Options {
            AllowedOrigins: []string {
                // JSON RPCServer
                "http://127.0.0.1:5000/rpc",
                "http://localhost:5000/rpc",

                // Browser (127.0.0.1:8080) <-> Request 127.0.0.1:5000 (VirtualBox Network Port Forward) <-> 10.0.x.x:5000
                "http://127.0.0.1:8080",
                "http://localhost:8080",
            },
            //AllowCredentials: true,
        })
        cors_handler := cors.Handler( _router )
        http.ListenAndServe( HTTP_JSONRPC_SERVER_HOST_PORT, cors_handler )
        //*/
    }

}

// Goroutine
/*
func run_aaa_main() {
    //rpc_server.Test_aaa__main_func()
}

// Goroutine
func run_bbb__main() {
    //rpc_server.Test_bbb__main_func()
}

// Goroutine
func run_ccc__main() {
    //rpc_server.Test_ccc__main_func()
}

func run_init_db() {
    //rpc_server.Init_db_func()
}

func run_release_db() {
    //rpc_server.Release_db_func()
}
*/



// --------------------------------------------------------------------



func main() {
    log.Println( "HOST: " + URL )

    /*
    // Initialize Database
    if __init_db() != true {
        log.Println( "main(): DB init...", "false" )
        return
    }

    // Release Database
    if gDB != nil {
        defer gDB.Close()
    }
    if gDBMemory != nil {
        defer gDBMemory.Close()
    }

    //__test_db()
    */


    //! FIXME: logs
    // -------------------------------------
    // Initialize Database instance for ...
    //run_init_db()

    var GOROUTINE_TOTAL = 4
    gWG.Add( GOROUTINE_TOTAL )
    //ctx, cancel := context.WithCancel( context.Background() )
    //ctx = context.WithValue( ctx, "key", "val" )

    // Start Data caching
    //go run_worker_cache()

    // Start HTTP RPC Server
    //go run_http_rpc_server()

    // Start HTTP JSON-RPC Server
    go run_http_jsonrpc_server()

    gWG.Wait()



    // Release Database instance for {transactions fetcher, updates balances, ...}
    //run_release_db()
}
