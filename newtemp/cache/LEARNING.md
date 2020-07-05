## 缓存

局部性原理 举例：
mysql 
innodb 引擎
页形式 组织 数据， 数据页 = > 内存







* 多级缓存

cpu 和内存之间的 瓶颈



* 缓存淘汰策略

LRU Least Recently used

mysql lru 的变体：

3/8 旧数据，其余热数据



Redis  随机 选 5个 比对内部时钟 进行淘汰


Lrf: 
Least Frequently Used(LFU)淘汰一定时期内被访问次数最少的页面，以次数作为参考


* 缓存安全