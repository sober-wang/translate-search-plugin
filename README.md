# deepin 全局搜索翻译插件

用于使用全局搜索功能翻译简单的词语与词组

依赖：
Golang version: 1.26.4

安装方式

```shell
cd build 
chmod +x install.sh 
killall dde-grand-search-daemon
```


# 路线目标
- [x] 联网调用翻译链接API
- [ ] deb 打包构建
- [ ] 离线词库依赖诸如：kd, clitrans 等工具
- [ ] 离线词库无依赖版本。
- [ ] AI翻译，支持配置 MCP 接口
- [ ] 多路并行，谁先返回展示谁的结果。


# 鸣谢

[bing_api参考](https://github.com/Shawyeok/bing-dict)

[deepin官方计算器插件样例](https://github.com/linuxdeepin/dde-grand-search/tree/master/examples/calculator-search-plugin)

[dde-grand-search plugin api 文档](https://github.com/linuxdeepin/dde-grand-search/blob/master/docs/plugin-development-guide.md)


## License

translate-search-plugin for dde-grand-search is licensed under [GPL-3.0-only](LICENSE)
