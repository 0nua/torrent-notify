FROM ubuntu:16.04

RUN mkdir -p /var/www/torrent-notify
RUN mkdir -p /data/db

RUN apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv EA312927
RUN echo "deb http://repo.mongodb.org/apt/ubuntu xenial/mongodb-org/3.2 multiverse" | tee /etc/apt/sources.list.d/mongodb-org-3.2.list

RUN apt-get update
RUN apt-get install -y software-properties-common
RUN add-apt-repository ppa:longsleep/golang-backports
RUN apt-get update
RUN apt-get install -y mongodb-org golang-go git

RUN export GOPATH /var/www/torrent-notify/

RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/telegram-bot-api.v4