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
WORKDIR /go/src/githu.com/olesiapoz
RUN cd .. 
RUN git clone ${devsource} 
WORKDIR user
RUN git fetch origin azure-pipelines && git checkout azure-pipelines  && pwd && ls
#RUN go get -v -t -d ${sourcesdir}
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh 
RUN pwd &&  dep ensure && go build -v .

ENTRYPOINT user
EXPOSE 8084
