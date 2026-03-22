# Linux 常用命令详解

## 一、文件和目录操作

### 1. 基本操作

```bash
# 查看当前目录
pwd

# 列出文件和目录
ls              # 简单列表
ls -l           # 详细信息
ls -la          # 包括隐藏文件
ls -lh          # 人类可读的文件大小
ls -lt          # 按修改时间排序
ls -lS          # 按文件大小排序

# 切换目录
cd /path/to/dir     # 切换到指定目录
cd ~                # 切换到家目录
cd ..               # 返回上级目录
cd -                # 返回上次所在目录

# 创建目录
mkdir dirname           # 创建单个目录
mkdir -p dir1/dir2/dir3 # 递归创建多级目录

# 删除目录
rmdir dirname           # 删除空目录
rm -r dirname           # 递归删除目录及内容
rm -rf dirname          # 强制递归删除（危险！）

# 复制文件/目录
cp file1 file2          # 复制文件
cp -r dir1 dir2         # 递归复制目录
cp -p file1 file2       # 保留文件属性
cp -i file1 file2       # 交互式复制（覆盖前询问）

# 移动/重命名
mv file1 file2          # 重命名文件
mv file1 /path/to/dir/  # 移动文件
mv -i file1 file2       # 交互式移动

# 删除文件
rm file                 # 删除文件
rm -f file              # 强制删除
rm -i file              # 交互式删除

# 创建空文件或更新时间戳
touch filename

# 查看文件类型
file filename
```


## 二、文件查看和编辑

### 1. 查看文件内容

```bash
# 查看整个文件
cat filename            # 显示全部内容
cat -n filename         # 显示行号

# 分页查看
more filename           # 向下翻页（空格键）
less filename           # 上下翻页（推荐）

# 查看文件头部
head filename           # 默认显示前10行
head -n 20 filename     # 显示前20行

# 查看文件尾部
tail filename           # 默认显示后10行
tail -n 20 filename     # 显示后20行
tail -f filename        # 实时查看文件更新（常用于日志）
tail -f -n 100 filename # 显示最后100行并实时更新
```

### 2. 编辑配置文件

#### vim 编辑器

```bash
# 打开文件
vim filename

# vim 基本操作
i           # 进入插入模式
Esc         # 退出插入模式
:w          # 保存
:q          # 退出
:wq 或 :x   # 保存并退出
:q!         # 强制退出不保存
:set nu     # 显示行号
/keyword    # 搜索关键字
n           # 下一个搜索结果
N           # 上一个搜索结果
dd          # 删除当前行
yy          # 复制当前行
p           # 粘贴
u           # 撤销
Ctrl+r      # 重做
gg          # 跳到文件开头
G           # 跳到文件末尾
:行号       # 跳到指定行
```

#### nano 编辑器（更简单）

```bash
# 打开文件
nano filename

# nano 基本操作
Ctrl+O      # 保存
Ctrl+X      # 退出
Ctrl+W      # 搜索
Ctrl+K      # 剪切当前行
Ctrl+U      # 粘贴
Ctrl+G      # 帮助
```

## 三、搜索和查找

### 1. 查找文件

```bash
# find 命令
find /path -name "filename"         # 按名称查找
find /path -name "*.log"            # 使用通配符
find /path -type f                  # 查找文件
find /path -type d                  # 查找目录
find /path -size +100M              # 查找大于100M的文件
find /path -mtime -7                # 查找7天内修改的文件
find /path -user username           # 按用户查找
find /path -name "*.log" -exec rm {} \;  # 查找并删除

# locate 命令（更快，但需要更新数据库）
locate filename
updatedb                            # 更新数据库

# which 命令（查找命令位置）
which python
which java
```

### 2. 搜索文件内容

```bash
# grep 命令
grep "keyword" filename             # 搜索关键字
grep -i "keyword" filename          # 忽略大小写
grep -r "keyword" /path             # 递归搜索目录
grep -n "keyword" filename          # 显示行号
grep -v "keyword" filename          # 反向匹配（不包含）
grep -c "keyword" filename          # 统计匹配行数
grep -A 5 "keyword" filename        # 显示匹配行及后5行
grep -B 5 "keyword" filename        # 显示匹配行及前5行
grep -C 5 "keyword" filename        # 显示匹配行及前后5行
grep -E "pattern1|pattern2" file    # 使用正则表达式

# 常用组合
grep -rn "error" /var/log/          # 递归搜索日志中的错误
grep -i "exception" app.log | tail -20  # 搜索异常并显示最后20条
```

## 四、系统状态监控

### 1. CPU 监控

```bash
# 查看 CPU 信息
cat /proc/cpuinfo
lscpu

# 实时监控 CPU 使用率
top                     # 实时系统监控
top -u username         # 监控指定用户的进程
htop                    # 更友好的 top（需要安装）

# top 常用快捷键
P           # 按 CPU 使用率排序
M           # 按内存使用率排序
k           # 终止进程
q           # 退出

# 查看平均负载
uptime
w
```

### 2. 内存监控

```bash
# 查看内存使用情况
free -h                 # 人类可读格式
free -m                 # 以 MB 为单位

# 详细内存信息
cat /proc/meminfo

# 查看进程内存使用
ps aux --sort=-%mem | head -10      # 内存使用最多的10个进程
```

### 3. 磁盘监控

```bash
# 查看磁盘使用情况
df -h                   # 人类可读格式
df -i                   # 查看 inode 使用情况

# 查看目录大小
du -sh /path            # 显示目录总大小
du -h --max-depth=1 /path   # 显示一级子目录大小
du -sh * | sort -hr     # 按大小排序

# 查看磁盘 I/O
iostat                  # 需要安装 sysstat
iostat -x 1             # 每秒更新一次
```

### 4. 网络监控

```bash
# 查看网络接口
ifconfig                # 传统命令
ip addr                 # 新命令
ip a                    # 简写

# 查看网络连接
netstat -tuln           # 查看监听端口
netstat -tunp           # 查看所有连接和进程
ss -tuln                # 更快的 netstat 替代

# 查看路由表
route -n
ip route

# 测试网络连接
ping hostname
ping -c 4 hostname      # 发送4个包

# 查看网络流量
iftop                   # 实时流量监控（需要安装）
nethogs                 # 按进程显示流量（需要安装）
```


## 五、端口和网络连接

### 1. 查看端口监听

```bash
# 查看所有监听端口
netstat -tuln
ss -tuln
lsof -i -P -n | grep LISTEN

# 查看指定端口
netstat -tuln | grep :8080
ss -tuln | grep :8080
lsof -i :8080

# 查看端口被哪个进程占用
lsof -i :8080
netstat -tunlp | grep :8080
ss -tunlp | grep :8080
```

### 2. 查看网络连接状态

```bash
# 查看所有连接
netstat -an
ss -an

# 查看 TCP 连接
netstat -ant
ss -ant

# 查看 UDP 连接
netstat -anu
ss -anu

# 统计连接状态
netstat -an | awk '/^tcp/ {print $6}' | sort | uniq -c
ss -ant | awk '{print $1}' | sort | uniq -c

# 查看建立的连接
netstat -antp | grep ESTABLISHED
ss -antp | grep ESTABLISHED
```

## 六、进程管理

```bash
# 查看进程
ps aux                  # 查看所有进程
ps aux | grep nginx     # 查找特定进程
ps -ef                  # 另一种格式

# 查看进程树
pstree
pstree -p               # 显示 PID

# 实时监控进程
top
htop

# 终止进程
kill PID                # 正常终止
kill -9 PID             # 强制终止
killall process_name    # 按名称终止
pkill process_name      # 按名称终止

# 后台运行程序
command &               # 后台运行
nohup command &         # 后台运行且不受终端关闭影响
nohup command > output.log 2>&1 &   # 重定向输出

# 查看后台任务
jobs
fg %1                   # 将后台任务调到前台
bg %1                   # 继续后台任务
```

## 七、部署和启动程序

### 1. 服务管理（systemd）

```bash
# 启动服务
systemctl start service_name
systemctl start nginx

# 停止服务
systemctl stop service_name

# 重启服务
systemctl restart service_name

# 重新加载配置
systemctl reload service_name

# 查看服务状态
systemctl status service_name

# 开机自启
systemctl enable service_name

# 禁用开机自启
systemctl disable service_name

# 查看所有服务
systemctl list-units --type=service

# 查看失败的服务
systemctl --failed
```

### 2. 传统服务管理（service）

```bash
service nginx start
service nginx stop
service nginx restart
service nginx status
```

### 3. 部署应用程序

```bash
# 上传文件到服务器
scp local_file user@host:/remote/path
scp -r local_dir user@host:/remote/path

# 从服务器下载文件
scp user@host:/remote/file local_path

# 解压程序包
tar -xzf app.tar.gz
unzip app.zip

# 赋予执行权限
chmod +x app

# 运行程序
./app
nohup ./app > app.log 2>&1 &

# 查看程序是否运行
ps aux | grep app
netstat -tuln | grep :port
```

## 八、日志查看

### 1. 系统日志

```bash
# 查看系统日志
tail -f /var/log/syslog         # Debian/Ubuntu
tail -f /var/log/messages       # CentOS/RHEL

# 使用 journalctl（systemd）
journalctl                      # 查看所有日志
journalctl -u nginx             # 查看指定服务日志
journalctl -f                   # 实时查看日志
journalctl -n 100               # 查看最后100行
journalctl --since "1 hour ago" # 查看最近1小时的日志
journalctl --since "2024-01-01" # 查看指定日期后的日志
journalctl -p err               # 只看错误日志
```

### 2. 应用日志

```bash
# 实时查看日志
tail -f /var/log/nginx/access.log
tail -f /var/log/nginx/error.log
tail -f app.log

# 查看最后N行
tail -n 100 app.log

# 查看并搜索
tail -n 1000 app.log | grep "ERROR"
tail -f app.log | grep --line-buffered "ERROR"

# 多文件同时查看
tail -f /var/log/nginx/*.log

# 查看压缩日志
zcat app.log.gz | grep "ERROR"
zless app.log.gz
```

### 3. 搜索日志

```bash
# 搜索关键字
grep "ERROR" app.log
grep -i "error" app.log             # 忽略大小写
grep -r "ERROR" /var/log/           # 递归搜索

# 搜索并显示上下文
grep -C 5 "ERROR" app.log           # 前后5行
grep -A 5 "ERROR" app.log           # 后5行
grep -B 5 "ERROR" app.log           # 前5行

# 搜索多个关键字
grep -E "ERROR|WARN" app.log
grep "ERROR\|WARN" app.log

# 统计出现次数
grep -c "ERROR" app.log

# 搜索时间范围内的日志
sed -n '/2024-01-01 10:00/,/2024-01-01 11:00/p' app.log

# 使用 awk 过滤
awk '/ERROR/ {print}' app.log
awk '/2024-01-01.*ERROR/ {print}' app.log
```

## 九、压缩和解压

### 1. tar 命令

```bash
# 打包压缩（.tar.gz）
tar -czf archive.tar.gz files/      # 压缩目录
tar -czf archive.tar.gz file1 file2 # 压缩文件

# 解压
tar -xzf archive.tar.gz             # 解压到当前目录
tar -xzf archive.tar.gz -C /path    # 解压到指定目录

# 查看压缩包内容
tar -tzf archive.tar.gz

# 打包压缩（.tar.bz2）更高压缩率
tar -cjf archive.tar.bz2 files/

# 解压 .tar.bz2
tar -xjf archive.tar.bz2

# 仅打包不压缩
tar -cf archive.tar files/

# 解包
tar -xf archive.tar
```

### 2. zip/unzip 命令

```bash
# 压缩
zip archive.zip file1 file2
zip -r archive.zip directory/       # 递归压缩目录

# 解压
unzip archive.zip
unzip archive.zip -d /path          # 解压到指定目录

# 查看压缩包内容
unzip -l archive.zip

# 解压指定文件
unzip archive.zip file1
```

### 3. gzip/gunzip 命令

```bash
# 压缩（会删除原文件）
gzip file

# 保留原文件
gzip -c file > file.gz

# 解压
gunzip file.gz
gzip -d file.gz

# 查看压缩文件内容
zcat file.gz
zless file.gz
```


## 十、用户和权限管理

### 1. 用户管理

```bash
# 查看当前用户
whoami
id

# 查看所有用户
cat /etc/passwd
cut -d: -f1 /etc/passwd

# 添加用户
useradd username
useradd -m username             # 创建家目录
useradd -m -s /bin/bash username    # 指定 shell

# 设置密码
passwd username

# 删除用户
userdel username
userdel -r username             # 同时删除家目录

# 修改用户
usermod -aG groupname username  # 添加用户到组
usermod -s /bin/bash username   # 修改 shell

# 切换用户
su username
su - username                   # 切换并加载环境变量
sudo -u username command        # 以指定用户执行命令
```

### 2. 组管理

```bash
# 查看所有组
cat /etc/group

# 查看用户所属组
groups username
id username

# 创建组
groupadd groupname

# 删除组
groupdel groupname

# 将用户添加到组
usermod -aG groupname username
gpasswd -a username groupname

# 从组中删除用户
gpasswd -d username groupname
```

### 3. sudo 权限

```bash
# 编辑 sudo 配置
visudo

# 给用户 sudo 权限
usermod -aG sudo username       # Debian/Ubuntu
usermod -aG wheel username      # CentOS/RHEL

# 使用 sudo 执行命令
sudo command
sudo -i                         # 切换到 root
sudo su -                       # 切换到 root

# 查看 sudo 权限
sudo -l
```

## 十一、文件权限

### 1. 查看权限

```bash
# 查看文件权限
ls -l filename
ls -la                          # 包括隐藏文件

# 权限说明
# -rwxrwxrwx
# 第1位：文件类型（- 文件，d 目录，l 链接）
# 2-4位：所有者权限（r读 w写 x执行）
# 5-7位：组权限
# 8-10位：其他用户权限
```

### 2. 修改权限（chmod）

```bash
# 数字方式
chmod 755 file                  # rwxr-xr-x
chmod 644 file                  # rw-r--r--
chmod 777 file                  # rwxrwxrwx（不推荐）

# 权限数字对照
# r=4, w=2, x=1
# 7=4+2+1=rwx
# 6=4+2=rw-
# 5=4+1=r-x
# 4=4=r--

# 符号方式
chmod u+x file                  # 所有者添加执行权限
chmod g-w file                  # 组去除写权限
chmod o+r file                  # 其他用户添加读权限
chmod a+x file                  # 所有人添加执行权限
chmod u=rwx,g=rx,o=r file       # 设置具体权限

# 递归修改
chmod -R 755 directory/

# 常用权限设置
chmod 755 script.sh             # 可执行脚本
chmod 644 config.conf           # 配置文件
chmod 600 private.key           # 私钥文件
chmod 700 ~/.ssh                # SSH 目录
```

### 3. 修改所有者（chown）

```bash
# 修改所有者
chown username file
chown username:groupname file   # 同时修改所有者和组
chown -R username directory/    # 递归修改

# 只修改组
chgrp groupname file
chgrp -R groupname directory/

# 常用示例
chown www-data:www-data /var/www/html
chown -R nginx:nginx /usr/share/nginx/html
```

### 4. 特殊权限

```bash
# SUID（Set User ID）
chmod u+s file                  # 执行时以文件所有者身份运行
chmod 4755 file

# SGID（Set Group ID）
chmod g+s directory             # 目录中创建的文件继承目录的组
chmod 2755 directory

# Sticky Bit
chmod +t directory              # 只有文件所有者可以删除文件
chmod 1777 directory            # 如 /tmp
```

## 十二、Redis 命令行操作

### 1. 连接 Redis

```bash
# 连接本地 Redis
redis-cli

# 连接远程 Redis
redis-cli -h hostname -p 6379

# 带密码连接
redis-cli -h hostname -p 6379 -a password

# 连接后认证
redis-cli
AUTH password

# 选择数据库
SELECT 0
```

### 2. 基本操作

```bash
# 测试连接
PING

# 查看所有键
KEYS *
KEYS user:*

# 设置和获取
SET key value
GET key

# 删除键
DEL key

# 查看键类型
TYPE key

# 设置过期时间
EXPIRE key 3600
TTL key

# 查看数据库信息
INFO
INFO server
INFO memory

# 清空数据库
FLUSHDB                         # 清空当前数据库
FLUSHALL                        # 清空所有数据库

# 退出
EXIT
QUIT
```

### 3. 性能测试

```bash
# 基准测试
redis-benchmark
redis-benchmark -n 10000        # 执行10000次请求
redis-benchmark -q              # 简化输出

# 监控命令
redis-cli MONITOR

# 查看慢查询
redis-cli SLOWLOG GET 10
```

## 十三、MySQL 命令行操作

### 1. 连接 MySQL

```bash
# 连接本地 MySQL
mysql -u root -p

# 连接远程 MySQL
mysql -h hostname -u username -p

# 指定数据库连接
mysql -u username -p database_name

# 执行 SQL 文件
mysql -u username -p database_name < script.sql

# 导出数据库
mysqldump -u username -p database_name > backup.sql

# 导出所有数据库
mysqldump -u username -p --all-databases > all_backup.sql
```

### 2. 数据库操作

```sql
-- 查看所有数据库
SHOW DATABASES;

-- 创建数据库
CREATE DATABASE dbname;
CREATE DATABASE dbname CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 删除数据库
DROP DATABASE dbname;

-- 选择数据库
USE dbname;

-- 查看当前数据库
SELECT DATABASE();
```

### 3. 表操作

```sql
-- 查看所有表
SHOW TABLES;

-- 查看表结构
DESC tablename;
DESCRIBE tablename;
SHOW COLUMNS FROM tablename;

-- 查看建表语句
SHOW CREATE TABLE tablename;

-- 创建表
CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 删除表
DROP TABLE tablename;

-- 清空表
TRUNCATE TABLE tablename;
```

### 4. 数据操作

```sql
-- 插入数据
INSERT INTO users (name, email) VALUES ('张三', 'zhangsan@example.com');

-- 查询数据
SELECT * FROM users;
SELECT name, email FROM users WHERE id = 1;

-- 更新数据
UPDATE users SET email = 'new@example.com' WHERE id = 1;

-- 删除数据
DELETE FROM users WHERE id = 1;
```

### 5. 用户和权限

```sql
-- 查看所有用户
SELECT user, host FROM mysql.user;

-- 创建用户
CREATE USER 'username'@'localhost' IDENTIFIED BY 'password';
CREATE USER 'username'@'%' IDENTIFIED BY 'password';

-- 授予权限
GRANT ALL PRIVILEGES ON dbname.* TO 'username'@'localhost';
GRANT SELECT, INSERT ON dbname.* TO 'username'@'localhost';

-- 刷新权限
FLUSH PRIVILEGES;

-- 查看用户权限
SHOW GRANTS FOR 'username'@'localhost';

-- 撤销权限
REVOKE ALL PRIVILEGES ON dbname.* FROM 'username'@'localhost';

-- 删除用户
DROP USER 'username'@'localhost';

-- 修改密码
ALTER USER 'username'@'localhost' IDENTIFIED BY 'new_password';
```

### 6. 常用查询

```sql
-- 查看 MySQL 版本
SELECT VERSION();

-- 查看当前用户
SELECT USER();

-- 查看数据库大小
SELECT 
    table_schema AS 'Database',
    ROUND(SUM(data_length + index_length) / 1024 / 1024, 2) AS 'Size (MB)'
FROM information_schema.tables
GROUP BY table_schema;

-- 查看表大小
SELECT 
    table_name AS 'Table',
    ROUND(((data_length + index_length) / 1024 / 1024), 2) AS 'Size (MB)'
FROM information_schema.tables
WHERE table_schema = 'dbname'
ORDER BY (data_length + index_length) DESC;

-- 查看进程
SHOW PROCESSLIST;

-- 查看状态
SHOW STATUS;

-- 查看变量
SHOW VARIABLES;
SHOW VARIABLES LIKE 'max_connections';
```

### 7. 退出 MySQL

```sql
EXIT;
QUIT;
\q
```

## 十四、常用组合命令

```bash
# 查找并删除大文件
find /path -type f -size +100M -exec rm -i {} \;

# 查看目录下最大的10个文件
du -ah /path | sort -rh | head -10

# 批量重命名
for file in *.txt; do mv "$file" "${file%.txt}.bak"; done

# 查看端口占用并终止进程
lsof -ti:8080 | xargs kill -9

# 实时监控日志中的错误
tail -f app.log | grep --line-buffered "ERROR"

# 统计日志中各状态码数量
awk '{print $9}' access.log | sort | uniq -c | sort -rn

# 查找并压缩旧日志
find /var/log -name "*.log" -mtime +30 -exec gzip {} \;

# 批量修改文件权限
find /path -type f -exec chmod 644 {} \;
find /path -type d -exec chmod 755 {} \;
```

## 十五、实用技巧

```bash
# 命令历史
history                         # 查看命令历史
!100                            # 执行第100条命令
!!                              # 执行上一条命令
!$                              # 上一条命令的最后一个参数

# 快捷键
Ctrl+C                          # 终止当前命令
Ctrl+Z                          # 暂停当前命令
Ctrl+D                          # 退出当前 shell
Ctrl+L                          # 清屏
Ctrl+R                          # 搜索历史命令
Ctrl+A                          # 光标移到行首
Ctrl+E                          # 光标移到行尾

# 管道和重定向
command > file                  # 输出重定向到文件（覆盖）
command >> file                 # 输出追加到文件
command 2> file                 # 错误输出重定向
command > file 2>&1             # 标准输出和错误都重定向
command1 | command2             # 管道

# 后台运行
command &                       # 后台运行
nohup command &                 # 不挂断运行
screen                          # 会话管理
tmux                            # 终端复用器
```
