#!/bin/bash


function init_network() {
  sed -i 's/BOOTPROTO=dhcp/BOOTPROTO=static/g' /etc/sysconfig/network-scripts/ifcfg-eth0
  sed -i 's/ONBOOT=no/ONBOOT=yes/g' /etc/sysconfig/network-scripts/ifcfg-eth0
  echo -e "\nIPADDR=$1\nNETMASK=255.255.255.0\nGATEWAY=192.168.2.1" >> /etc/sysconfig/network-scripts/ifcfg-eth0
  systemctl restart network
  echo "Wait for the network to restart..."
  sleep 5
  ip addr
}

function init_yum() {
  if [ -f /etc/yum.repos.d/CentOS-Base.repo ]
  then
    mv /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo.bak
  fi
  ping -c 5 repo.huaweicloud.com
  curl https://repo.huaweicloud.com/repository/conf/CentOS-7-reg.repo -o /etc/yum.repos.d/CentOS-Base.repo
  yum clean all
  yum makecache
  yum install -y epel-release
  sleep 2
  yum -y install htop
  yum -y install vim
  yum -y install net-tools
  yum -y install wget
}

function install_go_env() {
    echo "Start installing golang..."
    cd /home
    wget https://mirrors.aliyun.com/golang/go1.23.8.linux-amd64.tar.gz
    tar -zxvf /home/go1.23.8.linux-amd64.tar.gz
    mv /home/go /usr/local/go
    echo -e "\nexport PATH=\$PATH:/usr/local/go/bin\n" >> /etc/profile
    echo -e "export GOPATH=/opt/gopath\n" >> /etc/profile
    echo -e "export PATH=\$GOPATH/bin:\$PATH\n" >> /etc/profile
    source /etc/profile
    go env -w GO111MODULE=on
    go env -w GOPROXY='https://goproxy.cn,direct'
    go env
    echo "The Golang development environment is installed!"
}

function install_java_env() {

}

function install_py3_env() {

}

function setup_cron() {

}

if [ $# -eq 1 ];then
  echo "ip address: $1"
else
  echo "please enter the ip address!"
  exit 1
fi

init_network $1
init_yum

#install_go_env
#install_java_env
#install_py3_env