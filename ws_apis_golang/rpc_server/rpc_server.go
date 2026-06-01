/* --------------------------------------------------------------
Project:    websocket test
Purpose:
Author:     Ho-Jung Kim (godmode2k@hotmail.com)
Date:       Since May 4, 2026
Filename:   rpc_server.go

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
    //"math/big"
    //"encoding/json"

    "runtime"
    //"regexp"

    // HTTP JSON-RPC Server
    //"net/http"
    //"github.com/gorilla/mux"
    //"github.com/gorilla/rpc"
    //"github.com/gorilla/rpc/json"

    // $ go get -u github.com/go-sql-driver/mysql
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/mattn/go-sqlite3"
)



//! Definition
// --------------------------------------------------------------------

//var test = ""



//! Implementation
// --------------------------------------------------------------------

type RPCServerAPIs struct {
    Db *sql.DB
    DbMemory *sql.DB
}



// ---------------------------------------------------------------



// Source: https://stackoverflow.com/questions/25927660/how-to-get-the-current-function-name
func __trace() {
    pc := make([]uintptr, 10)  // at least 1 entry needed
    runtime.Callers(2, pc)
    f := runtime.FuncForPC(pc[0])
    file, line := f.FileLine(pc[0])
    log.Printf("%s:%d %s\n", file, line, f.Name())
}
