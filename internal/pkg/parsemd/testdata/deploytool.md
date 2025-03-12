# 解析理应不受这段文字影响
```Dockerfile
# [to]: deploytool
FROM alpine:3.18

# Install dependencies
RUN apk add --no-cache \
    bash \
    curl \
    git \
    jq \
    openssh-client \
    gettext \
    docker-cli \    
    && rm -rf /var/cache/apk/*

RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" \
    && install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl \
    && rm kubectl
```
在这个markdown中,你可以一些其他信息