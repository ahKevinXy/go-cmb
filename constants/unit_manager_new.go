package constants

// Code generated from CMB official 交易管家（新） documentation. DO NOT EDIT BY HAND.
// Source: https://openbiz.cmbchina.com/doc-detail?bizkey=DCCT20231226155549458&subclass=1&treeID=2026031714524636293

const (
	CmbUnitManagerNewNTDMAADD    = "NTDMAADD"    // 1.1新增记账子单元NTDMAADD
	CmbUnitManagerNewNTDMADLT    = "NTDMADLT"    // 1.2关闭记账子单元NTDMADLT
	CmbUnitManagerNewNTDMAMNT    = "NTDMAMNT"    // 1.3修改记账子单元NTDMAMNT
	CmbUnitManagerNewNTDMALST    = "NTDMALST"    // 1.4查询记账子单元信息NTDMALST
	CmbUnitManagerNewNTOPRDMR    = "NTOPRDMR"    // 2.1 记账子单元原路退回NTOPRDMR
	CmbUnitManagerNewNTDMATRX    = "NTDMATRX"    // 2.2 记账子单元内部调配NTDMATRX
	CmbUnitManagerNewNTHKPACT    = "NTHKPACT"    // 2.3手工记账NTHKPACT
	CmbUnitManagerNewNTHCLACT    = "NTHCLACT"    // 2.4记账子单元冲账NTHCLACT
	CmbUnitManagerNewBHHDZCTZ    = "BHHDZCTZ"    // 3.1交易查询BHHDZCTZ
	CmbUnitManagerNewNTDMTQRYPDF = "NTDMTQRYPDF" // 3.2记账子单元交易明细对账单获取NTDMTQRYPDF
	CmbUnitManagerNewDCTASKID    = "DCTASKID"    // 3.3记账子单元交易明细对账单处理结果查询请求DCTASKID
	CmbUnitManagerNewBHHDZDBQ    = "BHHDZDBQ"    // 3.4退票交易查询BHHDZDBQ
	CmbUnitManagerNewNTDMABAL    = "NTDMABAL"    // 4.1查询记账子单元余额NTDMABAL
	CmbUnitManagerNewNTDMAHBD    = "NTDMAHBD"    // 4.2查询所有记账子单元的某日余额NTDMAHBD
	CmbUnitManagerNewNTDMAHAD    = "NTDMAHAD"    // 4.3查询单个记账子单元的历史余额NTDMAHAD
	CmbUnitManagerNewNTMTHAUB    = "NTMTHAUB"    // 4.4查询记账子单元历史月积数NTMTHAUB
	CmbUnitManagerNewNTDUMRED    = "NTDUMRED"    // 5.1按业务参考号查询结果NTDUMRED
	CmbUnitManagerNewNTDUMINF    = "NTDUMINF"    // 5.2详情查询NTDUMINF
	CmbUnitManagerNewNTDMAWLS    = "NTDMAWLS"    // 6.1设置结算户入账分类条件NTDMAWLS
	CmbUnitManagerNewNTDMAWLD    = "NTDMAWLD"    // 6.2删除结算户入账分类条件NTDMAWLD
	CmbUnitManagerNewBHHDZWQC    = "BHHDZWQC"    // 6.3 查询结算户入账分类条件BHHDZWQC
	CmbUnitManagerNewBHHZWMNT    = "BHHZWMNT"    // 6.4修改不满足结算户入账分类条件的入账方式BHHZWMNT
	CmbUnitManagerNewBHHDNWCA    = "BHHDNWCA"    // 7.1新增结算户出账分类条件BHHDNWCA
	CmbUnitManagerNewBHHDNWCC    = "BHHDNWCC"    // 7.2删除结算户出账分类条件BHHDNWCC
	CmbUnitManagerNewBHHDNWCQ    = "BHHDNWCQ"    // 7.3查询结算户出账分类条件BHHDNWCQ
	CmbUnitManagerNewNTDMARLT    = "NTDMARLT"    // 8.1设置记账子单元入账条件NTDMARLT
	CmbUnitManagerNewNTDMARLD    = "NTDMARLD"    // 8.2 删除记账子单元入账条件NTDMARLD
	CmbUnitManagerNewNTDMARLQ    = "NTDMARLQ"    // 8.3查询记账子单元入账条件NTDMARLQ
	CmbUnitManagerNewNTDMATMN    = "NTDMATMN"    // 8.4 修改不满足记账子单元入账条件时的入账方式NTDMATMN
	CmbUnitManagerNewBHHDNDSA    = "BHHDNDSA"    // 9.1新增个人账户白名单BHHDNDSA
	CmbUnitManagerNewBHHDNDSD    = "BHHDNDSD"    // 9.2删除个人账户白名单BHHDNDSD
	CmbUnitManagerNewBHHDNDSQ    = "BHHDNDSQ"    // 9.3查询个人账户白名单BHHDNDSQ
	CmbUnitManagerNewNTDCDADD    = "NTDCDADD"    // 10.1设置记账子单元关联公司卡NTDCDADD
	CmbUnitManagerNewNTDCDDLT    = "NTDCDDLT"    // 10.2解除记账子单元与公司卡关联NTDCDDLT
	CmbUnitManagerNewNTDCDRLT    = "NTDCDRLT"    // 10.3查询记账子单元关联公司卡列表NTDCDRLT
	CmbUnitManagerNewNTDUMCTL    = "NTDUMCTL"    // 11.1设置记账子单元收款/出款限制NTDUMCTL
	CmbUnitManagerNewNTDUMCLE    = "NTDUMCLE"    // 11.2解除记账子单元收款/出款限制NTDUMCLE
	CmbUnitManagerNewBHHJINAS    = "BHHJINAS"    // 12.1 单笔设置子单元利率BHHJINAS
	CmbUnitManagerNewBHHJINCS    = "BHHJINCS"    // 12.2 单笔删除子单元利率BHHJINCS
	CmbUnitManagerNewBHHJINQD    = "BHHJINQD"    // 12.3 查询子单元利率BHHJINQD
	CmbUnitManagerNewBHHDINRP    = "BHHDINRP"    // 12.4 子单元单笔提前结息BHHDINRP
	CmbUnitManagerNewBHHJINRS    = "BHHJINRS"    // 12.5 子单元单笔手工结息BHHJINRS
	CmbUnitManagerNewBHHJINSD    = "BHHJINSD"    // 12.6 查询子单元结息结果BHHJINSD
	CmbUnitManagerNewBHHDINDT    = "BHHDINDT"    // 12.7合并结息明细查询BHHDINDT
	CmbUnitManagerNewNTDMAINT    = "NTDMAINT"    // 12.8记账子单元利息试算NTDMAINT
	CmbUnitManagerNewNTDMIQRY    = "NTDMIQRY"    // 12.9记账子单元交易利息试算NTDMIQRY
	CmbUnitManagerNewBHHDINBB    = "BHHDINBB"    // 12.10按指定金额做利息试算（批量）BHHDINBB
	CmbUnitManagerNewBHHDLMAD    = "BHHDLMAD"    // 13.1单笔新增支付限额BHHDLMAD
	CmbUnitManagerNewBHHDLMMD    = "BHHDLMMD"    // 13.2 单笔修改支付限额BHHDLMMD
	CmbUnitManagerNewBHHDLMCD    = "BHHDLMCD"    // 13.3 单笔关闭支付限额BHHDLMCD
	CmbUnitManagerNewBHHDLMTP    = "BHHDLMTP"    // 13.4 单笔设置/关闭临时限额BHHDLMTP
	CmbUnitManagerNewBHHDLMQY    = "BHHDLMQY"    // 13.5 查询支付限额BHHDLMQY
	CmbUnitManagerNewBHHDGPON    = "BHHDGPON"    // 14.1. 开启多层级架构BHHDGPON
	CmbUnitManagerNewBHHDGPAD    = "BHHDGPAD"    // 14.2新增普通分组BHHDGPAD
	CmbUnitManagerNewBHHDGPMN    = "BHHDGPMN"    // 14.3 修改普通分组BHHDGPMN
	CmbUnitManagerNewBHHDGPSC    = "BHHDGPSC"    // 14.4关闭普通分组BHHDGPSC
	CmbUnitManagerNewBHHDGPBL    = "BHHDGPBL"    // 14.5分组余额查询BHHDGPBL
	CmbUnitManagerNewBHHDZCTG    = "BHHDZCTG"    // 14.6分组交易查询BHHDZCTG
	CmbUnitManagerNewBHHDGNAD    = "BHHDGNAD"    // 14.7分组下新增下级子单元BHHDGNAD
	CmbUnitManagerNewBHHDGNCD    = "BHHDGNCD"    // 14.8分组下删除下级子单元BHHDGNCD
	CmbUnitManagerNewBHHDGNAG    = "BHHDGNAG"    // 14.9分组下新增下级分组BHHDGNAG
	CmbUnitManagerNewBHHDGNCG    = "BHHDGNCG"    // 14.10分组下删除下级分组BHHDGNCG
	CmbUnitManagerNewBHHDGNGS    = "BHHDGNGS"    // 14.11查询下级节点信息BHHDGNGS
	CmbUnitManagerNewBHHDGNGT    = "BHHDGNGT"    // 14.12 查询上级节点信息BHHDGNGT
)

var CmbUnitManagerNewFuncodes = map[string]string{
	CmbUnitManagerNewNTDMAADD:    "1.1新增记账子单元NTDMAADD",
	CmbUnitManagerNewNTDMADLT:    "1.2关闭记账子单元NTDMADLT",
	CmbUnitManagerNewNTDMAMNT:    "1.3修改记账子单元NTDMAMNT",
	CmbUnitManagerNewNTDMALST:    "1.4查询记账子单元信息NTDMALST",
	CmbUnitManagerNewNTOPRDMR:    "2.1 记账子单元原路退回NTOPRDMR",
	CmbUnitManagerNewNTDMATRX:    "2.2 记账子单元内部调配NTDMATRX",
	CmbUnitManagerNewNTHKPACT:    "2.3手工记账NTHKPACT",
	CmbUnitManagerNewNTHCLACT:    "2.4记账子单元冲账NTHCLACT",
	CmbUnitManagerNewBHHDZCTZ:    "3.1交易查询BHHDZCTZ",
	CmbUnitManagerNewNTDMTQRYPDF: "3.2记账子单元交易明细对账单获取NTDMTQRYPDF",
	CmbUnitManagerNewDCTASKID:    "3.3记账子单元交易明细对账单处理结果查询请求DCTASKID",
	CmbUnitManagerNewBHHDZDBQ:    "3.4退票交易查询BHHDZDBQ",
	CmbUnitManagerNewNTDMABAL:    "4.1查询记账子单元余额NTDMABAL",
	CmbUnitManagerNewNTDMAHBD:    "4.2查询所有记账子单元的某日余额NTDMAHBD",
	CmbUnitManagerNewNTDMAHAD:    "4.3查询单个记账子单元的历史余额NTDMAHAD",
	CmbUnitManagerNewNTMTHAUB:    "4.4查询记账子单元历史月积数NTMTHAUB",
	CmbUnitManagerNewNTDUMRED:    "5.1按业务参考号查询结果NTDUMRED",
	CmbUnitManagerNewNTDUMINF:    "5.2详情查询NTDUMINF",
	CmbUnitManagerNewNTDMAWLS:    "6.1设置结算户入账分类条件NTDMAWLS",
	CmbUnitManagerNewNTDMAWLD:    "6.2删除结算户入账分类条件NTDMAWLD",
	CmbUnitManagerNewBHHDZWQC:    "6.3 查询结算户入账分类条件BHHDZWQC",
	CmbUnitManagerNewBHHZWMNT:    "6.4修改不满足结算户入账分类条件的入账方式BHHZWMNT",
	CmbUnitManagerNewBHHDNWCA:    "7.1新增结算户出账分类条件BHHDNWCA",
	CmbUnitManagerNewBHHDNWCC:    "7.2删除结算户出账分类条件BHHDNWCC",
	CmbUnitManagerNewBHHDNWCQ:    "7.3查询结算户出账分类条件BHHDNWCQ",
	CmbUnitManagerNewNTDMARLT:    "8.1设置记账子单元入账条件NTDMARLT",
	CmbUnitManagerNewNTDMARLD:    "8.2 删除记账子单元入账条件NTDMARLD",
	CmbUnitManagerNewNTDMARLQ:    "8.3查询记账子单元入账条件NTDMARLQ",
	CmbUnitManagerNewNTDMATMN:    "8.4 修改不满足记账子单元入账条件时的入账方式NTDMATMN",
	CmbUnitManagerNewBHHDNDSA:    "9.1新增个人账户白名单BHHDNDSA",
	CmbUnitManagerNewBHHDNDSD:    "9.2删除个人账户白名单BHHDNDSD",
	CmbUnitManagerNewBHHDNDSQ:    "9.3查询个人账户白名单BHHDNDSQ",
	CmbUnitManagerNewNTDCDADD:    "10.1设置记账子单元关联公司卡NTDCDADD",
	CmbUnitManagerNewNTDCDDLT:    "10.2解除记账子单元与公司卡关联NTDCDDLT",
	CmbUnitManagerNewNTDCDRLT:    "10.3查询记账子单元关联公司卡列表NTDCDRLT",
	CmbUnitManagerNewNTDUMCTL:    "11.1设置记账子单元收款/出款限制NTDUMCTL",
	CmbUnitManagerNewNTDUMCLE:    "11.2解除记账子单元收款/出款限制NTDUMCLE",
	CmbUnitManagerNewBHHJINAS:    "12.1 单笔设置子单元利率BHHJINAS",
	CmbUnitManagerNewBHHJINCS:    "12.2 单笔删除子单元利率BHHJINCS",
	CmbUnitManagerNewBHHJINQD:    "12.3 查询子单元利率BHHJINQD",
	CmbUnitManagerNewBHHDINRP:    "12.4 子单元单笔提前结息BHHDINRP",
	CmbUnitManagerNewBHHJINRS:    "12.5 子单元单笔手工结息BHHJINRS",
	CmbUnitManagerNewBHHJINSD:    "12.6 查询子单元结息结果BHHJINSD",
	CmbUnitManagerNewBHHDINDT:    "12.7合并结息明细查询BHHDINDT",
	CmbUnitManagerNewNTDMAINT:    "12.8记账子单元利息试算NTDMAINT",
	CmbUnitManagerNewNTDMIQRY:    "12.9记账子单元交易利息试算NTDMIQRY",
	CmbUnitManagerNewBHHDINBB:    "12.10按指定金额做利息试算（批量）BHHDINBB",
	CmbUnitManagerNewBHHDLMAD:    "13.1单笔新增支付限额BHHDLMAD",
	CmbUnitManagerNewBHHDLMMD:    "13.2 单笔修改支付限额BHHDLMMD",
	CmbUnitManagerNewBHHDLMCD:    "13.3 单笔关闭支付限额BHHDLMCD",
	CmbUnitManagerNewBHHDLMTP:    "13.4 单笔设置/关闭临时限额BHHDLMTP",
	CmbUnitManagerNewBHHDLMQY:    "13.5 查询支付限额BHHDLMQY",
	CmbUnitManagerNewBHHDGPON:    "14.1. 开启多层级架构BHHDGPON",
	CmbUnitManagerNewBHHDGPAD:    "14.2新增普通分组BHHDGPAD",
	CmbUnitManagerNewBHHDGPMN:    "14.3 修改普通分组BHHDGPMN",
	CmbUnitManagerNewBHHDGPSC:    "14.4关闭普通分组BHHDGPSC",
	CmbUnitManagerNewBHHDGPBL:    "14.5分组余额查询BHHDGPBL",
	CmbUnitManagerNewBHHDZCTG:    "14.6分组交易查询BHHDZCTG",
	CmbUnitManagerNewBHHDGNAD:    "14.7分组下新增下级子单元BHHDGNAD",
	CmbUnitManagerNewBHHDGNCD:    "14.8分组下删除下级子单元BHHDGNCD",
	CmbUnitManagerNewBHHDGNAG:    "14.9分组下新增下级分组BHHDGNAG",
	CmbUnitManagerNewBHHDGNCG:    "14.10分组下删除下级分组BHHDGNCG",
	CmbUnitManagerNewBHHDGNGS:    "14.11查询下级节点信息BHHDGNGS",
	CmbUnitManagerNewBHHDGNGT:    "14.12 查询上级节点信息BHHDGNGT",
}
