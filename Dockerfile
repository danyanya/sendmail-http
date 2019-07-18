FROM interlegis/alpine-postfix 

LABEL maintainer="Danya Sliusar <danya.brain@gmail.com>"

ADD ./bin/sendmail-http /
ADD start.sh /

VOLUME ["/var/spool/postfix"]

EXPOSE 25 1001

ENTRYPOINT ["/start.sh"]