/* --------------------------------------------------------------
Project:    websocket test
Purpose:
Author:     Ho-Jung Kim (godmode2k@hotmail.com)
Date:       Since May 4, 2026
Filename:   ws_server.go

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
	//"flag"
    //"fmt"
    "log"
    "time"
    "strings"
    //"strconv"
    "encoding/json"
    //"math/big"

    //"runtime"
    //"regexp"

    "net/http"

    // not work
    // CORS
    //"github.com/rs/cors"

    // UUID
    "github.com/google/uuid"

    // HTTP JSON-RPC Server
    //"github.com/gorilla/mux"
    //"github.com/gorilla/rpc"
    //"github.com/gorilla/rpc/json"

    // WebSocket
	"github.com/gorilla/websocket"

    // $ go get -u github.com/go-sql-driver/mysql
    //"database/sql"
    //_ "github.com/go-sql-driver/mysql"
    //_ "github.com/mattn/go-sqlite3"
)



//! Definition
// --------------------------------------------------------------------

//var test = ""

type WSServer struct {
    //Db *sql.DB
    //DbMemory *sql.DB


    // websocket
    ws_sids map[string]*websocket.Conn
    //ws_event_handler map[string]func(map[string]interface{}, *string) error
    ws_event_handler map[string]func(map[string]interface{})
    ws_addr string // "host:port"
    ws_upgrader websocket.Upgrader
    ws_is_started bool
    ws_is_closed bool

    title string
    room_id string
    // { "<user_id>": "<cid>" }
    //sid_list map[string]string
    // { "user_id": "<user_id>", "cid": "<cid>", "action": { "submit": "", "history": {} } }
    //players_list map[string]interface{}
    players_list []*wsServer_player
    n_max_players int
    n_players int
    timer wsServer_timer
}

type wsServer_timer struct {
    Now_utc_datetime string
    Now_datetime string

    Initial_start_timestamp int64
    Initial_start_datetime string
    Initial_start_prep_time_s int

    Interval_s int

    utc_start_prep_time_s int
    end_prep_time_s int

    // UTC
    utc_start_timeatamp int64
    utc_start_datetime string
    utc_end_timeatamp int64
    utc_end_datetime string

    // Local
    local_start_timeatamp int64
    local_start_datetime string
    local_end_timeatamp int64
    local_end_datetime string
}

//type WSServer_sid_list struct {
//    sid string
//    username string
//}

// { "user_id": "<user_id>", "username": "<username>", "cid": "<cid>",
//   "action": { "submit": "", "history": {} } }
type wsServer_player struct {
    // sid: websocket.Conn (socket id)
    // cid: string uuid for sid

    user_id string
    username string
    cid string
    reconnected_count int

    // action
    //{
    //  "submit": "<player answer>",
    //  "history": {"question_number": {"answser": "<player answer>", "correct": true} }
    //}
    //
    //{ "action": {"submit": "", "history": {"0": {"a": "", "correct": true}, ...} } }
    action map[string]interface{}
}

type wsServer_quiz_list struct {
    //{ "quiz_num":
    //  {
    //      "q": "question...",
    //      "a": "answer",
    //      "t": type(int): 0: text, 1: image,
    //      "desc": "descriptions",
    //      "dur_s": duration_seconds(int)
    //  }
    //}
    //{ "quiz": { "0": {"q": "question...", "a": "answer", "t": int(0), "desc": "descriptions", "dur_s": int(0)} } }
    quiz_list map[string]interface{}
}



//! Implementation
// --------------------------------------------------------------------

func (t *WSServer) wsServer_test(response *string) error {
    log.Println( "wsServer_test()" )

    //*response = fmt.Sprintf( "" )
    //result, _ := json.Marshal( _result )
    //*response = string(result)
    return nil
}


func (t *WSServer) Get_host_port() []string {
    return strings.Split( t.ws_addr, ":" )
}

func (t *WSServer) Get_room_id() string {
    return t.room_id
}

func (t *WSServer) Get_n_max_players() int {
    return t.n_max_players
}

func (t *WSServer) Get_n_players() int {
    //t.n_players = len(t.sid_list)
    t.n_players = len(t.players_list)
    log.Println( "n_players = ", t.n_players )
    return t.n_players
}

func (t *WSServer) Get_is_started() bool {
    return t.ws_is_started
}

func (t *WSServer) Get_is_closed() bool {
    return t.ws_is_closed
}

func (t *WSServer) Get_timer() wsServer_timer {
    return t.timer
}

func (t *WSServer) player_new() *wsServer_player {
    //{
    //  "submit": "<player answer>",
    //  "history": {"question_number": {"answser": "<player answer>", "correct": true} }
    //}
    //
    //{ "action": {"submit": "", "history": {"0": {"a": "", "correct": true}, ...} } }

    player := &wsServer_player {
        user_id: "",
        username: "",
        cid: "",
        action: map[string]interface{} {
            "action": map[string]interface{} {
                "submit": "", //map[string]interface{} {},
                "history": map[string]interface{} {
                    //"0": map[string]interface{} { "answer": "", "correct": true }
                },
            },
        },
    }
    t.players_list = append( t.players_list, player )
    return player
}

func (t *WSServer) player_get(cid string) *wsServer_player {
    if len(t.players_list) > 0 {
        for _, v := range t.players_list {
            if v == nil { continue }
            if v.cid == cid {
                return v
            }
        }
    }

    return nil
}

func (t *WSServer) player_del(cid string) {
    idx := int(0)

    if len(t.players_list) > 0 {
        for _, v := range t.players_list {
            if v == nil { continue }
            if v.cid == cid {
                t.players_list[idx] = nil
                t.players_list = append( t.players_list[:idx], t.players_list[idx+1:]... )
            }
        }
    }
}

func (t *WSServer) quiz_list_init_new() *wsServer_quiz_list {
    val := map[string]interface{} {
            "quiz": map[string]interface{} {
                "0": map[string]interface{} { "q": "", "a": "", "t": 0, "desc": "", "dur_s": 0, },
                "1": map[string]interface{} { "q": "", "a": "", "t": 0, "desc": "", "dur_s": 0, },
                "2": map[string]interface{} { "q": "", "a": "", "t": 0, "desc": "", "dur_s": 0, },
                "3": map[string]interface{} { "q": "", "a": "", "t": 0, "desc": "", "dur_s": 0, },
                "4": map[string]interface{} { "q": "", "a": "", "t": 0, "desc": "", "dur_s": 0, },
            },
        }

    return t.quiz_list_init_new_ex(&val)
}

func (t *WSServer) quiz_list_init_new_ex(val *map[string]interface{}) *wsServer_quiz_list {
    /*
    return &WSServer_quiz_list {
        quiz_list: map[string]interface{} {
            "quiz": map[string]interface{} {
                "0": map[string]interface{} { "q": "", "a": "", "t": 0, "desc": "", "dur_s": 0, },
                "1": map[string]interface{} { "q": "", "a": "", "t": 0, "desc": "", "dur_s": 0, },
                "2": map[string]interface{} { "q": "", "a": "", "t": 0, "desc": "", "dur_s": 0, },
                "3": map[string]interface{} { "q": "", "a": "", "t": 0, "desc": "", "dur_s": 0, },
                "4": map[string]interface{} { "q": "", "a": "", "t": 0, "desc": "", "dur_s": 0, },
            },
        },
    }
    */

    return &wsServer_quiz_list { quiz_list: *val, }
}

func (t *WSServer) WSServer_init(
    room_id string,
    host string,
    port string,
    n_max_players int,
    timer map[string]interface{},
) {
    log.Println( "WSServer_init()" )

    //var ws_addr = flag.String("addr", "localhost:8080", "http service address")

    t.ws_sids = make(map[string]*websocket.Conn)
    //t.ws_event_handler = make(map[string]func(map[string]interface{}, *string) error)
    t.ws_event_handler = make(map[string]func(map[string]interface{}))

    t.ws_addr = host + ":" + port
    t.ws_is_started = true
    t.ws_is_closed = false

    t.room_id = room_id
    t.n_max_players = n_max_players
    t.n_players = n_max_players

    t.wsServer_timer_reset( true, timer )

    log.Println( "WSServer_init(): Starting ws: ", t.ws_addr )
    log.Println( "WSServer_init(): room_id = ", t.room_id )
    log.Println( "WSServer_init(): n_max_players = ", t.n_max_players )


    go go_thread_run( t )




    //t.ws_upgrader = websocket.Upgrader{} // use default options
    //
    // with CORS
    t.ws_upgrader = websocket.Upgrader{
        CheckOrigin: func(r *http.Request) bool {
            origin := r.Header.Get("Origin")
            allowed_origins := []string {
                "",
                "http://localhost:8080",
                "http://127.0.0.1:8080",
            }

            log.Println( "origin = ", origin )
            for _, allowed := range allowed_origins {
                if origin == allowed {
                    return true
                }
            }

            return false
        },
    }

	http.HandleFunc( "/test", t.wsServer_ws_handler )
    //
    // not work
    //
    // with CORS
    //mux := http.NewServeMux()
	//mux.HandleFunc( "/test", t.wsServer_ws_handler )

    /*
    // example
    http.HandleFunc( "/events", func(w http.ResponseWriter, r *http.Request) {
        ctx := r.Context()
        fmt.Println("[/events] connected...")
        fmt.Println("[/events] start...")
        //
        // do...
        //evt := __GET_EVENT_FROM_request__
        fmt.Println("[EVENT] start...")
        fmt.Println("[EVENT] end...")
        //
        <-ctx.Done()
        fmt.Println("[/events] end...")
        fmt.Println("[/events] disconnected...")
    })
    */

    // register event handler
    t.ws_event_handler["c"] = t.wsServer_ws_handler_connect
    t.ws_event_handler["dc"] = t.wsServer_ws_handler_disconnect
    t.ws_event_handler["sd"] = t.WSServer_ws_handler_shutdown
    t.ws_event_handler["s"] = t.wsServer_ws_handler_states

	log.Fatal( http.ListenAndServe(t.ws_addr, nil) )
    //
    // not work
    //
    // with CORS
	//log.Fatal( http.ListenAndServe(t.ws_addr, enable_cors(mux)) )
    //
    //handler_cors := cors.Default().Handler(mux)
	//log.Fatal( http.ListenAndServe(t.ws_addr, handler_cors) )
}

func (t *WSServer) WSServer_release() {
    log.Println( "WSServer_release()" )
    log.Println( "WSServer_release(): room_id = ", t.room_id, t.ws_addr )

    t.ws_is_started = false
    t.ws_is_closed = true
}

func (t *WSServer) wsServer_timer_reset(initial bool, timer map[string]interface{}) {
    log.Println( "wsServer_timer_reset()" )

    datetime_str := timer["datetime"].(string)
    prep_time_s := timer["prep_time_s"].(int)

    timestamp := time.Now().Unix()
    //timestamp := time.Now().UnixMilli()
    //timestamp := time.Now().UnixNano()
    log.Println( "wsServer_timer_reset(): timestamp = ", timestamp )

    now_utc_datetime := time.Now().UTC().Format(time.DateTime)
    now_datetime := time.Now().Format(time.DateTime)
    log.Println( "wsServer_timer_reset(): now_utc_datetime = ", now_utc_datetime )
    log.Println( "wsServer_timer_reset(): now_datetime = ", now_datetime )

    //time.Sleep( time.Second * time.Duration(UPDATES_INTERVAL) )
    //time.Sleep( time.Millisecond * 1000 * time.Duration(UPDATES_INTERVAL) )


    // input datetime (Local)
    datetime, _ := time.Parse( time.DateTime, datetime_str )
    log.Println( "wsServer_timer_reset(): datetime_str = ", datetime_str )
    log.Println( "wsServer_timer_reset(): datetime = ", datetime )

    // UTC (internal)
    datetime_utc := datetime.UTC()
    log.Println( "wsServer_timer_reset(): datetime_utc = ", datetime_utc )

    // Local (external)



    if initial == true {
        //t.timer.Initial_start_timestamp = int(timestamp)
        t.timer.Initial_start_datetime = datetime_str
        t.timer.Initial_start_prep_time_s = prep_time_s

        //t.timer.Interval_s = int(interval)

        /*
        utc_start_prep_time_s int
        end_prep_time_s int

        // UTC
        utc_start_timeatamp int64
        utc_start_datetime string
        utc_end_timeatamp int64
        utc_end_datetime string

        // Local
        local_start_timeatamp int64
        local_start_datetime string
        local_end_timeatamp int64
        local_end_datetime string
        */
    }
}

func (t *WSServer) WSServer_timer_now_reset() {
    //log.Println( "WSServer_timer_now_reset()" )

    //timestamp := time.Now().Unix()
    //timestamp := time.Now().UnixMilli()
    //timestamp := time.Now().UnixNano()
    //log.Println( "WSServer_timer_now_reset(): timestamp = ", timestamp )

    t.timer.Now_utc_datetime = time.Now().UTC().Format(time.DateTime)
    t.timer.Now_datetime = time.Now().Format(time.DateTime)
    //log.Println( "WSServer_timer_now_reset(): now_utc_datetime = ", t.Now_utc_datetime )
    //log.Println( "WSServer_timer_now_reset(): now_datetime = ", t.Now_datetime )
}


func (t *WSServer) WSServer_checks_max_players() bool {
    log.Println( "WSServer_checks_max_players()" )

    ret := true
    if t.n_players > t.n_max_players {
        ret = false
    }

    return ret
}

func (t *WSServer) WSServer_reconnectable_and_apply(user_id string) bool {
    log.Println( "WSServer_reconnectable_and_apply()" )

    ret := true
    return ret
}


// ws handler
//

/*
// not work
func enable_cors(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")

        if r.Method == "OPTIONS" {
            w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
            w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
            return
        }

        next.ServeHTTP(w, r)
    })
}
*/

func (t *WSServer) wsServer_ws_handler(w http.ResponseWriter, r *http.Request) {
    //log.Println( "wsServer_ws_handler()" )

	c, err := t.ws_upgrader.Upgrade(w, r, nil)
	if err != nil {
        log.Println("wsServer_ws_handler(): upgrade:", err)
		return
	}
	defer c.Close()

    // connected
    cid := uuid.New().String()
    t.ws_sids[cid] = c
    cid_data := map[string]interface{} { "cid": cid, "data": map[string]interface{} {}, }
    //t.wsServer_ws_handler_connect( cid_data, &result )
    //t.ws_event_handler["c"]( cid_data )

	for {
        if t.ws_is_closed == true {
            log.Println("wsServer_ws_handler(): closing...")
            break
        }

        // Read
		//mt, message, err := c.ReadMessage()
		_, message, err := c.ReadMessage()
		if err != nil {
            log.Println("wsServer_ws_handler(): read:", err)
            break
		}

        //log.Printf("wsServer_ws_handler(): [%s] recv: %s\n", cid, message)

        // Write
		//err = c.WriteMessage(mt, message)
		//if err != nil {
        //    log.Println("wsServer_ws_handler(): write:", err)
        //break
		//}


        {
            // switch events
            // message: { "e": "", "v": {} }
            event := ""

            //event_byte, _ := json.Marshal( message )
            //event_byte, err_marshal := json.Marshal( message )
            //if err_marshal != nil {
            //    //panic( err_marshal.Error() )
            //    log.Printf( "wsServer_ws_handler(): [ERROR]: %s\n", err_marshal.Error() )
            //}
            //event_str = string(event_byte)

            var event_json map[string]interface{}
            err_unmarshal := json.Unmarshal( []byte(message), &event_json )
            if err_unmarshal != nil {
                //panic( err_unmarshal.Error() )
                log.Printf( "wsServer_ws_handler(): [ERROR]: %s\n", err_unmarshal.Error() )
                continue
            }

            event = event_json["e"].(string)
            cid_data["data"] = event_json["v"]

            if len(event) > 0 {
                t.ws_event_handler[ event ](cid_data)
            }
        }
	} // for {}

    // disconnected
    delete( t.ws_sids, cid )
    //t.wsServer_ws_handler_disconnect( cid_data, &result )
    t.ws_event_handler["dc"]( cid_data )
}

func (t *WSServer) payload_cid(cid_data map[string]interface{}) string {
    //return cid_data["cid"].(string)

    ret := ""
    if v, has := cid_data["cid"]; has { ret = v.(string) }
    return ret
}

func (t *WSServer) payload_data(cid_data map[string]interface{}) interface{} {
    //return cid_data["data"]

    var ret interface{}
    if v, has := cid_data["data"]; has { ret = v.(interface{}) }
    return ret
}

func (t *WSServer) payload_data_get(cid_data map[string]interface{}, key string) interface{} {
    //return cid_data["data"]

    var ret interface{}
    if _v, _has := cid_data["data"]; _has {
        if v, has := _v.(map[string]interface{})[key]; has {
            ret = v.(interface{})
        }
    }
    return ret
}

func (t *WSServer) payload_get(cid_data map[string]interface{}, key string) interface{} {
    var ret interface{}
    if v, has := cid_data[key]; has { ret = v.(interface{}) }
    return ret
}

func (t *WSServer) emit_to(e string, m interface{}, to string) {
    //broadcast := false
    //t.emit_ex( e, m, to, broadcast )

    msg := map[string]interface{} { "e": e, "v": m, }

    c := t.ws_sids[to]

    mt := websocket.TextMessage
    //mt := websocket.BinaryMessage

    //msg_byte, _ := json.Marshal( msg )
    msg_byte, err_marshal := json.Marshal( msg )
    if err_marshal != nil {
        //panic( err_marshal.Error() )
        log.Printf( "emit_to(): [ERROR]: %s\n", err_marshal.Error() )
    }

    err := c.WriteMessage( mt, msg_byte )
    if err != nil {
        log.Println("emit_to(): write:", err)
    }
}

func (t *WSServer) emit_bc(e string, m interface{}) {
    //to := ""
    //broadcast := true
    //t.emit_ex( e, m, to, broadcast )

    // val: unused
    for key, _ := range t.ws_sids {
        // val: unused
        if _, has := t.ws_sids[key]; has {
            t.emit_to( e, m, key )
        }
    }
}

func (t *WSServer) emit_ex(e string, m interface{}, to string, broadcast bool) {
    msg := map[string]interface{} { "e": e, "v": m, }

    if broadcast == true {
    } else {
        is_room_id := false
        {
            // checks 'to' for room_id
        }

        if is_room_id == true {
        } else {
            c := t.ws_sids[to]

            mt := websocket.TextMessage
            //mt := websocket.BinaryMessage

            //msg_byte, _ := json.Marshal( msg )
            msg_byte, err_marshal := json.Marshal( msg )
            if err_marshal != nil {
                //panic( err_marshal.Error() )
                log.Printf( "emit_ex(): [ERROR]: %s\n", err_marshal.Error() )
            }

            err := c.WriteMessage( mt, msg_byte )
            if err != nil {
                log.Println("emit_ex(): write:", err)
            }
        }
    }
}

func (t *WSServer) wsServer_ws_handler_connect(
    data map[string]interface{},
) {
    cid := t.payload_cid(data) 
    //v := data["data"]
    log.Println( "wsServer_ws_handler_connect(): cid = ", cid )
    log.Println( "wsServer_ws_handler_connect(): data = ", t.payload_data(data) )

    user_id := t.payload_data_get( data, "user_id" ).(string)
    username := t.payload_data_get( data, "username" ).(string)

    if user_id == "" || username == "" {
        // disconnected
        delete( t.ws_sids, cid )
        t.ws_event_handler["dc"]( data )
    }

    player := t.player_new()
    player.user_id = user_id
    player.username = username
    player.cid = cid

    r := map[string]interface{} {} // initial response
    t.emit_to( "s", r, t.payload_cid(data) )
}

func (t *WSServer) wsServer_ws_handler_disconnect(
    data map[string]interface{},
) {
    cid := t.payload_cid(data) 
    log.Println( "wsServer_ws_handler_disconnect(): cid = ", cid )

    //! FIXME: reconnection
    //
    delete( t.ws_sids, cid )
    t.player_del( cid )
}

func (t *WSServer) WSServer_ws_handler_shutdown(
    data map[string]interface{},
) {
    cid := t.payload_cid(data) 
    log.Println( "wsServer_ws_handler_shutdown(): cid = ", cid )

    t.ws_is_closed = true

    t.emit_bc( "d", "" )
}

func (t *WSServer) wsServer_ws_handler_states(
    data map[string]interface{},
) {
    cid := t.payload_cid(data) 
    //log.Println( "wsServer_ws_handler_states(): cid_data = ", cid, t.payload_data(data) )

    //t.payload_get( data, "" )

    dummy := t.payload_data_get( data, "dummy")
    if dummy == nil { dummy = int(0) }

    user_id := ""
    username := ""
    player := t.player_get( cid )
    if player != nil {
        user_id = player.user_id
        username = player.username
    }

    t.Get_timer()

    r := map[string]interface{} {
        "dummy": dummy,
        "user_id": user_id,
        "username": username,
        "timer": map[string]interface{} {
            "init_datetime": t.Get_timer().Initial_start_datetime,
            "prep_time_s": t.Get_timer().Initial_start_prep_time_s,
            "datetime": t.Get_timer().Now_datetime,
        },
    }

    t.emit_to( "s", r, cid )
}

func (t *WSServer) wsServer_ws_handler_chat (
    data map[string]interface{},
) {
    //cid := t.payload_cid(data) 
    //log.Println( "wsServer_ws_handler_chat(): cid_data = ", cid, t.payload_data(data) )
    //t.emit_to( "e", "test test test", cid )
}


// ---------------------------------------------------------------

