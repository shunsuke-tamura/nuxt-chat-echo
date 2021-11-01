FROM postgres:12-alpine
ENV POSTGRES_USER=root
ENV POSTGRES_PASSWORD=root
ENV POSTGRES_DB=nuxtechochat
ENV POSTGRES_INITDB_ARGS=--encoding=UTF-8
# ENV POSTGRES_HOST_AUTH_METHOD=trust
ADD ./chat.sql /docker-entrypoint-initdb.d
EXPOSE 5432