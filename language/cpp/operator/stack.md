### 栈相关的操作
```C++
// 定义
stack<int> st;

// 插入元素
st.push(1);
st.push(2);
st.push(3);

// 获取栈顶元素
cout << "栈顶元素: " << st.top() << endl;

// 删除栈顶元素
st.pop();

// 获取栈元素个数
cout << "栈元素个数: " << st.size() << endl;

// 判断栈是否为空
cout << st.empty() << endl;

// 插入元素
st.emplace(3);

// 交换两个栈
stack<int> st2;
st2.push(7);
st2.push(8);
st2.push(9);
st.swap(st2);
```