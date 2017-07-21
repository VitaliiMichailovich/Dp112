FROM golang:onbuild

MAINTAINER Vitalii Radchenko <vitaliy@online.ua>

ADD . ./client
ADD . ./templates

EXPOSE 8080