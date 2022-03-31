
yellow-onekey-install-apk
----

## 介绍
golang写的可以一键把电脑上的几个apk安装到手机上的小工具

（第一个版本，比较粗糙、暴力）

## 使用

`yellow-onekey-install-apk.exe -adb "D:\adb.exe" -apk "D:\xxx.apk"`

|arg|必填|description|
|-|-|-|
|-adb|必填|`adb.exe`文件地址，没有需要去下载|
|-apk|必填|`apk`文件地址|
|-apk2|选填|第二个`apk`文件地址|
|-apk3|选填|第三个`apk`文件地址|


## 原理
借助 adb 的 `adb -s xxx install xx.apk` 命令实现的