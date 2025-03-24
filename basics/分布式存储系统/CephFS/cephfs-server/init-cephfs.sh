#!/bin/bash
set -e

echo "🚀 初始化 CephFS 中... 要先启动该容器，然后启动集群，方可完成初始化"

# ✅ 创建 pool（幂等）
ceph osd pool ls | grep cephfs_data || ceph osd pool create cephfs_data 32
ceph osd pool ls | grep cephfs_metadata || ceph osd pool create cephfs_metadata 32

# ✅ 设置副本数为 1（单节点部署必备）默认为3 如果不够 将无法进行挂载
ceph config set global mon_allow_pool_size_one true

ceph osd pool set cephfs_data size 1 --yes-i-really-mean-it
ceph osd pool set cephfs_metadata size 1 --yes-i-really-mean-it


# ✅ 创建 CephFS（幂等）
ceph fs ls | grep cephfs || ceph fs new cephfs cephfs_metadata cephfs_data

# ✅ 生成 mgr 密钥
mkdir -p /var/lib/ceph/mgr/ceph-mgr
ceph auth get-or-create mgr.mgr mon 'allow profile mgr' osd 'allow *' mds 'allow *' > /var/lib/ceph/mgr/ceph-mgr/keyring

# ✅ 生成 osd 密钥
mkdir -p /var/lib/ceph/osd/ceph-0
ceph auth get-or-create osd.0 mon 'allow profile osd' osd 'allow *' > /var/lib/ceph/osd/ceph-0/keyring

# ✅ 生成 mds 密钥
mkdir -p /var/lib/ceph/mds/ceph-mds-a
ceph auth get-or-create mds.mds-a mon 'allow profile mds' osd 'allow *' mds 'allow' > /var/lib/ceph/mds/ceph-mds-a/keyring

# ✅ 启动一个 active 的 mds 实例
ceph fs set cephfs standby_count_wanted 1

echo "✅ CephFS 初始化完成！"
