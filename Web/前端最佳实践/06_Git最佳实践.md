# Git 最佳实践

---

## 1. Commit 规范

```
<type>(<scope>): <subject>

feat: 新增用户登录功能
fix(auth): 修复token过期问题
docs: 更新API文档
style: 格式化代码（不影响功能）
refactor: 重构用户模块
perf: 性能优化
test: 添加单元测试
chore: 更新依赖/构建脚本
```

### 常用 type

| type | 说明 |
|------|------|
| feat | 新功能 |
| fix | bug修复 |
| docs | 文档 |
| style | 格式（不影响运行） |
| refactor | 重构 |
| perf | 性能优化 |
| test | 测试 |
| chore | 维护 |

---

## 2. 分支命名

```
feature/user-login        # 功能分支
feature/order-system

fix/token-expiry          # 修复分支
fix/login-validation

hotfix/critical-bug       # 紧急修复
hotfix/security-patch

release/v1.2.0           # 发布分支

develop                   # 开发主分支
main/master              # 生产分支
```

---

## 3. Git Flow 工作流

```
main/master    ─────●───────────────────●───────●──→ 生产环境
                    │                   ↑
                    │              release/v1.2
                    ↓                   │
develop    ●───●───●───●───●───●──●───●───●───●
                ↑           ↑       ↑
                │       feature/    │
         feature/        order-sys   │
         user-login                  │
                                      │
                              feature/payment
```

---

## 4. 日常操作

### 创建并切换分支
```bash
git checkout -b feature/user-login
git switch -c feature/user-login  # 现代写法
```

### 暂存修改
```bash
git add file.txt           # 单文件
git add src/               # 目录
git add .                  # 全部
git add -p                 # 部分暂存
```

### 提交
```bash
git commit -m "feat: add user login"
git commit -m "fix: resolve token issue"
```

### 拉取合并
```bash
git pull origin main                    # 拉取并合并
git pull --rebase origin main           # 变基合并（保持分支整洁）
```

### 推送
```bash
git push origin feature/user-login
```

---

## 5. 分支合并策略

```bash
# 切换到主分支
git checkout main

# 合并功能分支（推荐 --no-ff，保留合并历史）
git merge --no-ff feature/user-login

# 删除已合并分支
git branch -d feature/user-login
```

### 冲突处理
```bash
# 1. 拉取最新代码
git fetch origin

# 2. 变基到最新
git rebase origin/main

# 3. 解决冲突后
git add .
git rebase --continue

# 或放弃变基
git rebase --abort
```

---

## 6. 查看操作

```bash
git status                  # 查看状态
git log --oneline -10       # 最近10条提交
git log --graph --oneline   # 分支图
git diff                    # 工作区 vs 暂存区
git diff --cached           # 暂存区 vs 最新提交
git diff main..feature      # 两分支差异
git show <commit>           # 查看某次提交
```

---

## 7. 撤销操作

```bash
# 工作区撤销（未 git add）
git checkout -- file.txt
git restore file.txt        # 现代写法

# 暂存区撤回（已 git add）
git reset HEAD file.txt
git restore --staged file.txt

# 提交后撤销（未推送）
git commit --amend          # 修改提交信息
git reset --soft HEAD~1     # 保留更改在暂存区
git reset --mixed HEAD~1    # 保留更改在工作区
git reset --hard HEAD~1     # 丢弃更改（危险）

# 推送后撤销（已推送，需谨慎）
git revert HEAD             # 创建新提交逆转旧提交
```

---

## 8. 标签管理

```bash
# 创建标签
git tag v1.0.0
git tag -a v1.0.0 -m "版本1.0.0"

# 推送标签
git push origin v1.0.0
git push origin --tags      # 推送所有标签

# 查看标签
git tag -l
git show v1.0.0
```

---

## 9. 忽略文件

```bash
# .gitignore 示例
node_modules/
dist/
*.log
.env
.env.local
.DS_Store

# 已经被跟踪的文件需先移除
git rm --cached file.txt
```

---

## 10.  stash 暂存

```bash
git stash                   # 暂存当前修改
git stash pop               # 恢复并删除 stash
git stash apply             # 恢复但不删除
git stash list              # 查看 stash 列表
git stash drop              # 删除 stash
```

---

## 11. 常用别名

```bash
# ~/.gitconfig 配置
[alias]
    st = status
    co = checkout
    br = branch
    ci = commit
    unstage = reset HEAD --
    last = log -1 HEAD
    lg = log --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit
```

---

## 速查表

| 操作 | 命令 |
|------|------|
| 创建分支 | `git switch -c feature/x` |
| 暂存 | `git add .` |
| 提交 | `git commit -m "type: desc"` |
| 拉取合并 | `git pull --rebase origin main` |
| 推送 | `git push origin branch` |
| 合并 | `git merge --no-ff branch` |
| 查看状态 | `git status` |
| 查看日志 | `git log --oneline --graph` |
| 撤销工作区 | `git restore file` |
| 撤销暂存区 | `git restore --staged file` |
| stash | `git stash pop` |
