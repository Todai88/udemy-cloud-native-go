FROM golang:1.9-alpine
LABEL AUTHOR="J.-Bajoul Kakaei"

ENV SOURCES /go/src/github.com/todai88/udemy-cloud-native-go-api/
ENV PORT 5000

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go install -a
RUN ls ${SOURCES}
RUN chmod +x -R ${SOURCES}

EXPOSE ${PORT}

ENTRYPOINT udemy-cloud-native-go-api