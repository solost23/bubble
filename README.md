# bubble清单
一个基于gin+gorm开发的练手小项目，通过该项目可初识go web开发该有的姿势。
前端界面基于vue和lementUI开发，如果对前端不熟悉可以直接下载templates和static文件夹下的内容使用。
## 使用指南
### 下载
    git clone git@github.com:Solost23/bubble.git
### 配置MySQL
#### 1.在你的数据库中执行以下命令，创建本项目所用的数据库：
    CREATE DATABASE bubble DEFAULT CHARSET=utf8mb4;
#### 2.在config/config.yaml中配置数据库相关：
```yaml
connections:
  mysql:
    host: mysqldb
    user: root
    password: 123
    port: 3306
    db: bubble
    charset: utf8mb4
```

### 编译
    go build
### 执行
#### Mac/Unix:
    ./bubble
#### Windows:
    bubble.exe


## docker部署
```bash
docker build -t bubble:v1.0.0 .
```

```bash
docker run -d --rm -p 8080:8080 -v E:\Desktop\bubble:/app/bubble --name bubble --link MySQL:mysqldb bubble:v1.0.0
```

启动之后，使用浏览器打开http://127.0.0.1:8080/。
接口文档:http://127.0.0.1:8080/swagger/index.html
