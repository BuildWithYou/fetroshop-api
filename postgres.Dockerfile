FROM postgres:latest

WORKDIR /
COPY db/init.sql /docker-entrypoint-initdb.d/init.sql
COPY db/migrations /db/migrations

ENV PGDATA=/var/lib/postgresql/data/pgdata
ENV POSTGRES_USER=postgres_user
ENV POSTGRES_PASSWORD=postgres_password
ENV DB_NAME=postgres_db
ENV DB_NAME2=postgres_db2

EXPOSE 5432