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
RUN cd .. 
RUN git clone ${devsource} 
RUN cd user && git checkout azure-pipelines
#RUN go get -v -t -d ${sourcesdir}
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh 
RUN cd /go/src/${sourcesdir} && ls && dep init && dep ensure && go build -v .

ENTRYPOINT user
EXPOSE 8084
