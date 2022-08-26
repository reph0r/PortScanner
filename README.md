# PortScanner
实现TCP半开式、全开式扫描
- 全开式扫描：已解决，可以正常使用
- 半开式扫描：有问题，无法进行正常扫描

后期合并全开式和半开式扫描代码
------

## 全开式扫描使用方法：
### 使用说明
+ 支持多IP输入，中间采用** , ** 分割即可
*


```go
go run main.go "8.142.189.231,114.114.114.114" 21,22,23,80-1024
```
![图片](https://user-images.githubusercontent.com/102449999/186863131-b671b8f2-b754-433f-b869-88ebae1bf579.png)

