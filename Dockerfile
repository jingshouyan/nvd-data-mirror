FROM httpd:alpine



# Labels.
LABEL maintainer="jingshouyan@gmail.com"
LABEL name="jingshouyan/nvd-data-mirror"
LABEL org.label-schema.schema-version="1.0"
LABEL org.label-schema.build-date=$BUILD_DATE
LABEL org.label-schema.name="jingshouyan/nvd-data-mirror"
LABEL org.label-schema.description="NIST Data Mirror"
LABEL org.label-schema.url="https://github.com/jingshouyan/nvd-data-mirror"
LABEL org.label-schema.vcs-url="https://github.com/jingshouyan/nvd-data-mirror"
LABEL org.label-schema.vendor="jingshouyan"

ENV user=mirror

COPY nvd-data-mirror /  docker/ /


RUN apk update                                               && \
    apk add --no-cache dcron nss supervisor                  && \
    addgroup -S $user                                        && \
    adduser -S $user -G $user                                && \
    mkdir -p /tmp/nvd                                        && \
    chown -R $user:$user /tmp/nvd                            && \
    chown -R $user:$user /usr/local/apache2/htdocs           && \
    rm -v /usr/local/apache2/htdocs/index.html               && \
    chmod +x /mirror.sh


EXPOSE 80/tcp

CMD ["/usr/bin/supervisord", "-n", "-c", "/etc/supervisor/conf.d/supervisord.conf", "-l", "/var/log/supervisord.log", "-j", "/var/run/supervisord.pid"]
