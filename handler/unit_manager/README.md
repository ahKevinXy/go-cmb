# 交易管家（新）

本目录保留原交易管家接口，并新增官方文档“交易管家（新）”SDK。生成来源：https://openbiz.cmbchina.com/doc-detail?bizkey=DCCT20231226155549458&subclass=1&treeID=2026031714524636293

## 使用方式

```go
body := &models.UnitManagerNewNTDMAADDBody{
	Ntbusmody: []*models.UnitManagerNewNtbusmody{{Busmod: "00001"}},
	Ntdmaaddx: []*models.UnitManagerNewNtdmaaddx{{
		Accnbr: "755900000000000",
		Dmanbr: "000001",
		Dmanam: "测试子单元",
	}},
}
resp, err := unit_manager.UnitManagerNewNTDMAADD(userId, sm4Key, userPrivateKey, body)
```

也可以使用通用入口：

```go
resp, err := unit_manager.RequestUnitManagerNewRaw(userId, sm4Key, userPrivateKey, constants.CmbUnitManagerNewNTDMAADD, body)
```

## 已生成接口

| Funcode | Handler | Request Body | 说明 |
| --- | --- | --- | --- |
| NTDMAADD | UnitManagerNewNTDMAADD | models.UnitManagerNewNTDMAADDBody | 1.1新增记账子单元NTDMAADD |
| NTDMADLT | UnitManagerNewNTDMADLT | models.UnitManagerNewNTDMADLTBody | 1.2关闭记账子单元NTDMADLT |
| NTDMAMNT | UnitManagerNewNTDMAMNT | models.UnitManagerNewNTDMAMNTBody | 1.3修改记账子单元NTDMAMNT |
| NTDMALST | UnitManagerNewNTDMALST | models.UnitManagerNewNTDMALSTBody | 1.4查询记账子单元信息NTDMALST |
| NTOPRDMR | UnitManagerNewNTOPRDMR | models.UnitManagerNewNTOPRDMRBody | 2.1 记账子单元原路退回NTOPRDMR |
| NTDMATRX | UnitManagerNewNTDMATRX | models.UnitManagerNewNTDMATRXBody | 2.2 记账子单元内部调配NTDMATRX |
| NTHKPACT | UnitManagerNewNTHKPACT | models.UnitManagerNewNTHKPACTBody | 2.3手工记账NTHKPACT |
| NTHCLACT | UnitManagerNewNTHCLACT | models.UnitManagerNewNTHCLACTBody | 2.4记账子单元冲账NTHCLACT |
| BHHDZCTZ | UnitManagerNewBHHDZCTZ | models.UnitManagerNewBHHDZCTZBody | 3.1交易查询BHHDZCTZ |
| NTDMTQRYPDF | UnitManagerNewNTDMTQRYPDF | models.UnitManagerNewNTDMTQRYPDFBody | 3.2记账子单元交易明细对账单获取NTDMTQRYPDF |
| DCTASKID | UnitManagerNewDCTASKID | models.UnitManagerNewDCTASKIDBody | 3.3记账子单元交易明细对账单处理结果查询请求DCTASKID |
| BHHDZDBQ | UnitManagerNewBHHDZDBQ | models.UnitManagerNewBHHDZDBQBody | 3.4退票交易查询BHHDZDBQ |
| NTDMABAL | UnitManagerNewNTDMABAL | models.UnitManagerNewNTDMABALBody | 4.1查询记账子单元余额NTDMABAL |
| NTDMAHBD | UnitManagerNewNTDMAHBD | models.UnitManagerNewNTDMAHBDBody | 4.2查询所有记账子单元的某日余额NTDMAHBD |
| NTDMAHAD | UnitManagerNewNTDMAHAD | models.UnitManagerNewNTDMAHADBody | 4.3查询单个记账子单元的历史余额NTDMAHAD |
| NTMTHAUB | UnitManagerNewNTMTHAUB | models.UnitManagerNewNTMTHAUBBody | 4.4查询记账子单元历史月积数NTMTHAUB |
| NTDUMRED | UnitManagerNewNTDUMRED | models.UnitManagerNewNTDUMREDBody | 5.1按业务参考号查询结果NTDUMRED |
| NTDUMINF | UnitManagerNewNTDUMINF | models.UnitManagerNewNTDUMINFBody | 5.2详情查询NTDUMINF |
| NTDMAWLS | UnitManagerNewNTDMAWLS | models.UnitManagerNewNTDMAWLSBody | 6.1设置结算户入账分类条件NTDMAWLS |
| NTDMAWLD | UnitManagerNewNTDMAWLD | models.UnitManagerNewNTDMAWLDBody | 6.2删除结算户入账分类条件NTDMAWLD |
| BHHDZWQC | UnitManagerNewBHHDZWQC | models.UnitManagerNewBHHDZWQCBody | 6.3 查询结算户入账分类条件BHHDZWQC |
| BHHZWMNT | UnitManagerNewBHHZWMNT | models.UnitManagerNewBHHZWMNTBody | 6.4修改不满足结算户入账分类条件的入账方式BHHZWMNT |
| BHHDNWCA | UnitManagerNewBHHDNWCA | models.UnitManagerNewBHHDNWCABody | 7.1新增结算户出账分类条件BHHDNWCA |
| BHHDNWCC | UnitManagerNewBHHDNWCC | models.UnitManagerNewBHHDNWCCBody | 7.2删除结算户出账分类条件BHHDNWCC |
| BHHDNWCQ | UnitManagerNewBHHDNWCQ | models.UnitManagerNewBHHDNWCQBody | 7.3查询结算户出账分类条件BHHDNWCQ |
| NTDMARLT | UnitManagerNewNTDMARLT | models.UnitManagerNewNTDMARLTBody | 8.1设置记账子单元入账条件NTDMARLT |
| NTDMARLD | UnitManagerNewNTDMARLD | models.UnitManagerNewNTDMARLDBody | 8.2 删除记账子单元入账条件NTDMARLD |
| NTDMARLQ | UnitManagerNewNTDMARLQ | models.UnitManagerNewNTDMARLQBody | 8.3查询记账子单元入账条件NTDMARLQ |
| NTDMATMN | UnitManagerNewNTDMATMN | models.UnitManagerNewNTDMATMNBody | 8.4 修改不满足记账子单元入账条件时的入账方式NTDMATMN |
| BHHDNDSA | UnitManagerNewBHHDNDSA | models.UnitManagerNewBHHDNDSABody | 9.1新增个人账户白名单BHHDNDSA |
| BHHDNDSD | UnitManagerNewBHHDNDSD | models.UnitManagerNewBHHDNDSDBody | 9.2删除个人账户白名单BHHDNDSD |
| BHHDNDSQ | UnitManagerNewBHHDNDSQ | models.UnitManagerNewBHHDNDSQBody | 9.3查询个人账户白名单BHHDNDSQ |
| NTDCDADD | UnitManagerNewNTDCDADD | models.UnitManagerNewNTDCDADDBody | 10.1设置记账子单元关联公司卡NTDCDADD |
| NTDCDDLT | UnitManagerNewNTDCDDLT | models.UnitManagerNewNTDCDDLTBody | 10.2解除记账子单元与公司卡关联NTDCDDLT |
| NTDCDRLT | UnitManagerNewNTDCDRLT | models.UnitManagerNewNTDCDRLTBody | 10.3查询记账子单元关联公司卡列表NTDCDRLT |
| NTDUMCTL | UnitManagerNewNTDUMCTL | models.UnitManagerNewNTDUMCTLBody | 11.1设置记账子单元收款/出款限制NTDUMCTL |
| NTDUMCLE | UnitManagerNewNTDUMCLE | models.UnitManagerNewNTDUMCLEBody | 11.2解除记账子单元收款/出款限制NTDUMCLE |
| BHHJINAS | UnitManagerNewBHHJINAS | models.UnitManagerNewBHHJINASBody | 12.1 单笔设置子单元利率BHHJINAS |
| BHHJINCS | UnitManagerNewBHHJINCS | models.UnitManagerNewBHHJINCSBody | 12.2 单笔删除子单元利率BHHJINCS |
| BHHJINQD | UnitManagerNewBHHJINQD | models.UnitManagerNewBHHJINQDBody | 12.3 查询子单元利率BHHJINQD |
| BHHDINRP | UnitManagerNewBHHDINRP | models.UnitManagerNewBHHDINRPBody | 12.4 子单元单笔提前结息BHHDINRP |
| BHHJINRS | UnitManagerNewBHHJINRS | models.UnitManagerNewBHHJINRSBody | 12.5 子单元单笔手工结息BHHJINRS |
| BHHJINSD | UnitManagerNewBHHJINSD | models.UnitManagerNewBHHJINSDBody | 12.6 查询子单元结息结果BHHJINSD |
| BHHDINDT | UnitManagerNewBHHDINDT | models.UnitManagerNewBHHDINDTBody | 12.7合并结息明细查询BHHDINDT |
| NTDMAINT | UnitManagerNewNTDMAINT | models.UnitManagerNewNTDMAINTBody | 12.8记账子单元利息试算NTDMAINT |
| NTDMIQRY | UnitManagerNewNTDMIQRY | models.UnitManagerNewNTDMIQRYBody | 12.9记账子单元交易利息试算NTDMIQRY |
| BHHDINBB | UnitManagerNewBHHDINBB | models.UnitManagerNewBHHDINBBBody | 12.10按指定金额做利息试算（批量）BHHDINBB |
| BHHDLMAD | UnitManagerNewBHHDLMAD | models.UnitManagerNewBHHDLMADBody | 13.1单笔新增支付限额BHHDLMAD |
| BHHDLMMD | UnitManagerNewBHHDLMMD | models.UnitManagerNewBHHDLMMDBody | 13.2 单笔修改支付限额BHHDLMMD |
| BHHDLMCD | UnitManagerNewBHHDLMCD | models.UnitManagerNewBHHDLMCDBody | 13.3 单笔关闭支付限额BHHDLMCD |
| BHHDLMTP | UnitManagerNewBHHDLMTP | models.UnitManagerNewBHHDLMTPBody | 13.4 单笔设置/关闭临时限额BHHDLMTP |
| BHHDLMQY | UnitManagerNewBHHDLMQY | models.UnitManagerNewBHHDLMQYBody | 13.5 查询支付限额BHHDLMQY |
| BHHDGPON | UnitManagerNewBHHDGPON | models.UnitManagerNewBHHDGPONBody | 14.1. 开启多层级架构BHHDGPON |
| BHHDGPAD | UnitManagerNewBHHDGPAD | models.UnitManagerNewBHHDGPADBody | 14.2新增普通分组BHHDGPAD |
| BHHDGPMN | UnitManagerNewBHHDGPMN | models.UnitManagerNewBHHDGPMNBody | 14.3 修改普通分组BHHDGPMN |
| BHHDGPSC | UnitManagerNewBHHDGPSC | models.UnitManagerNewBHHDGPSCBody | 14.4关闭普通分组BHHDGPSC |
| BHHDGPBL | UnitManagerNewBHHDGPBL | models.UnitManagerNewBHHDGPBLBody | 14.5分组余额查询BHHDGPBL |
| BHHDZCTG | UnitManagerNewBHHDZCTG | models.UnitManagerNewBHHDZCTGBody | 14.6分组交易查询BHHDZCTG |
| BHHDGNAD | UnitManagerNewBHHDGNAD | models.UnitManagerNewBHHDGNADBody | 14.7分组下新增下级子单元BHHDGNAD |
| BHHDGNCD | UnitManagerNewBHHDGNCD | models.UnitManagerNewBHHDGNCDBody | 14.8分组下删除下级子单元BHHDGNCD |
| BHHDGNAG | UnitManagerNewBHHDGNAG | models.UnitManagerNewBHHDGNAGBody | 14.9分组下新增下级分组BHHDGNAG |
| BHHDGNCG | UnitManagerNewBHHDGNCG | models.UnitManagerNewBHHDGNCGBody | 14.10分组下删除下级分组BHHDGNCG |
| BHHDGNGS | UnitManagerNewBHHDGNGS | models.UnitManagerNewBHHDGNGSBody | 14.11查询下级节点信息BHHDGNGS |
| BHHDGNGT | UnitManagerNewBHHDGNGT | models.UnitManagerNewBHHDGNGTBody | 14.12 查询上级节点信息BHHDGNGT |
