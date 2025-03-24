#!/bin/bash
set -e

echo "➡️  初始化 mon 数据目录"
mkdir -p /var/lib/ceph/mon/ceph-$(hostname -s)

if [ ! -f /etc/ceph/monmap ]; then
  echo "📦 创建 monmap"
  ceph-authtool --create-keyring /etc/ceph/ceph.mon.keyring \
    --gen-key -n mon. --cap mon 'allow *'
  monmaptool --create --add $(hostname -s) $(hostname -i) --fsid=${FSID} /etc/ceph/monmap
fi

if [ ! -d /var/lib/ceph/mon/ceph-$(hostname -s)/store.db ]; then
  echo "🔧 初始化 mon 数据目录"
  ceph-mon --mkfs -i $(hostname -s) --monmap /etc/ceph/monmap \
    --keyring /etc/ceph/ceph.mon.keyring
fi

echo "🚀 启动 ceph-mon 服务"
exec ceph-mon -i $(hostname -s) --keyring /etc/ceph/ceph.mon.keyring --foreground
