#!/bin/bash


function initNetwork() {
  sed -i 's/BOOTPROTO=dhcp/BOOTPROTO=static/g' /etc/sysconfig/network-scripts/ifcfg-eth0
  sed -i 's/ONBOOT=no/ONBOOT=yes/g' /etc/sysconfig/network-scripts/ifcfg-eth0
  echo -e "\nIPADDR=$1\nNETMASK=255.255.255.0\nGATEWAY=192.168.2.1" >> /etc/sysconfig/network-scripts/ifcfg-eth0
  systemctl restart network
  echo "Wait for the network to restart..."
  sleep 5
  ip addr
}

function initYum() {
  if [ -f /etc/yum.repos.d/CentOS-Base.repo ]
  then
    mv /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo.bak
  fi
  ping -c 5 repo.huaweicloud.com
  curl https://repo.huaweicloud.com/repository/conf/CentOS-7-reg.repo -o /etc/yum.repos.d/CentOS-Base.repo
  yum clean all
  yum makecache
  yum install -y epel-release
}


if [ $# -eq 1 ];then
  echo "ip address: $1"
else
  echo "please enter the ip address!"
  exit 1
fi
initNetwork $1
initYum