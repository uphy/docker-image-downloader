FROM golang:1.9

RUN curl -o /download.sh https://raw.githubusercontent.com/moby/moby/master/contrib/download-frozen-image-v2.sh && \
    chmod +x /download.sh && \
    curl -Lo /bin/jq https://github.com/stedolan/jq/releases/download/jq-1.5/jq-linux64 && \
    chmod +x /bin/jq

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT [ "/entrypoint.sh" ]
