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

groupadd mysql

useradd -r -g mysql -s /bin/false mysql

mkdir -p ${MYSQL_DIR}/data

mkdir -p ${MYSQL_DIR}/logs

touch ${MYSQL_DIR}/logs/error.log

chown -R mysql:mysql ${MYSQL_DIR}



${MYSQL_DIR}/bin/mysqld --initialize --user=mysql --basedir=${MYSQL_DIR} --datadir=${MYSQL_DATA_DIR}

# 生成临时密码
${MYSQL_DIR}/bin/mysql_ssl_rsa_setup  --datadir=${MYSQL_DATA_DIR}

# 启动mysql
/opt/mysql/support-files/mysql.server start

cat ${MYSQL_DIR}/logs/error.log |grep "temporary password" |awk '{print $NF}' >> ${MYSQL_DIR}/tmp_password.txt

chown -R mysql:mysql ${MYSQL_DIR}









