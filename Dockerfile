FROM postgres:latest

# will be overriden
ENV POSTGRES_DB def
ENV POSTGRES_USER def
ENV POSTGRES_PASSWORD def

COPY ./db/postgres/init.sql /docker-entrypoint-initdb.d/init.sql