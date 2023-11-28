# 使用 Node.js 镜像作为基础镜像
FROM node:14

# 设置工作目录
WORKDIR /app

# 将 package.json 和 package-lock.json 复制到工作目录
COPY package*.json ./

# 安装项目依赖
RUN npm install

# 将项目文件复制到工作目录
COPY . .

# 构建项目
RUN npm run build

# 使用 Nginx 镜像作为基础镜像
FROM nginx:stable-alpine

# 将构建的文件从构建阶段复制到 Nginx 镜像
COPY --from=0 /app/dist /usr/share/nginx/html

# 将 Nginx 配置文件复制到 Nginx 镜像
COPY ./nginx/default.conf /etc/nginx/conf.d/default.conf

# 暴露端口 8009
EXPOSE 8009

# 启动 Nginx
CMD ["nginx", "-g", "daemon off;"]

# docker network create my-network
# docker build -t vue-nginx .
# docker run -d --name=vue-nginx --network=my-network -p 8009:8009 vue-nginx

