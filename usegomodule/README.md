
```
go mod init gomodule        # 產生相依關聯檔案
go mod tidy                 # 移除沒用到的套件
go clean -modcache			# 移除所有下載第三方套件的內容
git config --global url."https://xxxxx:xxxxx@gitlab.com" .insteadOf "https://gitlab.com"
export GOPRIVATE=gitlab.com  
```
