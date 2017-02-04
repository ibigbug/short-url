FROM hub.c.163.com/library/debian:jessie-slim

MAINTAINER Yuwei Ba <akabyw@gmail.com>

# TODO: split bootstrap and build to spped up image build
ADD . /src/
WORKDIR /src/
RUN sh ./scripts/bootstrap.sh

CMD ["sh", "./scripts/start.sh"]
