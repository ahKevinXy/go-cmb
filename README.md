# go-cmb



[![招商银行logo](./docs/static/img/cmb_logo.jpeg)](./docs/static/img/cmb_logo.jpeg)

版本变更日志参见 [`CHANGELOG.md`](CHANGELOG.md)

**特别感谢JetBrains对开源项目大力支持**

[![JetBrains Logo (Main) logo](https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg "logo")](https://jb.gg/OpenSourceSupport)

## 介绍

招商银行银企直联服务

## 分支说明

`v1` 版本基于  **2023年**招商银行之前对接文档开发 

`master` 基于**最新**招商银行银企直连官方文档开发

## 招商银行文档 [官网](https://openbiz.cmbchina.com/developer/UI/business/Index.aspx)


## 使用

```bash
go get -u github.com/ahKevinXy/go-cmb
```

## TODO 

- [ ] 更新 models 为 proto 兼容其他语言
- [ ] 优化请求参数
- [ ] 完善功能
- [ ]更新版本

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
* `models` 请求参数 和 返回参数
* `pkg` 第三方包
* `testdata` 本地测试数据


### 账务管理

### 交易管家


### 代发代扣


### 交易流水


## 开发调试

###  常用工具

[json 格式化工具 ](https://tool.lu/json/)

[json 转 go  struct ](https://mholt.github.io/json-to-go/)

[goland 注释工具 Goanno](https://plugins.jetbrains.com/plugin/14988-goanno)

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
const IsDebug = true // 判断是否是debug 模式 打印相关请求参数

```

## 接口清单

* 可经办业务模式查询  `DCLISMOD`
* 账户详细信息查询  `NTQACINF`
* 查询交易概要信息  `NTAGCINN`
* 电子回单查询 `DCAGPPDF`
* 批量查询余额 `NTQADINF`
* 新增记账子单元 `NTDMAADD`
* 查询记账子单元当天交易 `NTDMTQRD`
* 查询记账子单元历史交易 `NTDMTHLS`
* 查询记账子单元余额 `NTDMABAL`
* 记账子单元内部转账 `NTDMITRX`
* 超网代发其他 `NTAGCAPI`
* 查询交易明细信息 `NTAGDINF`
* 代发明细对账单处理结果查询请求 `DCTASKID`
* 查询记账子单元信息 `NTDMALST`
* 关闭子单元 `NTDMADLT`
* 设置记账子单元关联付款方信息 `NTDMARLT`
* 删除记账子单元关联付款方信息 `NTDMARLD`
* 修改记账子单元关联付款方信息 `NTDMATMN`
* 查询记账子单元关联付款信息 `NTDMARLQ`

### 账务查询（官方文档 2026-07-02）

| 接口 | Funcode | Handler | Request Body |
| --- | --- | --- | --- |
| 可经办业务模式查询 | `DCLISMOD` | `account.AccountNewDCLISMOD` | `models.AccountNewDCLISMODBody` |
| 查询可经办的账户列表 | `DCLISACC` | `account.AccountNewDCLISACC` | `models.AccountNewDCLISACCBody` |
| 结算账户详细信息查询 | `NTQACINF` | `account.AccountNewNTQACINF` | `models.AccountNewNTQACINFBody` |
| 查询分行号信息 | `NTACCBBK` | `account.AccountNewNTACCBBK` | `models.AccountNewNTACCBBKBody` |
| 非结算账户详细信息查询 | `QryNonSettleAccounts` | `account.AccountNewQryNonSettleAccounts` | `models.AccountNewQryNonSettleAccountsBody` |
| 综合账户信息查询 | `TotalAccountDirect` | `account.AccountNewTotalAccountDirect` | `models.AccountNewTotalAccountDirectBody` |
| 查询账户历史余额 | `NTQABINF` | `account.AccountNewNTQABINF` | `models.AccountNewNTQABINFBody` |
| 历史冻结余额等余额信息查询 | `NTQRYHSB` | `account.AccountNewNTQRYHSB` | `models.AccountNewNTQRYHSBBody` |
| 批量查询余额 | `NTQADINF` | `account.AccountNewNTQADINF` | `models.AccountNewNTQADINFBody` |
| 账户交易信息查询 | `trsQryByBreakPoint` | `account.AccountNewTrsQryByBreakPoint` | `models.AccountNewTrsQryByBreakPointBody` |
| 非结算户交易流水导出申请 | `NonSettleTranOpr` | `account.AccountNewNonSettleTranOpr` | `models.AccountNewNonSettleTranOprBody` |
| 查询关联非结算户交易流水 | `NonSettleTranQry` | `account.AccountNewNonSettleTranQry` | `models.AccountNewNonSettleTranQryBody` |
| PDF文件对账单获取 | `DCTRSPDF` | `account.AccountNewDCTRSPDF` | `models.AccountNewDCTRSPDFBody` |
| OFD文件对账单获取 | `issueBillOfd` | `account.AccountNewIssueBillOfd` | `models.AccountNewIssueBillOfdBody` |
| OFD文件对账单获取结果查询 | `queryBillOfd` | `account.AccountNewQueryBillOfd` | `models.AccountNewQueryBillOfdBody` |
| 单笔回单查询 | `DCSIGREC` | `account.AccountNewDCSIGREC` | `models.AccountNewDCSIGRECBody` |
| 电子回单异步查询 | `ASYCALHD` | `account.AccountNewASYCALHD` | `models.AccountNewASYCALHDBody` |
| 异步打印结果查询 | `DCTASKID` | `account.AccountNewDCTASKID` | `models.AccountNewDCTASKIDBody` |

### 代发代扣（官方文档 2026-07-02）

| 接口 | Funcode | Handler | Request Body |
| --- | --- | --- | --- |
| 卡BIN匹配识别标准银联卡 | `CdcMatchCardBin` | `payroll.PayrollNewCdcMatchCardBin` | `models.PayrollNewCdcMatchCardBinBody` |
| 代发经办 | `BB6BTHHL` | `payroll.PayrollNewBB6BTHHL` | `models.PayrollNewBB6BTHHLBody` |
| 批次与明细查询 | `BB6BPDQY` | `payroll.PayrollNewBB6BPDQY` | `models.PayrollNewBB6BPDQYBody` |
| 代发代扣综合查询 | `BB6INTQY` | `payroll.PayrollNewBB6INTQY` | `models.PayrollNewBB6INTQYBody` |
| 代发批次查询 | `BB6BTHQY` | `payroll.PayrollNewBB6BTHQY` | `models.PayrollNewBB6BTHQYBody` |
| 代发明细查询 | `BB6DTLQY` | `payroll.PayrollNewBB6DTLQY` | `models.PayrollNewBB6DTLQYBody` |
| 代发退票查询 | `BB6RFDQY` | `payroll.PayrollNewBB6RFDQY` | `models.PayrollNewBB6RFDQYBody` |
| 代发类型查询 | `BB6AGTQY` | `payroll.PayrollNewBB6AGTQY` | `models.PayrollNewBB6AGTQYBody` |
| 代扣经办 | `BB6ACLAK` | `payroll.PayrollNewBB6ACLAK` | `models.PayrollNewBB6ACLAKBody` |
| 代扣批次查询 | `BB6CDCQY` | `payroll.PayrollNewBB6CDCQY` | `models.PayrollNewBB6CDCQYBody` |
| 代扣明细查询 | `BB6ACLIF` | `payroll.PayrollNewBB6ACLIF` | `models.PayrollNewBB6ACLIFBody` |
| 代扣类型查询 | `BB6TRSQY` | `payroll.PayrollNewBB6TRSQY` | `models.PayrollNewBB6TRSQYBody` |
| 代发明细对账单查询请求 | `DCAGPPDF` | `payroll.PayrollNewDCAGPPDF` | `models.PayrollNewDCAGPPDFBody` |
| 代发明细对账单处理结果查询请求 | `DCTASKID` | `payroll.PayrollNewDCTASKID` | `models.PayrollNewDCTASKIDBody` |

