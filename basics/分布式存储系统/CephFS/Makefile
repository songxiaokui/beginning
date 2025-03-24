.PHONY: up down ps clean log init

ARCH ?= arm64

up:
	@echo "Starting CephFS server for architecture: $(ARCH)"
	@docker compose -f cephfs-server/docker-compose.$(ARCH).yml up -d

down:
	@echo "Starting CephFS server for architecture: $(ARCH)"
	@docker compose -f cephfs-server/docker-compose.$(ARCH).yml down -v
	@rm -rf cephfs-server/data cephfs-server/osd

init:
	@docker run --rm -it \
       -v $(CURDIR)/cephfs-server/data:/var/lib/ceph \
       -v $(CURDIR)/cephfs-server/osd:/var/lib/ceph/osd \
       -v $(CURDIR)/cephfs-server/init-cephfs.sh:/init.sh \
       -v $(CURDIR)/cephfs-server/ceph.conf:/etc/ceph/ceph.conf \
       --network host \
       quay.io/ceph/ceph:v17.2.6 \
       bash /init.sh
ps:
	@docker compose -f cephfs-server/docker-compose.$(ARCH).yml ps

clean:
	@docker compose -f cephfs-server/docker-compose.$(ARCH).yml down -v
	@rm -rf cephfs-server/data cephfs-server/osd

log:
	@echo "Starting CephFS server for all architectures"
	@docker ps -a
	@docker logs ceph-mon
	@docker logs ceph-mgr
	@docker logs ceph-osd
	@docker logs ceph-mds
