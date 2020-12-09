FROM ubuntu:xenial

LABEL "install-type"="mounted"

RUN apt-get update && apt-get install -y -q --no-install-recommends \
    curl \
    ca-certificates \
    pkg-config \
    build-essential \
    libfontconfig \
    libzmq3-dev \
 && curl -s -S -o /tmp/setup-node.sh https://deb.nodesource.com/setup_6.x \
 && chmod 755 /tmp/setup-node.sh \
 && /tmp/setup-node.sh \
 && apt-get install nodejs -y -q \
 && rm /tmp/setup-node.sh \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

RUN \
 if [ ! -z $HTTP_PROXY ] && [ -z $http_proxy ]; then \
  http_proxy=$HTTP_PROXY; \
 fi; \
 if [ ! -z $HTTPS_PROXY ] && [ -z $https_proxy ]; then \
  https_proxy=$HTTPS_PROXY; \
 fi; \
 if [ ! -z $http_proxy ]; then \
  npm config set proxy $http_proxy; \
 fi; \
 if [ ! -z $https_proxy ]; then \
  npm config set https-proxy $https_proxy; \
 fi

RUN npm install -g \
    prebuild-install \
    phantomjs-prebuilt

RUN mkdir -p /project/sawtooth-sdk-javascript/ \
 && mkdir -p /var/log/sawtooth \
 && mkdir -p /var/lib/sawtooth \
 && mkdir -p /etc/sawtooth \
 && mkdir -p /etc/sawtooth/keys

ENV PATH=$PATH:/project/sawtooth-sdk-javascript/bin:/node_modules/phantomjs-prebuilt/bin

WORKDIR /app

RUN apt-get update -qq && \
    apt-get install -y -qq libzmq3-dev;

ADD ./package.json /tmp/package.json
RUN cd /tmp && npm install

RUN cd /app && cp -a /tmp/node_modules /app
COPY . /app

EXPOSE 4004

ENTRYPOINT ["node", "index.js"]

