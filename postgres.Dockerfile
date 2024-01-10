FROM postgres:latest

WORKDIR /
COPY db/migrations /db/migrations
COPY db/init.sql /docker-entrypoint-initdb.d/init.sql


EXPOSE 5432