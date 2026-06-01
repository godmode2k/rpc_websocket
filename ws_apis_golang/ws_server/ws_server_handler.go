
/* --------------------------------------------------------------
Project:    websocket test
Purpose:
Author:     Ho-Jung Kim (godmode2k@hotmail.com)
Date:       Since May 4, 2026
Filename:   ws_server_handler.go

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
Dependencies:
-------------------------------------------------------------- */
package ws_server



//! Header
// ---------------------------------------------------------------

import (
    //"fmt"
    "log"
    "time"
    //"math/big"
    //"encoding/json"

    //"runtime"
    //"regexp"

    // HTTP JSON-RPC Server
    //"net/http"
    //"github.com/gorilla/mux"
    //"github.com/gorilla/rpc"
    //"github.com/gorilla/rpc/json"

    // $ go get -u github.com/go-sql-driver/mysql
    //"database/sql"
    //_ "github.com/go-sql-driver/mysql"
    //_ "github.com/mattn/go-sqlite3"
)



//! Definition
// --------------------------------------------------------------------

//var test = ""
//var UPDATES_INTERVAL = int(1) // 1s
var UPDATES_INTERVAL = float64(0.2) // 200ms



//! Implementation
// --------------------------------------------------------------------
/*
func thread_run() {
    // To prevent this channel from blocking, size is set to 1.
    ca := make(chan string, 1)

    go func() {
        // This function must not touch the Context.
        // Do some long running operations...

        //json.Unmarshal( []byte(""), &response )
        ca <- ""
    }()

    select {
    case result := <-ca:
        return nil
    //case <- :
    //    // timeout
    //    return nil
    }
}
*/

//func go_thread_run(ctx context.Context, ch chan int) {
func go_thread_run(p *WSServer) {
    log.Println( "go_thread_run()" )

    for {
        p.WSServer_timer_now_reset()

        if p.Get_is_started() {
        } else {
        }

        if p.Get_is_closed() {
            log.Println( "go_thread_run(): closed..." )
            break
        }


        //log.Println( "go_thread_run(): running..." )





        //time.Sleep( time.Second * time.Duration(UPDATES_INTERVAL) )
        time.Sleep( time.Millisecond * 1000 * time.Duration(UPDATES_INTERVAL) )
    } // for ()


    /*
    for {
        select {
        case <-ctx.Done():
            log.Println( "()", "context: Done" )
            close( ch )
            gWG.Done()
            break
        case <-ch:
            log.Println( "()", "chan: ", <-ch )
            // ch: 0, 1, 2, ...
        }
    }
    */

}



// ---------------------------------------------------------------
