FROM mysql@sha256:e9c9e3680bbadd5230a62c5548793bd8e59cbcc868032781e48bd53e888bd82f

ENV TZ=UTC \
  MYSQL_DATABASE=test \
  MYSQL_USER=test1234 \
  MYSQL_PASSWORD=test1234 \
  MYSQL_ROOT_PASSWORD=password

COPY ./my.cnf /etc/mysql/conf.d/my.cnf
RUN chmod 644 /etc/mysql/conf.d/my.cnf