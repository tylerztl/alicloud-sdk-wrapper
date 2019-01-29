FROM library/golang

ENV APP_DIR $GOPATH/src/zig-cloud
RUN mkdir -p $APP_DIR
ADD . $APP_DIR
WORKDIR $APP_DIR

RUN go get github.com/kardianos/govendor \
    && govendor sync \
    && go build $APP_DIR

ENTRYPOINT  ./zig-cloud
EXPOSE 8080
