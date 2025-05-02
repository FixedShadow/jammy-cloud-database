#!/bin/bash


initNetwork() {
  sed 's/BOOTPROTO=dhcp/BOOTPROTO=static/g' /etc/sysconfig/network-scripts/ifcfg-eth0
  sed 's/ONBOOT=no/ONBOOT=yes/g' /etc/sysconfig/network-scripts/ifcfg-eth0
  echo -e "\nIPADDR=$1\nNETMASK=255.255.255.0\nGATEWAY=192.168.2.1" >> /etc/sysconfig/network-scripts/ifcfg-eth0
  systemctl restart network
  ip addr
}

function initYum() {
  mv /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo.bak
  curl https://repo.huaweicloud.com/repository/conf/CentOS-7-reg.repo -o /etc/yum.repos.d/CentOS-Base.repo
  yum clean all
  yum makecache
  yum install -y epel-release
}

initNetwork $1
initYum