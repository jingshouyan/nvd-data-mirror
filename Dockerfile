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

RUN apk update                                               && \
    apk add --no-cache dcron nss supervisor                  && \
    addgroup -S $user                                        && \
    adduser -S $user -G $user                                && \
    mkdir -p /tmp/nvd                                        && \
    chown -R $user:$user /tmp/nvd                            && \
    chown -R $user:$user /usr/local/apache2/htdocs           && \
    rm -v /usr/local/apache2/htdocs/index.html

COPY ["/docker/conf/supervisord.conf", "/etc/supervisor/conf.d/supervisord.conf"]
COPY ["/docker/scripts/mirror.sh", "/mirror.sh"]
COPY ["/docker/crontab/mirror", "/etc/crontabs/mirror"]
COPY ["/docker/conf/mirror.conf", "/usr/local/apache2/conf"]
COPY ["/nvd-data-mirror", "/"]

EXPOSE 80/tcp

CMD ["/usr/bin/supervisord", "-n", "-c", "/etc/supervisor/conf.d/supervisord.conf", "-l", "/var/log/supervisord.log", "-j", "/var/run/supervisord.pid"]
