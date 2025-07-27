# .ide/Dockerfile

# 可将 golang 替换为需要的基础镜像
FROM golang:1.24

# 2. 定义你想要安装的 Node.js 主版本
#    例如：20 (Node.js 20.x), 18 (Node.js 18.x), lts (最新LTS), current (最新当前版本)
ARG NODE_MAJOR_VERSION=lts

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
RUN apt-get update && apt-get install -y git wget unzip openssh-server python3 python3-dev \
    curl -fsSL https://deb.nodesource.com/setup_${NODE_MAJOR_VERSION}.x | bash - && \
    apt-get install -y nodejs && \
    # 清理不再需要的包和缓存，进一步减小镜像大小
    apt-get autoremove -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# 指定字符集支持命令行输入中文（根据需要选择字符集）
ENV LANG C.UTF-8
ENV LANGUAGE C.UTF-8