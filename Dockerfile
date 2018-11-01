FROM golang:1.7-alpine
ENV sourcesdir github.com/olesiapoz/user/
ENV MONGO_HOST mytestdb:27017
ENV HATEAOS user
ENV USER_DATABASE mongodb

COPY . ${sourcesdir}
RUN apk update
RUN apk add git
RUN go get -v -t -d ${sourcesdir}
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh 
RUN cd /go/src/${sourcesdir} && dep ensure && go build -v .

ENTRYPOINT user
EXPOSE 8084
