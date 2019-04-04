#!/bin/bash

SERVER_ADDR=${SERVER_ADDR:-:1001} /sendmail-http &

/opt/monit/bin/monit-start.sh