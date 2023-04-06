FROM node:18-alpine AS node-builder

WORKDIR /work
COPY ./templates /work/templates
COPY ./tailwind.config.js /work/

RUN npx tailwindcss -o ./built.css --minify

FROM golang:alpine AS go-builder

ENV GOPROXY https://goproxy.io,direct
ENV CGO_ENABLED 0

WORKDIR /work

COPY ./go.mod /work/
COPY ./go.sum /work/
RUN go mod download

COPY ./main.go /work/
COPY ./api /work/api
COPY ./posts /work/posts
COPY ./templates /work/templates
RUN go build main.go

FROM scratch

ENV GIN_MODE release
ENV PORT 13109

WORKDIR /output
COPY ./assets /output/assets
COPY --from=node-builder /work/built.css /output/assets/
COPY --from=go-builder /work/main /output/

EXPOSE ${PORT}
ENTRYPOINT [ "./main" ]