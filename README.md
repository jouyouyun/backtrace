![](https://github.com/zhanghanyun/backtrace/workflows/Go/badge.svg?branch=main)
![GitHub](https://img.shields.io/github/license/zhanghanyun/backtrace?color=blueviolet)
![](https://tokei.rs/b1/github/zhanghanyun/backtrace?category=code)


## 三网回程路由测试

``` shell
$ sudo ./backtrace
2023/01/29 17:19:53 项目地址：github.com/jouyouyun/backtrace
2023/01/29 17:19:53 正在测试路由...
2023/01/29 17:19:54 [北京电信]	 219.141.136.12 	 电信CN2 [优质线路]
2023/01/29 17:19:54 [北京联通]	 202.106.50.1   	 电信CN2 [优质线路]
2023/01/29 17:19:54 [北京移动]	 221.179.155.161	 电信CN2 [优质线路]
2023/01/29 17:19:54 [上海电信]	 202.96.209.133 	 电信CN2 [优质线路]
2023/01/29 17:19:54 [上海联通]	 210.22.97.1    	 电信CN2 [优质线路]
2023/01/29 17:19:54 [上海移动]	 211.136.112.200	 电信CN2 [优质线路]
2023/01/29 17:19:54 [广州电信]	 58.60.188.222  	 电信CN2 [优质线路]
2023/01/29 17:19:54 [广州联通]	 210.21.196.6   	 电信CN2 [优质线路]
2023/01/29 17:19:54 [广州移动]	 120.196.165.24 	 电信CN2 [优质线路]
2023/01/29 17:19:54 [成都电信]	 61.139.2.69    	 电信CN2 [优质线路]
2023/01/29 17:19:54 [成都联通]	 119.6.6.6      	 电信CN2 [优质线路]
2023/01/29 17:19:54 [成都移动]	 211.137.96.205 	 电信CN2 [优质线路]
```

## 去程路由测试

``` shell
$ sudo ./backtrace -t idc -n no
2023/01/29 17:17:14 项目地址：github.com/jouyouyun/backtrace
2023/01/29 17:17:14 正在测试路由...
2023/01/29 17:17:15 [野草云 HK BGP]	 154.204.56.254 	 电信163 [普通线路]
2023/01/29 17:17:15 [野草云 HK CTGNet+BGP]	 193.134.211.254	 电信CTGNet [优质线路]
2023/01/29 17:17:15 [UUUVPS SanJose CN2]	 45.205.1.225   	 电信CN2 [优质线路]
2023/01/29 17:17:15 [AkkoCloud SanJose CN2 GIA]	 45.9.10.146    	 电信CN2 [优质线路]
2023/01/29 17:17:15 [AkkoCloud London CN2 GIA]	 185.217.110.10 	 电信CN2 [优质线路]
```

## 使用
终端下运行
```shell
curl https://raw.githubusercontent.com/zhanghanyun/backtrace/main/install.sh -sSf | sh
```
