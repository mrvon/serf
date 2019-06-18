#!/bin/bash

# user event
serf event deploy "hello world"
# user query
serf query load "hello world"
# member leave
serf leave -rpc-addr=127.0.0.1:7375
serf leave -rpc-addr=127.0.0.1:7376
