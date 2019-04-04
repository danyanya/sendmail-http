FROM interlegis/alpine-postfix 

ADD ./bin/sendmail-http /
ADD start.sh /

EXPOSE 25 1001

ENTRYPOINT ["/start.sh"]