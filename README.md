# srv
#简单的游戏服务器

#各文件夹介绍
* admin 配置脚本
* bin 可执行文件
* conf 服务器配置文件, 暂时只有mainserver和mysql是起作用的
* log 游戏运行日志
* pkg 编译生成的中间文件
* script 启动脚本等
* src 源代码, 各文件作用, 见注释

#编译
./admin make生成main_server和mysql两个可执行文件
* mysql 初始化数据库
* main_server 游戏主服务
