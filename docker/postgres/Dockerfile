FROM postgres

RUN mkdir -p /docker-entrypoint-initdb.d

ADD db_schema_setup.sql /docker-entrypoint-initdb.d
