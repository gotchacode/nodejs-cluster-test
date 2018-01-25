## Benchmarking multicore Node.js


A very simple node.js server trying to benchmark single-core vs multi-core performance


### Benchmark single core

```
node singlecore.js
ab -n 10000 -c 100  http://127.0.0.1:8000/
```

### Benchmark Multi core

```
node multicore.js
ab -n 10000 -c 100  http://127.0.0.1:8000/
```
