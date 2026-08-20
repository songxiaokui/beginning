### 1、tmux 解决 SSH 断开后同步任务中断的问题

* 安装（ubuntu22.04，默认已经自带）

  ```shell
  sudo apt update
  sudo apt install tmux -y
  tmux -V
  ```
  
* 创建会话

  ```shell
  tmux new -s test_sleep
  
  # 在新的 session 中执行命令
  sleep 600s
  
  # 脱离会话
  Ctrl + b 然后按 d
  # [detached (from session test_sleep)]
  
  # 查询所有会话
  tmux ls
  
  # test_sleep: 1 windows (created Thu May 14 09:51:34 2026)
  
  # 进入会话
  tmux attach -t test_sleep
  ```
  
  