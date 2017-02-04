FROM hub.c.163.com/library/debian:jessie-slim

MAINTAINER Yuwei Ba <akabyw@gmail.com>

ADD . /src/

WORKDIR /src/

RUN sh ./scripts/bootstrap.sh

CMD ["sh", "./scripts/start.sh"]
