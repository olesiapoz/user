FROM golang:1.7-alpine
ENV sourcesdir /go/src/github.com/olesiapoz/user/
ENV MONGO_HOST mytestdb:27017
ENV HATEAOS user
ENV USER_DATABASE mongodb

COPY . ${sourcesdir}
RUN apk update
RUN apk add git
RUN go version 1.7
RUN go get -v -t -d ./...
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh 
RUN cd ${sourcesdir} && dep ensure && go build -v .

ENTRYPOINT user
EXPOSE 8084
