#define FUSE_USE_VERSION 26
#include <fuse.h>
#include <cstring>
#include <map>
#include <string>
#include <vector>
#include <iostream>

// 模拟内存中的文件系统
static std::map<std::string, std::vector<char>> filesystem;

// 获取文件或目录的属性
static int my_getattr(const char *path, struct stat *stbuf) {
    std::cout << "my_getattr called for path: " << path << std::endl;
    memset(stbuf, 0, sizeof(struct stat));
    std::string spath(path);

    if (spath == "/") {
        stbuf->st_mode = S_IFDIR | 0755;
        stbuf->st_nlink = 2;
        return 0;
    } else if (filesystem.find(spath) != filesystem.end()) {
        stbuf->st_mode = S_IFREG | 0644;
        stbuf->st_nlink = 1;
        stbuf->st_size = filesystem[spath].size();
        return 0;
    }
    return -ENOENT;
}

// 读取目录内容
static int my_readdir(const char *path, void *buf, fuse_fill_dir_t filler, off_t offset,
                      struct fuse_file_info *fi) {
    std::cout << "my_readdir called for path: " << path << std::endl;
    if (strcmp(path, "/") != 0) {
        return -ENOENT;
    }

    filler(buf, ".", NULL, 0);
    filler(buf, "..", NULL, 0);

    for (const auto &entry : filesystem) {
        std::string filename = entry.first.substr(1);
        filler(buf, filename.c_str(), NULL, 0);
    }
    return 0;
}

// 创建文件（支持 touch）
static int my_mknod(const char *path, mode_t mode, dev_t dev) {
    std::cout << "my_mknod called for path: " << path << std::endl;
    std::string spath(path);

    // 只支持普通文件
    if (S_ISREG(mode)) {
        if (filesystem.find(spath) == filesystem.end()) {
            filesystem[spath] = std::vector<char>(); // 创建空文件
            std::cout << "File created via mknod: " << spath << std::endl;
            return 0;
        }
        return -EEXIST; // 文件已存在
    }
    return -EPERM; // 不支持其他类型（如设备文件）
}

// 打开文件
static int my_open(const char *path, struct fuse_file_info *fi) {
    std::cout << "my_open called for path: " << path << std::endl;
    std::string spath(path);
    if (filesystem.find(spath) == filesystem.end()) {
        if (fi->flags & O_CREAT) {
            filesystem[spath] = std::vector<char>();
            std::cout << "File created via open: " << spath << std::endl;
        } else {
            return -ENOENT;
        }
    }
    return 0;
}

// 读取文件内容
static int my_read(const char *path, char *buf, size_t size, off_t offset,
                   struct fuse_file_info *fi) {
    std::cout << "my_read called for path: " << path << " with size: " << size
              << " offset: " << offset << std::endl;
    std::string spath(path);
    if (filesystem.find(spath) == filesystem.end()) {
        return -ENOENT;
    }

    const std::vector<char> &data = filesystem[spath];
    if (static_cast<size_t>(offset) >= data.size()) {
        return 0;
    }

    size_t len = data.size() - static_cast<size_t>(offset);
    if (len > size) len = size;
    memcpy(buf, data.data() + offset, len);
    return len;
}

// 写入文件内容
static int my_write(const char *path, const char *buf, size_t size, off_t offset,
                    struct fuse_file_info *fi) {
    std::cout << "my_write called for path: " << path << " with size: " << size
              << " offset: " << offset << std::endl;
    std::string spath(path);
    if (filesystem.find(spath) == filesystem.end()) {
        return -ENOENT;
    }

    std::vector<char> &data = filesystem[spath];
    if (static_cast<size_t>(offset) + size > data.size()) {
        data.resize(static_cast<size_t>(offset) + size);
    }
    memcpy(data.data() + offset, buf, size);
    return size;
}

// 修改文件权限（支持 chmod）
static int my_chmod(const char *path, mode_t mode) {
    std::cout << "my_chmod called for path: " << path << " with mode: " << mode << std::endl;
    std::string spath(path);

    if (spath == "/") {
        // 根目录权限不修改，直接返回成功（简单实现）
        return 0;
    } else if (filesystem.find(spath) != filesystem.end()) {
        // 文件存在，模拟修改权限（当前内存实现不实际存储 mode）
        return 0;
    }
    return -ENOENT;
}

// FUSE 操作结构体
static struct fuse_operations my_operations = {
    .getattr = my_getattr,
    .readlink = NULL,
    .getdir = NULL,
    .mknod = my_mknod,        // 新增：支持文件创建
    .mkdir = NULL,
    .unlink = NULL,
    .rmdir = NULL,
    .symlink = NULL,
    .rename = NULL,
    .link = NULL,
    .chmod = my_chmod,        // 新增：支持权限修改
    .chown = NULL,
    .truncate = NULL,
    .utime = NULL,
    .open = my_open,
    .read = my_read,
    .write = my_write,
    .statfs = NULL,
    .flush = NULL,
    .release = NULL,
    .fsync = NULL,
    .readdir = my_readdir,
};

// 主函数
int main(int argc, char *argv[]) {
    std::cout << "Starting FUSE filesystem..." << std::endl;
    return fuse_main(argc, argv, &my_operations, NULL);
}