FROM alpine
LABEL authors="Deepesh"

COPY . /app
WORKDIR /app
RUN chmod +x src/*.sh Password-Manager.sh

ENTRYPOINT ["./Password-Manager.sh"]