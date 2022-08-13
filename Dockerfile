FROM alpine as certs
RUN apk update && apk add ca-certificates

FROM busybox:stable

ARG BUILD_DATE
ARG BUILD_VERSION

# Labels.
LABEL maintainer="jingshouyan@gmail.com"
LABEL name="jingshouyan/nvd-data-mirror"
LABEL version=$BUILD_VERSION
LABEL org.label-schema.schema-version="1.0"
LABEL org.label-schema.build-date=$BUILD_DATE
LABEL org.label-schema.name="jingshouyan/nvd-data-mirror"
LABEL org.label-schema.description="NIST Data Mirror"
LABEL org.label-schema.url="https://github.com/jingshouyan/nvd-data-mirror"
LABEL org.label-schema.vcs-url="https://github.com/jingshouyan/nvd-data-mirror"
LABEL org.label-schema.vendor="jingshouyan"
LABEL org.label-schema.version=$BUILD_VERSION

COPY nvd-data-mirror /  docker/entry-point.sh /
COPY --from=certs /etc/ssl/certs /etc/ssl/certs

RUN chmod +x /entry-point.sh

VOLUME [ "/data" ]

EXPOSE 80/tcp

ENTRYPOINT [ "/entry-point.sh" ]
