### 架設環境

## Step 1
先安裝`gitlab`
default account root
```
docker run --detach \
  -p 80:80 \
  -p 12222:22 \
  --name gitlab \
  --restart always \
  --volume /etc/localtime:/etc/localtime:ro \
  --volume /Users/arieswang/Documents/gitlab/config:/etc/gitlab \
  --volume /Users/arieswang/Documents/gitlab/logs:/var/log/gitlab \
  --volume /Users/arieswang/Documents/gitlab/data:/var/opt/gitlab \
  gitlab/gitlab-ce:latest
```