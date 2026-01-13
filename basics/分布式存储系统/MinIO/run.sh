docker rm -f minio-storage

mkdir -pv ./.rootfs

docker run -p 9001:9000 -p 9002:9002 \
  --name "minio-storage" \
  -itd \
  --restart=always \
  -e "MINIO_ROOT_USER=AdminSXK" \
  -e "MINIO_ROOT_PASSWORD=Dgjf993921abc" \
  -v ./.rootfs:/data \
  quay.io/minio/minio server /data --console-address ":9002"