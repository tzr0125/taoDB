
# 基于Bitcask的KV型数据库

## 数据结构
1. DB： 数据库主体
2. storage： 数据存储文件
3. index： 索引，存key对应value的存储位置
4. entry：基础KV单元


## 目前在做
1. 最基本的，实现storage顺序存日志，涉及到根据key、value生成entry，存放在storage，在index注册三个部分。

## 预期
1. 支持shell交互，从而真正变成一个数据库
2. 支持并发控制
3. 支持主从复制


## 模块梳理

