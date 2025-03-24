#!/bin/bash

MON_IP=${1:-"172.16.56.4"}
MNT_PATH="/mnt/cephfs"

sudo apt update
sudo apt install -y ceph-common

sudo mkdir -p $MNT_PATH
mkdir -pv /etc/ceph
echo "AQBXZTZkqLYaBxAAz4JtHtW8DzDiRjX/7FSbpg==" | sudo tee /etc/ceph/real.secret > /dev/null
sudo chmod 600 /etc/ceph/real.secret
echo "🛰️  尝试内核挂载 CephFS..."
sudo mount -t ceph ${MON_IP}:6789:/ ${MNT_PATH} -o name=admin,secretfile=/etc/ceph/real.secret

echo "✅ 挂载完成，内容如下："
ls -l $MNT_PATH


