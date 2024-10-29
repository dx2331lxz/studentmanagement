<h1 align="center">Student Management System</h1>
<p align='center'>
<img src="https://img.shields.io/badge/made%20by-daoxuan-blue">
<img src="https://img.shields.io/badge/go-1.23.0-blue">
</p>
<img width="1374" alt="image" src="https://github.com/user-attachments/assets/8f7a2504-fefb-40be-afbf-adb81d88bd21">
<h2 align="center"><a  href="http://81.70.143.162:7879/">Live Demo</a></h2>

## Description
本项目为海大数据库作业，实现一个学生管理系统，按照课程要求，项目实现了对于学生的增删改查
## How to Start
首先确保你有go环境
1. 将代码clone到本地
2. 在项目跟目录运行命令,安装项目依赖
``` bash
go mod tidy
```
3. 配置数据库
找到数据库配置文件mysql.conf修改为自己的数据库
> [!IMPORTANT]
> 这里需要注意，本项目不包含数据库的创建和迁移部分
创建命令可以参考：
```
CREATE TABLE `Student` (
  `Sno` char(20) NOT NULL,
  `Sname` char(20) DEFAULT NULL,
  `Ssex` char(2) DEFAULT NULL,
  `Sage` smallint(6) DEFAULT NULL,
  `Sdept` char(20) DEFAULT NULL,
  PRIMARY KEY (`Sno`),
  UNIQUE KEY `Sname` (`Sname`)
)
```
5. 此时你可以直接使用如下命令进行编译
```
go build -o student
```
* student为可执行文件名称

6. 可以运行`./student`来运行项目
7. 当然你也可以使用bee工具
```
go install github.com/beego/bee/v2@latest
```
8. 然后在项目跟目录
```
bee run
```
## 贡献者

<a href="https://github.com/dx2331lxz/studentmanagement/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=dx2331lxz/studentmanagement" />
</a>

## LICENSE
[MIT](https://opensource.org/license/mit/)

  
  
