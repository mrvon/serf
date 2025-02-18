#!/bin/bash
SERF="./bin/serf"
HANDLER="./scripts/handler/handler"

# Set secret, blank for disabled
SECRET="1FzgH8LsTtr0Wopn4934OQ=="
#SECRET=""

# Set protocol version
PROTO="5"

for i in {0..5}
do
    BIND=`expr 2 + $i`
    PORT=`expr 7373 + $i`
    echo Starting Serf agent $i on 127.0.0.1:$BIND, RPC on port $PORT
    $SERF agent -node=node$i -rpc-addr=127.0.0.1:$PORT \
        -bind=127.0.0.$BIND -log-level=warn -join=127.0.0.2 -encrypt=$SECRET \
        -protocol=$PROTO -event-handler=$HANDLER &
    if [ $i -eq 0 ]
    then
        sleep 1;
    fi
done
