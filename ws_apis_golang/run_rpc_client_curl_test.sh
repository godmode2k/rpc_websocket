#!/bin/bash



#echo -ne ' ' >> /dev/null



HOST="http://127.0.0.1:5000/rpc"
CONTENT_TYPE="Content-Type: application/json"
REQ_METHOD=$1
REQ_PARAMS=$2




# (cURL)

if [ -z $REQ_METHOD ] || [ "$REQ_METHOD" == "--help" ]; then
    echo "Usage: $ sh $0 [request method] [params]"
    echo
    echo "request methods:"
    echo "    create, close, list"
    echo "params:"
    echo "    create:"
    echo '        {"host": "0.0.0.0", "port": "5001", "n_max_players": 50,'
    echo '         "datetime": "2026-01-01 10:00:00", "prep_time_s": 30}'
    echo "    close:"
    echo '        {}'
    echo "    list:"
    echo '        {}'
    exit
fi

if [ "$REQ_METHOD" == "create" ]; then
    # room: create
    #-d '{ "method":"RPCServerAPIs.JSONRPC_room_create", "params":[{}], "id":0}'
    REQ_PARAMS='
        "params":[{"req": {"host": "0.0.0.0", "port": "5001", "n_max_players": 50,
        "datetime": "2026-01-01 10:00:00","prep_time_s": 30}}]
    '
    curl -H "$CONTENT_TYPE" \
        -d "{ \"method\":\"RPCServerAPIs.JSONRPC_room_create\", ${REQ_PARAMS}, \"id\":0 }" \
        $HOST
elif [ "$REQ_METHOD" == "close" ]; then
    # room: close
    #-d '{ "method":"RPCServerAPIs.JSONRPC_room_close", "params":[{}], "id":0}'
    #
    REQ_PARAMS='"params":[{"req": {"room_id": "", "port": ""}}]'

    curl -H "$CONTENT_TYPE" \
        -d "{ \"method\":\"RPCServerAPIs.JSONRPC_room_close\", ${REQ_PARAMS}, \"id\":0 }" \
        $HOST
elif [ "$REQ_METHOD" == "list" ]; then
    # room: list
    #-d '{ "method":"RPCServerAPIs.JSONRPC_room_list", "params":[{}], "id":0}'
    REQ_PARAMS='"params":[{"req": {}}]'
    curl -H "$CONTENT_TYPE" \
        -d "{ \"method\":\"RPCServerAPIs.JSONRPC_room_list\", ${REQ_PARAMS}, \"id\":0 }" \
        $HOST
else
    echo "error..."
    echo "for help: use '--help'"
    exit
fi


