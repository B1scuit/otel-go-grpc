FROM golang:1.18-alpine AS build

ADD ./ /go/app

WORKDIR /go/app

ARG SERVICE=service_one
ARG PORT=8080

ENV PORT=${PORT}

RUN apk update && apk upgrade && \
    apk add --no-cache bash git curl make gcc openssh tzdata ca-certificates && \
    # Update CA certificates
    update-ca-certificates 2>/dev/null || true

RUN make ${SERVICE}

#
## Create final image
FROM scratch

COPY --from=build /go/app/main /opt/main
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /opt

ENTRYPOINT ["./main"]

EXPOSE ${PORT}