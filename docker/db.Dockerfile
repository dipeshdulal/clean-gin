FROM mysql:5.7.22

COPY ./docker/custom.cnf /etc/mysql/conf.d/custom.cnf