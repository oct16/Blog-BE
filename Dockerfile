FROM golang

MAINTAINER Oct16 mail@fengfan.me

# Bundle app source
COPY . /go/src/echo-blog

WORKDIR /go/src/echo-blog
RUN mv depends/* ../

EXPOSE 3016

CMD ["./gorun.sh"]
