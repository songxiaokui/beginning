### 队列相关操作
> Queue是一种先进先出的队列
> 官方文档: https://cplusplus.com/reference/queue/queue/?kw=queue

* 基本操作
  - 创建队列
    ```C++
    queue<int> mq;
    ```
  - 插入元素
    ```C++
    mq.push(1);
    ```
  - 删除头部元素
    ```C++
    mq.pop();
    ```
  - 获取尾部元素
    ```C++
    mq.back();
    ```
  - 获取头部元素
    ```C++
    mq.front();
    ```
  - 获取元素个数
    ```C++
    mq.size();
    ```
  - 判断队列是否为空
    ```C++
    mq.empty();
    ```
  - 两个队列交换
    ```C++
    // 创建新队列
    queue<int> sq;
    mq.swap(sq);
    ```
  - 打印队列
    ```C++
    template<typename T>
    void print(queue<T> &q) {
    while (!q.empty()) {
    cout << q.front() << "->";
    q.pop();
    }
    cout << "End" << endl;
    }
    ```
