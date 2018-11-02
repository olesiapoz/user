FROM golang:1.7-alpine
ENV sourcesdir github.com/olesiapoz/user
ENV devsource https://github.com/olesiapoz/user.git
ENV MONGO_HOST mytestdb:27017
ENV HATEAOS user
ENV USER_DATABASE mongodb

COPY . ${sourcesdir}
RUN apk update
RUN apk add git
RUN apk add curl
RUN mkdir -p /go/src/${sourcesdir} 
WORKDIR /go/src/github.com/olesiapoz
RUN git clone ${devsource} 
WORKDIR user
RUN git fetch origin azure-pipelines && git checkout azure-pipelines
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh 
RUN dep ensure && go build -v .

ENTRYPOINT user
EXPOSE 8084
