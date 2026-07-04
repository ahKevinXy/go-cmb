# 账户管理

## 新版账务查询接口（官方文档 2026-07-02）

来源：https://openbiz.cmbchina.com/doc-detail?bizkey=DCCT20201214145038074&subclass=1&treeID=2026062909220930543

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

## 调用示例

```go
body := &models.AccountNewNTQACINFBody{
	Ntqacinfx: []*models.AccountNewNtqacinfx{{
		Bbknbr: "69",
		Accnbr: "755900000000000",
	}},
}

resp, err := account.AccountNewNTQACINF(userId, sm4PrivateKey, userPrivateKey, body)
```

## 获取支付模式

## 单个支付

## 批量支付

## 交易流水
