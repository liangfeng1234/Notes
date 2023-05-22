## 实验2-1 openMP程序的编译和运行 

### Linux下编译和运行

```
[root@VM-4-14-centos openMP]# gcc -fopenmp -O2 -o hellomp.out hellomp.c
[root@VM-4-14-centos openMP]# ./hellomp.out
Hello World from OMP thread 0
Number of threads is 8
Hello World from OMP thread 7
Hello World from OMP thread 1
Hello World from OMP thread 2
Hello World from OMP thread 3
Hello World from OMP thread 4
Hello World from OMP thread 5
Hello World from OMP thread 6
```

```
[root@VM-4-14-centos openMP]# OMP_NUM_THREADS=10
[root@VM-4-14-centos openMP]# export OMP_NUM_THREADS
[root@VM-4-14-centos openMP]# vim hellomp.c
[root@VM-4-14-centos openMP]# gcc -fopenmp -O2 -o hellomp.out hellomp.c
[root@VM-4-14-centos openMP]# ./hellomp.out
Hello World from OMP thread 7
Hello World from OMP thread 1
Hello World from OMP thread 0
Number of threads is 10
Hello World from OMP thread 2
Hello World from OMP thread 3
Hello World from OMP thread 4
Hello World from OMP thread 8
Hello World from OMP thread 9
Hello World from OMP thread 5
Hello World from OMP thread 6
```

### windows下编译和运行

```
D:\C++Projects\openMP\cmake-build-debug\openMP.exe
Hello World from OMP thread 1
Hello World from OMP thread 0
Number of threads is 8       
Hello World from OMP thread 3
Hello World from OMP thread 4
Hello World from OMP thread 5
Hello World from OMP thread 6
Hello World from OMP thread 7
Hello World from OMP thread 2
```

## 实验2-2 矩阵乘法的openMP实现和性能分析

## 2-2-1 矩阵乘法

### Linux下编译和运行

我的Linux只有2核2G，可以看到使用并行计算反而比线性计算性能更低

```
[root@VM-4-14-centos openMP]# gcc -fopenmp -O2 -o matrixTest.out matrixTest.c 
[root@VM-4-14-centos openMP]# ./matrixTest.out 
Thread_pnum = 2
serial matrix multiply time: 53.810000
parell matrix multiply time: 124.680000
```

### windows下编译和运行

我的windows是8核16G，并行计算的速度是线性计算的3倍左右

```
Thread_pnum = 8
serial matrix multiply time: 39.106000
parell matrix multiply time: 12.698000
```

## 求解圆周率pi

### 串行程序

```C++
#include <iostream>
#include <chrono>

using namespace std;

static long num_step = 100000000;
double step;
int main (int argc, char *argv[])
{
  int i;
  double x, pi, sum = 0.0;
  step = 1.0/(double) num_step;
  chrono::milliseconds start_time = chrono::duration_cast<chrono::milliseconds>(chrono::system_clock::now().time_since_epoch());
  for (i = 0; i < num_step; ++i) {
    x = (i + 0.5) * step;
    // x = (i + 0.4) * step;
    sum = sum + 4.0/(1.0 + x * x);
  }
  chrono::milliseconds end_time = chrono::duration_cast<chrono::milliseconds>(chrono::system_clock::now().time_since_epoch());
  pi = step * sum;
  
  std::cout <<"The result: "<< pi << std::endl;
  cout << "cost time: " <<  chrono::milliseconds(end_time).count() - chrono::milliseconds(start_time).count() << "ms" << endl;
  return 0;
}
```

执行结果：

```
D:\C++Projects\openMP_C++\cmake-build-debug\openMP_C__.exe
The result: 3.14159
cost time: 369ms
```

### 并行程序

```
#include <iostream>
#include <chrono>
#include <omp.h>

using namespace std;

static long num_step = 100000000;
double step;
#define NUM_THREADS 8
int main (int argc, char *argv[])
{

    double pi = 0.0, sum = 0.0;
    step = 1.0/(double) num_step;
    int number;


    chrono::milliseconds start_time = chrono::duration_cast<chrono::milliseconds>(chrono::system_clock::now().time_since_epoch());
    omp_set_num_threads(NUM_THREADS);
#pragma omp parallel
    {
        int i, id, nthreads;
        double x, partial_sum;
        id = omp_get_thread_num();
        nthreads = omp_get_num_threads();
        if (id == 0) {
            number = nthreads;
        }
        partial_sum = 0.0;
#pragma omp for
        for (i = 0; i < num_step; ++i) {
            x = (i + 0.5) * step;
            partial_sum += 4.0/(1.0 + x * x);
        }
#pragma omp critical
        sum += partial_sum;
    }
    pi = step * sum;

    chrono::milliseconds end_time = chrono::duration_cast<chrono::milliseconds>(chrono::system_clock::now().time_since_epoch());

    std::cout <<"The result: "<< pi << std::endl;
    cout << "cost time: " <<  chrono::milliseconds(end_time).count() - chrono::milliseconds(start_time).count() << "ms" << endl;
    return 0;
}
```

执行结果

```
D:\C++Projects\openMP_C++\cmake-build-debug\openMP_C__.exe
The result: 3.14159
cost time: 84ms
```

### 性能比较

并行计算的速度是串行计算的4.3倍左右

