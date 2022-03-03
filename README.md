# GO Datastruct

使用 Go 实现常用的数据结构，单元测试覆盖率100%。

## 使用

```bash
git clone git@github.com:someu/go-datastruct.git
```

```bash
# 运行测试
go test -v -coverprofile=cover.out
# 查看单元测试覆盖率
go tool cover -html=cover.out -o cover.html

```

## 数据结构

1. 链表：单向链表、双向链表、循环链表
2. 栈
3. 队列
4. 树：二叉搜索树、平衡二叉树