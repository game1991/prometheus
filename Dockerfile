# syntax=docker/dockerfile:1
##
## Build
##
ARG IMAGE
FROM ${IMAGE} AS builder
WORKDIR /src
ADD . .
ARG APP_NAME
RUN apt-get update \
    && apt-get install net-tools -y \
    && make build APP_NAME=${APP_NAME} \
    && mkdir -p /src/data


##
## Deploy
##
FROM debian:bullseye-slim
WORKDIR /service
ARG PROJECT_NAME
COPY --from=builder /src/${PROJECT_NAME} /service/app
COPY --from=builder /src/conf        /service/conf
COPY --from=builder /src/data        /service/data

ENTRYPOINT [ "/service/app" ]
CMD [ "/bin/bash" ]

