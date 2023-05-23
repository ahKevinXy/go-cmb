# go-cmb


[![招商银行logo](./docs/static/img/cmb_logo.jpeg)](./docs/static/img/cmb_logo.jpeg)

版本变更日志参见 [`CHANGELOG.md`](CHANGELOG.md)


[![JetBrains Logo (Main) logo](https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg "logo")](https://jb.gg/OpenSourceSupport)

## 介绍

招商银行银企直联服务

## 招商银行文档 [官网](https://openbiz.cmbchina.com/developer/UI/business/Index.aspx)


## 使用

```bash
go get -u github.com/ahKevinXy/go-cmb

```

## TODO 

- [ ] 更新 models 为 proto 兼容其他语言
- [ ] 优化请求参数
- [ ] 完善功能

### 目录说明

* `cmb_error` 错误参数
* `config`  配置文件
* `constants` 常量文件
* `docs` 文档文件
* `handler` 业务处理
    * `account` 账户处理
    * `notice` 通知
    * `payment` 支付
    * `payroll` 代发
    * `payrool_old` 代发(旧)
    * `unit_manager` 交易管家
* `help` 工具类函数
* `models` 实例
* `pkg` 第三方包
* `testdata` 本地测试数据


### 账务管理

### 交易管家


### 代发代扣


### 交易流水


## 开发调试

### 生成 testdata 参数数据
```bash
# 创建 testdata目录
 mkdir testdata
 
# 生成 testdata.go 文件
 cp -rf  testdata.go.template  testdata/testdata.go
 
```

### 录入本地测试数据

```go
package testdata

// 商户配置
const (
	UserId     = "" // 用户ID
	Account    = "" //账户
	UserKey    = "" // 用户秘钥
	AseKey     = "" // 对称秘钥
	BankLinkNo = "" // 分行号
)

// Sass平台配置
const (
	CmbSassName       = "" // Sass 名称
	CmbSassPrivateKey = "" // 平台私钥
	CmbSigdatDefault  = "" // 默认值
	CmbUrl            = "" // 银企直联地址
)

```

## 微信群

[![微信群](./docs/static/img/webchat.png)](./docs/static/img/webchat.png)