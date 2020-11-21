### 架設環境

## Step 1
安裝`docker`
```
yum install -y yum-utils
yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo
yum install -y docker-ce docker-ce-cli containerd.io    
systemctl start docker
systemctl enable docker
```

## Step 2
安裝`gitlab`
default account root
```
docker run --detach \
  -p 80:80 \
  -p 12222:22 \
  --name gitlab \
  --restart always \
  --volume /etc/localtime:/etc/localtime:ro \
  --volume /root/gitlab/config:/etc/gitlab \
  --volume /root/gitlab/logs:/var/log/gitlab \
  --volume /root/gitlab/data:/var/opt/gitlab \
  gitlab/gitlab-ce:latest
```
