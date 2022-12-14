FROM ttbb/base:go AS build
COPY . /opt/sh/compile
WORKDIR /opt/sh/compile/pkg
RUN go build -o nginx_mate .


FROM ttbb/openresty:nake

ENV OPENRESTY_HOME /opt/sh/openresty

COPY docker-build /opt/sh/openresty/mate

COPY --from=build /opt/sh/compile/pkg/nginx_mate /opt/sh/openresty/mate/nginx_mate

WORKDIR /opt/sh/openresty

CMD ["/usr/bin/dumb-init", "bash", "-vx", "/opt/sh/openresty/mate/scripts/start.sh"]
