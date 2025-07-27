# .ide/Dockerfile

# 可将 golang 替换为需要的基础镜像
FROM golang:1.24

# 定义你想要安装的版本
ARG NODE_VERSION=22.17.1

# 安装 code-server 和 vscode 常用插件
RUN curl -fsSL https://code-server.dev/install.sh | sh \
  && code-server --install-extension cnbcool.cnb-welcome \
  && code-server --install-extension redhat.vscode-yaml \
  && code-server --install-extension dbaeumer.vscode-eslint \
  && code-server --install-extension waderyan.gitblame \
  && code-server --install-extension mhutchie.git-graph \
  && code-server --install-extension donjayamanne.githistory \
  && code-server --install-extension golang.go \
  && code-server --install-extension ms-python.python \
  && code-server --install-extension tencent-cloud.coding-copilot \
  && echo done
  
# 安装 ssh 服务，用于支持 VSCode 等客户端通过 Remote-SSH 访问开发环境（也可按需安装其他软件）
RUN apt-get update && apt-get install -y git wget unzip openssh-server python3 python3-dev curl && \
    wget https://nodejs.org/dist/v${NODE_VERSION}/node-v${NODE_VERSION}-linux-x64.tar.gz && \
    tar -xzf node-v${NODE_VERSION}-linux-x64.tar.gz -C /usr/local --strip-components=1 && \
    rm node-v${NODE_VERSION}-linux-x64.tar.gz

# 指定字符集支持命令行输入中文（根据需要选择字符集）
ENV LANG C.UTF-8
ENV LANGUAGE C.UTF-8