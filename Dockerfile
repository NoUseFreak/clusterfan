FROM golang:alpine AS gobuilder

WORKDIR /go/src/clusterfan

COPY go.* ./
RUN go mod download

COPY . ./
RUN go build ./cmd/clusterfan

FROM alpine
RUN addgroup -S appgroup && adduser -S appuser -G appgroup \
    && mkdir /data \
    && chown -R appuser:appgroup /data

COPY --from=gobuilder /go/src/clusterfan/clusterfan /clusterfan

USER appuser
CMD [ "/clusterfan" ]