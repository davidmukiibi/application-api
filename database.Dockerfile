FROM postgres:alpine

# ADD dbsetup.sql /docker-entrypoint-initdb.d

# RUN apk --no-cache add sudo1

ADD dbsetup.sh /docker-entrypoint-initdb.d

RUN chmod +x /docker-entrypoint-initdb.d/dbsetup.sh

USER postgres

EXPOSE 5432