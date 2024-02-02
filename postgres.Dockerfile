FROM postgres:latest

WORKDIR /
COPY db/init.sql /docker-entrypoint-initdb.d/init.sql
COPY db/migrations /db/migrations


EXPOSE 5432