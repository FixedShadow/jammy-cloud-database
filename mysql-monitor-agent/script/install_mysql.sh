#!/bin/bash

echo "installing mysql5.7"

export LANG=en_US.UTF-8
export LANGUAGE=en_US:

TAR_FILE=/opt/mysql-5.7.20-linux-glibc2.12-x86_64.tar.gz
WD_HOME=/opt
MYSQL_DIR=/opt/mysql
MYSQL_DATA_DIR=/opt/mysql/data

tar -zxvf ${TAR_FILE} -C ${WD_HOME}

mv /${WD_HOME}/mysql-5.7.20-linux-glibc2.12-x86_64 ${MYSQL_DIR}

mkdir -p ${MYSQL_DATA_DIR}

cd ${MYSQL_DIR}

bin/mysqld --initialize --user=mysql --basedir=${MYSQL_DIR} --datadir=${MYSQL_DATA_DIR}

# 生成临时密码
bin/mysql_ssl_rsa_setup  --datadir=${MYSQL_DATA_DIR}

cp ${MYSQL_DIR}/support-files/mysql.server /etc/init.d/mysql

# 启动mysql
/etc/init.d/mysql start






