FROM library/alpine
MAINTAINER <Zhengyao Xie>

RUN apk update && apk add tzdata diffutils curl && cp -r -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && apk del tzdata

WORKDIR /tantan-simplify
COPY ./bin/tantan-simplify /tantan-simplify/bin/tantan-simplify
COPY ./conf /tantan-simplify/conf
COPY ./log /tantan-simplify/log
COPY ./run-in-docker.sh /tantan-simplify/run-in-docker.sh
RUN chmod 0750 /tantan-simplify/run-in-docker.sh
CMD ["./run-in-docker.sh"]