FROM mysql/mysql-server:8.0

COPY ./docker/custom.cnf /etc/mysql/conf.d/custom.cnf