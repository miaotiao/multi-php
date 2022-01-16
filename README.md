# php version manager

## usage

- pvm add <php_path> 增加一个 php 环境（本地已经下载好的）
- pvm ls 列出所有可用的 php 版本，前面使用 * 号 ，后面显示 php 版本（将存取的 php.txt 读取一边判断 php 是否真的存在）
- pvm use <php_version> 激活某个 php 版本

## uninstall
- 删除当前文件夹的所有内容

## todo
- 将自己加入 path
  - 卸载功能
    - 将自己从 path 里删除，将本文件夹内所有的内容都删除了
- 显示帮助优化
- 优化报错问题
- 可以获取 php 版本
    - pvm get <php_version> 获取某个 php 版本，从 搜狗 下载或者 php 官网
- pvm config
- pvm 升级自己功能
- 记录日志
- 卸载功能
- 刷新 path // https://stackoverflow.com/questions/17794507/powershell-reload-the-path-in-powershell