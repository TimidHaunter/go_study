##初始化文件夹
go mod init gin

##安装gin
go get -u github.com/gin-gonic/gin

##查看Go env配置项
go env

##修改env配置项
go env -w GO111MODULE="on"
go env -w GOPROXY="https://goproxy.io,direct"