FROM golang:1.21.2-alpine3.18 AS build-stage
WORKDIR /app
COPY ./ /app
RUN mkdir -p /app/build
RUN go mod download
RUN go build -v -o /app/build/api ./cmd/api


FROM gcr.io/distroless/static-debian11
COPY --from=build-stage /app/build/api /api
COPY --from=build-stage /app/template /template
COPY --from=build-stage /app/.env /
EXPOSE 3000
CMD ["/api"]