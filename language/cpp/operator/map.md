### Map基本操作
* 创建 map
  > map<string, int> mp;
* 添加元素
  > mp["one"] = 1;
  > mp.insert(pair("four", 4));
* 获取已知元素的值
  > mp["one"];
* 判断元素是否存在在 map 中
  ```C++
    // 判断元素是否存在
    // count统计 key 出现的次数
    if (mp.count("x") > 0)
    {
        cout << "元素 x 存在" << endl;
    } else {
        cout << "元素 x 不存在" << endl;
    }
  ```
* Map遍历
  ```C++
  // 正向遍历
  for (auto iter = mp.begin(); iter != mp.end(); iter++)
  {
    cout << "key: " << iter->first << " Value: " << iter->second << endl;
  }
  
  // 逆序遍历
  for (auto riter = mp.rbegin(); iter != mp.rend(); riter++)
  {
    cout << "key: " << riter->first << " Value: " << riter->second << endl;
  }
  
  // for循环遍历
  for (const auto& iter: mp)
  {
    cout << "key: " << iter.first << " Value: " << iter.second << endl;
  }
  ```
* 删除元素
  ```C++
  // 按 key 删除
  mp.erase("one");
  
  // 按位置删除
  mp.erase(mp.find("one"));
  ```
* map 元素个数获取
  > mp.size();
* map 空判断
  > mp.empty();
* find 查找元素
  ```C++
  auto iter = mp.find("one");
  if (iter != mp.end())
  {
    cout << "元素存在"<< endl;
  }
  else
  {
    cout << "元素不存在" << endl;
  }
  ```
  