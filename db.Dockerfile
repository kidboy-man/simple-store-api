FROM postgres
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD postgrespw
ENV POSTGRES_DB db_simple_store
# COPY init.sql /docker-entrypoint-initdb.d/
EXPOSE 5432