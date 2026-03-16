# Git 基本操作文档

## 1. Git 简介

Git 是一个分布式版本控制系统，用于跟踪代码变更、协作开发和管理项目历史。

## 2. 基本配置

首次使用 Git 需要配置用户信息：

```bash
# 配置用户名
git config --global user.name "你的名字"

# 配置邮箱
git config --global user.email "your.email@example.com"

# 查看配置
git config --list
```

## 3. 初始化仓库

### 3.1 创建新仓库

```bash
# 在当前目录初始化 Git 仓库
git init

# 在指定目录初始化仓库
git init <目录名>

# 初始化并指定默认分支名
git init -b main
```

### 3.2 配置远程仓库地址

```bash
# 添加远程仓库
git remote add origin <仓库URL>

# 示例
git remote add origin https://github.com/username/repository.git

# 查看远程仓库
git remote -v

# 修改远程仓库地址
git remote set-url origin <新的仓库URL>

# 删除远程仓库
git remote remove origin

# 重命名远程仓库
git remote rename origin new-origin
```

### 3.3 首次推送到远程

```bash
# 添加文件
git add .

# 提交
git commit -m "Initial commit"

# 推送到远程并设置上游分支
git push -u origin main
```

## 4. Clone（克隆仓库）

从远程仓库克隆项目到本地：

```bash
# 克隆远程仓库
git clone <仓库URL>

# 示例
git clone https://github.com/username/repository.git

# 克隆到指定目录
git clone <仓库URL> <目录名>

# 浅克隆（只克隆最近的提交历史）
git clone --depth 1 <仓库URL>

# 浅克隆指定深度
git clone --depth <数量> <仓库URL>

# 克隆指定分支
git clone -b <分支名> <仓库URL>

# 浅克隆指定分支
git clone -b <分支名> --depth 1 <仓库URL>

# 递归克隆（包含子模块）
git clone --recursive <仓库URL>
```

### 4.1 浅克隆的优势

- 下载速度更快
- 占用磁盘空间更少
- 适合只需要最新代码的场景（如 CI/CD）

### 4.2 将浅克隆转换为完整克隆

```bash
# 获取完整历史
git fetch --unshallow

# 或者
git pull --unshallow
```

## 5. Commit（提交更改）

### 5.1 查看状态

```bash
# 查看工作区状态
git status

# 查看简洁状态
git status -s
```

### 5.2 添加文件到暂存区

```bash
# 添加指定文件
git add <文件名>

# 添加所有更改
git add .

# 添加所有 .txt 文件
git add *.txt
```

### 5.3 提交更改

```bash
# 提交暂存区的更改
git commit -m "提交说明"

# 添加并提交（跳过 git add）
git commit -am "提交说明"

# 修改最后一次提交
git commit --amend
```

### 5.4 查看提交历史

```bash
# 查看提交历史
git log

# 查看简洁历史
git log --oneline

# 查看图形化历史
git log --graph --oneline --all
```

## 6. Branch（分支管理）

### 6.1 查看分支

```bash
# 查看本地分支
git branch

# 查看所有分支（包括远程）
git branch -a

# 查看远程分支
git branch -r
```

### 6.2 创建分支

```bash
# 创建新分支
git branch <分支名>

# 创建并切换到新分支
git checkout -b <分支名>

# 或使用新命令
git switch -c <分支名>
```

### 6.3 切换分支

```bash
# 切换到指定分支
git checkout <分支名>

# 或使用新命令
git switch <分支名>
```

### 6.4 删除分支

```bash
# 删除本地分支
git branch -d <分支名>

# 强制删除分支
git branch -D <分支名>

# 删除远程分支
git push origin --delete <分支名>
```

## 7. Merge（合并分支）

### 7.1 合并分支

```bash
# 将指定分支合并到当前分支
git merge <分支名>

# 示例：将 feature 分支合并到 main
git checkout main
git merge feature
```

### 7.2 处理合并冲突

当合并出现冲突时：

1. 查看冲突文件：`git status`
2. 手动编辑冲突文件，解决冲突标记（`<<<<<<<`, `=======`, `>>>>>>>`）
3. 添加解决后的文件：`git add <文件名>`
4. 完成合并：`git commit`

### 7.3 取消合并

```bash
# 取消合并
git merge --abort
```

## 8. 远程仓库操作

### 8.1 查看远程仓库

```bash
# 查看远程仓库
git remote -v

# 查看远程仓库详细信息
git remote show origin
```

### 8.2 推送到远程

```bash
# 推送到远程分支
git push origin <分支名>

# 推送当前分支
git push

# 首次推送并设置上游分支
git push -u origin <分支名>
```

### 8.3 拉取远程更新

```bash
# 拉取并合并
git pull

# 拉取指定分支
git pull origin <分支名>

# 仅拉取不合并
git fetch
```

## 9. 常用工作流程

### 9.1 从零开始创建项目流程

```bash
# 1. 创建项目目录
mkdir my-project
cd my-project

# 2. 初始化 Git 仓库
git init

# 3. 创建文件
echo "# My Project" > README.md

# 4. 添加并提交
git add .
git commit -m "Initial commit"

# 5. 在 GitHub/GitLab 创建远程仓库后，添加远程地址
git remote add origin https://github.com/username/my-project.git

# 6. 推送到远程
git push -u origin main
```

### 9.2 日常开发流程

```bash
# 1. 克隆项目
git clone <仓库URL>

# 2. 创建功能分支
git checkout -b feature/new-feature

# 3. 进行开发并提交
git add .
git commit -m "添加新功能"

# 4. 推送到远程
git push -u origin feature/new-feature

# 5. 切换回主分支
git checkout main

# 6. 合并功能分支
git merge feature/new-feature

# 7. 推送主分支
git push origin main
```

### 9.3 协作开发流程

```bash
# 1. 更新本地代码
git pull origin main

# 2. 创建功能分支
git checkout -b feature/my-feature

# 3. 开发并提交
git add .
git commit -m "完成功能开发"

# 4. 推送分支
git push -u origin feature/my-feature

# 5. 在 GitHub/GitLab 创建 Pull Request/Merge Request
# 6. 代码审查通过后合并
```

## 10. 其他常用命令

```bash
# 查看文件差异
git diff

# 撤销工作区修改
git checkout -- <文件名>

# 撤销暂存区文件
git reset HEAD <文件名>

# 回退到指定提交
git reset --hard <commit-id>

# 暂存当前工作
git stash

# 恢复暂存的工作
git stash pop

# 查看暂存列表
git stash list
```

## 11. 最佳实践

1. **频繁提交**：保持小而频繁的提交，便于追踪和回滚
2. **清晰的提交信息**：使用有意义的提交说明
3. **使用分支**：为每个功能或修复创建独立分支
4. **定期同步**：经常从主分支拉取更新
5. **代码审查**：使用 Pull Request 进行代码审查
6. **不提交敏感信息**：使用 `.gitignore` 排除敏感文件

## 12. .gitignore 示例

```gitignore
# 依赖目录
node_modules/
vendor/

# 编译输出
*.exe
*.dll
*.so
*.dylib

# 日志文件
*.log

# 环境配置
.env
.env.local

# IDE 配置
.vscode/
.idea/
*.swp
*.swo

# 操作系统文件
.DS_Store
Thumbs.db
```
