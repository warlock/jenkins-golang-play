FROM golang:alpine AS builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o app .

FROM scratch
WORKDIR /app
COPY --from=builder /build/app .
ENTRYPOINT [ "./app" ]