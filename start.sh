#!/bin/bash

/etc/postfix/postfix-service.sh start &

SERVER_ADDR=${SERVER_ADDR:-:1001} /sendmail-http
