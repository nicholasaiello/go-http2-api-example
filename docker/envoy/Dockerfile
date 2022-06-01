FROM envoyproxy/envoy-dev:latest

COPY docker/envoy/envoy.yaml /etc/envoy.yaml
COPY server.crt server.key /etc/ssl/

RUN chmod go+r /etc/envoy.yaml
CMD /usr/local/bin/envoy -c /etc/envoy.yaml -l debug
