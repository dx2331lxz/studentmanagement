# 使用 Alpine 作为基础镜像
FROM alpine:latest

# 设置工作目录（可选）
WORKDIR /www/data

# 复制文件到容器中
COPY . /www/data

# 运行可执行文件
CMD ["./studentmanagement"]