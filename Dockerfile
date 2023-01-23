FROM node:11.15.0 AS node-builder

WORKDIR /work
COPY ./client /work/client
COPY ./package.json /work/
COPY ./package-lock.json /work/
COPY ./config/webpack.config.js /work/config/

RUN npm install
RUN ./node_modules/webpack/bin/webpack.js --config config/webpack.config.js

FROM golang:alpine AS go-builder

ENV GO111MODULE on
ENV GOPROXY https://goproxy.io
ENV CGO_ENABLED 0

WORKDIR /work
COPY ./main.go /work/
COPY ./server /work/server
COPY ./posts/*.go /work/posts/
COPY ./config/embed.go /work/config/
COPY ./go.mod /work/
COPY ./go.sum /work/

RUN go mod download
RUN go build main.go

FROM scratch

ENV GIN_MODE release
ENV INSIDE_DOCKER true

WORKDIR /output
COPY ./config/server.json /output/config/
COPY ./assets /output/assets
COPY ./server/templates /output/server/templates
COPY ./posts /output/posts
COPY --from=node-builder /work/client/public/bundle.js /output/client/public/
COPY --from=go-builder /work/main /output/

EXPOSE 13109
ENTRYPOINT [ "./main" ]