### **1. 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。**

```shell
root@robot:~# redis-benchmark -d 10 -t get,set
====== SET ======
  100000 requests completed in 0.87 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1

100.00% <= 0 milliseconds
115606.94 requests per second

====== GET ======
  100000 requests completed in 0.83 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1

100.00% <= 0 milliseconds
120192.30 requests per second

root@robot:~# redis-benchmark -d 20 -t get,set
====== SET ======
  100000 requests completed in 0.86 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1

100.00% <= 0 milliseconds
116414.43 requests per second

====== GET ======
  100000 requests completed in 0.84 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1

100.00% <= 0 milliseconds
119617.22 requests per second

root@robot:~# redis-benchmark -d 50 -t get,set
====== SET ======
  100000 requests completed in 0.89 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1

100.00% <= 0 milliseconds
112233.45 requests per second

====== GET ======
  100000 requests completed in 0.84 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1

100.00% <= 0 milliseconds
119189.52 requests per second

root@robot:~# redis-benchmark -d 100 -t get,set
====== SET ======
  100000 requests completed in 0.87 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1

99.84% <= 1 milliseconds
99.97% <= 2 milliseconds
100.00% <= 2 milliseconds
115606.94 requests per second

====== GET ======
  100000 requests completed in 0.86 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1

100.00% <= 0 milliseconds
116144.02 requests per second

root@robot:~# redis-benchmark -d 200 -t get,set
====== SET ======
  100000 requests completed in 0.87 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1

100.00% <= 0 milliseconds
114678.90 requests per second

====== GET ======
  100000 requests completed in 0.81 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1

100.00% <= 0 milliseconds
123152.71 requests per second

root@robot:~# redis-benchmark -d 1024 -t get,set
====== SET ======
  100000 requests completed in 0.89 seconds
  50 parallel clients
  1024 bytes payload
  keep alive: 1

99.95% <= 3 milliseconds
99.97% <= 4 milliseconds
100.00% <= 4 milliseconds
112739.57 requests per second

====== GET ======
  100000 requests completed in 0.80 seconds
  50 parallel clients
  1024 bytes payload
  keep alive: 1

100.00% <= 0 milliseconds
125786.16 requests per second

root@robot:~# redis-benchmark -d 5120 -t get,set
====== SET ======
  100000 requests completed in 0.92 seconds
  50 parallel clients
  5120 bytes payload
  keep alive: 1

100.00% <= 0 milliseconds
108342.37 requests per second

====== GET ======
  100000 requests completed in 0.83 seconds
  50 parallel clients
  5120 bytes payload
  keep alive: 1

100.00% <= 0 milliseconds
120336.95 requests per second

```



