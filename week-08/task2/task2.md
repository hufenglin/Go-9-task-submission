####  Before 插入之前

```
# Memory
...
used_memory:1402440
used_memory_human:1.34M
used_memory_rss:11956224
used_memory_rss_human:11.40M
used_memory_peak:414964784
used_memory_peak_human:395.74M
used_memory_peak_perc:0.34%
used_memory_overhead:851120
used_memory_startup:809824
used_memory_dataset:551320
used_memory_dataset_perc:93.03%
...
```



#### 插入 100000 条数据（每个数据大小为 100 bytes) 后的 memory 信息

```
# Memory
...
used_memory:18726696
used_memory_human:17.86M
used_memory_rss:29011968
used_memory_rss_human:27.67M
used_memory_peak:414964784
used_memory_peak_human:395.74M
used_memory_peak_perc:4.51%
used_memory_overhead:5375368
used_memory_startup:809824
used_memory_dataset:13351328
used_memory_dataset_perc:74.52%
...
```



每个 key 的平均占用内存空间 (18726696 - 1402440) / 100000 - 100 = 73.2(byte)

#### 插入 100000 条数据（每个数据大小为 5120 bytes) 后的 memory 信息

```
# Memory
...
used_memory:622451136
used_memory_human:593.62M
used_memory_rss:637300736
used_memory_rss_human:607.78M
used_memory_peak:622472136
used_memory_peak_human:593.64M
used_memory_peak_perc:100.00%
used_memory_overhead:5899656
used_memory_startup:809824
used_memory_dataset:616551480
used_memory_dataset_perc:99.18%
...
```



每个 key 的平均占用内存空间 (622451136 - 1402440) / 100000 - 5120 = 1090.5(byte)

| 指标                     | 含义                                                         |
| ------------------------ | ------------------------------------------------------------ |
| used_memory              | 由 Redis 分配器分配的内存总量，以字节（byte）为单位，即当前redis使用内存大小。 |
| used_memory_human        | 已更直观的单位展示分配的内存总量。                           |
| used_memory_rss          | 向操作系统申请的内存大小，即redis使用的物理内存大小。        |
| used_memory_rss_human    | 已更直观的单位展示向操作系统申请的内存大小。                 |
| used_memory_peak         | redis的内存消耗峰值，以字节为单位，即历史使用记录中redis使用内存峰值。 |
| used_memory_peak_human   | 以更直观的格式返回redis的内存消耗峰值                        |
| used_memory_peak_perc    | 使用内存达到峰值内存的百分比                                 |
| used_memory_overhead     | Redis为了维护数据集的内部机制所需的内存开销，包括所有客户端输出缓冲区、查询缓冲区、AOF重写缓冲区和主从复制的backlog。 |
| used_memory_startup      | Redis服务器启动时消耗的内存                                  |
| used_memory_dataset      | 数据实际占用的内存大小，即 used_memory-used_memory_overhead  |
| used_memory_dataset_perc | 数据占用的内存大小的百分比                                   |