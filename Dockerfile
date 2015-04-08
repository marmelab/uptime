FROM ubuntu:14.04

RUN apt-get update --yes

RUN apt-get install --yes golang

WORKDIR /srv

CMD bash
