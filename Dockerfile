FROM library/golang

ENV APP_DIR $GOPATH/src/alicloud-sdk-wrapper
RUN mkdir -p $APP_DIR
ADD . $APP_DIR
WORKDIR $APP_DIR

RUN go get github.com/kardianos/govendor \
    && govendor sync \
    && go build $APP_DIR

ENTRYPOINT  ./alicloud-sdk-wrapper
EXPOSE 8080
