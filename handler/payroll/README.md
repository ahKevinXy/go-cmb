# 代发代扣（官方 2026-07-02）

本目录保留原代发代扣接口，并新增官方文档版本 SDK。生成来源：https://openbiz.cmbchina.com/doc-detail?bizkey=DCCT20220309170929467&subclass=1&treeID=2026062909291024364

## 使用方式

```go
body := &models.PayrollNewBB6BTHHLBody{
	Bb6busmod: []*models.PayrollNewBb6busmod{{
		Buscod: "N03010",
		Busmod: "00001",
	}},
	Bb6cdcbhx1: []*models.PayrollNewBb6cdcbhx1{{
		Begtag: "Y",
		Endtag: "Y",
		Yurref: "your-unique-ref",
	}},
	Bb6cdcdlx1: []*models.PayrollNewBb6cdcdlx1{{
		Trxseq: "00000001",
		Accnbr: "6210000000000000000",
		Accnam: "张三",
		Trsamt: "100.00",
	}},
}
resp, err := payroll.PayrollNewBB6BTHHL(userId, sm4Key, userPrivateKey, body)
```

也可以使用通用入口：

```go
resp, err := payroll.RequestPayrollNewRaw(userId, sm4Key, userPrivateKey, constants.CmbPayrollNewBB6BTHHL, body)
```

## 已生成接口

| Funcode | Handler | Request Body | 说明 |
| --- | --- | --- | --- |
| CdcMatchCardBin | PayrollNewCdcMatchCardBin | models.PayrollNewCdcMatchCardBinBody | 1.卡BIN匹配识别标准银联卡CdcMatchCardBin |
| BB6BTHHL | PayrollNewBB6BTHHL | models.PayrollNewBB6BTHHLBody | 2.代发经办BB6BTHHL |
| BB6BPDQY | PayrollNewBB6BPDQY | models.PayrollNewBB6BPDQYBody | 3.批次与明细查询 BB6BPDQY |
| BB6INTQY | PayrollNewBB6INTQY | models.PayrollNewBB6INTQYBody | 4. 代发代扣综合查询 BB6INTQY |
| BB6BTHQY | PayrollNewBB6BTHQY | models.PayrollNewBB6BTHQYBody | 5. 代发批次查询 BB6BTHQY |
| BB6DTLQY | PayrollNewBB6DTLQY | models.PayrollNewBB6DTLQYBody | 6. 代发明细查询 BB6DTLQY |
| BB6RFDQY | PayrollNewBB6RFDQY | models.PayrollNewBB6RFDQYBody | 7.代发退票查询 BB6RFDQY |
| BB6AGTQY | PayrollNewBB6AGTQY | models.PayrollNewBB6AGTQYBody | 8.代发类型查询 BB6AGTQY |
| BB6ACLAK | PayrollNewBB6ACLAK | models.PayrollNewBB6ACLAKBody | 9.代扣经办BB6ACLAK |
| BB6CDCQY | PayrollNewBB6CDCQY | models.PayrollNewBB6CDCQYBody | 10. 代扣批次查询 BB6CDCQY |
| BB6ACLIF | PayrollNewBB6ACLIF | models.PayrollNewBB6ACLIFBody | 11.代扣明细查询BB6ACLIF |
| BB6TRSQY | PayrollNewBB6TRSQY | models.PayrollNewBB6TRSQYBody | 12.代扣类型查询BB6TRSQY |
| DCAGPPDF | PayrollNewDCAGPPDF | models.PayrollNewDCAGPPDFBody | 13.代发明细对账单查询请求DCAGPPDF |
| DCTASKID | PayrollNewDCTASKID | models.PayrollNewDCTASKIDBody | 14.代发明细对账单处理结果查询请求DCTASKID |
