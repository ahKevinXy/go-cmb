package models

import "encoding/json"

// Code generated from CMB official 交易管家（新） documentation. DO NOT EDIT BY HAND.
// Source: https://openbiz.cmbchina.com/doc-detail?bizkey=DCCT20231226155549458&subclass=1&treeID=2026031714524636293
// Version: 1.0.0, updated: 2026-03-18 16:22:46

type UnitManagerNewRequest struct {
	Request   UnitManagerNewRequestData `json:"request"`
	Signature Signature                 `json:"signature"`
}

type UnitManagerNewRequestData struct {
	Body interface{} `json:"body,omitempty"`
	Head Head        `json:"head"`
}

type UnitManagerNewResponse struct {
	Response UnitManagerNewResponseData `json:"response,omitempty"`
}

type UnitManagerNewResponseData struct {
	Body UnitManagerNewRawBody `json:"body,omitempty"`
	Head HeadResponse          `json:"head,omitempty"`
}

type UnitManagerNewRawBody map[string]json.RawMessage

type UnitManagerNewBody map[string]interface{}

// UnitManagerNewBhducmody1 输入接口 bhducmody1.
type UnitManagerNewBhducmody1 struct {
	// 业务类型 String(6) required:Y N36090
	YBusCod string `json:"yBusCod,omitempty"`
	// 业务模式 String(5) required:Y
	YBusMod string `json:"yBusMod,omitempty"`
}

// UnitManagerNewBHDUCMODY1 输入接口 BHDUCMODY1.
type UnitManagerNewBHDUCMODY1 struct {
	// 业务模式 String(5) required:Y
	YBusMod string `json:"yBusMod,omitempty"`
	// 业务类型 String(6) required:Y N36090：交易管家
	YBusCod string `json:"yBusCod,omitempty"`
	// 保留字段 String(50)
	YRsv50z string `json:"yRsv50z,omitempty"`
}

// UnitManagerNewBHHCDSRTNZ 详细信息 BHHCDSRTNZ.
type UnitManagerNewBHHCDSRTNZ struct {
	// 公司卡号 String(35)
	ZCdsNbr string `json:"zCdsNbr,omitempty"`
	// 公司卡户口名 String(62)
	ZCdsNam string `json:"zCdsNam,omitempty"`
	// 持卡人姓名 String(200)
	ZCltNam string `json:"zCltNam,omitempty"`
}

// UnitManagerNewBHHDAADDZ 详细信息 BHHDAADDZ.
type UnitManagerNewBHHDAADDZ struct {
	// 活期结算账户 String(35)
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20)
	Dmanbr string `json:"dmanbr,omitempty"`
	// 记账子单元名称 Z(82)
	Dmanam string `json:"dmanam,omitempty"`
	// 额度控制标志 String(1)
	Ovrctl string `json:"ovrctl,omitempty"`
	// 退票处理方式 String(1)
	Bcktyp string `json:"bcktyp,omitempty"`
	// 余额关闭标志 String(1)
	Clstyp string `json:"clstyp,omitempty"`
	// 业务参考号 String(30)
	Yurref string `json:"yurref,omitempty"`
	// 额度标志 String(1)
	Lmtflg string `json:"lmtflg,omitempty"`
	// 余额上限额度 BigDecimal(15,2)
	Ballmt string `json:"ballmt,omitempty"`
}

// UnitManagerNewBHHDADLTZ 详细信息 BHHDADLTZ.
type UnitManagerNewBHHDADLTZ struct {
	// 账号 String(35)
	Accnbr string `json:"accnbr,omitempty"`
	// 业务参考号 String(30)
	Yurref string `json:"yurref,omitempty"`
	// 记账子单元编号 String(20)
	Dmanbr string `json:"dmanbr,omitempty"`
	// 序号 String(22)
	ZRdtSeqNbr string `json:"zRdtSeqNbr,omitempty"`
	// 请求状态 String(3)
	ZRdtReqStq string `json:"zRdtReqStq,omitempty"`
	// 处理结果 String(1)
	ZRdtRtnFlg string `json:"zRdtRtnFlg,omitempty"`
	// 错误描述 String(300)
	ZRdtErrTxt string `json:"zRdtErrTxt,omitempty"`
}

// UnitManagerNewBHHDAUASZ 详细信息 BHHDAUASZ.
type UnitManagerNewBHHDAUASZ struct {
	// 账号 String(35)
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20)
	Dmanbr string `json:"dmanbr,omitempty"`
	// 业务参考号 String(30)
	Yurref string `json:"yurref,omitempty"`
	// 用户号 String(10)
	Usrnbr string `json:"usrnbr,omitempty"`
	// 操作权限位图 String(20)
	Autmap string `json:"autmap,omitempty"`
	// 序号 String(22)
	ZRdtSeqNbr string `json:"zRdtSeqNbr,omitempty"`
	// 请求状态 String(3)
	ZRdtReqStq string `json:"zRdtReqStq,omitempty"`
	// 处理结果 String(1)
	ZRdtRtnFlg string `json:"zRdtRtnFlg,omitempty"`
	// 错误描述 String(300)
	ZRdtErrTxt string `json:"zRdtErrTxt,omitempty"`
}

// UnitManagerNewBHHDAUMSZ 详细信息 BHHDAUMSZ.
type UnitManagerNewBHHDAUMSZ struct {
	// 账号 String(35)
	Accnbr string `json:"accnbr,omitempty"`
	// 业务参考号 String(30)
	Yurref string `json:"yurref,omitempty"`
	// 用户号 String(10)
	Usrnbr string `json:"usrnbr,omitempty"`
	// 序号 String(22)
	ZRdtSeqNbr string `json:"zRdtSeqNbr,omitempty"`
	// 请求状态 String(3)
	ZRdtReqStq string `json:"zRdtReqStq,omitempty"`
	// 处理结果 String(1)
	ZRdtRtnFlg string `json:"zRdtRtnFlg,omitempty"`
	// 错误描述 String(300)
	ZRdtErrTxt string `json:"zRdtErrTxt,omitempty"`
}

// UnitManagerNewBHHDAUULZ 详细信息 BHHDAUULZ.
type UnitManagerNewBHHDAUULZ struct {
	// 账号 String(35)
	Accnbr string `json:"accnbr,omitempty"`
	// 用户号 String(10)
	Usrnbr string `json:"usrnbr,omitempty"`
	// 操作权限位图 String(20) 第一位：查询，第二位：记账，第三位：修改 Y：有权限，N：无权限
	Autmap string `json:"autmap,omitempty"`
	// 业务参考号 String(30)
	Yurref string `json:"yurref,omitempty"`
	// 序号 String(22)
	ZRdtSeqNbr string `json:"zRdtSeqNbr,omitempty"`
	// 请求状态 String(3)
	ZRdtReqStq string `json:"zRdtReqStq,omitempty"`
	// 处理结果 String(1)
	ZRdtRtnFlg string `json:"zRdtRtnFlg,omitempty"`
	// 错误描述 String(300)
	ZRdtErrTxt string `json:"zRdtErrTxt,omitempty"`
}

// UnitManagerNewBHHDAUUSZ 详细信息 BHHDAUUSZ.
type UnitManagerNewBHHDAUUSZ struct {
	// 账号 String(35)
	Accnbr string `json:"accnbr,omitempty"`
	// 用户号 String(10)
	Usrnbr string `json:"usrnbr,omitempty"`
	// 业务参考号 String(30)
	Yurref string `json:"yurref,omitempty"`
	// 记账子单元编号 String(20)
	Dmanbr string `json:"dmanbr,omitempty"`
	// 操作权限位图 String(20)
	Autmap string `json:"autmap,omitempty"`
	// 序号 String(22)
	ZRdtSeqNbr string `json:"zRdtSeqNbr,omitempty"`
	// 请求状态 String(3)
	ZRdtReqStq string `json:"zRdtReqStq,omitempty"`
	// 处理结果 String(1)
	ZRdtRtnFlg string `json:"zRdtRtnFlg,omitempty"`
	// 错误描述 String(300)
	ZRdtErrTxt string `json:"zRdtErrTxt,omitempty"`
}

// UnitManagerNewBHHDGNADX1 输入接口 BHHDGNADX1.
type UnitManagerNewBHHDGNADX1 struct {
	// 户口号 String(35) required:Y
	XEacNbr string `json:"xEacNbr,omitempty"`
	// 分组编号 String(20) required:Y
	XGrpNbr string `json:"xGrpNbr,omitempty"`
	// 记账子单元编号 String(20) required:Y
	XDumNbr string `json:"xDumNbr,omitempty"`
	// 业务参考号 String(30) required:Y
	XYurRef string `json:"xYurRef,omitempty"`
}

// UnitManagerNewBHHDGNADZ1 详细信息 BHHDGNADZ1.
type UnitManagerNewBHHDGNADZ1 struct {
	// 外部分组编号 String(20)
	ZGrpNbr string `json:"zGrpNbr,omitempty"`
	// 记账子单元编号 String(20)
	ZDumNbr string `json:"zDumNbr,omitempty"`
}

// UnitManagerNewBHHDGNAGX1 输入接口 BHHDGNAGX1.
type UnitManagerNewBHHDGNAGX1 struct {
	// 户口号 String(35) required:Y
	XEacNbr string `json:"xEacNbr,omitempty"`
	// 上级分组编号 String(20) required:Y
	XTopGrp string `json:"xTopGrp,omitempty"`
	// 下级分组编号 String(20) required:Y
	XSubGrp string `json:"xSubGrp,omitempty"`
	// 业务参考号 String(30) required:Y
	XYurRef string `json:"xYurRef,omitempty"`
}

// UnitManagerNewBHHDGNAGZ1 详细信息 BHHDGNAGZ1.
type UnitManagerNewBHHDGNAGZ1 struct {
	// 上级外部分组号 String(20)
	ZTopGrp string `json:"zTopGrp,omitempty"`
	// 下级外部分组号 String(20)
	ZSubGrp string `json:"zSubGrp,omitempty"`
}

// UnitManagerNewBHHDGNCDX1 输入接口 BHHDGNCDX1.
type UnitManagerNewBHHDGNCDX1 struct {
	// 户口号 String(35) required:Y
	XEacNbr string `json:"xEacNbr,omitempty"`
	// 分组编号 String(20) required:Y
	XGrpNbr string `json:"xGrpNbr,omitempty"`
	// 记账子单元编号 String(20) required:Y
	XDumNbr string `json:"xDumNbr,omitempty"`
	// 业务参考号 String(30) required:Y
	XYurRef string `json:"xYurRef,omitempty"`
}

// UnitManagerNewBHHDGNCDZ1 详细信息 BHHDGNCDZ1.
type UnitManagerNewBHHDGNCDZ1 struct {
	// 外部分组编号 String(20)
	ZGrpNbr string `json:"zGrpNbr,omitempty"`
	// 记账子单元编号 String(20)
	ZDumNbr string `json:"zDumNbr,omitempty"`
}

// UnitManagerNewBHHDGNCGX1 输入接口 BHHDGNCGX1.
type UnitManagerNewBHHDGNCGX1 struct {
	// 户口号 String(35) required:Y
	XEacNbr string `json:"xEacNbr,omitempty"`
	// 上级分组编号 String(20) required:Y
	XTopGrp string `json:"xTopGrp,omitempty"`
	// 下级分组编号 String(20) required:Y
	XSubGrp string `json:"xSubGrp,omitempty"`
	// 业务参考号 String(30) required:Y
	XYurRef string `json:"xYurRef,omitempty"`
}

// UnitManagerNewBHHDGNCGZ1 详细信息 BHHDGNCGZ1.
type UnitManagerNewBHHDGNCGZ1 struct {
	// 上级外部分组号 String(20)
	ZTopGrp string `json:"zTopGrp,omitempty"`
	// 下级外部分组号 String(20)
	ZSubGrp string `json:"zSubGrp,omitempty"`
}

// UnitManagerNewBHHDGNGSX1 输入接口 BHHDGNGSX1.
type UnitManagerNewBHHDGNGSX1 struct {
	// 户口号 String(35) required:Y
	XEacNbr string `json:"xEacNbr,omitempty"`
	// 上级节点类型 String(3) required:Y GRP：分组
	XTopGrp string `json:"xTopGrp,omitempty"`
	// 上级节点编号 String(20) required:Y 分组编号
	XTopNbr string `json:"xTopNbr,omitempty"`
	// 下级节点类型 String(3) GRP：分组 DUM：子单元 空：全部
	XSubTyp string `json:"xSubTyp,omitempty"`
	// 下级节点编号 String(20)
	XSubNbr string `json:"xSubNbr,omitempty"`
	// 查询条数 BigDecimal(3,0) required:Y 1-200
	XQryCnt string `json:"xQryCnt,omitempty"`
}

// UnitManagerNewBHHDGNGSY1 续传接口 BHHDGNGSY1.
type UnitManagerNewBHHDGNGSY1 struct {
	// 续传键值 String(80) 当返回报文有续传接口，取续传接口放到下一次请求中
	YCtnKey string `json:"yCtnKey,omitempty"`
}

// UnitManagerNewBHHDGNGSZ1 输出接口 BHHDGNGSZ1.
type UnitManagerNewBHHDGNGSZ1 struct {
	// 上级节点类型 String(3） GRP：分组
	ZTopTyp string `json:"zTopTyp,omitempty"`
	// 上级节点编号 String(20)
	ZTopNbr string `json:"zTopNbr,omitempty"`
	// 下级节点类型 String(3) GRP：分组 DUM：子单元
	ZSubTyp string `json:"zSubTyp,omitempty"`
	// 下级节点编号 String(20)
	ZSubNbr string `json:"zSubNbr,omitempty"`
}

// UnitManagerNewBHHDGNGTX1 输入接口 BHHDGNGTX1.
type UnitManagerNewBHHDGNGTX1 struct {
	// 户口号 String(35) required:Y
	XEacNbr string `json:"xEacNbr,omitempty"`
	// 下级节点类型 String(3) required:Y GRP：分组 DUM：子单元
	XSubTyp string `json:"xSubTyp,omitempty"`
	// 下级节点编号 String(20) required:Y
	XSubNbr string `json:"xSubNbr,omitempty"`
}

// UnitManagerNewBHHDGNGTZ1 输出接口 BHHDGNGTZ1.
type UnitManagerNewBHHDGNGTZ1 struct {
	// 上级节点类型 String(3） GRP：分组
	ZTopTyp string `json:"zTopTyp,omitempty"`
	// 上级节点编号 String(20)
	ZTopNbr string `json:"zTopNbr,omitempty"`
	// 下级节点类型 String(3) GRP：分组 DUM：子单元
	ZSubTyp string `json:"zSubTyp,omitempty"`
	// 下级节点编号 String(20)
	ZSubNbr string `json:"zSubNbr,omitempty"`
}

// UnitManagerNewBHHDGPADX1 输入接口 BHHDGPADX1.
type UnitManagerNewBHHDGPADX1 struct {
	// 户口号 String(35) required:Y
	XAccNbr string `json:"xAccNbr,omitempty"`
	// 分组编号 String(20) required:Y 字母、数字组合
	XGrpNbr string `json:"xGrpNbr,omitempty"`
	// 分组名称 String(200) required:Y
	XGrpNam string `json:"xGrpNam,omitempty"`
	// 业务参考号 String(30) required:Y
	XYurRef string `json:"xYurRef,omitempty"`
	// 保留字50 String(50)
	XRsv50z string `json:"xRsv50z,omitempty"`
}

// UnitManagerNewBHHDGPADZ1 详细信息 BHHDGPADZ1.
type UnitManagerNewBHHDGPADZ1 struct {
	// 外部分组编号 String(20)
	ZGrpNbr string `json:"zGrpNbr,omitempty"`
	// 分组名称 String(200)
	ZGrpNam string `json:"zGrpNam,omitempty"`
}

// UnitManagerNewBHHDGPBLX1 输入接口 BHHDGPBLX1.
type UnitManagerNewBHHDGPBLX1 struct {
	// 户口号 String(35) required:Y
	XAccNbr string `json:"xAccNbr,omitempty"`
	// 分组编号 String(20) required:Y
	XGrpNbr string `json:"xGrpNbr,omitempty"`
}

// UnitManagerNewBHHDGPBLZ1 输出接口 BHHDGPBLZ1.
type UnitManagerNewBHHDGPBLZ1 struct {
	// 分组编号 String(20）
	GrpNbr string `json:"grpNbr,omitempty"`
	// 分组名称 String(20)
	GrpNam string `json:"grpNam,omitempty"`
	// 分组联机余额 BigDecimal(15,2)
	OnlBal string `json:"onlBal,omitempty"`
}

// UnitManagerNewBHHDGPMNX1 输入接口 BHHDGPMNX1.
type UnitManagerNewBHHDGPMNX1 struct {
	// 户口号 String(35) required:Y
	XAccNbr string `json:"xAccNbr,omitempty"`
	// 分组编号 String(20) required:Y
	XGrpNbr string `json:"xGrpNbr,omitempty"`
	// 分组名称 String(200) required:Y
	XGrpNam string `json:"xGrpNam,omitempty"`
	// 业务参考号 String(30) required:Y
	XYurRef string `json:"xYurRef,omitempty"`
	// 保留字50 String(50)
	XRsv50z string `json:"xRsv50z,omitempty"`
}

// UnitManagerNewBHHDGPMNZ1 详细信息 BHHDGPMNZ1.
type UnitManagerNewBHHDGPMNZ1 struct {
	// 外部分组编号 String(20)
	ZGrpNbr string `json:"zGrpNbr,omitempty"`
	// 分组名称 String(200)
	ZGrpNam string `json:"zGrpNam,omitempty"`
}

// UnitManagerNewBHHDGPONX1 输入接口 BHHDGPONX1.
type UnitManagerNewBHHDGPONX1 struct {
	// 户口号 String(35) required:Y
	XAccNbr string `json:"xAccNbr,omitempty"`
	// 业务参考号 String(30) required:Y
	XYurRef string `json:"xYurRef,omitempty"`
	// 保留字50 String(50)
	XRsv50z string `json:"xRsv50z,omitempty"`
}

// UnitManagerNewBHHDGPSCX1 输入接口 BHHDGPSCX1.
type UnitManagerNewBHHDGPSCX1 struct {
	// 户口号 String(35) required:Y
	XAccNbr string `json:"xAccNbr,omitempty"`
	// 分组编号 String(20) required:Y
	XGrpNbr string `json:"xGrpNbr,omitempty"`
	// 业务参考号 String(30) required:Y
	XYurRef string `json:"xYurRef,omitempty"`
	// 保留字50 String(50)
	XRsv50z string `json:"xRsv50z,omitempty"`
}

// UnitManagerNewBHHDGPSCZ1 详细信息 BHHDGPSCZ1.
type UnitManagerNewBHHDGPSCZ1 struct {
	// 外部分组编号 String(20)
	ZGrpNbr string `json:"zGrpNbr,omitempty"`
}

// UnitManagerNewBHHDINABZ 详细信息 BHHDINABZ.
type UnitManagerNewBHHDINABZ struct {
	// 账号 String(35)
	Accnbr string `json:"accnbr,omitempty"`
	// 操作类型 String(1) L：清单导入 D：直接设置
	Oprcod string `json:"oprcod,omitempty"`
	// 业务参考号 String(30)
	Yurref string `json:"yurref,omitempty"`
	// 开始记账子单元编号 String(20)
	Begdma string `json:"begdma,omitempty"`
	// 结束记账子单元编号 String(20)
	Enddma string `json:"enddma,omitempty"`
	// 结息周期 String(1) S：季度 M：月度
	Intcyc string `json:"intcyc,omitempty"`
	// 利率类型 String(3) HQI：活期 SLF：自定义
	Settyp string `json:"settyp,omitempty"`
	// 利率值 BigDecimal(11,7) 利率类型为Slf返回
	Intval string `json:"intval,omitempty"`
	// 记账子单元编号 String(20)
	Dmanbr string `json:"dmanbr,omitempty"`
	// 序号 String(22)
	ZRdtSeqNbr string `json:"zRdtSeqNbr,omitempty"`
	// 请求状态 String(3)
	ZRdtReqStq string `json:"zRdtReqStq,omitempty"`
	// 处理结果 String(1)
	ZRdtRtnFlg string `json:"zRdtRtnFlg,omitempty"`
	// 错误描述 String(300)
	ZRdtErrTxt string `json:"zRdtErrTxt,omitempty"`
}

// UnitManagerNewBHHDINABZ1 详细信息 BHHDINABZ1.
type UnitManagerNewBHHDINABZ1 struct {
	// 序号 String(22)
	ZRdtSeqNbr string `json:"zRdtSeqNbr,omitempty"`
	// 计息子单元 String(20)
	ZDumNbr string `json:"zDumNbr,omitempty"`
	// 收息子单元 String(20)
	ZTrsDum string `json:"zTrsDum,omitempty"`
	// 代收息方式 String(1)
	ZTrsTyp string `json:"zTrsTyp,omitempty"`
}

// UnitManagerNewBHHDINASX1 输入接口 BHHDINASX1.
type UnitManagerNewBHHDINASX1 struct {
	// 账号 String(35) required:Y
	XAccNbr string `json:"xAccNbr,omitempty"`
	// 虚拟户编号 String(20) required:Y
	XDmaNbr string `json:"xDmaNbr,omitempty"`
	// 结息周期 String(1) required:Y S：季度 M：月度
	XIntCyc string `json:"xIntCyc,omitempty"`
	// 利率类型 String(3) required:Y HQI：活期 ，SLF：自定义
	XSetTyp string `json:"xSetTyp,omitempty"`
	// 利率值 BigDecimal(11,7) 利率类型为SLF时必输，必须大于0小于等于100 示例：0.11， 表示0.11%
	XIntVal string `json:"xIntVal,omitempty"`
	// 业务参考号 String(30) required:Y
	XYurRef string `json:"xYurRef,omitempty"`
	// 收息子单元 String(20) 选择代收息，该字段非空；否则该字段必须为空
	XTrsDum string `json:"xTrsDum,omitempty"`
	// 代收息方式 String(1) 1：逐笔，2：合并 （xTrsDum非空时，xTrsTyp有效）
	XTrsTyp string `json:"xTrsTyp,omitempty"`
	// 结息日 String(2) 结息周期为季度：必须为空， 结息周期为月度：必输，范围[01,28]
	XIntDay string `json:"xIntDay,omitempty"`
	// 保留字段50 String(15)
	XRsv50z string `json:"xRsv50z,omitempty"`
}

// UnitManagerNewBHHDINASZ 详细信息 BHHDINASZ.
type UnitManagerNewBHHDINASZ struct {
	// 账号 String(35)
	ZAccNbr string `json:"zAccNbr,omitempty"`
	// 虚拟户编号 String(20)
	ZDmaNbr string `json:"zDmaNbr,omitempty"`
	// 结息周期 String(1) S：季度 M：月度
	ZIntCyc string `json:"zIntCyc,omitempty"`
	// 利率类型 String(3) HQI：活期 SLF：自定义
	ZSetTyp string `json:"zSetTyp,omitempty"`
	// 利率值 BigDecimal(11,7) 利率类型为Slf返回
	ZIntVal string `json:"zIntVal,omitempty"`
	// 业务参考号 String(30)
	ZYurRef string `json:"zYurRef,omitempty"`
}

// UnitManagerNewBHHDINASZ1 详细信息 BHHDINASZ1.
type UnitManagerNewBHHDINASZ1 struct {
	// 收息子单元 String(20)
	ZTrsDum string `json:"zTrsDum,omitempty"`
	// 代收息方式 String(1)
	ZTrsTyp string `json:"zTrsTyp,omitempty"`
	// 结息日 String(2)
	ZIntDay string `json:"zIntDay,omitempty"`
	// 保留字段 String(48)
	ZRsv50z string `json:"zRsv50z,omitempty"`
}

// UnitManagerNewBHHDINBBX1 输入接口 BHHDINBBX1.
type UnitManagerNewBHHDINBBX1 struct {
	// 账号 String(35) required:Y
	XAccNbr string `json:"xAccNbr,omitempty"`
	// 利率类型 String(3) required:Y SLF 自定义，HQI 活期
	XRatTyp string `json:"xRatTyp,omitempty"`
	// 利率值 BigDecimal(11,7) 利率类型为SLF时必输
	XRatVal string `json:"xRatVal,omitempty"`
}

// UnitManagerNewBHHDINBBX2 输入接口 BHHDINBBX2.
type UnitManagerNewBHHDINBBX2 struct {
	// 交易金额 BigDecimal(15,2)
	XTrsAmt string `json:"xTrsAmt,omitempty"`
	// 开始日期 String(8) required:Y
	XBegDat string `json:"xBegDat,omitempty"`
	// 结束日期 String(8) required:Y
	XEndDat string `json:"xEndDat,omitempty"`
}

// UnitManagerNewBHHDINBBZ1 详细信息 BHHDINBBZ1.
type UnitManagerNewBHHDINBBZ1 struct {
	// 币种 String(2)
	ZCcyNbr string `json:"zCcyNbr,omitempty"`
	// 计息积数 BigDecimal(17,2)
	ZIntAub string `json:"zIntAub,omitempty"`
	// 利息 BigDecimal(15,2)
	ZIntVal string `json:"zIntVal,omitempty"`
	// 利率 BigDecimal(11,7)
	ZRatVal string `json:"zRatVal,omitempty"`
}

// UnitManagerNewBHHDINCBZ 详细信息 BHHDINCBZ.
type UnitManagerNewBHHDINCBZ struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 业务参考号 String(30)
	Yurref string `json:"yurref,omitempty"`
	// 记账子单元编号 String(20) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
	// 序号 String(22)
	ZRdtSeqNbr string `json:"zRdtSeqNbr,omitempty"`
	// 请求状态 String(3)
	ZRdtReqStq string `json:"zRdtReqStq,omitempty"`
	// 处理结果 String(1)
	ZRdtRtnFlg string `json:"zRdtRtnFlg,omitempty"`
	// 错误描述 String(300)
	ZRdtErrTxt string `json:"zRdtErrTxt,omitempty"`
}

// UnitManagerNewBHHDINCSX1 输入接口 BHHDINCSX1.
type UnitManagerNewBHHDINCSX1 struct {
	// 账号 String(35) required:Y
	XAccNbr string `json:"xAccNbr,omitempty"`
	// 虚拟户编号 String(20) required:Y
	XDmaNbr string `json:"xDmaNbr,omitempty"`
	// 业务参考号 String(30) required:Y
	XYurRef string `json:"xYurRef,omitempty"`
	// 保留字段50 String(20)
	XRsv50z string `json:"xRsv50z,omitempty"`
}

// UnitManagerNewBHHDINCSZ 详细信息 BHHDINCSZ.
type UnitManagerNewBHHDINCSZ struct {
	// 账号 String(35)
	ZAccNbr string `json:"zAccNbr,omitempty"`
	// 虚拟户编号 String(20)
	ZDmaNbr string `json:"zDmaNbr,omitempty"`
	// 业务参考号 String(30)
	ZYurRef string `json:"zYurRef,omitempty"`
}

// UnitManagerNewBHHDINDTX1 查询条件 BHHDINDTX1.
type UnitManagerNewBHHDINDTX1 struct {
	// 户口 String(35) required:Y
	XEacNbr string `json:"xEacNbr,omitempty"`
	// 收息子单元编号 String(20) required:Y
	XDumNbr string `json:"xDumNbr,omitempty"`
	// 结息日 String(8) required:Y 对应结息结束日期+1
	XIntDat string `json:"xIntDat,omitempty"`
}

// UnitManagerNewBHHDINDTY1 续传接口 BHHDINDTY1.
type UnitManagerNewBHHDINDTY1 struct {
	// 续传标记 String(1) required:Y Y：续传，N: 不续传
	XCtnFlg string `json:"xCtnFlg,omitempty"`
	// 续传子单元编号 String(20) xCtnFlg=Y：必须非空 xCtnFlg=N：必须为空
	XCtnDum string `json:"xCtnDum,omitempty"`
}

// UnitManagerNewBHHDINDTZ1 详细信息 BHHDINDTZ1.
type UnitManagerNewBHHDINDTZ1 struct {
	// 协议编号 String(10)
	ZCstNbr string `json:"zCstNbr,omitempty"`
	// 主户口 String(35)
	ZEacNbr string `json:"zEacNbr,omitempty"`
	// 收息记账子单元编号 String(20)
	ZDumNbr string `json:"zDumNbr,omitempty"`
	// 计息记账子单元编号 String(20)
	ZIntDum string `json:"zIntDum,omitempty"`
	// 计息方式 String(1) 1:自动，2:手动
	ZIntTyp string `json:"zIntTyp,omitempty"`
	// 计息周期 String(1) M:月，S:季
	ZIntCyc string `json:"zIntCyc,omitempty"`
	// 分段标志 String(1) Y：分段，N：不分段
	ZSubFlg string `json:"zSubFlg,omitempty"`
	// 利率类型 String(3) HQI：活期，SLF：自定义
	ZRatTyp string `json:"zRatTyp,omitempty"`
	// 利率 BigDecimal(11, 7)
	ZRatVal string `json:"zRatVal,omitempty"`
	// 计息年份 String(4)
	ZCurYer string `json:"zCurYer,omitempty"`
	// 计息期数 String(2)
	ZCurPer string `json:"zCurPer,omitempty"`
	// 计息日 String(8)
	ZCurDat string `json:"zCurDat,omitempty"`
	// 计息开始日期 String(8)
	ZCurBeg string `json:"zCurBeg,omitempty"`
	// 计息结束日期 String(8)
	ZCurEnd string `json:"zCurEnd,omitempty"`
	// 积数 BigDecimal(17, 2)
	ZAubBal string `json:"zAubBal,omitempty"`
	// 利息 BigDecimal(15, 2)
	ZTrsAmt string `json:"zTrsAmt,omitempty"`
}

// UnitManagerNewBHHDINQDX1 查询条件 BHHDINQDX1.
type UnitManagerNewBHHDINQDX1 struct {
	// 账号 String(35) required:Y
	XAccNbr string `json:"xAccNbr,omitempty"`
	// 虚拟户编号 String(20) 空：从头开始查询 非空：有续传值时填入该字段，从指定的子单元开始查询
	XDmaNbr string `json:"xDmaNbr,omitempty"`
	// 保留字段50 String(30)
	XRsv50z string `json:"xRsv50z,omitempty"`
}

// UnitManagerNewBHHDINQDZ1 详细信息 BHHDINQDZ1.
type UnitManagerNewBHHDINQDZ1 struct {
	// 虚拟户编号 String(20)
	ZDmaNbr string `json:"zDmaNbr,omitempty"`
	// 虚拟户户名 Z(82)
	ZDmaNam string `json:"zDmaNam,omitempty"`
	// 结息周期 String(1) S：季度 M：月度
	ZIntCyc string `json:"zIntCyc,omitempty"`
	// 利率类型 String(3) HQI：活期，SLF：自定义
	ZSetTyp string `json:"zSetTyp,omitempty"`
	// 利率值 BigDecimal(11,7)
	ZIntVal string `json:"zIntVal,omitempty"`
	// 下次结息年份 String(4)
	ZNxtYer string `json:"zNxtYer,omitempty"`
	// 下次结息期数 String(2)
	ZNxtDur string `json:"zNxtDur,omitempty"`
	// 下次结息日 String(8)
	ZNxtDat string `json:"zNxtDat,omitempty"`
	// 代收息标志 String(1) Y：代收息，N：非代收息
	ZRepFlg string `json:"zRepFlg,omitempty"`
	// 收息子单元 String(20) zRepFlg=Y：有值
	ZTrsDum string `json:"zTrsDum,omitempty"`
	// 代收息方式 String(1) 1：逐笔，2：合并 空=非代收息
	ZTrsTyp string `json:"zTrsTyp,omitempty"`
	// 保留字段 String(8)
	ZRsv30z string `json:"zRsv30z,omitempty"`
}

// UnitManagerNewBHHDINQDZ2 续传接口 BHHDINQDZ2.
type UnitManagerNewBHHDINQDZ2 struct {
	// 续传标志 String(1) Y：续传 N：非续传
	ZCtnFlg string `json:"zCtnFlg,omitempty"`
	// 续传虚拟户编号 String(20) zCtnFlg=Y：该字段有值
	ZDmaNbr string `json:"zDmaNbr,omitempty"`
	// 保留字段50 String(50)
	ZRsv50z string `json:"zRsv50z,omitempty"`
}

// UnitManagerNewBHHDINRBZ 详细信息 BHHDINRBZ.
type UnitManagerNewBHHDINRBZ struct {
	// 账号 String(35)
	Accnbr string `json:"accnbr,omitempty"`
	// 业务参考号 String(30)
	ZYurRef string `json:"zYurRef,omitempty"`
	// 记账子单元编号 String(20)
	Dmanbr string `json:"dmanbr,omitempty"`
	// 结息流水号 String(20)
	Seqnbr string `json:"seqnbr,omitempty"`
	// 利息 M(15)
	Trsamt string `json:"trsamt,omitempty"`
	// 序号 String(22)
	ZRdtSeqNbr string `json:"zRdtSeqNbr,omitempty"`
	// 请求状态 String(3)
	ZRdtReqStq string `json:"zRdtReqStq,omitempty"`
	// 处理结果 String(1)
	ZRdtRtnFlg string `json:"zRdtRtnFlg,omitempty"`
	// 错误描述 String(300)
	ZRdtErrTxt string `json:"zRdtErrTxt,omitempty"`
}

// UnitManagerNewBHHDINRPX1 输入接口 BHHDINRPX1.
type UnitManagerNewBHHDINRPX1 struct {
	// 主户口 String(35) required:Y
	XEacNbr string `json:"xEacNbr,omitempty"`
	// 计息子单元 String(20) required:Y
	XDumNbr string `json:"xDumNbr,omitempty"`
	// 结息区间结束日 String(8) required:Y 结束日不能大于等于当天
	XEndDat string `json:"xEndDat,omitempty"`
	// 业务参考号 String(30) required:Y
	XYurRef string `json:"xYurRef,omitempty"`
}

// UnitManagerNewBHHDINRPZ1 详细信息 BHHDINRPZ1.
type UnitManagerNewBHHDINRPZ1 struct {
	// 主户口 String(35)
	ZEacNbr string `json:"zEacNbr,omitempty"`
	// 计息子单元 String(20)
	ZDumNbr string `json:"zDumNbr,omitempty"`
	// 结息区间结束日 String(8)
	ZEndDat string `json:"zEndDat,omitempty"`
}

// UnitManagerNewBHHDINRSX1 输入接口 BHHDINRSX1.
type UnitManagerNewBHHDINRSX1 struct {
	// 账号 String(35) required:Y
	XAccNbr string `json:"xAccNbr,omitempty"`
	// 虚拟户编号 String(20) required:Y
	XDmaNbr string `json:"xDmaNbr,omitempty"`
	// 结息流水号 String(20) required:Y
	XSeqNbr string `json:"xSeqNbr,omitempty"`
	// 利息 BigDecimal(15,2) required:Y 非0
	XTrsAmt string `json:"xTrsAmt,omitempty"`
	// 业务参考号 String(30) required:Y
	XYurRef string `json:"xYurRef,omitempty"`
	// 保留字段50 String(20)
	XRsv50z string `json:"xRsv50z,omitempty"`
}

// UnitManagerNewBHHDINRSZ 详细信息 BHHDINRSZ.
type UnitManagerNewBHHDINRSZ struct {
	// 账号 String(35)
	ZAccNbr string `json:"zAccNbr,omitempty"`
	// 虚拟户编号 String(20)
	ZDmaNbr string `json:"zDmaNbr,omitempty"`
	// 结息流水号 String(20)
	ZSeqNbr string `json:"zSeqNbr,omitempty"`
	// 利息 BigDecimal(15,2)
	ZTrsAmt string `json:"zTrsAmt,omitempty"`
	// 业务参考号 String(30)
	ZYurRef string `json:"zYurRef,omitempty"`
}

// UnitManagerNewBHHDINRTZ1 详细信息 BHHDINRTZ1.
type UnitManagerNewBHHDINRTZ1 struct {
	// 主户口 String(35)
	ZEacNbr string `json:"zEacNbr,omitempty"`
	// 计息子单元 String(20)
	ZDumNbr string `json:"zDumNbr,omitempty"`
	// 结息区间结束日 String(8)
	ZEndDat string `json:"zEndDat,omitempty"`
	// 利息 BigDecimal(15,2)
	ZIntAmt string `json:"zIntAmt,omitempty"`
	// 请求状态 String(3)
	ZReqSta string `json:"zReqSta,omitempty"`
	// 处理结果 String(1)
	ZRtnFlg string `json:"zRtnFlg,omitempty"`
	// 错误描述 String(300)
	ZErrTxt string `json:"zErrTxt,omitempty"`
}

// UnitManagerNewBHHDINSDX1 查询条件 BHHDINSDX1.
type UnitManagerNewBHHDINSDX1 struct {
	// 账号 String(35) required:Y
	XAccNbr string `json:"xAccNbr,omitempty"`
	// 年份 String(4) required:Y
	XPrdYer string `json:"xPrdYer,omitempty"`
	// 虚拟户编号 String(20) 续传字段
	XDmaNbr string `json:"xDmaNbr,omitempty"`
	// 结息流水号 String(20) 续传字段
	XSeqNbr string `json:"xSeqNbr,omitempty"`
	// 结息子序号 String(5) 续传字段
	XSubSeq string `json:"xSubSeq,omitempty"`
	// 保留字段50 String(50)
	XRsv50z string `json:"xRsv50z,omitempty"`
}

// UnitManagerNewBHHDINSDZ1 详细信息 BHHDINSDZ1.
type UnitManagerNewBHHDINSDZ1 struct {
	// 结息流水号 String(20)
	ZSeqNbr string `json:"zSeqNbr,omitempty"`
	// 结息子序号 String(5)
	ZSubSeq string `json:"zSubSeq,omitempty"`
	// 虚拟户编号 String(20)
	ZDmaNbr string `json:"zDmaNbr,omitempty"`
	// 币种 String(2)
	ZCcyNbr string `json:"zCcyNbr,omitempty"`
	// 结息周期 String(1) S：季度 M：月度
	ZCycTyp string `json:"zCycTyp,omitempty"`
	// 结息方式 String(1) 1：自动，2：手工
	ZIntTyp string `json:"zIntTyp,omitempty"`
	// 当前结息年份 String(4)
	ZPrcYer string `json:"zPrcYer,omitempty"`
	// 当前结息期数 String(2)
	ZPrcDur string `json:"zPrcDur,omitempty"`
	// 结息开始日期 String(8)
	ZBegDat string `json:"zBegDat,omitempty"`
	// 结息结束日期 String(8)
	ZEndDat string `json:"zEndDat,omitempty"`
	// 分段标志 String(1) Y：分段，N：不分段
	ZSubFlg string `json:"zSubFlg,omitempty"`
	// 积数 BigDecimal(17, 2)
	ZMpuBal string `json:"zMpuBal,omitempty"`
	// 利率 BigDecimal(11, 7)
	ZMpuRat string `json:"zMpuRat,omitempty"`
	// 交易金额 BigDecimal(15, 2)
	ZIntAmt string `json:"zIntAmt,omitempty"`
	// 交易日期 String(8)
	ZTrxDat string `json:"zTrxDat,omitempty"`
	// 交易时间 String(6)
	ZTrxTim string `json:"zTrxTim,omitempty"`
	// 付息交易流水号 String(15)
	ZDbtTrx string `json:"zDbtTrx,omitempty"`
	// 收息交易流水号 String(15)
	ZCrdTrx string `json:"zCrdTrx,omitempty"`
	// 状态 String(1) S：成功，F：失败，I：处理中
	ZStsCod string `json:"zStsCod,omitempty"`
	// 失败原因 Z(192)
	ZRtnTxt string `json:"zRtnTxt,omitempty"`
	// 计息子单元 String(20)
	ZIntDum string `json:"zIntDum,omitempty"`
	// 代收息方式 String(1) 1：逐笔，2：合并 空=非代收息
	ZTrsTyp string `json:"zTrsTyp,omitempty"`
	// 保留字段 String(9)
	ZRsv30z string `json:"zRsv30z,omitempty"`
}

// UnitManagerNewBHHDINSDZ2 续传接口 BHHDINSDZ2.
type UnitManagerNewBHHDINSDZ2 struct {
	// 续传标志 String(1) Y：续传，N：不续传
	ZCtnFlg string `json:"zCtnFlg,omitempty"`
	// 续传虚拟户编号 String(20)
	ZDmaNbr string `json:"zDmaNbr,omitempty"`
	// 续传结息流水号 String(20)
	ZSeqNbr string `json:"zSeqNbr,omitempty"`
	// 续传结息子序号 String(5)
	ZSubSeq string `json:"zSubSeq,omitempty"`
	// 保留字段50 String(50)
	ZRsv50z string `json:"zRsv50z,omitempty"`
}

// UnitManagerNewBHHDLMADX1 输入接口 BHHDLMADX1.
type UnitManagerNewBHHDLMADX1 struct {
	// 户口号 String(35) required:Y
	XAccNbr string `json:"xAccNbr,omitempty"`
	// 子单元编号 String(20) required:Y
	XDumNbr string `json:"xDumNbr,omitempty"`
	// 单笔限额-设置标志 String(1) required:Y Y：设置 N：不设置
	SigSetFlag string `json:"sigSetFlag,omitempty"`
	// 单笔限额-设置额度 BigDecimal(15,2)
	SigSetAmt string `json:"sigSetAmt,omitempty"`
	// 每日限额-设置标志 String(1) required:Y Y：设置 N：不设置
	DaySetFlag string `json:"daySetFlag,omitempty"`
	// 每日限额-设置额度 BigDecimal(15,2)
	DaySetAmt string `json:"daySetAmt,omitempty"`
	// 每周限额-设置标志 String(1) required:Y Y：设置 N：不设置
	WekSetFlag string `json:"wekSetFlag,omitempty"`
	// 每周限额-设置额度 BigDecimal(15,2)
	WekSetAmt string `json:"wekSetAmt,omitempty"`
	// 每周限额-重置日 String(1)
	WekRstDay string `json:"wekRstDay,omitempty"`
	// 每月限额-设置标志 String(1) required:Y Y：设置 N：不设置
	MthSetFlag string `json:"mthSetFlag,omitempty"`
	// 每月限额-设置额度 BigDecimal(15,2)
	MthSetAmt string `json:"mthSetAmt,omitempty"`
	// 每月限额-重置日 String(2)
	MthRstDay string `json:"mthRstDay,omitempty"`
	// 每季限额-设置标志 String(1) required:Y Y：设置 N：不设置
	QtrSetFlag string `json:"qtrSetFlag,omitempty"`
	// 每季限额-设置额度 BigDecimal(15,2)
	QtrSetAmt string `json:"qtrSetAmt,omitempty"`
	// 每季限额-重置月 String(1)
	QtrRstMth string `json:"qtrRstMth,omitempty"`
	// 每季限额-重置日 String(2)
	QtrRstDay string `json:"qtrRstDay,omitempty"`
	// 每年限额-设置标志 String(1) required:Y Y：设置 N：不设置
	YerSetFlag string `json:"yerSetFlag,omitempty"`
	// 每年限额-设置额度 BigDecimal(15,2)
	YerSetAmt string `json:"yerSetAmt,omitempty"`
	// 每年限额-重置月 String(2)
	YerRstMth string `json:"yerRstMth,omitempty"`
	// 每年限额-重置日 String(2)
	YerRstDay string `json:"yerRstDay,omitempty"`
	// 业务参考号 String(30) required:Y
	XYurRef string `json:"xYurRef,omitempty"`
	// 保留字50 String(50)
	XRsv50z string `json:"xRsv50z,omitempty"`
}

// UnitManagerNewBHHDLMADZ1 详细信息 BHHDLMADZ1.
type UnitManagerNewBHHDLMADZ1 struct {
	// 户口号 String(35)
	AccNbr string `json:"accNbr,omitempty"`
	// 子单元编号 String(20)
	DumNbr string `json:"dumNbr,omitempty"`
	// 单笔限额-设置标志 String(1) Y：设置 N：不设置
	SigSetFlag string `json:"sigSetFlag,omitempty"`
	// 单笔限额-设置额度 BigDecimal(15,2)
	SigSetAmt string `json:"sigSetAmt,omitempty"`
	// 每日限额-设置标志 String(1) Y：设置 N：不设置
	DaySetFlag string `json:"daySetFlag,omitempty"`
	// 每日限额-设置额度 BigDecimal(15,2)
	DaySetAmt string `json:"daySetAmt,omitempty"`
	// 每周限额-设置标志 String(1) Y：设置 N：不设置
	WekSetFlag string `json:"wekSetFlag,omitempty"`
	// 每周限额-设置额度 BigDecimal(15,2)
	WekSetAmt string `json:"wekSetAmt,omitempty"`
	// 每周限额-重置日 String(1)
	WekRstDay string `json:"wekRstDay,omitempty"`
	// 每月限额-设置标志 String(1) Y：设置 N：不设置
	MthSetFlag string `json:"mthSetFlag,omitempty"`
	// 每月限额-设置额度 BigDecimal(15,2)
	MthSetAmt string `json:"mthSetAmt,omitempty"`
	// 每月限额-重置日 String(2)
	MthRstDay string `json:"mthRstDay,omitempty"`
	// 每季限额-设置标志 String(1) Y：设置 N：不设置
	QtrSetFlag string `json:"qtrSetFlag,omitempty"`
	// 每季限额-设置额度 BigDecimal(15,2)
	QtrSetAmt string `json:"qtrSetAmt,omitempty"`
	// 每季限额-重置月 String(2)
	QtrRstMth string `json:"qtrRstMth,omitempty"`
	// 每季限额-重置日 String(2)
	QtrRstDay string `json:"qtrRstDay,omitempty"`
	// 每年限额-设置标志 String(1) Y：设置 N：不设置
	YerSetFlag string `json:"yerSetFlag,omitempty"`
	// 每年限额-设置额度 BigDecimal(15,2)
	YerSetAmt string `json:"yerSetAmt,omitempty"`
	// 每年限额-重置月 String(2)
	YerRstMth string `json:"yerRstMth,omitempty"`
	// 每年限额-重置日 String(2)
	YerRstDay string `json:"yerRstDay,omitempty"`
	// 业务参考号 String(30)
	YurRef string `json:"yurRef,omitempty"`
	// 保留字50 String(50)
	Rsv50z string `json:"rsv50z,omitempty"`
}

// UnitManagerNewBHHDLMCDX1 输入接口 BHHDLMCDX1.
type UnitManagerNewBHHDLMCDX1 struct {
	// 户口号 String(35) required:Y
	XAccNbr string `json:"xAccNbr,omitempty"`
	// 子单元编号 String(20) required:Y
	XDumNbr string `json:"xDumNbr,omitempty"`
	// 业务参考号 String(30) required:Y
	XYurRef string `json:"xYurRef,omitempty"`
	// 保留字50 String(50)
	XRsv50z string `json:"xRsv50z,omitempty"`
}

// UnitManagerNewBHHDLMCDZ1 详细信息 BHHDLMCDZ1.
type UnitManagerNewBHHDLMCDZ1 struct {
	// 户口号 String(35)
	AccNbr string `json:"accNbr,omitempty"`
	// 子单元编号 String(20)
	DumNbr string `json:"dumNbr,omitempty"`
	// 业务参考号 String(30)
	YurRef string `json:"yurRef,omitempty"`
	// 保留字50 String(50)
	Rsv50z string `json:"rsv50z,omitempty"`
}

// UnitManagerNewBHHDLMMDX1 输入接口 BHHDLMMDX1.
type UnitManagerNewBHHDLMMDX1 struct {
	// 户口号 String(35) required:Y
	XAccNbr string `json:"xAccNbr,omitempty"`
	// 子单元编号 String(20) required:Y
	XDumNbr string `json:"xDumNbr,omitempty"`
	// 单笔限额-是否维护 String(1) required:Y Y：维护 N：不维护
	SigMntTag string `json:"sigMntTag,omitempty"`
	// 单笔限额-设置标志 String(1) Y：设置 N：不设置
	SigSetFlag string `json:"sigSetFlag,omitempty"`
	// 单笔限额-设置额度 BigDecimal(15,2)
	SigSetAmt string `json:"sigSetAmt,omitempty"`
	// 每日限额-是否维护 String(1) required:Y Y：维护 N：不维护
	DayMntTag string `json:"dayMntTag,omitempty"`
	// 每日限额-设置标志 String(1) Y：设置 N：不设置
	DaySetFlag string `json:"daySetFlag,omitempty"`
	// 每日限额-设置额度 BigDecimal(15,2)
	DaySetAmt string `json:"daySetAmt,omitempty"`
	// 每周限额-是否维护 String(1) required:Y Y：维护 N：不维护
	WekMntTag string `json:"wekMntTag,omitempty"`
	// 每周限额-设置标志 String(1) Y：设置 N：不设置
	WekSetFlag string `json:"wekSetFlag,omitempty"`
	// 每周限额-设置额度 BigDecimal(15,2)
	WekSetAmt string `json:"wekSetAmt,omitempty"`
	// 每周限额-重置日 String(1)
	WekRstDay string `json:"wekRstDay,omitempty"`
	// 每月限额-是否维护 String(1) required:Y Y：维护 N：不维护
	MthMntTag string `json:"mthMntTag,omitempty"`
	// 每月限额-设置标志 String(1) Y：设置 N：不设置
	MthSetFlag string `json:"mthSetFlag,omitempty"`
	// 每月限额-设置额度 BigDecimal(15,2)
	MthSetAmt string `json:"mthSetAmt,omitempty"`
	// 每月限额-重置日 String(2)
	MthRstDay string `json:"mthRstDay,omitempty"`
	// 每季限额-是否维护 String(1) required:Y Y：维护 N：不维护
	QtrMntTag string `json:"qtrMntTag,omitempty"`
	// 每季限额-设置标志 String(1) Y：设置 N：不设置
	QtrSetFlag string `json:"qtrSetFlag,omitempty"`
	// 每季限额-设置额度 BigDecimal(15,2)
	QtrSetAmt string `json:"qtrSetAmt,omitempty"`
	// 每季限额-重置月 String(1)
	QtrRstMth string `json:"qtrRstMth,omitempty"`
	// 每季限额-重置日 String(2)
	QtrRstDay string `json:"qtrRstDay,omitempty"`
	// 每年限额-是否维护 String(1) required:Y Y：维护 N：不维护
	YerMntTag string `json:"yerMntTag,omitempty"`
	// 每年限额-设置标志 String(1) Y：设置 N：不设置
	YerSetFlag string `json:"yerSetFlag,omitempty"`
	// 每年限额-设置额度 BigDecimal(15,2)
	YerSetAmt string `json:"yerSetAmt,omitempty"`
	// 每年限额-重置月 String(2)
	YerRstMth string `json:"yerRstMth,omitempty"`
	// 每年限额-重置日 String(2)
	YerRstDay string `json:"yerRstDay,omitempty"`
	// 临时额度-是否维护 String(1) required:Y Y：维护 N：不维护
	TmpMntTag string `json:"tmpMntTag,omitempty"`
	// 临时额度-设置标志 String(1) Y：设置 N：不设置
	TmpSetFlag string `json:"tmpSetFlag,omitempty"`
	// 临时额度-设置额度 BigDecimal(15,2)
	TmpSetAmt string `json:"tmpSetAmt,omitempty"`
	// 临时额度-生效日期 String(8)
	TmpEftDate string `json:"tmpEftDate,omitempty"`
	// 临时额度-失效日期 String(8)
	TmpExpDate string `json:"tmpExpDate,omitempty"`
	// 业务参考号 String(30) required:Y
	XYurRef string `json:"xYurRef,omitempty"`
	// 保留字50 String(50)
	XRsv50z string `json:"xRsv50z,omitempty"`
}

// UnitManagerNewBHHDLMMDZ1 详细信息 BHHDLMMDZ1.
type UnitManagerNewBHHDLMMDZ1 struct {
	// 户口号 String(35)
	AccNbr string `json:"accNbr,omitempty"`
	// 子单元编号 String(20)
	DumNbr string `json:"dumNbr,omitempty"`
	// 单笔限额-是否维护 String(1) Y：维护 N：不维护
	SigMntTag string `json:"sigMntTag,omitempty"`
	// 单笔限额-设置标志 String(1) Y：设置 N：不设置
	SigSetFlag string `json:"sigSetFlag,omitempty"`
	// 单笔限额-设置额度 BigDecimal(15,2)
	SigSetAmt string `json:"sigSetAmt,omitempty"`
	// 每日限额-是否维护 String(1) Y：维护 N：不维护
	DayMntTag string `json:"dayMntTag,omitempty"`
	// 每日限额-设置标志 String(1) Y：设置 N：不设置
	DaySetFlag string `json:"daySetFlag,omitempty"`
	// 每日限额-设置额度 BigDecimal(15,2)
	DaySetAmt string `json:"daySetAmt,omitempty"`
	// 每周限额-是否维护 String(1) Y：维护 N：不维护
	WekMntTag string `json:"wekMntTag,omitempty"`
	// 每周限额-设置标志 String(1) Y：设置 N：不设置
	WekSetFlag string `json:"wekSetFlag,omitempty"`
	// 每周限额-设置额度 BigDecimal(15,2)
	WekSetAmt string `json:"wekSetAmt,omitempty"`
	// 每周限额-重置日 String(1)
	WekRstDay string `json:"wekRstDay,omitempty"`
	// 每月限额-是否维护 String(1) Y：维护 N：不维护
	MthMntTag string `json:"mthMntTag,omitempty"`
	// 每月限额-设置标志 String(1) Y：设置 N：不设置
	MthSetFlag string `json:"mthSetFlag,omitempty"`
	// 每月限额-设置额度 BigDecimal(15,2)
	MthSetAmt string `json:"mthSetAmt,omitempty"`
	// 每月限额-重置日 String(2)
	MthRstDay string `json:"mthRstDay,omitempty"`
	// 每季限额-是否维护 String(1) Y：维护 N：不维护
	QtrMntTag string `json:"qtrMntTag,omitempty"`
	// 每季限额-设置标志 String(1) Y：设置 N：不设置
	QtrSetFlag string `json:"qtrSetFlag,omitempty"`
	// 每季限额-设置额度 BigDecimal(15,2)
	QtrSetAmt string `json:"qtrSetAmt,omitempty"`
	// 每季限额-重置月 String(1)
	QtrRstMth string `json:"qtrRstMth,omitempty"`
	// 每季限额-重置日 String(2)
	QtrRstDay string `json:"qtrRstDay,omitempty"`
	// 每年限额-是否维护 String(1) Y：维护 N：不维护
	YerMntTag string `json:"yerMntTag,omitempty"`
	// 每年限额-设置标志 String(1) Y：设置 N：不设置
	YerSetFlag string `json:"yerSetFlag,omitempty"`
	// 每年限额-设置额度 BigDecimal(15,2)
	YerSetAmt string `json:"yerSetAmt,omitempty"`
	// 每年限额-重置月 String(2)
	YerRstMth string `json:"yerRstMth,omitempty"`
	// 每年限额-重置日 String(2)
	YerRstDay string `json:"yerRstDay,omitempty"`
	// 临时额度-是否维护 String(1) Y：维护 N：不维护
	TmpMntTag string `json:"tmpMntTag,omitempty"`
	// 临时额度-设置标志 String(1) Y：设置 N：不设置
	TmpSetFlag string `json:"tmpSetFlag,omitempty"`
	// 临时额度-设置额度 BigDecimal(15,2)
	TmpSetAmt string `json:"tmpSetAmt,omitempty"`
	// 临时额度-生效日期 String(8)
	TmpEftDate string `json:"tmpEftDate,omitempty"`
	// 临时额度-失效日期 String(8)
	TmpExpDate string `json:"tmpExpDate,omitempty"`
	// 业务参考号 String(30)
	YurRef string `json:"yurRef,omitempty"`
	// 保留字50 String(50)
	Rsv50z string `json:"rsv50z,omitempty"`
}

// UnitManagerNewBHHDLMQYY1 查询信息 BHHDLMQYY1.
type UnitManagerNewBHHDLMQYY1 struct {
	// 户口号 String(35) required:Y
	XAccNbr string `json:"xAccNbr,omitempty"`
	// 记账子单元编号 String(20) 传空查所有，传子单元则查指定子单元
	XDumNbr string `json:"xDumNbr,omitempty"`
	// 续传键值 String(80) 首次传空，当返回报文有此接口，取此接口放到下次查询中
	XCtnKey string `json:"xCtnKey,omitempty"`
}

// UnitManagerNewBHHDLMQYZ1 详细信息 BHHDLMQYZ1.
type UnitManagerNewBHHDLMQYZ1 struct {
	// 协议编号 String(10)
	CstNbr string `json:"cstNbr,omitempty"`
	// 户口号 String(35)
	EacNbr string `json:"eacNbr,omitempty"`
	// 记账子单元编号 String(20)
	DumNbr string `json:"dumNbr,omitempty"`
	// 记账子单元名称 String(200)
	DumNam string `json:"dumNam,omitempty"`
	// 单笔限额-设置标志 String(1) Y：设置 N：不设置
	SigSetFlag string `json:"sigSetFlag,omitempty"`
	// 单笔限额-设置额度 BigDecimal(15,2)
	SigSetAmt string `json:"sigSetAmt,omitempty"`
	// 每日限额-设置标志 String(1) Y：设置 N：不设置
	DaySetFlag string `json:"daySetFlag,omitempty"`
	// 每日限额-设置额度 BigDecimal(15,2)
	DaySetAmt string `json:"daySetAmt,omitempty"`
	// 每日限额-当期已用额度 BigDecimal(15,2)
	DayCurUsedAmt string `json:"dayCurUsedAmt,omitempty"`
	// 每日限额-当期可用额度 BigDecimal(15,2)
	DayCurAviAmt string `json:"dayCurAviAmt,omitempty"`
	// 每周限额-设置标志 String(1) Y：设置 N：不设置
	WekSetFlag string `json:"wekSetFlag,omitempty"`
	// 每周限额-设置额度 BigDecimal(15,2)
	WekSetAmt string `json:"wekSetAmt,omitempty"`
	// 每周限额-重置日 String(1)
	WekRstDay string `json:"wekRstDay,omitempty"`
	// 每周限额-当期已用额度 BigDecimal(15,2)
	WekCurUsedAmt string `json:"wekCurUsedAmt,omitempty"`
	// 每周限额-当期可用额度 BigDecimal(15,2)
	WekCurAviAmt string `json:"wekCurAviAmt,omitempty"`
	// 每月限额-设置标志 String(1) Y：设置 N：不设置
	MthSetFlag string `json:"mthSetFlag,omitempty"`
	// 每月限额-设置额度 BigDecimal(15,2)
	MthSetAmt string `json:"mthSetAmt,omitempty"`
	// 每月限额-重置日 String(2)
	MthRstDay string `json:"mthRstDay,omitempty"`
	// 每月限额-当期已用额度 BigDecimal(15,2)
	MthCurUsedAmt string `json:"mthCurUsedAmt,omitempty"`
	// 每月限额-当期可用额度 BigDecimal(15,2)
	MthCurAviAmt string `json:"mthCurAviAmt,omitempty"`
	// 每季限额-设置标志 String(1) Y：设置 N：不设置
	QtrSetFlag string `json:"qtrSetFlag,omitempty"`
	// 每季限额-重置月 String(1)
	QtrRstMth string `json:"qtrRstMth,omitempty"`
	// 每季限额-重置日 String(2)
	QtrRstDay string `json:"qtrRstDay,omitempty"`
	// 每季限额-当期已用额度 BigDecimal(15,2)
	QtrCurUsedAmt string `json:"qtrCurUsedAmt,omitempty"`
	// 每季限额-当期可用额度 BigDecimal(15,2)
	QtrCurAviAmt string `json:"qtrCurAviAmt,omitempty"`
	// 每年限额-设置标志 String(1) Y：设置 N：不设置
	YerSetFlag string `json:"yerSetFlag,omitempty"`
	// 每年限额-设置额度 BigDecimal(15,2)
	YerSetAmt string `json:"yerSetAmt,omitempty"`
	// 每年限额-重置月 String(2)
	YerRstMth string `json:"yerRstMth,omitempty"`
	// 每年限额-重置日 String(2)
	YerRstDay string `json:"yerRstDay,omitempty"`
	// 每年限额-当期已用额度 BigDecimal(15,2)
	YerCurUsedAmt string `json:"yerCurUsedAmt,omitempty"`
	// 每年限额-当期可用额度 BigDecimal(15,2)
	YerCurAviAmt string `json:"yerCurAviAmt,omitempty"`
	// 临时额度-设置标志 String(1) Y：设置 N：不设置
	TmpSetFlag string `json:"tmpSetFlag,omitempty"`
	// 临时额度-设置额度 BigDecimal(15,2)
	TmpSetAmt string `json:"tmpSetAmt,omitempty"`
	// 临时额度-生效日期 String(8)
	TmpEftDate string `json:"tmpEftDate,omitempty"`
	// 临时额度-失效日期 String(8)
	TmpExpDate string `json:"tmpExpDate,omitempty"`
	// 临时额度-已用额度 BigDecimal(15,2)
	TmpUsedAmt string `json:"tmpUsedAmt,omitempty"`
	// 临时额度-可用额度 BigDecimal(15,2)
	TmpAviAmt string `json:"tmpAviAmt,omitempty"`
}

// UnitManagerNewBHHDLMTPX1 输入接口 BHHDLMTPX1.
type UnitManagerNewBHHDLMTPX1 struct {
	// 户口号 String(35) required:Y
	XAccNbr string `json:"xAccNbr,omitempty"`
	// 子单元编号 String(20) required:Y
	XDumNbr string `json:"xDumNbr,omitempty"`
	// 临时限额-设置标志 String(1) required:Y Y：设置 N：取消
	TmpSetFlag string `json:"tmpSetFlag,omitempty"`
	// 临时限额-设置额度 BigDecimal(15,2)
	TmpSetAmt string `json:"tmpSetAmt,omitempty"`
	// 临时限额-生效日期 String(8)
	TmpEftDate string `json:"tmpEftDate,omitempty"`
	// 临时限额-失效日期 String(8)
	TmpExpDate string `json:"tmpExpDate,omitempty"`
	// 业务参考号 String(30) required:Y
	XYurRef string `json:"xYurRef,omitempty"`
	// 保留字50 String(50)
	XRsv50z string `json:"xRsv50z,omitempty"`
}

// UnitManagerNewBHHDLMTPZ1 详细信息 BHHDLMTPZ1.
type UnitManagerNewBHHDLMTPZ1 struct {
	// 户口号 String(35)
	AccNbr string `json:"accNbr,omitempty"`
	// 子单元编号 String(20)
	DumNbr string `json:"dumNbr,omitempty"`
	// 临时额度-设置标志 String(1) Y：设置 N：不设置
	TmpSetFlag string `json:"tmpSetFlag,omitempty"`
	// 临时额度-设置额度 BigDecimal(15,2)
	TmpSetAmt string `json:"tmpSetAmt,omitempty"`
	// 临时额度-生效日期 String(8)
	TmpEftDate string `json:"tmpEftDate,omitempty"`
	// 临时额度-失效日期 String(8)
	TmpExpDate string `json:"tmpExpDate,omitempty"`
	// 业务参考号 String(30)
	YurRef string `json:"yurRef,omitempty"`
	// 保留字50 String(50)
	Rsv50z string `json:"rsv50z,omitempty"`
}

// UnitManagerNewBHHDNAABZ 详细信息 BHHDNAABZ.
type UnitManagerNewBHHDNAABZ struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20)
	Dmanbr string `json:"dmanbr,omitempty"`
	// 记账子单元名称 Z(82)
	Dmanam string `json:"dmanam,omitempty"`
	// 透支控制标志 String(1)
	Ovrctl string `json:"ovrctl,omitempty"`
	// 退票处理方式 String(1)
	Bcktyp string `json:"bcktyp,omitempty"`
	// 余额关闭方式 String(1)
	Clstyp string `json:"clstyp,omitempty"`
	// 收款额度标志 String(1)
	Lmtflg string `json:"lmtflg,omitempty"`
	// 余额上限 BigDecimal(15,2)
	Ballmt string `json:"ballmt,omitempty"`
	// 序号 String(22)
	ZRdtSeqNbr string `json:"zRdtSeqNbr,omitempty"`
	// 请求状态 String(3)
	ZRdtReqStq string `json:"zRdtReqStq,omitempty"`
	// 处理结果 String(1)
	ZRdtRtnFlg string `json:"zRdtRtnFlg,omitempty"`
	// 错误描述 String(300)
	ZRdtErrTxt string `json:"zRdtErrTxt,omitempty"`
}

// UnitManagerNewBHHDNAAMZ 详细信息 BHHDNAAMZ.
type UnitManagerNewBHHDNAAMZ struct {
	// 账号 String(35)
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20)
	Dmanbr string `json:"dmanbr,omitempty"`
	// 记账子单元名称 Z(82)
	Dmanam string `json:"dmanam,omitempty"`
	// 透支控制标志 String(1)
	Ovrctl string `json:"ovrctl,omitempty"`
	// 退票处理方式 String(1)
	Bcktyp string `json:"bcktyp,omitempty"`
	// 余额关闭方式 String(1)
	Clstyp string `json:"clstyp,omitempty"`
	// 业务参考号 String(30)
	Yurref string `json:"yurref,omitempty"`
	// 收款额度标志 String(1)
	Lmtflg string `json:"lmtflg,omitempty"`
	// 余额上限 BigDecimal(15,2)
	Ballmt string `json:"ballmt,omitempty"`
}

// UnitManagerNewBHHDNCSAZ 详细信息 BHHDNCSAZ.
type UnitManagerNewBHHDNCSAZ struct {
	// 账号 String(35)
	Eacnbr string `json:"eacnbr,omitempty"`
	// 记账子单元编号识别后缀长度 String(2)
	Dumlen string `json:"dumlen,omitempty"`
	// 自定义记账子单元编号 String(1)
	Slfdef string `json:"slfdef,omitempty"`
	// 是否使用记账子单元 String(1)
	Dumflg string `json:"dumflg,omitempty"`
	// 入账检查 String(1)
	Rcvchk string `json:"rcvchk,omitempty"`
	// 业务参考号 String(30)
	Yurref string `json:"yurref,omitempty"`
}

// UnitManagerNewBHHDNCSBZ 详细信息 BHHDNCSBZ.
type UnitManagerNewBHHDNCSBZ struct {
	// 账号 String(35)
	Eacnbr string `json:"eacnbr,omitempty"`
	// 记账子单元编号识别后缀长度 String(2)
	Dumlen string `json:"dumlen,omitempty"`
	// 入账检查 String(1)
	Rcvchk string `json:"rcvchk,omitempty"`
	// 业务参考号 String(30)
	Yurref string `json:"yurref,omitempty"`
}

// UnitManagerNewBHHDNCSUZ 详细信息 BHHDNCSUZ.
type UnitManagerNewBHHDNCSUZ struct {
	// 账号 String(35)
	ZEacNbr string `json:"zEacNbr,omitempty"`
	// 自定义记账子单元编号 String(1)
	ZSlfDef string `json:"zSlfDef,omitempty"`
	// 入账检查 String(1)
	ZRcvChk string `json:"zRcvChk,omitempty"`
	// 业务参考号 String(30)
	ZYurRef string `json:"zYurRef,omitempty"`
}

// UnitManagerNewBHHDNCTZX1 输入接口 BHHDNCTZX1.
type UnitManagerNewBHHDNCTZX1 struct {
	// 主户口 String(35) required:Y
	XEacNbr string `json:"xEacNbr,omitempty"`
	// 开始日期 String(8) required:Y 开始日期必须小于等于结束日期 只能查13个月内的交易
	XBegDat string `json:"xBegDat,omitempty"`
	// 结束日期 String(8) required:Y 结束日期必须小于等于当前日期 开始结束日期不能超过100天
	XEndDat string `json:"xEndDat,omitempty"`
	// 查询类型 String(1) required:Y 1：查所有子单元，3：查指定子单元（传X2接口） 传空默认按照1查询。
	XQryTyp string `json:"xQryTyp,omitempty"`
	// 续传子单元 String(20) 第一次传空，当返回接口有续传接口并且该接口的续传字段有值，说明需要续传，取出续传字段放到下次请求报文中继续查询
	XCtnDum string `json:"xCtnDum,omitempty"`
	// 续传日期 String(8)
	XCtnDat string `json:"xCtnDat,omitempty"`
	// 续传交易序号 String(9)
	XCtnSeq string `json:"xCtnSeq,omitempty"`
}

// UnitManagerNewBHHDNCTZX2 输入接口 BHHDNCTZX2.
type UnitManagerNewBHHDNCTZX2 struct {
	// 记账子单元编号 String(20) required:Y
	XDumNbr string `json:"xDumNbr,omitempty"`
}

// UnitManagerNewBHHDNCTZZ1 详细信息 BHHDNCTZZ1.
type UnitManagerNewBHHDNCTZZ1 struct {
	// 协议编号 String(10)
	ZCstNbr string `json:"zCstNbr,omitempty"`
	// 主户口 String(35)
	ZEacNbr string `json:"zEacNbr,omitempty"`
	// 记账子单元编号 String(20)
	ZDumNbr string `json:"zDumNbr,omitempty"`
	// 记账子单元名称 String(200)
	ZDumNam string `json:"zDumNam,omitempty"`
	// 交易币种 String(2) 见云直联文档中心附录A.3货币代码表
	ZCcyNbr string `json:"zCcyNbr,omitempty"`
	// 交易金额 BigDecimal(15,2)
	ZTrsAmt string `json:"zTrsAmt,omitempty"`
	// 交易方向 String(1) C：贷 （收入） ，D：借 （支出） 冲账交易与原交易借贷相反
	ZTrsDir string `json:"zTrsDir,omitempty"`
	// 交易日期 String(8)
	ZTrsDat string `json:"zTrsDat,omitempty"`
	// 交易时间 String(6)
	ZTrsTim string `json:"zTrsTim,omitempty"`
	// 交易序号 String(9)
	ZTrsSeq string `json:"zTrsSeq,omitempty"`
	// 收/付方账号 String(35) 与交易方向对应
	ZRpyAcc string `json:"zRpyAcc,omitempty"`
	// 收/付方名称 String(400)
	ZRpyNam string `json:"zRpyNam,omitempty"`
	// 收/付方开户行行号 String(12)
	ZRpyBbn string `json:"zRpyBbn,omitempty"`
	// 收/付方银行名称 String(200)
	ZRpyBnk string `json:"zRpyBnk,omitempty"`
	// 收/付方银行地址 String(200)
	ZRpyAdr string `json:"zRpyAdr,omitempty"`
	// 交易类型 String(1) 已废弃，忽略该字段
	ZTrsTyp string `json:"zTrsTyp,omitempty"`
	// 摘要代码 String(8) 见云直联文档中心附录A.15交易代码表
	ZTxtCod string `json:"zTxtCod,omitempty"`
	// 交易摘要 String(200)
	ZTrsTxt string `json:"zTrsTxt,omitempty"`
	// 交易套号 String(15)
	ZTrsSet string `json:"zTrsSet,omitempty"`
	// 交易流水号 String(15)
	ZTrsNbr string `json:"zTrsNbr,omitempty"`
	// 关联交易套号 String(15) 退款、关联内部调配的交易返回，（可能会延迟返回） 可关联原交易
	ZRltSet string `json:"zRltSet,omitempty"`
	// 关联交易流水号 String(15) 退款、关联内部调配、记账/冲账中的【冲账】交易返回（记账/冲账会产生2笔交易，一笔冲账、一笔补账，只有冲账的那一笔交易会返回这个字段）（可能会延迟返回） 可关联原交易
	ZRltTrs string `json:"zRltTrs,omitempty"`
	// 冲正标志 String(1) N/空：正常交易，A：补账，R：冲账
	ZRvsTag string `json:"zRvsTag,omitempty"`
	// 记账方式 String(1) 1：自动，2：调账，3：内部调配，4：关联公司卡，5：结算户入账分类条件
	ZAutFlg string `json:"zAutFlg,omitempty"`
	// 记账标识 String(20) 若收款账号包含子单元编号该字段为子单元编号；若子单元关联公司卡且收款账号为公司卡时，则该字段为公司卡
	ZNarInn string `json:"zNarInn,omitempty"`
	// 匹配标志 String(1) Y：匹配，N：不匹配，空： 空表示未设置关联付方账号
	ZMthFlg string `json:"zMthFlg,omitempty"`
	// 业务类型 String(3) INT：退款退息的利息交易；空/ OTH：其他
	ZTrsMod string `json:"zTrsMod,omitempty"`
	// 业务状态 String(3) 退款交易返回最后一次退款状态 OPR：退款中 RED：退款成功 空：其他
	ZBusSts string `json:"zBusSts,omitempty"`
	// 交易状态 String(1) A：生效
	ZTrsSts string `json:"zTrsSts,omitempty"`
	// 余额 BigDecimal(15,2)
	ZOnlBal string `json:"zOnlBal,omitempty"`
	// 余额是否有效 String(1) Y：余额有效
	ZBalFlg string `json:"zBalFlg,omitempty"`
	// 母子公司账号 String(35) 资金池、财务公司代理清算、企业识别码 交易相关字段
	ZSonAcc string `json:"zSonAcc,omitempty"`
	// 母子公司名称 String(200)
	ZSonNam string `json:"zSonNam,omitempty"`
	// 母子公司开户行名称 String(200)
	ZSonBnk string `json:"zSonBnk,omitempty"`
	// 母子公司开户行地址 String(200)
	ZSonAdr string `json:"zSonAdr,omitempty"`
	// 业务名称 String(60) 网银业务相关字段
	ZBusNam string `json:"zBusNam,omitempty"`
	// 业务摘要 String(400)
	ZBusTxt string `json:"zBusTxt,omitempty"`
	// 网银流程实例号 String(20)
	ZReqNbr string `json:"zReqNbr,omitempty"`
	// 网银业务参考号 String(30)
	ZYurRef string `json:"zYurRef,omitempty"`
	// 商务支付订单号 String(50) 由商务支付订单产生
	ZOrdNbr string `json:"zOrdNbr,omitempty"`
}

// UnitManagerNewBHHDNCTZZ2 详细信息 BHHDNCTZZ2.
type UnitManagerNewBHHDNCTZZ2 struct {
	// 续传标志 String(1) Y：续传，N：不续传
	ZCtnFlg string `json:"zCtnFlg,omitempty"`
	// 续传子单元 String(20)
	ZCtnDum string `json:"zCtnDum,omitempty"`
	// 续传日期 String(8)
	ZCtnDat string `json:"zCtnDat,omitempty"`
	// 续传交易序号 String(9)
	ZCtnSeq string `json:"zCtnSeq,omitempty"`
}

// UnitManagerNewBHHDNDCBZ 详细信息 BHHDNDCBZ.
type UnitManagerNewBHHDNDCBZ struct {
	// 账号 String(35)
	ZAccNbr string `json:"zAccNbr,omitempty"`
	// 关联类型 String(1) 1：账号 2：户名
	ZChkTyp string `json:"zChkTyp,omitempty"`
	// 关联信息 Z（200）
	ZRltInf string `json:"zRltInf,omitempty"`
	// 业务参考号 String(30)
	ZYurRef string `json:"zYurRef,omitempty"`
	// 序号 String(22)
	ZRdtSeqNbr string `json:"zRdtSeqNbr,omitempty"`
	// 请求状态 String(3)
	ZRdtReqStq string `json:"zRdtReqStq,omitempty"`
	// 处理结果 String(1)
	ZRdtRtnFlg string `json:"zRdtRtnFlg,omitempty"`
	// 错误描述 String(270)
	ZRdtErrTxt string `json:"zRdtErrTxt,omitempty"`
	// 校验类型 String(1)
	ZTrsTyp string `json:"zTrsTyp,omitempty"`
}

// UnitManagerNewBHHDNDDBZ 详细信息 BHHDNDDBZ.
type UnitManagerNewBHHDNDDBZ struct {
	// 账号 String(35)
	ZAccNbr string `json:"zAccNbr,omitempty"`
	// 关联信息 Z（200）
	ZRltInf string `json:"zRltInf,omitempty"`
	// 业务参考号 String(30)
	ZYurRef string `json:"zYurRef,omitempty"`
	// 序号 String(22)
	ZRdtSeqNbr string `json:"zRdtSeqNbr,omitempty"`
	// 请求状态 String(3)
	ZRdtReqStq string `json:"zRdtReqStq,omitempty"`
	// 处理结果 String(1)
	ZRdtRtnFlg string `json:"zRdtRtnFlg,omitempty"`
	// 错误描述 String(300)
	ZRdtErrTxt string `json:"zRdtErrTxt,omitempty"`
}

// UnitManagerNewBHHDNDRAZ 详细信息 BHHDNDRAZ.
type UnitManagerNewBHHDNDRAZ struct {
	// 分行号 String(2)
	Bbknbr string `json:"bbknbr,omitempty"`
	// 账号 String(35)
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20)
	Dmanbr string `json:"dmanbr,omitempty"`
	// 入账方式 String(1)
	Tlyopr string `json:"tlyopr,omitempty"`
	// 付方账号 String(35)
	Dbtacc string `json:"dbtacc,omitempty"`
	// 付方户名 Z（200）
	Dbtnam string `json:"dbtnam,omitempty"`
	// 付方金额 BigDecimal(15,2)
	Dbtamt string `json:"dbtamt,omitempty"`
	// 业务参考号 String(30)
	Yurref string `json:"yurref,omitempty"`
	// 校验类型 String(1)
	Chktyp string `json:"chktyp,omitempty"`
}

// UnitManagerNewBHHDNDRBZ 详细信息 BHHDNDRBZ.
type UnitManagerNewBHHDNDRBZ struct {
	// 账号 String(35)
	Accnbr string `json:"accnbr,omitempty"`
	// 入账方式 String(1) R：拒绝入账 Y：入指定子单元 N：入默认子单元
	Tlyopr string `json:"tlyopr,omitempty"`
	// 业务参考号 String(30)
	Yurref string `json:"yurref,omitempty"`
	// 记账子单元编号 String(20)
	Dmanbr string `json:"dmanbr,omitempty"`
	// 付方账号 String(35)
	Dbtacc string `json:"dbtacc,omitempty"`
	// 付方户名 Z(200)
	Dbtnam string `json:"dbtnam,omitempty"`
	// 付方金额 BigDecimal(15,2)
	Dbtamt string `json:"dbtamt,omitempty"`
	// 序号 String(22)
	ZRdtSeqNbr string `json:"zRdtSeqNbr,omitempty"`
	// 请求状态 String(3)
	ZRdtReqStq string `json:"zRdtReqStq,omitempty"`
	// 处理结果 String(1)
	ZRdtRtnFlg string `json:"zRdtRtnFlg,omitempty"`
	// 错误描述 String(270)
	ZRdtErrTxt string `json:"zRdtErrTxt,omitempty"`
	// 校验类型 String(1)
	Chktyp string `json:"chktyp,omitempty"`
}

// UnitManagerNewBHHDNDRDZ 详细信息 BHHDNDRDZ.
type UnitManagerNewBHHDNDRDZ struct {
	// 分行号 String(2)
	Bbknbr string `json:"bbknbr,omitempty"`
	// 账号 String(35)
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20)
	Dmanbr string `json:"dmanbr,omitempty"`
	// 付方账号 String(35)
	Dbtacc string `json:"dbtacc,omitempty"`
	// 付方户名 Z（200）
	Dbtnam string `json:"dbtnam,omitempty"`
	// 业务参考号 String(30)
	Yurref string `json:"yurref,omitempty"`
	// 校验类型 String(1)
	Chktyp string `json:"chktyp,omitempty"`
}

// UnitManagerNewBHHDNDRMZ 详细信息 BHHDNDRMZ.
type UnitManagerNewBHHDNDRMZ struct {
	// 账号 String(35)
	Accnbr string `json:"accnbr,omitempty"`
	// 业务参考号 String(30)
	Yurref string `json:"yurref,omitempty"`
	// 记账子单元编号 String(20)
	Dmanbr string `json:"dmanbr,omitempty"`
	// 入账方式 String(1) R：拒绝入账 Y：入指定子单元 N：入默认子单元
	Tlyopr string `json:"tlyopr,omitempty"`
	// 序号 String(22)
	ZRdtSeqNbr string `json:"zRdtSeqNbr,omitempty"`
	// 请求状态 String(3)
	ZRdtReqStq string `json:"zRdtReqStq,omitempty"`
	// 处理结果 String(1)
	ZRdtRtnFlg string `json:"zRdtRtnFlg,omitempty"`
	// 错误描述 String(300)
	ZRdtErrTxt string `json:"zRdtErrTxt,omitempty"`
}

// UnitManagerNewBHHDNDRNZ 详细信息 BHHDNDRNZ.
type UnitManagerNewBHHDNDRNZ struct {
	// 分行号 String(2)
	ZBbkNbr string `json:"zBbkNbr,omitempty"`
	// 账号 String(35)
	ZAccNbr string `json:"zAccNbr,omitempty"`
	// 记账子单元编号 String(20)
	ZDmaNbr string `json:"zDmaNbr,omitempty"`
	// 入账方式 String(1)
	ZTlyOpr string `json:"zTlyOpr,omitempty"`
	// 业务参考号 String(30)
	ZYurRef string `json:"zYurRef,omitempty"`
}

// UnitManagerNewBHHDNDSAX1 输入接口 BHHDNDSAX1.
type UnitManagerNewBHHDNDSAX1 struct {
	// 户口号 String(35) required:Y
	XAccNbr string `json:"xAccNbr,omitempty"`
	// 付方信息 Z(200) required:Y
	XRltInf string `json:"xRltInf,omitempty"`
	// 白名单类型 String(1) required:Y 1：账号 2：户名
	XChkTyp string `json:"xChkTyp,omitempty"`
	// 业务参考号 String(30) required:Y
	XYurRef string `json:"xYurRef,omitempty"`
}

// UnitManagerNewBHHDNDSAZ 详细信息 BHHDNDSAZ.
type UnitManagerNewBHHDNDSAZ struct {
	// 账号 String(35)
	ZAccNbr string `json:"zAccNbr,omitempty"`
	// 关联类型 String(1)
	ZChkTyp string `json:"zChkTyp,omitempty"`
	// 关联信息 Z（200）
	ZRltInf string `json:"zRltInf,omitempty"`
	// 业务参考号 String(30)
	ZYurRef string `json:"zYurRef,omitempty"`
	// 交易类型 String(1)
	ZTrsTyp string `json:"zTrsTyp,omitempty"`
}

// UnitManagerNewBHHDNDSDX1 输入接口 BHHDNDSDX1.
type UnitManagerNewBHHDNDSDX1 struct {
	// 户口号 String(35) required:Y
	XAccNbr string `json:"xAccNbr,omitempty"`
	// 付方信息 Z(200) required:Y
	XRltInf string `json:"xRltInf,omitempty"`
	// 业务参考号 String(30) required:Y
	XYurRef string `json:"xYurRef,omitempty"`
}

// UnitManagerNewBHHDNDSDZ 详细信息 BHHDNDSDZ.
type UnitManagerNewBHHDNDSDZ struct {
	// 账号 String(35)
	ZAccNbr string `json:"zAccNbr,omitempty"`
	// 关联信息 Z（200）
	ZRltInf string `json:"zRltInf,omitempty"`
	// 业务参考号 String(30)
	ZYurRef string `json:"zYurRef,omitempty"`
}

// UnitManagerNewBHHDNDSMZ 详细信息 BHHDNDSMZ.
type UnitManagerNewBHHDNDSMZ struct {
	// 账号 String(35)
	ZAccNbr string `json:"zAccNbr,omitempty"`
	// 关联类型 String(1) P：个人，A：所有
	ZChkTyp string `json:"zChkTyp,omitempty"`
	// 关联信息 Z（200）
	ZRltInf string `json:"zRltInf,omitempty"`
	// 业务参考号 String(30)
	ZYurRef string `json:"zYurRef,omitempty"`
}

// UnitManagerNewBHHDNDSQY1 查询条件 BHHDNDSQY1.
type UnitManagerNewBHHDNDSQY1 struct {
	// 户口号 String(35) required:Y
	YAccNbr string `json:"yAccNbr,omitempty"`
	// 白名单类型 String(1) 1：账号 2：户名 （空查询所有类型）
	YChkTyp string `json:"yChkTyp,omitempty"`
	// 开始关联付方信息 Z(200) 续传字段，首次传空
	YRltInf string `json:"yRltInf,omitempty"`
}

// UnitManagerNewBHHDNDSQZ1 详细信息 BHHDNDSQZ1.
type UnitManagerNewBHHDNDSQZ1 struct {
	// 户口号 String(35）
	ZAccNbr string `json:"zAccNbr,omitempty"`
	// 户名 Z（200）
	ZAccNam string `json:"zAccNam,omitempty"`
	// 关联类型 String(3） PCS：个人账户白名单
	ZRltTyp string `json:"zRltTyp,omitempty"`
	// 关联付方信息 Z（200）
	ZRltInf string `json:"zRltInf,omitempty"`
	// 白名单类型 String（1） 1：账号 2：户名
	ZChkTyp string `json:"zChkTyp,omitempty"`
	// 不满足个人账户白名单的入账方式 String（1） R：拒绝入账
	ZTlyOpr string `json:"zTlyOpr,omitempty"`
	// 状态 String（1） A；生效 W：银行处理中
	ZStsCod string `json:"zStsCod,omitempty"`
}

// UnitManagerNewBhhdnrefz2 详细信息 bhhdnrefz2.
type UnitManagerNewBhhdnrefz2 struct {
	// 流程实例号 String(10)
	Reqnbr string `json:"reqnbr,omitempty"`
	// 错误结果描述 String(192)
	Rtnnar string `json:"rtnnar,omitempty"`
	// 保留字段 String(200)
	Rsv200 string `json:"rsv200,omitempty"`
}

// UnitManagerNewBHHDNRLBZ 详细信息 BHHDNRLBZ.
type UnitManagerNewBHHDNRLBZ struct {
	// 账号 String(35)
	Accnbr string `json:"accnbr,omitempty"`
	// 业务参考号 String(30)
	Yurref string `json:"yurref,omitempty"`
	// 记账子单元编号 String(20)
	Dmanbr string `json:"dmanbr,omitempty"`
	// 付方账号 String(35)
	Sndacc string `json:"sndacc,omitempty"`
	// 付方户名 Z(200)
	Sndnam string `json:"sndnam,omitempty"`
	// 白名单校验类型 String(1) 1：账号 2：户名
	Chktyp string `json:"chktyp,omitempty"`
	// 入账方式 String(1) R：拒绝入账 N：入默认子单元
	Rcvtyp string `json:"rcvtyp,omitempty"`
	// 序号 String(22)
	ZRdtSeqNbr string `json:"zRdtSeqNbr,omitempty"`
	// 请求状态 String(3)
	ZRdtReqStq string `json:"zRdtReqStq,omitempty"`
	// 处理结果 String(1)
	ZRdtRtnFlg string `json:"zRdtRtnFlg,omitempty"`
	// 错误描述 String(300)
	ZRdtErrTxt string `json:"zRdtErrTxt,omitempty"`
}

// UnitManagerNewBHHDNWCAX1 输入接口 BHHDNWCAX1.
type UnitManagerNewBHHDNWCAX1 struct {
	// 账号 String(35) required:Y
	XAccNbr string `json:"xAccNbr,omitempty"`
	// 虚拟户编号 String(20) required:Y
	XDumNbr string `json:"xDumNbr,omitempty"`
	// 分类条件校验类型 String(1) required:Y 1：帐号 2：户名
	XChkTyp string `json:"xChkTyp,omitempty"`
	// 收方键值 String(200) required:Y
	XRcvKey string `json:"xRcvKey,omitempty"`
	// 业务参考号 String(30) required:Y
	XYurRef string `json:"xYurRef,omitempty"`
}

// UnitManagerNewBHHDNWCAZ1 详细信息 BHHDNWCAZ1.
type UnitManagerNewBHHDNWCAZ1 struct {
	// 虚拟户编号 String(20)
	ZDumNbr string `json:"zDumNbr,omitempty"`
	// 分类条件校验类型 String(1)
	ZChkTyp string `json:"zChkTyp,omitempty"`
	// 收方键值 String(200)
	ZRcvKey string `json:"zRcvKey,omitempty"`
	// 保留字 String(50)
	ZRsv50z string `json:"zRsv50z,omitempty"`
}

// UnitManagerNewBHHDNWCCX1 输入接口 BHHDNWCCX1.
type UnitManagerNewBHHDNWCCX1 struct {
	// 账号 String(35) required:Y
	XAccNbr string `json:"xAccNbr,omitempty"`
	// 虚拟户编号 String(20) required:Y
	XDumNbr string `json:"xDumNbr,omitempty"`
	// 收方键值 String(200) required:Y
	XRcvKey string `json:"xRcvKey,omitempty"`
	// 业务参考号 String(30) required:Y
	XYurRef string `json:"xYurRef,omitempty"`
}

// UnitManagerNewBHHDNWCCZ1 详细信息 BHHDNWCCZ1.
type UnitManagerNewBHHDNWCCZ1 struct {
	// 虚拟户编号 String(20)
	ZDumNbr string `json:"zDumNbr,omitempty"`
	// 收方键值 String(200)
	ZRcvKey string `json:"zRcvKey,omitempty"`
	// 业务参考号 String(30)
	ZYurRef string `json:"zYurRef,omitempty"`
	// 保留字 String(50)
	ZRsv50z string `json:"zRsv50z,omitempty"`
}

// UnitManagerNewBHHDNWCQX1 输入接口 BHHDNWCQX1.
type UnitManagerNewBHHDNWCQX1 struct {
	// 账号 String(35) required:Y
	XAccNbr string `json:"xAccNbr,omitempty"`
	// 记账子单元 String(20)
	XDumNbr string `json:"xDumNbr,omitempty"`
}

// UnitManagerNewBHHDNWCQY1 续传接口 BHHDNWCQY1.
type UnitManagerNewBHHDNWCQY1 struct {
	// 续传键值 String(300)
	YCtnKey string `json:"yCtnKey,omitempty"`
}

// UnitManagerNewBHHDNWCQZ1 输出接口 BHHDNWCQZ1.
type UnitManagerNewBHHDNWCQZ1 struct {
	// 记账子单元编号 String(20)
	ZDumNbr string `json:"zDumNbr,omitempty"`
	// 分类条件校验类型 String(1) 1：账号，2：户名
	ZChkTyp string `json:"zChkTyp,omitempty"`
	// 收方键值 String(200)
	ZRcvKey string `json:"zRcvKey,omitempty"`
	// 出账方式 String(1) 目前只有N，N：默认子单元，
	ZRcvTyp string `json:"zRcvTyp,omitempty"`
	// 关联日期 String(8)
	ZRltDat string `json:"zRltDat,omitempty"`
	// 关联时间 String(6)
	ZRltTim string `json:"zRltTim,omitempty"`
	// 状态 String(1) A：生效 W：银行处理中
	ZRltSts string `json:"zRltSts,omitempty"`
}

// UnitManagerNewBHHDNWQCY1 查询条件 BHHDNWQCY1.
type UnitManagerNewBHHDNWQCY1 struct {
	// 账号 String(35) required:Y
	YAccNbr string `json:"yAccNbr,omitempty"`
	// 子单元编号 String(20) 为空查全部，传子单元则查指定子单元
	YDmaNbr string `json:"yDmaNbr,omitempty"`
	// 续传键值 String(300) 第一次传空，当返回接口有本接口并且该接口的续传字段有值，说明需要续传，取出续传字段放到下次请求报文中继续查询
	YCtnKey string `json:"yCtnKey,omitempty"`
}

// UnitManagerNewBHHDNWQCZ1 详细信息 BHHDNWQCZ1.
type UnitManagerNewBHHDNWQCZ1 struct {
	// 记账子单元编号 String(20) required:Y
	ZDmaNbr string `json:"zDmaNbr,omitempty"`
	// 付方账号 String(35)
	ZSndAcc string `json:"zSndAcc,omitempty"`
	// 付方户名/客户摘要 Z(200)
	ZSndNam string `json:"zSndNam,omitempty"`
	// 分类条件校验类型 String(1) required:Y 1：帐号 2：户名 4、客户摘要
	ZChkTyp string `json:"zChkTyp,omitempty"`
	// 关联状态 String(1) required:Y A：生效 W：银行处理中
	ZRltSts string `json:"zRltSts,omitempty"`
	// 不满足结算户入账分类条件的入账方式 String(1) required:Y N：非关联收款入账默认子单元 R：非关联收款拒绝入账
	ZRcvTyp string `json:"zRcvTyp,omitempty"`
}

// UnitManagerNewBHHDRADDZ 详细信息 BHHDRADDZ.
type UnitManagerNewBHHDRADDZ struct {
	// 账号 String(35)
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20)
	Dmanbr string `json:"dmanbr,omitempty"`
	// 业务参考号 String(30)
	Yurref string `json:"yurref,omitempty"`
	// 公司卡号 String(35)
	Cdsnbr string `json:"cdsnbr,omitempty"`
	// 序号 String(22)
	ZRdtSeqNbr string `json:"zRdtSeqNbr,omitempty"`
	// 请求状态 String(3)
	ZRdtReqStq string `json:"zRdtReqStq,omitempty"`
	// 处理结果 String(1)
	ZRdtRtnFlg string `json:"zRdtRtnFlg,omitempty"`
	// 错误描述 String(300)
	ZRdtErrTxt string `json:"zRdtErrTxt,omitempty"`
}

// UnitManagerNewBHHDRDLTZ 详细信息 BHHDRDLTZ.
type UnitManagerNewBHHDRDLTZ struct {
	// 账号 String(35)
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20)
	Dmanbr string `json:"dmanbr,omitempty"`
	// 业务参考号 String(30)
	Yurref string `json:"yurref,omitempty"`
	// 公司卡号 String(35)
	Cdsnbr string `json:"cdsnbr,omitempty"`
	// 序号 String(22)
	ZRdtSeqNbr string `json:"zRdtSeqNbr,omitempty"`
	// 请求状态 String(3)
	ZRdtReqStq string `json:"zRdtReqStq,omitempty"`
	// 处理结果 String(1)
	ZRdtRtnFlg string `json:"zRdtRtnFlg,omitempty"`
	// 错误描述 String(300)
	ZRdtErrTxt string `json:"zRdtErrTxt,omitempty"`
}

// UnitManagerNewBHHDTAJCZ 详细信息 BHHDTAJCZ.
type UnitManagerNewBHHDTAJCZ struct {
	// 账号 String(35)
	Accnbr string `json:"accnbr,omitempty"`
	// 关联虚拟户编号 String(20)
	Dmanbr string `json:"dmanbr,omitempty"`
	// 业务参考号 String(30)
	Yurref string `json:"yurref,omitempty"`
	// 交易日期 D(8)
	Trxdat string `json:"trxdat,omitempty"`
	// 记账流水号 String(15)
	Trxnbr string `json:"trxnbr,omitempty"`
	// 序号 String(22)
	ZRdtSeqNbr string `json:"zRdtSeqNbr,omitempty"`
	// 请求状态 String(3)
	ZRdtReqStq string `json:"zRdtReqStq,omitempty"`
	// 处理结果 String(1)
	ZRdtRtnFlg string `json:"zRdtRtnFlg,omitempty"`
	// 错误描述 String(300)
	ZRdtErrTxt string `json:"zRdtErrTxt,omitempty"`
}

// UnitManagerNewBHHDTAJDZ 详细信息 BHHDTAJDZ.
type UnitManagerNewBHHDTAJDZ struct {
	// 账号 String(35)
	Accnbr string `json:"accnbr,omitempty"`
	// 业务参考号 String(30)
	Yurref string `json:"yurref,omitempty"`
	// 原虚拟户编号 String(20)
	Dmanbr string `json:"dmanbr,omitempty"`
	// 交易日期 D(8)
	Trxdat string `json:"trxdat,omitempty"`
	// 记账流水号 String(15)
	Trxnbr string `json:"trxnbr,omitempty"`
	// 序号 String(22)
	ZRdtSeqNbr string `json:"zRdtSeqNbr,omitempty"`
	// 请求状态 String(3)
	ZRdtReqStq string `json:"zRdtReqStq,omitempty"`
	// 处理结果 String(1)
	ZRdtRtnFlg string `json:"zRdtRtnFlg,omitempty"`
	// 错误描述 String(300)
	ZRdtErrTxt string `json:"zRdtErrTxt,omitempty"`
}

// UnitManagerNewBHHDTITXZ 详细信息 BHHDTITXZ.
type UnitManagerNewBHHDTITXZ struct {
	// 账号 String(35)
	Accnbr string `json:"accnbr,omitempty"`
	// 借方子单元编号 String(20)
	Dmadbt string `json:"dmadbt,omitempty"`
	// 贷方子单元编号 String(20)
	Dmacrt string `json:"dmacrt,omitempty"`
	// 交易金额 M(15)
	Trxamt string `json:"trxamt,omitempty"`
	// 交易摘要 Z(42)
	Trxtxt string `json:"trxtxt,omitempty"`
	// 业务参考号 String(30)
	Yurref string `json:"yurref,omitempty"`
}

// UnitManagerNewBhhdurtnz1 输出接口 bhhdurtnz1.
type UnitManagerNewBhhdurtnz1 struct {
	// 流水号 String(10)
	ZSqrNbr string `json:"zSqrNbr,omitempty"`
	// 流程实例号 String(10)
	ZReqNbr string `json:"zReqNbr,omitempty"`
	// 请求状态 String(3) BNK：银行处理中 FIN：完成
	ZReqSts string `json:"zReqSts,omitempty"`
	// 待处理操作序列 String(3)
	ZOprSqn string `json:"zOprSqn,omitempty"`
	// 操作别名 String(32)
	ZOprAls string `json:"zOprAls,omitempty"`
	// 业务处理结果 String(1) S：成功 F：失败
	ZRtnFlg string `json:"zRtnFlg,omitempty"`
	// 错误码 String(7)
	ZErrCod string `json:"zErrCod,omitempty"`
	// 错误文本 String(92)
	ZErrTxt string `json:"zErrTxt,omitempty"`
	// 保留字30 String(30)
	ZRsv30z string `json:"zRsv30z,omitempty"`
}

// UnitManagerNewBHHDURTNZ1 输出接口 BHHDURTNZ1.
type UnitManagerNewBHHDURTNZ1 struct {
	// 流水号 String(10)
	ZSqrNbr string `json:"zSqrNbr,omitempty"`
	// 流程实例号 String(10)
	ZReqNbr string `json:"zReqNbr,omitempty"`
	// 请求状态 String(3) BNK：银行处理中 FIN：完成
	ZReqSts string `json:"zReqSts,omitempty"`
	// 待处理操作序列 String(3)
	ZOprSqn string `json:"zOprSqn,omitempty"`
	// 操作别名 String(32)
	ZOprAls string `json:"zOprAls,omitempty"`
	// 业务处理结果 String(1) S：成功 F：失败
	ZRtnFlg string `json:"zRtnFlg,omitempty"`
	// 错误码 String(7)
	ZErrCod string `json:"zErrCod,omitempty"`
	// 错误文本 String(92)
	ZErrTxt string `json:"zErrTxt,omitempty"`
	// 保留字30 String(30)
	ZRsv30z string `json:"zRsv30z,omitempty"`
}

// UnitManagerNewBhhdwmntx1 输入接口 bhhdwmntx1.
type UnitManagerNewBhhdwmntx1 struct {
	// 账号 String(35) required:Y
	XEacNbr string `json:"xEacNbr,omitempty"`
	// 不满足结算户入账分类条件的入账方式 String(1) required:Y N：非关联收款入账默认子单元 R：非关联收款拒绝入账
	XRcvTyp string `json:"xRcvTyp,omitempty"`
	// 业务参考号 String(30) required:Y 客户保证唯一
	XYurRef string `json:"xYurRef,omitempty"`
}

// UnitManagerNewBHHDWMNTZ 详细信息 BHHDWMNTZ.
type UnitManagerNewBHHDWMNTZ struct {
	// 账号 String(35)
	XEacNbr string `json:"xEacNbr,omitempty"`
	// 入账方式 String(1)
	XRcvTyp string `json:"xRcvTyp,omitempty"`
	// 业务参考号 String(30)
	XYurRef string `json:"xYurRef,omitempty"`
}

// UnitManagerNewBHHDYBADZ1 详细信息 BHHDYBADZ1.
type UnitManagerNewBHHDYBADZ1 struct {
	// 账号 String(35)
	ZAccNbr string `json:"zAccNbr,omitempty"`
	// 记账子单元编号 String(20)
	ZDumNbr string `json:"zDumNbr,omitempty"`
	// 公司卡号 String(35)
	ZCdsNbr string `json:"zCdsNbr,omitempty"`
}

// UnitManagerNewBHHDYDLBZ1 详细信息 BHHDYDLBZ1.
type UnitManagerNewBHHDYDLBZ1 struct {
	// 账号 String(35)
	ZAccNbr string `json:"zAccNbr,omitempty"`
	// 记账子单元编号 String(20)
	ZDumNbr string `json:"zDumNbr,omitempty"`
	// 公司卡号 String(35)
	ZCdsNbr string `json:"zCdsNbr,omitempty"`
}

// UnitManagerNewBHHDYRDBZ1 详细信息 BHHDYRDBZ1.
type UnitManagerNewBHHDYRDBZ1 struct {
	// 账号 String(35)
	ZAccNbr string `json:"zAccNbr,omitempty"`
	// 记账子单元编号 String(20)
	ZDumNbr string `json:"zDumNbr,omitempty"`
	// 关联信息 String(200)
	ZRltInf string `json:"zRltInf,omitempty"`
}

// UnitManagerNewBHHDZABMZ2 详细信息 BHHDZABMZ2.
type UnitManagerNewBHHDZABMZ2 struct {
	// 账号 String(35)
	ZAccNbr string `json:"zAccNbr,omitempty"`
	// 记账子单元编号 String(20)
	ZDumNbr string `json:"zDumNbr,omitempty"`
	// 是否修改透支控制 String(1) Y：修改，N：不修改
	ZOvrFlg string `json:"zOvrFlg,omitempty"`
	// 透支控制标志 String(1)
	ZOvrCtl string `json:"zOvrCtl,omitempty"`
	// 是否修改退票方式 String(1) Y：修改，N：不修改
	ZBckFlg string `json:"zBckFlg,omitempty"`
	// 退票处理方式 String(1)
	ZBckTyp string `json:"zBckTyp,omitempty"`
	// 是否修改关闭方式 String(1) Y：修改，N：不修改
	ZClsFlg string `json:"zClsFlg,omitempty"`
	// 余额关闭方式 String(1)
	ZClsTyp string `json:"zClsTyp,omitempty"`
	// 请求状态 String(3)
	ZReqSta string `json:"zReqSta,omitempty"`
	// 处理结果 String(1)
	ZRtnFlg string `json:"zRtnFlg,omitempty"`
	// 错误描述 String(300)
	ZErrTxt string `json:"zErrTxt,omitempty"`
}

// UnitManagerNewBHHDZCTBZ1 详细信息 BHHDZCTBZ1.
type UnitManagerNewBHHDZCTBZ1 struct {
	// 账号 String(35)
	ZAccNbr string `json:"zAccNbr,omitempty"`
	// 记账子单元编号 String(20)
	ZDumNbr string `json:"zDumNbr,omitempty"`
	// 请求状态 String(3)
	ZReqSta string `json:"zReqSta,omitempty"`
	// 处理结果 String(1)
	ZRtnFlg string `json:"zRtnFlg,omitempty"`
	// 错误描述 String(300)
	ZErrTxt string `json:"zErrTxt,omitempty"`
	// 收款限制 String(1) Y：设置/取消 N：不设置/不取消 空默认Y
	ZRcvFlg string `json:"zRcvFlg,omitempty"`
	// 付款限制 String(1) Y：设置/取消 N：不设置/不取消 空默认N
	ZPacFlg string `json:"zPacFlg,omitempty"`
}

// UnitManagerNewBHHDZCTGX1 输入接口 BHHDZCTGX1.
type UnitManagerNewBHHDZCTGX1 struct {
	// 户口号 String(35) required:Y
	XEacNbr string `json:"xEacNbr,omitempty"`
	// 分组编号 String(20) required:Y
	XGrpOut string `json:"xGrpOut,omitempty"`
	// 开始日期 String(8) required:Y
	XBegDat string `json:"xBegDat,omitempty"`
	// 结束日期 String(8) required:Y
	XEndDat string `json:"xEndDat,omitempty"`
}

// UnitManagerNewBHHDZCTGY1 续传接口 BHHDZCTGY1.
type UnitManagerNewBHHDZCTGY1 struct {
	// 续传键值 String(80)
	YCtnKey string `json:"yCtnKey,omitempty"`
}

// UnitManagerNewBHHDZCTGZ1 输出接口 BHHDZCTGZ1.
type UnitManagerNewBHHDZCTGZ1 struct {
	// 协议编号 String(10)
	ZCstNbr string `json:"zCstNbr,omitempty"`
	// 主户口 String(35)
	ZEacNbr string `json:"zEacNbr,omitempty"`
	// 记账子单元编号 String(20)
	ZDumNbr string `json:"zDumNbr,omitempty"`
	// 记账子单元名称 String(200)
	ZDumNam string `json:"zDumNam,omitempty"`
	// 交易币种 String(2) 见云直联文档中心附录A.3货币代码表
	ZCcyNbr string `json:"zCcyNbr,omitempty"`
	// 交易金额 BigDecimal(15,2)
	ZTrsAmt string `json:"zTrsAmt,omitempty"`
	// 交易方向 String(1) C：贷 D：借
	ZTrsDir string `json:"zTrsDir,omitempty"`
	// 交易日期 String(8)
	ZTrsDat string `json:"zTrsDat,omitempty"`
	// 交易时间 String(6)
	ZTrsTim string `json:"zTrsTim,omitempty"`
	// 交易序号 String(9)
	ZTrsSeq string `json:"zTrsSeq,omitempty"`
	// 收/付方账号 String(35)
	ZRpyAcc string `json:"zRpyAcc,omitempty"`
	// 收/付方名称 String(400)
	ZRpyNam string `json:"zRpyNam,omitempty"`
	// 收/付方开户行行号 String(12)
	ZRpyBbn string `json:"zRpyBbn,omitempty"`
	// 收/付方银行名称 String(200)
	ZRpyBnk string `json:"zRpyBnk,omitempty"`
	// 收/付方银行地址 String(200)
	ZRpyAdr string `json:"zRpyAdr,omitempty"`
	// 交易类型 String(1) F：金融交易
	ZTrsTyp string `json:"zTrsTyp,omitempty"`
	// 摘要代码 String(8)
	ZTxtCod string `json:"zTxtCod,omitempty"`
	// 交易摘要 String(200)
	ZTrsTxt string `json:"zTrsTxt,omitempty"`
	// 交易套号 String(15)
	ZTrsSet string `json:"zTrsSet,omitempty"`
	// 交易流水号 String(15)
	ZTrsNbr string `json:"zTrsNbr,omitempty"`
	// 关联交易套号 String(15) 退款、关联内部调配的交易返回， 可关联原交易
	ZRltSet string `json:"zRltSet,omitempty"`
	// 关联交易流水号 String(15) 退款、关联内部调配、记账/冲账中的【冲账】交易返回（记账/冲账会产生2笔交易，一笔冲账、一笔补账，只有冲账的那一笔交易会返回这个字段） 可关联原交易
	ZRltTrs string `json:"zRltTrs,omitempty"`
	// 冲正标志 String(1)
	ZRvsTag string `json:"zRvsTag,omitempty"`
	// 记账方式 String(1) 1：自动，2：手工，3：内部调配，4：关联公司卡，5：结算户入账分类条件
	ZAutFlg string `json:"zAutFlg,omitempty"`
	// 记账标识 String(20)
	ZNarInn string `json:"zNarInn,omitempty"`
	// 匹配标志 String(1) 空表示未设置关联付方账号
	ZMthFlg string `json:"zMthFlg,omitempty"`
	// 记账类型 String(3) INT：退款退息的利息交易；空/OTH：其他交易；
	ZTrsMod string `json:"zTrsMod,omitempty"`
	// 退款状态 String(3) 退款交易返回 OPR：退款中，RED：已退款，
	ZBusSts string `json:"zBusSts,omitempty"`
	// 交易状态 String(1) A：成功
	ZTrsSts string `json:"zTrsSts,omitempty"`
	// 余额 BigDecimal(15,2)
	ZOnlBal string `json:"zOnlBal,omitempty"`
	// 余额是否有效 String(1) Y：余额有效
	ZBalFlg string `json:"zBalFlg,omitempty"`
}

// UnitManagerNewBHHDZDBQX1 查询条件 BHHDZDBQX1.
type UnitManagerNewBHHDZDBQX1 struct {
	// 账号 String(35) required:Y
	XEacNbr string `json:"xEacNbr,omitempty"`
	// 记账子单元 String(20)
	XDumNbr string `json:"xDumNbr,omitempty"`
	// 退票开始日期 String(8) required:Y
	XRtnBeg string `json:"xRtnBeg,omitempty"`
	// 退票结束日期 String(8) required:Y
	XRtnEnd string `json:"xRtnEnd,omitempty"`
	// 业务参考号 String(30) 退款交易的业务参考号
	XYurRef string `json:"xYurRef,omitempty"`
	// 经办开始日期 String(8) 如果业务参考号非空，则必填
	XOprBeg string `json:"xOprBeg,omitempty"`
	// 经办结束日期 String(8) 如果业务参考号非空，则必填
	XOprEnd string `json:"xOprEnd,omitempty"`
}

// UnitManagerNewBHHDZDBQY1 续传接口 BHHDZDBQY1.
type UnitManagerNewBHHDZDBQY1 struct {
	// 续传退票日期 String(8) 如返回续传接口，取出该接口放入下次请求中继续查询。
	YCtnDat string `json:"yCtnDat,omitempty"`
	// 续传系统键值 String(40)
	YCtnKey string `json:"yCtnKey,omitempty"`
}

// UnitManagerNewBHHDZDBQZ1 详细信息 BHHDZDBQZ1.
type UnitManagerNewBHHDZDBQZ1 struct {
	// 协议编号 String(10) required:Y
	ZCstNbr string `json:"zCstNbr,omitempty"`
	// 主户口 String(35) required:Y
	ZEacNbr string `json:"zEacNbr,omitempty"`
	// 记账子单元编号 String(20) required:Y
	ZDumNbr string `json:"zDumNbr,omitempty"`
	// 对方系统代码 String(4)
	ZSysCod string `json:"zSysCod,omitempty"`
	// 业务参考号 String(30)
	ZYurRef string `json:"zYurRef,omitempty"`
	// 操作类型 String(3) required:Y DMP：记账子单元付款 DMR：记账子单元原路退回
	ZOprTyp string `json:"zOprTyp,omitempty"`
	// 经办日期 String(8) required:Y
	ZOprDat string `json:"zOprDat,omitempty"`
	// 退票日期 String(8) required:Y
	ZRtnDat string `json:"zRtnDat,omitempty"`
	// 退票时间 String(6) required:Y
	ZRtnTim string `json:"zRtnTim,omitempty"`
	// 退票原因 Z(300) required:Y
	ZErrTxt string `json:"zErrTxt,omitempty"`
	// 维护用户 String(30)
	ZMntUsr string `json:"zMntUsr,omitempty"`
	// 状态 String(1) required:Y B：退票
	ZRltSts string `json:"zRltSts,omitempty"`
	// 金额 BigDecimal(15, 2) required:Y
	ZTrsAmt string `json:"zTrsAmt,omitempty"`
	// 退票流水号 String(15)
	ZRtnIdn string `json:"zRtnIdn,omitempty"`
	// 退票套号 String(15) required:Y
	ZRtnSet string `json:"zRtnSet,omitempty"`
	// 退票状态 String(1) required:Y R：退票
	ZRtnSts string `json:"zRtnSts,omitempty"`
	// 收付方 账号 String(35) required:Y
	ZRpyAcc string `json:"zRpyAcc,omitempty"`
	// 收付方 户名 Z(400) required:Y
	ZRpyNam string `json:"zRpyNam,omitempty"`
	// 收付方 联行号 String(12)
	ZRpyBbn string `json:"zRpyBbn,omitempty"`
	// 收付方 银行名称 Z(400)
	ZRpyBnk string `json:"zRpyBnk,omitempty"`
	// 收付方 银行地址 Z(400)
	ZRpyAdr string `json:"zRpyAdr,omitempty"`
}

// UnitManagerNewBHHDZDRTZ1 详细信息 BHHDZDRTZ1.
type UnitManagerNewBHHDZDRTZ1 struct {
	// 账号 String(35)
	ZAccNbr string `json:"zAccNbr,omitempty"`
	// 记账子单元编号 String(20)
	ZDumNbr string `json:"zDumNbr,omitempty"`
	// 关联信息 Z(200) 有值：校验类型为：账号/户名 无值：校验类型为：金额
	ZRltInf string `json:"zRltInf,omitempty"`
	// 请求状态 String(3)
	ZRdtReqStq string `json:"zRdtReqStq,omitempty"`
	// 处理结果 String(1)
	ZRdtRtnFlg string `json:"zRdtRtnFlg,omitempty"`
	// 错误描述 String(300)
	ZRdtErrTxt string `json:"zRdtErrTxt,omitempty"`
}

// UnitManagerNewBHHDZDTBZ2 详细信息 BHHDZDTBZ2.
type UnitManagerNewBHHDZDTBZ2 struct {
	// 户口号 String(35)
	ZEacNbr string `json:"zEacNbr,omitempty"`
	// 付方记账子单元编号 String(20)
	ZDmaDbt string `json:"zDmaDbt,omitempty"`
	// 收方记账子单元编号 String(20)
	ZDmaCrt string `json:"zDmaCrt,omitempty"`
	// 交易金额 BigDecimal(15,2)
	ZTrsAmt string `json:"zTrsAmt,omitempty"`
	// 交易摘要 String(60)
	ZTrsTxt string `json:"zTrsTxt,omitempty"`
	// 请求状态 String(3)
	ZReqSta string `json:"zReqSta,omitempty"`
	// 处理结果 String(1)
	ZRtnFlg string `json:"zRtnFlg,omitempty"`
	// 错误描述 String(300)
	ZErrTxt string `json:"zErrTxt,omitempty"`
}

// UnitManagerNewBHHDZDTRZ2 详细信息 BHHDZDTRZ2.
type UnitManagerNewBHHDZDTRZ2 struct {
	// 户口号 String(35)
	ZEacNbr string `json:"zEacNbr,omitempty"`
	// 原交易流水号 String(15)
	ZOrgTrs string `json:"zOrgTrs,omitempty"`
	// 原交易日期 String(8)
	ZOrgDat string `json:"zOrgDat,omitempty"`
	// 原记账子单元 String(20)
	ZOrgDum string `json:"zOrgDum,omitempty"`
	// 目标记账子单元 String(20)
	ZTgtDum string `json:"zTgtDum,omitempty"`
	// 金额 BigDecimal(15,2)
	ZTrsAmt string `json:"zTrsAmt,omitempty"`
	// 交易摘要 String(60)
	ZTrsTxt string `json:"zTrsTxt,omitempty"`
	// 摘要是否自动拼原交易流水号 String(1) Y：自动拼接，N：非自动拼接
	ZCctTxt string `json:"zCctTxt,omitempty"`
}

// UnitManagerNewBHHDZUBSZ1 详细信息 BHHDZUBSZ1.
type UnitManagerNewBHHDZUBSZ1 struct {
	// 账号 String(35)
	ZAccNbr string `json:"zAccNbr,omitempty"`
	// 设置标识业务代码 String(6)
	ZBusCod string `json:"zBusCod,omitempty"`
	// 虚拟户标识 String(1) V：设置 N：取消
	ZVtlFlg string `json:"zVtlFlg,omitempty"`
}

// UnitManagerNewBHHTRNRTNZ 详细信息 BHHTRNRTNZ.
type UnitManagerNewBHHTRNRTNZ struct {
	// 子单元编号 String(20)
	ZDumNbr string `json:"zDumNbr,omitempty"`
	// 交易类型 String(1)
	ZTrsTyp string `json:"zTrsTyp,omitempty"`
	// 记账流水号 String(15)
	ZTrsNbr string `json:"zTrsNbr,omitempty"`
	// 交易日期 String(8)
	ZTrsDat string `json:"zTrsDat,omitempty"`
	// 交易时间 String(6)
	ZTrsTim string `json:"zTrsTim,omitempty"`
	// 交易金额 BigDecimal(17,2)
	ZTrsAmt string `json:"zTrsAmt,omitempty"`
	// 交易摘要 String(200)
	ZTrsTxt string `json:"zTrsTxt,omitempty"`
	// 收付方账号 String(35)
	ZRpyAcc string `json:"zRpyAcc,omitempty"`
	// 收付方名称 String(400)
	ZRpyNam string `json:"zRpyNam,omitempty"`
	// 收付方银行名称 String(200)
	ZRpyBnk string `json:"zRpyBnk,omitempty"`
	// 收付方银行地址 String(200)
	ZRpyAdr string `json:"zRpyAdr,omitempty"`
	// 收付方联行号 String(12)
	ZRpyBbn string `json:"zRpyBbn,omitempty"`
}

// UnitManagerNewBHHUSRRTNZ 详细信息 BHHUSRRTNZ.
type UnitManagerNewBHHUSRRTNZ struct {
	// 用户号 String(10)
	ZUsrNbr string `json:"zUsrNbr,omitempty"`
	// 登录名 String(90)
	ZLgnNam string `json:"zLgnNam,omitempty"`
	// 用户姓名 String(120)
	ZUsrNam string `json:"zUsrNam,omitempty"`
}

// UnitManagerNewNtbusmody 输入接口 ntbusmody.
type UnitManagerNewNtbusmody struct {
	// 业务模式 String(5) required:Y
	Busmod string `json:"busmod,omitempty"`
}

// UnitManagerNewNtdcdaddx1 输入接口 ntdcdaddx1.
type UnitManagerNewNtdcdaddx1 struct {
	// 实体结算账户 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
	// 业务参考号 String(30) required:Y 客户保证唯一
	Yurref string `json:"yurref,omitempty"`
}

// UnitManagerNewNtdcdaddx2 输入接口 ntdcdaddx2.
type UnitManagerNewNtdcdaddx2 struct {
	// 公司卡号 String(35) required:Y 见下方
	Cdsnbr string `json:"cdsnbr,omitempty"`
}

// UnitManagerNewNtdcdaddz1 详细信息 ntdcdaddz1.
type UnitManagerNewNtdcdaddz1 struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 虚拟户编号 String(20) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
	// 公司卡号 String(35) required:Y
	Cdsnbr string `json:"cdsnbr,omitempty"`
	// 操作类型 String(3) CDS：虚拟户关联公司卡 CDU：解除公司卡关联
	Oprtyp string `json:"oprtyp,omitempty"`
	// 处理结果 String(1) required:Y S：成功，F：失败，W：银行处理中
	Rtnflg string `json:"rtnflg,omitempty"`
	// 错误信息 Z(192)
	Errtxt string `json:"errtxt,omitempty"`
}

// UnitManagerNewNtdcddltx1 输入接口 ntdcddltx1.
type UnitManagerNewNtdcddltx1 struct {
	// 实体结算账户 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
	// 业务参考号 String(30) required:Y 客户保证唯一
	Yurref string `json:"yurref,omitempty"`
}

// UnitManagerNewNtdcddltx2 输入接口 ntdcddltx2.
type UnitManagerNewNtdcddltx2 struct {
	// 公司卡号 String(35) required:Y
	Cdsnbr string `json:"cdsnbr,omitempty"`
}

// UnitManagerNewNtdcdrltx1 查询条件 ntdcdrltx1.
type UnitManagerNewNtdcdrltx1 struct {
	// 实体结算账户 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20) 为空查全部，传子单元则指定查询该子单元
	Dmanbr string `json:"dmanbr,omitempty"`
	// 续传字段 String(50) 第一次传空，当返回接口有本接口并且该接口的续传字段有值，说明需要续传，取出续传字段放到下次请求报文中继续查询
	Rsv50z string `json:"rsv50z,omitempty"`
}

// UnitManagerNewNtdcdrltz1 详细信息 ntdcdrltz1.
type UnitManagerNewNtdcdrltz1 struct {
	// 记账子单元编号 String(20)
	Dmanbr string `json:"dmanbr,omitempty"`
	// 公司卡号 String(35)
	Cdsnbr string `json:"cdsnbr,omitempty"`
	// 凭证种类 String(40)
	Psbcod string `json:"psbcod,omitempty"`
	// 公司卡户口名 Z(62)
	Cdsnam string `json:"cdsnam,omitempty"`
	// 持卡人姓名 Z(200)
	Cltnam string `json:"cltnam,omitempty"`
	// 备注名称 Z(62)
	Txtnam string `json:"txtnam,omitempty"`
	// 开户日期 D
	Opndat string `json:"opndat,omitempty"`
	// 开户城市 String(6)
	Ctycod string `json:"ctycod,omitempty"`
	// 开户机构 String(6)
	Brnnbr string `json:"brnnbr,omitempty"`
	// 管理机构 String(6)
	Mgmbrn string `json:"mgmbrn,omitempty"`
	// 激活标志 String(1)
	Actflg string `json:"actflg,omitempty"`
	// 新旧户口标志 String(1) Y：新户口
	Dgrflg string `json:"dgrflg,omitempty"`
	// 公司卡户口状态 String(1)
	Stscod string `json:"stscod,omitempty"`
	// 公司卡关联付方账号校验标志 String(1) N：无关联（公司卡进账交易不校验付方账号） Y：有关联（公司卡进账交易要校验付方账号）
	Pdcflg string `json:"pdcflg,omitempty"`
	// 入账方式 String(1) 空：公司卡无关联付方账号 N：公司卡进账交易非匹配付方账号入默认记账子单元编号 Y：公司卡进账交易非匹配付方账号入关联记账子单元编号
	Tlyopr string `json:"tlyopr,omitempty"`
	// 关联状态 String(1) A：生效 W：银行处理中
	Rltsts string `json:"rltsts,omitempty"`
}

// UnitManagerNewNtdmaaddx 输入接口 ntdmaaddx.
type UnitManagerNewNtdmaaddx struct {
	// 账号 String（35） required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String（20） 自定义记账子单元编号：非空 系统自动生成记账子单元编号：空
	Dmanbr string `json:"dmanbr,omitempty"`
	// 记账子单元名称 Z（82） required:Y
	Dmanam string `json:"dmanam,omitempty"`
	// 是否可透支 String（1） Y：可透支 N：不可透支 空：默认Y X：不适用（收款识别开子单元此字段必须为X）
	Ovrctl string `json:"ovrctl,omitempty"`
	// 支付失败退回方式 String（1） Y ： 退款至该记账子单元 N：退款至结算账户（默认记账子单元） 空：默认N X：不适用（收款识别开子单元此字段必须为X）
	Bcktyp string `json:"bcktyp,omitempty"`
	// 关闭条件 String（1） Y：不限制， N：余额为零 空：默认Y X：不适用（收款识别开子单元此字段必须为X）
	Clstyp string `json:"clstyp,omitempty"`
	// 业务参考号 String（30） required:Y 客户保证唯一
	Yurref string `json:"yurref,omitempty"`
	// 是否收款限额 String（1） Y：是，N：否 默认N X：不适用（收款识别开子单元此字段必须为X）
	Lmtflg string `json:"lmtflg,omitempty"`
	// 余额上限 M（15） LMTFLG=Y时，BALLMT要大于0
	Ballmt string `json:"ballmt,omitempty"`
}

// UnitManagerNewNtdmabadz1 详细信息 ntdmabadz1.
type UnitManagerNewNtdmabadz1 struct {
	// 记账子单元编号 String(20)
	Dmanbr string `json:"dmanbr,omitempty"`
	// 业务处理结果 String(1) S：成功，F：失败，I/W：银行处理中
	Rtnsts string `json:"rtnsts,omitempty"`
	// 错误码 String(7)
	Errcod string `json:"errcod,omitempty"`
	// 错误信息 Z(192)
	Errtxt string `json:"errtxt,omitempty"`
}

// UnitManagerNewNtdmabalx 查询条件 ntdmabalx.
type UnitManagerNewNtdmabalx struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20) 填空表示查询该账号下所有记账子单元信息
	Dmanbr string `json:"dmanbr,omitempty"`
	// 续传字段 String(50) 第一次传空，当返回接口有本接口并且该接口的续传字段有值，说明需要续传，取出续传字段放到下次请求报文中继续查询
	Rsv50z string `json:"rsv50z,omitempty"`
}

// UnitManagerNewNtdmabalz 详细信息 ntdmabalz.
type UnitManagerNewNtdmabalz struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
	// 记账子单元名称 Z(82) required:Y
	Dmanam string `json:"dmanam,omitempty"`
	// 记账子单元币种 String(2) required:Y 见云直联文档中心附录A.3货币代码表
	Dmaccy string `json:"dmaccy,omitempty"`
	// 生效日期 D required:Y
	Eftdat string `json:"eftdat,omitempty"`
	// 余额 M required:Y
	Actbal string `json:"actbal,omitempty"`
	// 状态 String(1) required:Y A：生效 W：银行处理中
	Stscod string `json:"stscod,omitempty"`
	// 当月积数 F(17,2) 余额*天数 累加
	Mthaub string `json:"mthaub,omitempty"`
	// 余额是否有效 String(1) Y：余额有效
	Balflg string `json:"balflg,omitempty"`
}

// UnitManagerNewNtdmadltx1 输入接口 ntdmadltx1.
type UnitManagerNewNtdmadltx1 struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 业务参考号 String(30) required:Y 客户保证唯一
	Yurref string `json:"yurref,omitempty"`
}

// UnitManagerNewNtdmadltx2 输入接口 ntdmadltx2.
type UnitManagerNewNtdmadltx2 struct {
	// 记账子单元编号 String(20) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
}

// UnitManagerNewNtdmadltz 详细信息 ntdmadltz.
type UnitManagerNewNtdmadltz struct {
	// 记账子单元名称 Z(82) required:Y
	Dmanam string `json:"dmanam,omitempty"`
	// 记账子单元编号 String(20) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
	// 处理状态 String(1) required:Y S：成功 F：失败 I/W：银行处理中
	Prcsts string `json:"prcsts,omitempty"`
	// 返回码 String(7) SUC0000：成功 其他：失败
	Rtncod string `json:"rtncod,omitempty"`
	// 返回信息 Z(92) 返回码不是SUC0000时，会返回相关提示信息
	Rtntxt string `json:"rtntxt,omitempty"`
}

// UnitManagerNewNtdmahadx1 查询条件 ntdmahadx1.
type UnitManagerNewNtdmahadx1 struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
	// 开始日期 D required:Y 支持查询本年和去年的历史数据 一次查询间隔不能超过3个月
	Begdat string `json:"begdat,omitempty"`
	// 结束日期 D required:Y
	Enddat string `json:"enddat,omitempty"`
}

// UnitManagerNewNtdmahadz1 详细信息 ntdmahadz1.
type UnitManagerNewNtdmahadz1 struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元 String(20) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
	// 记账子单元名称 Z(82) required:Y
	Dmanam string `json:"dmanam,omitempty"`
	// 币种 String(2) required:Y 见云直联文档中心附录A.3货币代码表
	Ccynbr string `json:"ccynbr,omitempty"`
	// 日期 D required:Y
	Psedat string `json:"psedat,omitempty"`
	// 余额 M required:Y
	Dmabal string `json:"dmabal,omitempty"`
}

// UnitManagerNewNtdmahbdx1 查询条件 ntdmahbdx1.
type UnitManagerNewNtdmahbdx1 struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 查询日期 D required:Y 支持查询本年和去年的历史数据
	Qrydat string `json:"qrydat,omitempty"`
	// 续传记账子单元 String(20) 首次查询空，续传查询传上次查询返回的NTDMAHBDZ2的续传记账子单元dmaNbr
	Dmanbr string `json:"dmanbr,omitempty"`
}

// UnitManagerNewNtdmahbdz1 详细信息 ntdmahbdz1.
type UnitManagerNewNtdmahbdz1 struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元 String(20) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
	// 记账子单元名称 Z(82) required:Y
	Dmanam string `json:"dmanam,omitempty"`
	// 币种 String(2) required:Y 见云直联文档中心附录A.3货币代码表
	Ccynbr string `json:"ccynbr,omitempty"`
	// 日期 D required:Y
	Psedat string `json:"psedat,omitempty"`
	// 余额 M required:Y
	Dmabal string `json:"dmabal,omitempty"`
}

// UnitManagerNewNtdmahbdz2 续传接口 ntdmahbdz2.
type UnitManagerNewNtdmahbdz2 struct {
	// 续传标志 String(1) required:Y Y：有续传，N：无续传
	Ctnflg string `json:"ctnflg,omitempty"`
	// 续传记账子单元 String(20) 有续传时非空，无续传空
	Dmanbr string `json:"dmanbr,omitempty"`
}

// UnitManagerNewNtdmaintx1 查询条件 ntdmaintx1.
type UnitManagerNewNtdmaintx1 struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
	// 周期类型 String(1) required:Y Q:季度 M:月份 D：天
	Prdtyp string `json:"prdtyp,omitempty"`
	// 年份 String(4) required:Y
	Prdyer string `json:"prdyer,omitempty"`
	// 第几期 N(2) 可空，周期类型为Q/M指定具体的季度或月份取值要与PRDTYP对应： Q:01-04 M:01-12
	Prdnbr string `json:"prdnbr,omitempty"`
	// 计息利率 F(11,7) required:Y 百分比，例如：3.15表示3.15%
	Intrat string `json:"intrat,omitempty"`
	// 开始日期 D 格式：月日 可空，周期类型为D：必输 比如0930表示9月30号，与年份拼起来表示开始日期为2019年9月30号
	Begdat string `json:"begdat,omitempty"`
	// 结束日期 D 格式：月日 周期类型为D：必输 比如1030表示10月30号，与年份拼起来表示结束日期为2019年10月30号
	Enddat string `json:"enddat,omitempty"`
}

// UnitManagerNewNtdmaintz1 详细信息 ntdmaintz1.
type UnitManagerNewNtdmaintz1 struct {
	// 计息积数 F(17,2) required:Y
	Intaub string `json:"intaub,omitempty"`
	// 试算利息 M required:Y
	Intamt string `json:"intamt,omitempty"`
	// 利息方向 String(1) required:Y D:借方 C:贷方
	Trxdir string `json:"trxdir,omitempty"`
}

// UnitManagerNewNtdmalstx 查询条件 ntdmalstx.
type UnitManagerNewNtdmalstx struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20) 填空表示查询该账号下所有记账子单元信息
	Dmanbr string `json:"dmanbr,omitempty"`
	// 续传字段 String(50) 第一次传空，当返回接口有本接口并且该接口的续传字段有值，说明需要续传，取出续传字段放到下次请求报文中继续查询
	Rsv50z string `json:"rsv50z,omitempty"`
}

// UnitManagerNewNtdmalstz 详细信息 ntdmalstz.
type UnitManagerNewNtdmalstz struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
	// 记账子单元名称 Z(82) required:Y
	Dmanam string `json:"dmanam,omitempty"`
	// 记账子单元币种 String(2) required:Y 见云直联文档中心附录A.3货币代码表
	Dmaccy string `json:"dmaccy,omitempty"`
	// 生效日期 D required:Y
	Eftdat string `json:"eftdat,omitempty"`
	// 状态 String(1) required:Y A：生效 W：银行处理中
	Stscod string `json:"stscod,omitempty"`
	// 是否可透支 String(1) required:Y Y：可透支N：不可透支 X：不适用
	Ovrctl string `json:"ovrctl,omitempty"`
	// 付方校验标志 String(1) required:Y Y：有校验 N：不校验
	Rltchk string `json:"rltchk,omitempty"`
	// 入账方式 String(1) Y：非关联收款入账付方指定子单元 N：非关联收款入账默认子单元 R：非关联收款拒绝入账 （空表示未设置子单元入账条件）
	Tlyopr string `json:"tlyopr,omitempty"`
	// 支付失败退回方式 String(1) required:Y Y：退款至该记账子单元 N：退款至结算账户（默认记账子单元） X：不适用
	Bcktyp string `json:"bcktyp,omitempty"`
	// 关闭条件 String(1) required:Y Y：不限制 N：余额为零 X：不适用
	Clstyp string `json:"clstyp,omitempty"`
	// 停止收款标记 String(1) Y：是 N： 否
	Rcvlck string `json:"rcvlck,omitempty"`
	// 停止出款标记 String(1) Y：是 N： 否
	Paclck string `json:"paclck,omitempty"`
}

// UnitManagerNewNtdmalstz2 详细信息 ntdmalstz2.
type UnitManagerNewNtdmalstz2 struct {
	// 记账子单元编号 String(20) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
	// 是否收款限额 String(1) Y：是 N：否 X：不适用
	Lmtflg string `json:"lmtflg,omitempty"`
	// 余额上限 Ｍ（1 5）
	Ballmt string `json:"ballmt,omitempty"`
	// 收款剩余额度 Ｍ（1 5） 余额上限-联机余额
	Avilmt string `json:"avilmt,omitempty"`
}

// UnitManagerNewNtdmamntx1 输入接口 ntdmamntx1.
type UnitManagerNewNtdmamntx1 struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
	// 记账子单元名称 Z(82) required:Y
	Dmanam string `json:"dmanam,omitempty"`
	// 是否可透支 String(1) required:Y Y：可透支 N：不可透支 X：不适用（收款识别开子单元此字段必须为X）
	Ovrctl string `json:"ovrctl,omitempty"`
	// 支付失败退回方式 String(1) required:Y Y：退款至该记账子单元 N：退款至结算账户（默认子单元） X：不适用（收款识别开子单元此字段必须为X）
	Bcktyp string `json:"bcktyp,omitempty"`
	// 关闭条件 String(1) required:Y Y：不限制 N：余额为零 X：不适用（收款识别开子单元此字段必须为X）
	Clstyp string `json:"clstyp,omitempty"`
	// 业务参考号 String（30） required:Y 客户保证唯一
	Yurref string `json:"yurref,omitempty"`
	// 是否收款限额 String（1） Y：是，N：否 X：不适用（收款识别开子单元此字段必须为X）
	Lmtflg string `json:"lmtflg,omitempty"`
	// 余额上限 M（15） LMTFLG=Y时，BALLMT要大于0
	Ballmt string `json:"ballmt,omitempty"`
}

// UnitManagerNewNtdmaqryz1 详细信息 ntdmaqryz1.
type UnitManagerNewNtdmaqryz1 struct {
	// 分行号 String(2) required:Y
	Bbknbr string `json:"bbknbr,omitempty"`
	// 活期结算账户 String(35) required:Y
	Inbacc string `json:"inbacc,omitempty"`
	// 流程实例号 String(10) required:Y
	Reqnbr string `json:"reqnbr,omitempty"`
	// 业务编码 String(6) required:Y
	Buscod string `json:"buscod,omitempty"`
	// 业务模式 String(5) required:Y
	Busmod string `json:"busmod,omitempty"`
	// 操作类型 String(4) required:Y CSTA(新增协议）、CSTB（修改协议）、CSTU（收款识别升级）、DAUM（设置管理用户）、DAUA（按子单元设置权限）、DAUU（按用户设置权限）、DAUL（按用户设置所有权限）、AUTB（设置结算户支付权限）、 DUMA（新增记账子单元）、DMAA（新增记账子单元）、DMAB（新增子单元（批量））、DMAM（修
	Oprtyp string `json:"oprtyp,omitempty"`
	// 金额 M
	Trsamt string `json:"trsamt,omitempty"`
	// 经办日期 D required:Y
	Oprdat string `json:"oprdat,omitempty"`
	// 业务参考号 String(30) required:Y
	Yurref string `json:"yurref,omitempty"`
	// 请求状态 String(3) required:Y BNK：银行处理中 FIN：完成
	Reqsta string `json:"reqsta,omitempty"`
	// 业务处理结果 String(1) S：成功 F：失败 B：退票
	Rtnflg string `json:"rtnflg,omitempty"`
}

// UnitManagerNewNtdmarldx1 输入接口 ntdmarldx1.
type UnitManagerNewNtdmarldx1 struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
	// 付款方账号 String(35) required:Y
	Dbtacc string `json:"dbtacc,omitempty"`
	// 业务参考号 String(30) required:Y 客户保证唯一
	Yurref string `json:"yurref,omitempty"`
}

// UnitManagerNewNtdmarldx2 输入接口 ntdmarldx2.
type UnitManagerNewNtdmarldx2 struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
	// 付款方户名 Z（200） required:Y
	Dbtnam string `json:"dbtnam,omitempty"`
	// 业务参考号 String(30) required:Y 客户保证唯一
	Yurref string `json:"yurref,omitempty"`
}

// UnitManagerNewNtdmarldx3 输入接口 ntdmarldx3.
type UnitManagerNewNtdmarldx3 struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
	// 业务参考号 String(30) required:Y 客户保证唯一
	Yurref string `json:"yurref,omitempty"`
}

// UnitManagerNewNtdmarlqy1 查询条件 ntdmarlqy1.
type UnitManagerNewNtdmarlqy1 struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20) 为空查全部，传子单元则指定查询该子单元
	Dmanbr string `json:"dmanbr,omitempty"`
	// 续传字段 Z(80) 第一次传空，当返回接口有本接口并且该接口的续传字段有值，说明需要续传，取出续传字段放到下次请求报文中继续查询
	Cntkey string `json:"cntkey,omitempty"`
}

// UnitManagerNewNtdmarlqz1 详细信息 ntdmarlqz1.
type UnitManagerNewNtdmarlqz1 struct {
	// 账户名称 Z(200) required:Y
	Accnam string `json:"accnam,omitempty"`
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
	// 记账子单元名称 Z(82) required:Y
	Dmanam string `json:"dmanam,omitempty"`
	// 关联账号 String(35) required:Y
	Rltacc string `json:"rltacc,omitempty"`
	// 不满足条件时的 入账方式 String(1) required:Y Y：非关联收款入账付方指定子单元 N：非关联收款入账默认子单元 R：非关联收款拒绝入账
	Tlyopr string `json:"tlyopr,omitempty"`
	// 币种 String(2)
	Ccynbr string `json:"ccynbr,omitempty"`
	// 关联状态 String(1) A：生效 W：银行处理中
	Stscod string `json:"stscod,omitempty"`
}

// UnitManagerNewNtdmarlqz2 详细信息 ntdmarlqz2.
type UnitManagerNewNtdmarlqz2 struct {
	// 账户名称 Z(200) required:Y
	Accnam string `json:"accnam,omitempty"`
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
	// 记账子单元名称 Z(82) required:Y
	Dmanam string `json:"dmanam,omitempty"`
	// 关联户名 Z(200) required:Y
	Rltnam string `json:"rltnam,omitempty"`
	// 不满足条件时的 入账方式 String(1) required:Y Y：非关联收款入账付方指定子单元 N：非关联收款入账默认子单元 R：非关联收款拒绝入账
	Tlyopr string `json:"tlyopr,omitempty"`
	// 币种 String(2)
	Ccynbr string `json:"ccynbr,omitempty"`
	// 关联状态 String(1) A：生效 W：银行处理中
	Stscod string `json:"stscod,omitempty"`
}

// UnitManagerNewNtdmarlqz3 详细信息 ntdmarlqz3.
type UnitManagerNewNtdmarlqz3 struct {
	// 账户名称 Z(200) required:Y
	Accnam string `json:"accnam,omitempty"`
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
	// 记账子单元名称 Z(82) required:Y
	Dmanam string `json:"dmanam,omitempty"`
	// 关联金额 M（15） required:Y
	Rltamt string `json:"rltamt,omitempty"`
	// 不满足条件时的 入账方式 String(1) required:Y Y：非关联收款入账付方指定子单元 N：非关联收款入账默认子单元 R：非关联收款拒绝入账
	Tlyopr string `json:"tlyopr,omitempty"`
	// 币种 String(2)
	Ccynbr string `json:"ccynbr,omitempty"`
	// 关联状态 String(1) A：生效 W：银行处理中
	Stscod string `json:"stscod,omitempty"`
}

// UnitManagerNewNtdmarltx1 输入接口 ntdmarltx1.
type UnitManagerNewNtdmarltx1 struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
	// 不满足条件时的入账方式 String(1) required:Y 存量客户：Y：非关联收款入账付方指定子单元 N：非关联收款入账默认子单元 R：非关联收款拒绝入账 新开协议客户：R：非关联收款拒绝入账
	Tlyopr string `json:"tlyopr,omitempty"`
	// 付款方账号 String(35) required:Y
	Dbtacc string `json:"dbtacc,omitempty"`
	// 业务参考号 String(30) required:Y 客户保证唯一
	Yurref string `json:"yurref,omitempty"`
}

// UnitManagerNewNtdmarltx2 输入接口 ntdmarltx2.
type UnitManagerNewNtdmarltx2 struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
	// 不满足条件时的入账方式 String(1) required:Y 存量客户：Y：非关联收款入账付方指定子单元 N：非关联收款入账默认子单元 R：非关联收款拒绝入账 新开协议客户：R：非关联收款拒绝入账
	Tlyopr string `json:"tlyopr,omitempty"`
	// 付款方户名 Z（200） required:Y
	Dbtnam string `json:"dbtnam,omitempty"`
	// 业务参考号 String(30) required:Y 客户保证唯一
	Yurref string `json:"yurref,omitempty"`
}

// UnitManagerNewNtdmarltx3 输入接口 ntdmarltx3.
type UnitManagerNewNtdmarltx3 struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
	// 不满足条件时的入账方式 String(1) required:Y 存量客户：Y：非关联收款入账付方指定子单元 N：非关联收款入账默认子单元 R：非关联收款拒绝入账 新开协议客户：R：非关联收款拒绝入账
	Tlyopr string `json:"tlyopr,omitempty"`
	// 付款方金额 M(15) required:Y
	Dbtamt string `json:"dbtamt,omitempty"`
	// 业务参考号 String(30) required:Y 客户保证唯一
	Yurref string `json:"yurref,omitempty"`
}

// UnitManagerNewNtdmatmnx1 输入接口 ntdmatmnx1.
type UnitManagerNewNtdmatmnx1 struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
	// 不满足条件时的 入账方式 String(1) required:Y 存量客户：Y：非关联收款入账付方指定子单元 N：非关联收款入账默认子单元 R：非关联收款拒绝入账 新开协议客户：R：非关联收款拒绝入账
	Tlyopr string `json:"tlyopr,omitempty"`
	// 业务参考号 String(30) required:Y 客户保证唯一
	Yurref string `json:"yurref,omitempty"`
}

// UnitManagerNewNtdmatrxx1 输入接口 ntdmatrxx1.
type UnitManagerNewNtdmatrxx1 struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 付方记账子单元编号 String(20) required:Y
	Dmadbt string `json:"dmadbt,omitempty"`
	// 收方记账子单元编号 String(20) required:Y
	Dmacrt string `json:"dmacrt,omitempty"`
	// 交易金额 M required:Y
	Trxamt string `json:"trxamt,omitempty"`
	// 交易摘要 Z(42)
	Trxtxt string `json:"trxtxt,omitempty"`
	// 业务参考号 String(30) required:Y 客户保证唯一
	Yurref string `json:"yurref,omitempty"`
	// 是否需要审批 String(1) Y：需要审批，此时业务模式必需带审批岗 N/空：不需要审批，此时业务模式必须为0审批
	Apvflg string `json:"apvflg,omitempty"`
}

// UnitManagerNewNtdmatrxz1 详细信息 ntdmatrxz1.
type UnitManagerNewNtdmatrxz1 struct {
	// 记账日期 D required:Y
	Trxdat string `json:"trxdat,omitempty"`
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
	// 记账流水号 String(15) required:Y
	Trxnbr string `json:"trxnbr,omitempty"`
	// 交易金额 M required:Y
	Trxamt string `json:"trxamt,omitempty"`
	// 交易方向 String(1) required:Y D:付方 C:收方
	Trxdir string `json:"trxdir,omitempty"`
}

// UnitManagerNewNtdmawldx1 输入接口 ntdmawldx1.
type UnitManagerNewNtdmawldx1 struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
	// 付方账号 String(35)
	Sndacc string `json:"sndacc,omitempty"`
	// 分类条件类型 String(1) required:Y 1：帐号 2：户名 4、客户摘要
	Chktyp string `json:"chktyp,omitempty"`
	// 业务参考号 String(30) required:Y 客户保证唯一
	Yurref string `json:"yurref,omitempty"`
}

// UnitManagerNewNtdmawldx2 输入接口 ntdmawldx2.
type UnitManagerNewNtdmawldx2 struct {
	// 付方户名/客户摘要 Z(200)
	Sndnam string `json:"sndnam,omitempty"`
}

// UnitManagerNewNtdmawlsx1 输入接口 ntdmawlsx1.
type UnitManagerNewNtdmawlsx1 struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
	// 付方账号 String(35)
	Sndacc string `json:"sndacc,omitempty"`
	// 付方户名/客户摘要 Z(200)
	Sndnam string `json:"sndnam,omitempty"`
	// 分类条件类型 String(１) required:Y 1：帐号 2：户名 4：客户摘要
	Chktyp string `json:"chktyp,omitempty"`
	// 业务参考号 String(30) required:Y 客户保证唯一
	Yurref string `json:"yurref,omitempty"`
	// 不满足结算户入账分类条件的入账方式 String(１) required:Y N：非关联收款入账默认子单元 R：非关联收款拒绝入账
	Rcvtyp string `json:"rcvtyp,omitempty"`
}

// UnitManagerNewNtdmhdtxz1 详细信息 ntdmhdtxz1.
type UnitManagerNewNtdmhdtxz1 struct {
	// 分行号 String(2)
	Bbknbr string `json:"bbknbr,omitempty"`
	// 主账号 String(35)
	Accnbr string `json:"accnbr,omitempty"`
	// 虚拟户编号 String(20)
	Dumnbr string `json:"dumnbr,omitempty"`
	// 交易金额 BigDecimal(15,2)
	Trsamt string `json:"trsamt,omitempty"`
	// 交易套号 String(15)
	Setnbr string `json:"setnbr,omitempty"`
	// 交易流水号 String(15)
	Trxnbr string `json:"trxnbr,omitempty"`
	// 币种 String(2)
	Ccynbr string `json:"ccynbr,omitempty"`
	// 利率 BigDecimal(11,7)
	Intrat string `json:"intrat,omitempty"`
	// 利息 BigDecimal(15,2)
	Intamt string `json:"intamt,omitempty"`
	// 结息日期 String(8)
	Prcdat string `json:"prcdat,omitempty"`
	// 支取交易流水号 String(15)
	Dtxtrx string `json:"dtxtrx,omitempty"`
}

// UnitManagerNewNtdmidtxz1 详细信息 ntdmidtxz1.
type UnitManagerNewNtdmidtxz1 struct {
	// 分行号 String(2)
	Bbknbr string `json:"bbknbr,omitempty"`
	// 账号 String(35)
	Accnbr string `json:"accnbr,omitempty"`
	// 借方虚拟户编号 String(20)
	Dmadbt string `json:"dmadbt,omitempty"`
	// 贷方虚拟户编号 String(20)
	Dmacrt string `json:"dmacrt,omitempty"`
	// 交易金额 BigDecimal(15,2)
	Trsamt string `json:"trsamt,omitempty"`
	// 交易摘要 Z(42)
	Trstxt string `json:"trstxt,omitempty"`
	// 交易日期 String(8)
	Trsdat string `json:"trsdat,omitempty"`
	// 交易时间 String(6)
	Trstim string `json:"trstim,omitempty"`
	// 交易套号 String(15)
	Trxset string `json:"trxset,omitempty"`
	// 借方交易流水号 String(15)
	Trdnbr string `json:"trdnbr,omitempty"`
	// 贷方交易流水号 String(15)
	Trcnbr string `json:"trcnbr,omitempty"`
}

// UnitManagerNewNtdmiqryx1 查询条件 ntdmiqryx1.
type UnitManagerNewNtdmiqryx1 struct {
	// 活期结算账户 String(35) required:Y
	Inbacc string `json:"inbacc,omitempty"`
	// 记账子单元编号 String(20) required:Y 原收款交易
	Dyanbr string `json:"dyanbr,omitempty"`
	// 利率类型 String(3) required:Y FIX:对公活期利率 OTH:自定义利率
	Inttyp string `json:"inttyp,omitempty"`
	// 利率 F(11,7) 利率类型为OTH时必输，必须大于0小于等于1（例如：0.0315）
	Intval string `json:"intval,omitempty"`
	// 结束日 D required:Y 利息试算区间[原交易入账日期，结束日期），结束日期不含在区间内
	Enddat string `json:"enddat,omitempty"`
	// 原交易流水号 String(15) required:Y 原收款交易
	Orgtrx string `json:"orgtrx,omitempty"`
	// 原交易套号 String(15) required:Y 原收款交易
	Orgset string `json:"orgset,omitempty"`
}

// UnitManagerNewNtdmiqryz1 详细信息 ntdmiqryz1.
type UnitManagerNewNtdmiqryz1 struct {
	// 活期结算账户 String(35) required:Y
	Inbacc string `json:"inbacc,omitempty"`
	// 记账子单元编号 String(20) required:Y
	Dyanbr string `json:"dyanbr,omitempty"`
	// 币种 String(2) required:Y
	Ccynbr string `json:"ccynbr,omitempty"`
	// 利息 M required:Y
	Intamt string `json:"intamt,omitempty"`
	// 利率 F(11,7) required:Y
	Intval string `json:"intval,omitempty"`
	// 计息金额 F(17,2)
	Intbal string `json:"intbal,omitempty"`
}

// UnitManagerNewNtdmtqryy1 查询条件 ntdmtqryy1.
type UnitManagerNewNtdmtqryy1 struct {
	// 活期结算账户 String(35） required:Y
	Inbacc string `json:"inbacc,omitempty"`
	// 记账子单元编号 String(20） 为空时查记账子单元所有交易
	Dyanbr string `json:"dyanbr,omitempty"`
	// 开始日期 D required:Y 开始日期和结束日期范围不能超过100天
	Begdat string `json:"begdat,omitempty"`
	// 结束日期 D required:Y
	Enddat string `json:"enddat,omitempty"`
	// 续传字段 String(80) 第一次传空，当有该接口返回时并且该续传字段有值，说明需要续传，将 ntdmtqryy1接口内容放入下次请求报文继续查询。
	Ctnkey string `json:"ctnkey,omitempty"`
	// 分行号 String(2) required:Y 见云直联文档中心附录A.1招行分行
	Bbknbr string `json:"bbknbr,omitempty"`
}

// UnitManagerNewNtdmtqryzx 任务ID接口 ntdmtqryzx.
type UnitManagerNewNtdmtqryzx struct {
	// 任务ID String(35) required:Y
	Taskid string `json:"taskid,omitempty"`
}

// UnitManagerNewNtdumadxz1 详细信息 ntdumadxz1.
type UnitManagerNewNtdumadxz1 struct {
	// 分行号 String(2)
	Bbknbr string `json:"bbknbr,omitempty"`
	// 活期结算账户 String(35)
	Inbacc string `json:"inbacc,omitempty"`
	// 记账子单元编号 String(20)
	Dyanbr string `json:"dyanbr,omitempty"`
	// 记账子单元名称 Z(200)
	Dyanam string `json:"dyanam,omitempty"`
	// 生效日期 String(8)
	Eftdat string `json:"eftdat,omitempty"`
	// 终止日期 String(8)
	Enddat string `json:"enddat,omitempty"`
	// 是否允许透支 String(1)
	Ovrctl string `json:"ovrctl,omitempty"`
	// 额度标志 String(1)
	Lmtflg string `json:"lmtflg,omitempty"`
	// 余额上限额度 BigDecimal(15,2)
	Ballmt string `json:"ballmt,omitempty"`
}

// UnitManagerNewNtdumclex1 输入接口 ntdumclex1.
type UnitManagerNewNtdumclex1 struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20) required:Y
	Dumnbr string `json:"dumnbr,omitempty"`
	// 业务参考号 String(30) required:Y 客户保证唯一
	Yurref string `json:"yurref,omitempty"`
	// 收款限制 String(1) required:Y Y:解除，N:不解除
	Rcvflg string `json:"rcvflg,omitempty"`
	// 出款限制 String(1) required:Y Y:解除，N:不解除
	Pacflg string `json:"pacflg,omitempty"`
}

// UnitManagerNewNtdumctlx1 输入接口 ntdumctlx1.
type UnitManagerNewNtdumctlx1 struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20) required:Y
	Dumnbr string `json:"dumnbr,omitempty"`
	// 业务参考号 String(30) required:Y 客户保证唯一
	Yurref string `json:"yurref,omitempty"`
	// 收款限制 String(1) required:Y Y:设置，N:不设置
	Rcvflg string `json:"rcvflg,omitempty"`
	// 出款限制 String(1) required:Y Y:设置，N:不设置
	Pacflg string `json:"pacflg,omitempty"`
}

// UnitManagerNewNtdumctlz1 详细信息 ntdumctlz1.
type UnitManagerNewNtdumctlz1 struct {
	// 账号 String(35)
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20)
	Dumnbr string `json:"dumnbr,omitempty"`
	// 收款限制 String(1) Y：设置/取消 N：不设置/不取消 空默认Y
	Rcvflg string `json:"rcvflg,omitempty"`
	// 出款限制 String(1) Y：设置/取消 N：不设置/不取消 空默认N
	Pacflg string `json:"pacflg,omitempty"`
}

// UnitManagerNewNtdumdlxz1 详细信息 ntdumdlxz1.
type UnitManagerNewNtdumdlxz1 struct {
	// 分行号 String(2)
	Bbknbr string `json:"bbknbr,omitempty"`
	// 账号 String(35)
	Inbacc string `json:"inbacc,omitempty"`
	// 记账子单元编号 String(20)
	Dyanbr string `json:"dyanbr,omitempty"`
}

// UnitManagerNewNtduminfx1 查询条件 ntduminfx1.
type UnitManagerNewNtduminfx1 struct {
	// 流程实例号 String(10) required:Y
	REQNBR string `json:"REQNBR,omitempty"`
}

// UnitManagerNewNtduminfz1 详细信息 ntduminfz1.
type UnitManagerNewNtduminfz1 struct {
	// 流程实例号 String(10) required:Y
	Reqnbr string `json:"reqnbr,omitempty"`
	// 业务类型 String(6) required:Y
	Buscod string `json:"buscod,omitempty"`
	// 业务模式 String(5) required:Y
	Busmod string `json:"busmod,omitempty"`
	// 操作类型 String(4) required:Y CSTA(新增协议）、CSTB（修改协议）、CSTU（收款识别升级）、DAUM（设置管理用户）、DAUA（按子单元设置权限）、DAUU（按用户设置权限）、DAUL（按用户设置所有权限）、AUTB（设置结算户支付权限）、DUMA（新增记账子单元）、DMAA（新增记账子单元）、DMAB（新增子单元（批量））、DMAM（修改
	Oprtyp string `json:"oprtyp,omitempty"`
	// 业务请求状态 String(3) required:Y BNK：银行处理中 FIN：完成
	Reqsta string `json:"reqsta,omitempty"`
	// 待处理操作序列 String(3)
	Oprsqn string `json:"oprsqn,omitempty"`
	// 操作别名 Z(32)
	Oprals string `json:"oprals,omitempty"`
	// 用户名 Z(32) required:Y
	Lgnnam string `json:"lgnnam,omitempty"`
	// 用户姓名 Z(20) required:Y
	Usrnam string `json:"usrnam,omitempty"`
	// 业务请求结果 String(1) S：成功 F：失败 B：退票
	Rtnflg string `json:"rtnflg,omitempty"`
	// 结果描述 Z(192)
	Rtnnar string `json:"rtnnar,omitempty"`
	// 是否有附件信息 String(1)
	Athflg string `json:"athflg,omitempty"`
	// 业务参考号 String(30) required:Y
	Yurref string `json:"yurref,omitempty"`
}

// UnitManagerNewNTDUMINFZ3 详细信息 NTDUMINFZ3.
type UnitManagerNewNTDUMINFZ3 struct {
	// 交易日期 String(8)
	ZTrsDat string `json:"zTrsDat,omitempty"`
	// 交易时间 String(6)
	ZTrsTim string `json:"zTrsTim,omitempty"`
	// 交易流水号 String(15)
	ZTrsNbr string `json:"zTrsNbr,omitempty"`
	// 保留字100 String(100)
	ZRsv100 string `json:"zRsv100,omitempty"`
}

// UnitManagerNewNtdumraxz1 详细信息 ntdumraxz1.
type UnitManagerNewNtdumraxz1 struct {
	// 账号 String(35)
	Accnbr string `json:"accnbr,omitempty"`
	// 虚拟户编号 String(20)
	Dumnbr string `json:"dumnbr,omitempty"`
	// 关联户口 String(35)
	Rltacc string `json:"rltacc,omitempty"`
	// 关联户口名称 Z（200）
	Rltnam string `json:"rltnam,omitempty"`
	// 入账方式 String(1)
	Rcvtyp string `json:"rcvtyp,omitempty"`
	// 校验类型 String(1)
	Chktyp string `json:"chktyp,omitempty"`
}

// UnitManagerNewNtdumrdxz1 详细信息 ntdumrdxz1.
type UnitManagerNewNtdumrdxz1 struct {
	// 账号 String(35)
	Accnbr string `json:"accnbr,omitempty"`
	// 虚拟户编号 String(20)
	Dumnbr string `json:"dumnbr,omitempty"`
	// 关联户口 String(35)
	Rltacc string `json:"rltacc,omitempty"`
	// 校验类型 String(1)
	Chktyp string `json:"chktyp,omitempty"`
}

// UnitManagerNewNtdumredx1 查询条件 ntdumredx1.
type UnitManagerNewNtdumredx1 struct {
	// 业务参考号 String(30) required:Y
	Yurref string `json:"yurref,omitempty"`
	// 起始日期 D required:Y 经办日期查询范围不能超过7天，只能查3个月内的请求
	Bgndat string `json:"bgndat,omitempty"`
	// 结束日期 D required:Y
	Enddat string `json:"enddat,omitempty"`
}

// UnitManagerNewNthcdadxz1 详细信息 nthcdadxz1.
type UnitManagerNewNthcdadxz1 struct {
	// 分行号 String(2)
	Bbknbr string `json:"bbknbr,omitempty"`
	// 账号 String(35)
	Accnbr string `json:"accnbr,omitempty"`
	// 虚拟户编号 String(20)
	Dmanbr string `json:"dmanbr,omitempty"`
	// 公司卡号 String(35)
	Cdsnbr string `json:"cdsnbr,omitempty"`
}

// UnitManagerNewNthcddtxz1 详细信息 nthcddtxz1.
type UnitManagerNewNthcddtxz1 struct {
	// 分行号 String(2)
	Bbknbr string `json:"bbknbr,omitempty"`
	// 账号 String(35)
	Accnbr string `json:"accnbr,omitempty"`
	// 虚拟户编号 String(20)
	Dmanbr string `json:"dmanbr,omitempty"`
	// 公司卡号 String(35)
	Cdsnbr string `json:"cdsnbr,omitempty"`
}

// UnitManagerNewNthclactx1 输入接口 nthclactx1.
type UnitManagerNewNthclactx1 struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 业务参考号 String(30) required:Y 客户保证唯一
	Yurref string `json:"yurref,omitempty"`
}

// UnitManagerNewNthclactx2 输入接口 nthclactx2.
type UnitManagerNewNthclactx2 struct {
	// 原记账子单元编号 String(20) required:Y 需冲账的交易关联的记账子单元。
	Dmanbr string `json:"dmanbr,omitempty"`
	// 交易日期 D required:Y 输入需要冲账的交易明细对应的交易日期
	Trxdat string `json:"trxdat,omitempty"`
	// 记账流水号 String(15) required:Y 输入需要冲账的交易明细对应的记账流水号。 交易日期和流水号一起才能唯一匹配一笔交易。 不能指定“记账方式”为“内部调配”的交易记录来冲账。
	Trxnbr string `json:"trxnbr,omitempty"`
}

// UnitManagerNewNthclactz 详细信息 nthclactz.
type UnitManagerNewNthclactz struct {
	// 交易日期 D required:Y
	Trxdat string `json:"trxdat,omitempty"`
	// 记账流水号 String(15) required:Y
	Trxnbr string `json:"trxnbr,omitempty"`
	// 原记账子单元编号 String(10) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
	// 交易金额 M required:Y
	Trxamt string `json:"trxamt,omitempty"`
	// 交易方向 String(1) required:Y
	Trxdir string `json:"trxdir,omitempty"`
	// 处理状态 String(1) required:Y S：成功 F：失败 I/W：银行处理中
	Prcsts string `json:"prcsts,omitempty"`
	// 返回码 String(7) required:Y SUC0000：成功 其他：失败
	Rtncod string `json:"rtncod,omitempty"`
	// 返回信息 Z(92) 返回码不是 SUC0000时，会返回相关提示信息
	Rtntxt string `json:"rtntxt,omitempty"`
}

// UnitManagerNewNthkpactx1 输入接口 nthkpactx1.
type UnitManagerNewNthkpactx1 struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 关联记账子单元编号 String(20) required:Y 指定交易要关联到哪个记账子单元。 不能为默认记账子单元编号。
	Dmanbr string `json:"dmanbr,omitempty"`
	// 业务参考号 String(30) required:Y 客户保证唯一
	Yurref string `json:"yurref,omitempty"`
}

// UnitManagerNewNthkpactx2 输入接口 nthkpactx2.
type UnitManagerNewNthkpactx2 struct {
	// 交易日期 D required:Y 输入需要关联的交易明细对应的交易日期
	Trxdat string `json:"trxdat,omitempty"`
	// 记账流水号 String(15) required:Y 输入需要关联的交易明细对应的记账流水号。 交易日期和流水号一起才能唯一匹配一笔交易。 不能指定“记账方式”为“内部调配”的交易记录进行关联。
	Trxnbr string `json:"trxnbr,omitempty"`
}

// UnitManagerNewNthkpactz 详细信息 nthkpactz.
type UnitManagerNewNthkpactz struct {
	// 交易日期 D required:Y
	Trxdat string `json:"trxdat,omitempty"`
	// 记账流水号 String(15) required:Y
	Trxnbr string `json:"trxnbr,omitempty"`
	// 关联记账子单元编号 String(20) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
	// 交易金额 M required:Y
	Trxamt string `json:"trxamt,omitempty"`
	// 交易方向 String(1) required:Y
	Trxdir string `json:"trxdir,omitempty"`
	// 处理状态 String(1) required:Y S：成功 F：失败 I/W：银行处理中
	Prcsts string `json:"prcsts,omitempty"`
	// 返回码 String(7) required:Y SUC0000：成功 其他：失败
	Rtncod string `json:"rtncod,omitempty"`
	// 返回信息 Z(92) 返回码不是 SUC0000时，会返回相关提示信息
	Rtntxt string `json:"rtntxt,omitempty"`
}

// UnitManagerNewNtmthaubx 查询条件 ntmthaubx.
type UnitManagerNewNtmthaubx struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20) 填空表示查询该账号下所有记账子单元信息（不能为默认记账子单元编号）
	Dmanbr string `json:"dmanbr,omitempty"`
	// 起始年月 N(6) required:Y 本年或上一年，小于当月
	Begmth string `json:"begmth,omitempty"`
	// 结束年月 N(6) required:Y 本年或上一年，小于当月，不小于起始年月
	Endmth string `json:"endmth,omitempty"`
	// 续传字段 String(50) 第一次传空，当返回接口有本接口并且该接口的续传字段有值，说明需要续传，取出续传字段放到下次请求报文中继续查询
	Rsv50z string `json:"rsv50z,omitempty"`
}

// UnitManagerNewNtmthaubz 详细信息 ntmthaubz.
type UnitManagerNewNtmthaubz struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20) required:Y
	Dmanbr string `json:"dmanbr,omitempty"`
	// 记账子单元名称 Z(82) required:Y
	Dmanam string `json:"dmanam,omitempty"`
	// 记账子单元币种 String(2) required:Y 见云直联文档中心附录A.3货币代码表
	Dmaccy string `json:"dmaccy,omitempty"`
	// 年份 N(4) required:Y
	Aubyer string `json:"aubyer,omitempty"`
	// 月份 N(2) required:Y
	Aubmth string `json:"aubmth,omitempty"`
	// 月积数 F(17,2) required:Y
	Mthaub string `json:"mthaub,omitempty"`
}

// UnitManagerNewNtoprdmrx1 输入接口 ntoprdmrx1.
type UnitManagerNewNtoprdmrx1 struct {
	// 原交易套号 String(15)
	Setnbr string `json:"setnbr,omitempty"`
	// 原交易流水号 String(15) required:Y
	Trxnbr string `json:"trxnbr,omitempty"`
	// 退款本金 M required:Y 不退本金可退利息，本金可为0
	Trsamt string `json:"trsamt,omitempty"`
	// 主账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 记账子单元编号 String(20) required:Y
	Dumnbr string `json:"dumnbr,omitempty"`
	// 退款日期 D required:Y
	Eptdat string `json:"eptdat,omitempty"`
	// 原付方账号 String(35) required:Y
	Rpyacc string `json:"rpyacc,omitempty"`
	// 原付方名称 Z(200) required:Y
	Rpynam string `json:"rpynam,omitempty"`
	// 是否退息 String(1) required:Y Y:退息，N：不退息
	Intflg string `json:"intflg,omitempty"`
	// 利息 M required:Y 退息时必须大于0，不退息时必须为0
	Intamt string `json:"intamt,omitempty"`
	// 用途 Z(62) 本字段会记录在记账子单元交易的摘要中 为空默认为“资金原路返回交易”
	Nusage string `json:"nusage,omitempty"`
	// 客户摘要 Z(200)
	Busnar string `json:"busnar,omitempty"`
	// 业务参考号 String(30) required:Y 可通过“按业务参考号查询结果NTDUMRED”获取请求最终状态
	Yurref string `json:"yurref,omitempty"`
	// 部分退款标志 String(1) Y：部分退款，N：全额退款 为空则默认为N：全额退款 （产品升级后支持多次部分退款，该字段失效，系统自动根据退款本金判断）
	Bckflg string `json:"bckflg,omitempty"`
	// 附加标志 String(1) 可空，附加标志传Y，必须输入NTOPRDMRX2接口。如果原交易的付方地址为空，允许客户通过NTOPRDMRX2补录地址、名称、行号
	Apdflg string `json:"apdflg,omitempty"`
	// 是否需要审批 String(1) Y：需要审批，此时业务模式必需带审批岗 N/空：不需要审批，此时业务模式必须为0审批
	Apvflg string `json:"apvflg,omitempty"`
}

// UnitManagerNewNtoprdmrx2 输入接口 ntoprdmrx2.
type UnitManagerNewNtoprdmrx2 struct {
	// 收付方开户行地址 Z(100) required:Y
	Rpyadr string `json:"rpyadr,omitempty"`
	// 收付方开户行行名称 Z(62)
	Rpybkn string `json:"rpybkn,omitempty"`
	// 收付方开户行行号 String(20) 行内可不传，行外转账：联行号
	Rpybbn string `json:"rpybbn,omitempty"`
}

// UnitManagerNewNtoprdmrx3 输入接口 ntoprdmrx3.
type UnitManagerNewNtoprdmrx3 struct {
	// 是否需要利息回单 String(1) Y：是 N：否
	Intprt string `json:"intprt,omitempty"`
	// 退款回单是否附加备注 String(1) Y：是 N：否（如是，显示详细的本金和利息）
	Dmrprt string `json:"dmrprt,omitempty"`
}

// UnitManagerNewNtoprdpxz1 详细信息 ntoprdpxz1.
type UnitManagerNewNtoprdpxz1 struct {
	// 转出分行号 String(2) 结算账户开户行
	Dbtbbk string `json:"dbtbbk,omitempty"`
	// 转出帐号 String(35) 结算账户
	Dbtacc string `json:"dbtacc,omitempty"`
	// 记账子单元编号 String(20)
	Dumnbr string `json:"dumnbr,omitempty"`
	// 币种 String(2) 见云直联文档中心附录A.3货币代码表
	Ccynbr string `json:"ccynbr,omitempty"`
	// 金额 M 子单元原路退回时此字段指退款本金
	Trsamt string `json:"trsamt,omitempty"`
	// 系统内标志 String(1) Y：行内 N：行外
	Bnkflg string `json:"bnkflg,omitempty"`
	// 收方分行号 String(3) BNKFLG为Y时非空
	Oppbbk string `json:"oppbbk,omitempty"`
	// 收方开户行行号 String(20) 联行号
	Oppbbn string `json:"oppbbn,omitempty"`
	// 联行号类型 String(3)
	Opptyp string `json:"opptyp,omitempty"`
	// 收方帐号 String(35)
	Crtacc string `json:"crtacc,omitempty"`
	// 收方户名 Z(200)
	Crtnam string `json:"crtnam,omitempty"`
	// 收方行名称 Z(62)
	Crtbnk string `json:"crtbnk,omitempty"`
	// 收方行地址 Z100)
	Crtadr string `json:"crtadr,omitempty"`
	// 期望日 D
	Eptdat string `json:"eptdat,omitempty"`
	// 期望时间 String(6)
	Epttim string `json:"epttim,omitempty"`
	// 结算通道 String(1)
	Stlchn string `json:"stlchn,omitempty"`
	// 用途 Z(62)
	Nusage string `json:"nusage,omitempty"`
	// 业务摘要 Z(200)
	Busnar string `json:"busnar,omitempty"`
	// 业务种类 String(6) 100001=普通汇兑 101001=慈善捐款 101002 =其他 默认100001
	Trstyp string `json:"trstyp,omitempty"`
}

// UnitManagerNewNtoprdrxz1 详细信息 ntoprdrxz1.
type UnitManagerNewNtoprdrxz1 struct {
	// 原交易套号 String(15)
	Setnbr string `json:"setnbr,omitempty"`
	// 原交易流水号 String(15)
	Trxnbr string `json:"trxnbr,omitempty"`
	// 是否退息 String(1) Y：退息 N：不退息
	Intflg string `json:"intflg,omitempty"`
	// 利息 BigDecimal(15,2)
	Intamt string `json:"intamt,omitempty"`
	// 部分退款标志 String(1)
	Bckflg string `json:"bckflg,omitempty"`
}

// UnitManagerNewNtoprrtnz 详细信息 ntoprrtnz.
type UnitManagerNewNtoprrtnz struct {
	// 流水号 String(10)
	Sqrnbr string `json:"sqrnbr,omitempty"`
	// 流程实例号 String(10)
	Reqnbr string `json:"reqnbr,omitempty"`
	// 请求状态 String(3) required:Y BNK：银行处理中 FIN：完成
	Reqsts string `json:"reqsts,omitempty"`
	// 待处理操作序列 String(3)
	Oprsqn string `json:"oprsqn,omitempty"`
	// 操作别名 Z(32)
	Oprals string `json:"oprals,omitempty"`
	// 业务处理结果 String(1) S：成功 F：失败
	Rtnflg string `json:"rtnflg,omitempty"`
	// 错误码 String(7) required:Y
	Errcod string `json:"errcod,omitempty"`
	// 错误文本 Z(92)
	Errtxt string `json:"errtxt,omitempty"`
}

// UnitManagerNewNTDMAADDBody 1.1新增记账子单元NTDMAADD request body.
type UnitManagerNewNTDMAADDBody struct {
	Ntbusmody []*UnitManagerNewNtbusmody `json:"ntbusmody,omitempty"`
	Ntdmaaddx []*UnitManagerNewNtdmaaddx `json:"ntdmaaddx,omitempty"`
}

// UnitManagerNewNTDMAADDResponseBody 1.1新增记账子单元NTDMAADD response body.
type UnitManagerNewNTDMAADDResponseBody struct {
	Ntdmabadz1 []*UnitManagerNewNtdmabadz1 `json:"ntdmabadz1,omitempty"`
}

type UnitManagerNewNTDMAADDResponse struct {
	Response struct {
		Body UnitManagerNewNTDMAADDResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewNTDMADLTBody 1.2关闭记账子单元NTDMADLT request body.
type UnitManagerNewNTDMADLTBody struct {
	Ntbusmody  []*UnitManagerNewNtbusmody  `json:"ntbusmody,omitempty"`
	Ntdmadltx1 []*UnitManagerNewNtdmadltx1 `json:"ntdmadltx1,omitempty"`
	Ntdmadltx2 []*UnitManagerNewNtdmadltx2 `json:"ntdmadltx2,omitempty"`
}

// UnitManagerNewNTDMADLTResponseBody 1.2关闭记账子单元NTDMADLT response body.
type UnitManagerNewNTDMADLTResponseBody struct {
	Ntdmadltz []*UnitManagerNewNtdmadltz `json:"ntdmadltz,omitempty"`
	Ntoprrtnz []*UnitManagerNewNtoprrtnz `json:"ntoprrtnz,omitempty"`
}

type UnitManagerNewNTDMADLTResponse struct {
	Response struct {
		Body UnitManagerNewNTDMADLTResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewNTDMAMNTBody 1.3修改记账子单元NTDMAMNT request body.
type UnitManagerNewNTDMAMNTBody struct {
	Ntbusmody  []*UnitManagerNewNtbusmody  `json:"ntbusmody,omitempty"`
	Ntdmamntx1 []*UnitManagerNewNtdmamntx1 `json:"ntdmamntx1,omitempty"`
}

// UnitManagerNewNTDMAMNTResponseBody 1.3修改记账子单元NTDMAMNT response body.
type UnitManagerNewNTDMAMNTResponseBody struct {
	Ntoprrtnz []*UnitManagerNewNtoprrtnz `json:"ntoprrtnz,omitempty"`
}

type UnitManagerNewNTDMAMNTResponse struct {
	Response struct {
		Body UnitManagerNewNTDMAMNTResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewNTDMALSTBody 1.4查询记账子单元信息NTDMALST request body.
type UnitManagerNewNTDMALSTBody struct {
	Ntdmalstx []*UnitManagerNewNtdmalstx `json:"ntdmalstx,omitempty"`
}

// UnitManagerNewNTDMALSTResponseBody 1.4查询记账子单元信息NTDMALST response body.
type UnitManagerNewNTDMALSTResponseBody struct {
	Ntdmalstz  []*UnitManagerNewNtdmalstz  `json:"ntdmalstz,omitempty"`
	Ntdmalstz2 []*UnitManagerNewNtdmalstz2 `json:"ntdmalstz2,omitempty"`
	Ntdmalstx  []*UnitManagerNewNtdmalstx  `json:"ntdmalstx,omitempty"`
}

type UnitManagerNewNTDMALSTResponse struct {
	Response struct {
		Body UnitManagerNewNTDMALSTResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewNTOPRDMRBody 2.1 记账子单元原路退回NTOPRDMR request body.
type UnitManagerNewNTOPRDMRBody struct {
	Ntbusmody  []*UnitManagerNewNtbusmody  `json:"ntbusmody,omitempty"`
	Ntoprdmrx1 []*UnitManagerNewNtoprdmrx1 `json:"ntoprdmrx1,omitempty"`
	Ntoprdmrx2 []*UnitManagerNewNtoprdmrx2 `json:"ntoprdmrx2,omitempty"`
	Ntoprdmrx3 []*UnitManagerNewNtoprdmrx3 `json:"ntoprdmrx3,omitempty"`
}

// UnitManagerNewNTOPRDMRResponseBody 2.1 记账子单元原路退回NTOPRDMR response body.
type UnitManagerNewNTOPRDMRResponseBody struct {
	Ntoprrtnz []*UnitManagerNewNtoprrtnz `json:"ntoprrtnz,omitempty"`
}

type UnitManagerNewNTOPRDMRResponse struct {
	Response struct {
		Body UnitManagerNewNTOPRDMRResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewNTDMATRXBody 2.2 记账子单元内部调配NTDMATRX request body.
type UnitManagerNewNTDMATRXBody struct {
	Ntbusmody  []*UnitManagerNewNtbusmody  `json:"ntbusmody,omitempty"`
	Ntdmatrxx1 []*UnitManagerNewNtdmatrxx1 `json:"ntdmatrxx1,omitempty"`
}

// UnitManagerNewNTDMATRXResponseBody 2.2 记账子单元内部调配NTDMATRX response body.
type UnitManagerNewNTDMATRXResponseBody struct {
	Ntdmatrxz1 []*UnitManagerNewNtdmatrxz1 `json:"ntdmatrxz1,omitempty"`
	Ntoprrtnz  []*UnitManagerNewNtoprrtnz  `json:"ntoprrtnz,omitempty"`
}

type UnitManagerNewNTDMATRXResponse struct {
	Response struct {
		Body UnitManagerNewNTDMATRXResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewNTHKPACTBody 2.3手工记账NTHKPACT request body.
type UnitManagerNewNTHKPACTBody struct {
	Ntbusmody  []*UnitManagerNewNtbusmody  `json:"ntbusmody,omitempty"`
	Nthkpactx1 []*UnitManagerNewNthkpactx1 `json:"nthkpactx1,omitempty"`
	Nthkpactx2 []*UnitManagerNewNthkpactx2 `json:"nthkpactx2,omitempty"`
}

// UnitManagerNewNTHKPACTResponseBody 2.3手工记账NTHKPACT response body.
type UnitManagerNewNTHKPACTResponseBody struct {
	Nthkpactz []*UnitManagerNewNthkpactz `json:"nthkpactz,omitempty"`
	Ntoprrtnz []*UnitManagerNewNtoprrtnz `json:"ntoprrtnz,omitempty"`
}

type UnitManagerNewNTHKPACTResponse struct {
	Response struct {
		Body UnitManagerNewNTHKPACTResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewNTHCLACTBody 2.4记账子单元冲账NTHCLACT request body.
type UnitManagerNewNTHCLACTBody struct {
	Ntbusmody  []*UnitManagerNewNtbusmody  `json:"ntbusmody,omitempty"`
	Nthclactx1 []*UnitManagerNewNthclactx1 `json:"nthclactx1,omitempty"`
	Nthclactx2 []*UnitManagerNewNthclactx2 `json:"nthclactx2,omitempty"`
}

// UnitManagerNewNTHCLACTResponseBody 2.4记账子单元冲账NTHCLACT response body.
type UnitManagerNewNTHCLACTResponseBody struct {
	Nthclactz []*UnitManagerNewNthclactz `json:"nthclactz,omitempty"`
	Ntoprrtnz []*UnitManagerNewNtoprrtnz `json:"ntoprrtnz,omitempty"`
}

type UnitManagerNewNTHCLACTResponse struct {
	Response struct {
		Body UnitManagerNewNTHCLACTResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHDZCTZBody 3.1交易查询BHHDZCTZ request body.
type UnitManagerNewBHHDZCTZBody struct {
	BHHDNCTZX1 []*UnitManagerNewBHHDNCTZX1 `json:"BHHDNCTZX1,omitempty"`
	BHHDNCTZX2 []*UnitManagerNewBHHDNCTZX2 `json:"BHHDNCTZX2,omitempty"`
}

// UnitManagerNewBHHDZCTZResponseBody 3.1交易查询BHHDZCTZ response body.
type UnitManagerNewBHHDZCTZResponseBody struct {
	BHHDNCTZZ1 []*UnitManagerNewBHHDNCTZZ1 `json:"BHHDNCTZZ1,omitempty"`
	BHHDNCTZZ2 []*UnitManagerNewBHHDNCTZZ2 `json:"BHHDNCTZZ2,omitempty"`
}

type UnitManagerNewBHHDZCTZResponse struct {
	Response struct {
		Body UnitManagerNewBHHDZCTZResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewNTDMTQRYPDFBody 3.2记账子单元交易明细对账单获取NTDMTQRYPDF request body.
type UnitManagerNewNTDMTQRYPDFBody struct {
	Ntdmtqryy1 []*UnitManagerNewNtdmtqryy1 `json:"ntdmtqryy1,omitempty"`
}

// UnitManagerNewNTDMTQRYPDFResponseBody 3.2记账子单元交易明细对账单获取NTDMTQRYPDF response body.
type UnitManagerNewNTDMTQRYPDFResponseBody struct {
	Ntdmtqryzx []*UnitManagerNewNtdmtqryzx `json:"ntdmtqryzx,omitempty"`
	Ntdmtqryy1 []*UnitManagerNewNtdmtqryy1 `json:"ntdmtqryy1,omitempty"`
}

type UnitManagerNewNTDMTQRYPDFResponse struct {
	Response struct {
		Body UnitManagerNewNTDMTQRYPDFResponseBody `json:"body,omitempty"`
		Head HeadResponse                          `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewDCTASKIDBody 3.3记账子单元交易明细对账单处理结果查询请求DCTASKID request body.
type UnitManagerNewDCTASKIDBody struct {
	// 查询ID string (100) required:Y 查询请求NTDMTQRYPDF时返回
	Taskid string `json:"taskid,omitempty"`
	// 文件地址类型 String true或者空：生成的url链接为互联网可访问 false：生成的url链接为专线访问地址，（请专线接入客户该值传false）
	Qwenab string `json:"qwenab,omitempty"`
}

// UnitManagerNewDCTASKIDResponseBody 3.3记账子单元交易明细对账单处理结果查询请求DCTASKID response body.
type UnitManagerNewDCTASKIDResponseBody struct {
	// 文件获取地址 string(300) 生成完成后，PDF文件获取地址链接，请注意该链接有效期1天，获取到后请立即下载保存，以防过期失效。
	Fileurl string `json:"fileurl,omitempty"`
	// 文件生成时间 string(20) 后台PDF文件生成完成时间
	Fintim string `json:"fintim,omitempty"`
}

type UnitManagerNewDCTASKIDResponse struct {
	Response struct {
		Body UnitManagerNewDCTASKIDResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHDZDBQBody 3.4退票交易查询BHHDZDBQ request body.
type UnitManagerNewBHHDZDBQBody struct {
	BHHDZDBQX1 []*UnitManagerNewBHHDZDBQX1 `json:"BHHDZDBQX1,omitempty"`
	BHHDZDBQY1 []*UnitManagerNewBHHDZDBQY1 `json:"BHHDZDBQY1,omitempty"`
}

// UnitManagerNewBHHDZDBQResponseBody 3.4退票交易查询BHHDZDBQ response body.
type UnitManagerNewBHHDZDBQResponseBody struct {
	BHHDZDBQZ1 []*UnitManagerNewBHHDZDBQZ1 `json:"BHHDZDBQZ1,omitempty"`
	BHHDZDBQY1 []*UnitManagerNewBHHDZDBQY1 `json:"BHHDZDBQY1,omitempty"`
}

type UnitManagerNewBHHDZDBQResponse struct {
	Response struct {
		Body UnitManagerNewBHHDZDBQResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewNTDMABALBody 4.1查询记账子单元余额NTDMABAL request body.
type UnitManagerNewNTDMABALBody struct {
	Ntdmabalx []*UnitManagerNewNtdmabalx `json:"ntdmabalx,omitempty"`
}

// UnitManagerNewNTDMABALResponseBody 4.1查询记账子单元余额NTDMABAL response body.
type UnitManagerNewNTDMABALResponseBody struct {
	Ntdmabalz []*UnitManagerNewNtdmabalz `json:"ntdmabalz,omitempty"`
	Ntdmabalx []*UnitManagerNewNtdmabalx `json:"ntdmabalx,omitempty"`
}

type UnitManagerNewNTDMABALResponse struct {
	Response struct {
		Body UnitManagerNewNTDMABALResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewNTDMAHBDBody 4.2查询所有记账子单元的某日余额NTDMAHBD request body.
type UnitManagerNewNTDMAHBDBody struct {
	Ntdmahbdx1 []*UnitManagerNewNtdmahbdx1 `json:"ntdmahbdx1,omitempty"`
}

// UnitManagerNewNTDMAHBDResponseBody 4.2查询所有记账子单元的某日余额NTDMAHBD response body.
type UnitManagerNewNTDMAHBDResponseBody struct {
	Ntdmahbdz1 []*UnitManagerNewNtdmahbdz1 `json:"ntdmahbdz1,omitempty"`
	Ntdmahbdz2 []*UnitManagerNewNtdmahbdz2 `json:"ntdmahbdz2,omitempty"`
}

type UnitManagerNewNTDMAHBDResponse struct {
	Response struct {
		Body UnitManagerNewNTDMAHBDResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewNTDMAHADBody 4.3查询单个记账子单元的历史余额NTDMAHAD request body.
type UnitManagerNewNTDMAHADBody struct {
	Ntdmahadx1 []*UnitManagerNewNtdmahadx1 `json:"ntdmahadx1,omitempty"`
}

// UnitManagerNewNTDMAHADResponseBody 4.3查询单个记账子单元的历史余额NTDMAHAD response body.
type UnitManagerNewNTDMAHADResponseBody struct {
	Ntdmahadz1 []*UnitManagerNewNtdmahadz1 `json:"ntdmahadz1,omitempty"`
}

type UnitManagerNewNTDMAHADResponse struct {
	Response struct {
		Body UnitManagerNewNTDMAHADResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewNTMTHAUBBody 4.4查询记账子单元历史月积数NTMTHAUB request body.
type UnitManagerNewNTMTHAUBBody struct {
	Ntmthaubx []*UnitManagerNewNtmthaubx `json:"ntmthaubx,omitempty"`
}

// UnitManagerNewNTMTHAUBResponseBody 4.4查询记账子单元历史月积数NTMTHAUB response body.
type UnitManagerNewNTMTHAUBResponseBody struct {
	Ntmthaubz []*UnitManagerNewNtmthaubz `json:"ntmthaubz,omitempty"`
	Ntmthaubx []*UnitManagerNewNtmthaubx `json:"ntmthaubx,omitempty"`
}

type UnitManagerNewNTMTHAUBResponse struct {
	Response struct {
		Body UnitManagerNewNTMTHAUBResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewNTDUMREDBody 5.1按业务参考号查询结果NTDUMRED request body.
type UnitManagerNewNTDUMREDBody struct {
	Ntdumredx1 []*UnitManagerNewNtdumredx1 `json:"ntdumredx1,omitempty"`
}

// UnitManagerNewNTDUMREDResponseBody 5.1按业务参考号查询结果NTDUMRED response body.
type UnitManagerNewNTDUMREDResponseBody struct {
	Ntdmaqryz1 []*UnitManagerNewNtdmaqryz1 `json:"ntdmaqryz1,omitempty"`
	Bhhdnrefz2 []*UnitManagerNewBhhdnrefz2 `json:"bhhdnrefz2,omitempty"`
}

type UnitManagerNewNTDUMREDResponse struct {
	Response struct {
		Body UnitManagerNewNTDUMREDResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewNTDUMINFBody 5.2详情查询NTDUMINF request body.
type UnitManagerNewNTDUMINFBody struct {
	Ntduminfx1 []*UnitManagerNewNtduminfx1 `json:"ntduminfx1,omitempty"`
}

// UnitManagerNewNTDUMINFResponseBody 5.2详情查询NTDUMINF response body.
type UnitManagerNewNTDUMINFResponseBody struct {
	Ntduminfz1 []*UnitManagerNewNtduminfz1 `json:"ntduminfz1,omitempty"`
	BHHDNCSAZ  []*UnitManagerNewBHHDNCSAZ  `json:"BHHDNCSAZ,omitempty"`
	BHHDNCSBZ  []*UnitManagerNewBHHDNCSBZ  `json:"BHHDNCSBZ,omitempty"`
	BHHDNCSUZ  []*UnitManagerNewBHHDNCSUZ  `json:"BHHDNCSUZ,omitempty"`
	BHHDAUMSZ  []*UnitManagerNewBHHDAUMSZ  `json:"BHHDAUMSZ,omitempty"`
	BHHUSRRTNZ []*UnitManagerNewBHHUSRRTNZ `json:"BHHUSRRTNZ,omitempty"`
	BHHDAUASZ  []*UnitManagerNewBHHDAUASZ  `json:"BHHDAUASZ,omitempty"`
	BHHDAUUSZ  []*UnitManagerNewBHHDAUUSZ  `json:"BHHDAUUSZ,omitempty"`
	BHHDAUULZ  []*UnitManagerNewBHHDAUULZ  `json:"BHHDAUULZ,omitempty"`
	BHHDZUBSZ1 []*UnitManagerNewBHHDZUBSZ1 `json:"BHHDZUBSZ1,omitempty"`
	Ntdumadxz1 []*UnitManagerNewNtdumadxz1 `json:"ntdumadxz1,omitempty"`
	BHHDAADDZ  []*UnitManagerNewBHHDAADDZ  `json:"BHHDAADDZ,omitempty"`
	BHHDNAABZ  []*UnitManagerNewBHHDNAABZ  `json:"BHHDNAABZ,omitempty"`
	BHHDNAAMZ  []*UnitManagerNewBHHDNAAMZ  `json:"BHHDNAAMZ,omitempty"`
	BHHDZABMZ2 []*UnitManagerNewBHHDZABMZ2 `json:"BHHDZABMZ2,omitempty"`
	BHHDADLTZ  []*UnitManagerNewBHHDADLTZ  `json:"BHHDADLTZ,omitempty"`
	Ntdumdlxz1 []*UnitManagerNewNtdumdlxz1 `json:"ntdumdlxz1,omitempty"`
	Ntdumctlz1 []*UnitManagerNewNtdumctlz1 `json:"ntdumctlz1,omitempty"`
	BHHDZCTBZ1 []*UnitManagerNewBHHDZCTBZ1 `json:"BHHDZCTBZ1,omitempty"`
	BHHDTAJCZ  []*UnitManagerNewBHHDTAJCZ  `json:"BHHDTAJCZ,omitempty"`
	BHHDTAJDZ  []*UnitManagerNewBHHDTAJDZ  `json:"BHHDTAJDZ,omitempty"`
	BHHTRNRTNZ []*UnitManagerNewBHHTRNRTNZ `json:"BHHTRNRTNZ,omitempty"`
	Ntdmidtxz1 []*UnitManagerNewNtdmidtxz1 `json:"ntdmidtxz1,omitempty"`
	BHHDTITXZ  []*UnitManagerNewBHHDTITXZ  `json:"BHHDTITXZ,omitempty"`
	BHHDZDTRZ2 []*UnitManagerNewBHHDZDTRZ2 `json:"BHHDZDTRZ2,omitempty"`
	BHHDZDTBZ2 []*UnitManagerNewBHHDZDTBZ2 `json:"BHHDZDTBZ2,omitempty"`
	Ntdumraxz1 []*UnitManagerNewNtdumraxz1 `json:"ntdumraxz1,omitempty"`
	Ntdumrdxz1 []*UnitManagerNewNtdumrdxz1 `json:"ntdumrdxz1,omitempty"`
	BHHDNRLBZ  []*UnitManagerNewBHHDNRLBZ  `json:"BHHDNRLBZ,omitempty"`
	BHHDYRDBZ1 []*UnitManagerNewBHHDYRDBZ1 `json:"BHHDYRDBZ1,omitempty"`
	BHHDWMNTZ  []*UnitManagerNewBHHDWMNTZ  `json:"BHHDWMNTZ,omitempty"`
	Nthcdadxz1 []*UnitManagerNewNthcdadxz1 `json:"nthcdadxz1,omitempty"`
	BHHCDSRTNZ []*UnitManagerNewBHHCDSRTNZ `json:"BHHCDSRTNZ,omitempty"`
	Nthcddtxz1 []*UnitManagerNewNthcddtxz1 `json:"nthcddtxz1,omitempty"`
	BHHDRADDZ  []*UnitManagerNewBHHDRADDZ  `json:"BHHDRADDZ,omitempty"`
	BHHDRDLTZ  []*UnitManagerNewBHHDRDLTZ  `json:"BHHDRDLTZ,omitempty"`
	BHHDYBADZ1 []*UnitManagerNewBHHDYBADZ1 `json:"BHHDYBADZ1,omitempty"`
	BHHDYDLBZ1 []*UnitManagerNewBHHDYDLBZ1 `json:"BHHDYDLBZ1,omitempty"`
	BHHDNDRAZ  []*UnitManagerNewBHHDNDRAZ  `json:"BHHDNDRAZ,omitempty"`
	BHHDNDRDZ  []*UnitManagerNewBHHDNDRDZ  `json:"BHHDNDRDZ,omitempty"`
	BHHDNDRNZ  []*UnitManagerNewBHHDNDRNZ  `json:"BHHDNDRNZ,omitempty"`
	BHHDNDRBZ  []*UnitManagerNewBHHDNDRBZ  `json:"BHHDNDRBZ,omitempty"`
	BHHDNDRMZ  []*UnitManagerNewBHHDNDRMZ  `json:"BHHDNDRMZ,omitempty"`
	BHHDZDRTZ1 []*UnitManagerNewBHHDZDRTZ1 `json:"BHHDZDRTZ1,omitempty"`
	BHHDNDSAZ  []*UnitManagerNewBHHDNDSAZ  `json:"BHHDNDSAZ,omitempty"`
	BHHDNDCBZ  []*UnitManagerNewBHHDNDCBZ  `json:"BHHDNDCBZ,omitempty"`
	BHHDNDSDZ  []*UnitManagerNewBHHDNDSDZ  `json:"BHHDNDSDZ,omitempty"`
	BHHDNDDBZ  []*UnitManagerNewBHHDNDDBZ  `json:"BHHDNDDBZ,omitempty"`
	BHHDNDSMZ  []*UnitManagerNewBHHDNDSMZ  `json:"BHHDNDSMZ,omitempty"`
	BHHDINRSZ  []*UnitManagerNewBHHDINRSZ  `json:"BHHDINRSZ,omitempty"`
	BHHDINRBZ  []*UnitManagerNewBHHDINRBZ  `json:"BHHDINRBZ,omitempty"`
	BHHDINASZ  []*UnitManagerNewBHHDINASZ  `json:"BHHDINASZ,omitempty"`
	BHHDINASZ1 []*UnitManagerNewBHHDINASZ1 `json:"BHHDINASZ1,omitempty"`
	BHHDINABZ  []*UnitManagerNewBHHDINABZ  `json:"BHHDINABZ,omitempty"`
	BHHDINABZ1 []*UnitManagerNewBHHDINABZ1 `json:"BHHDINABZ1,omitempty"`
	BHHDINCSZ  []*UnitManagerNewBHHDINCSZ  `json:"BHHDINCSZ,omitempty"`
	BHHDINCBZ  []*UnitManagerNewBHHDINCBZ  `json:"BHHDINCBZ,omitempty"`
	BHHDINRPZ1 []*UnitManagerNewBHHDINRPZ1 `json:"BHHDINRPZ1,omitempty"`
	BHHDINRTZ1 []*UnitManagerNewBHHDINRTZ1 `json:"BHHDINRTZ1,omitempty"`
	Ntoprdpxz1 []*UnitManagerNewNtoprdpxz1 `json:"ntoprdpxz1,omitempty"`
	Ntoprdrxz1 []*UnitManagerNewNtoprdrxz1 `json:"ntoprdrxz1,omitempty"`
	Ntdmhdtxz1 []*UnitManagerNewNtdmhdtxz1 `json:"ntdmhdtxz1,omitempty"`
	BHHDGPADZ1 []*UnitManagerNewBHHDGPADZ1 `json:"BHHDGPADZ1,omitempty"`
	BHHDGPMNZ1 []*UnitManagerNewBHHDGPMNZ1 `json:"BHHDGPMNZ1,omitempty"`
	BHHDGPSCZ1 []*UnitManagerNewBHHDGPSCZ1 `json:"BHHDGPSCZ1,omitempty"`
	BHHDGNADZ1 []*UnitManagerNewBHHDGNADZ1 `json:"BHHDGNADZ1,omitempty"`
	BHHDGNCDZ1 []*UnitManagerNewBHHDGNCDZ1 `json:"BHHDGNCDZ1,omitempty"`
	BHHDGNAGZ1 []*UnitManagerNewBHHDGNAGZ1 `json:"BHHDGNAGZ1,omitempty"`
	BHHDGNCGZ1 []*UnitManagerNewBHHDGNCGZ1 `json:"BHHDGNCGZ1,omitempty"`
	BHHDLMADZ1 []*UnitManagerNewBHHDLMADZ1 `json:"BHHDLMADZ1,omitempty"`
	BHHDLMMDZ1 []*UnitManagerNewBHHDLMMDZ1 `json:"BHHDLMMDZ1,omitempty"`
	BHHDLMCDZ1 []*UnitManagerNewBHHDLMCDZ1 `json:"BHHDLMCDZ1,omitempty"`
	BHHDLMTPZ1 []*UnitManagerNewBHHDLMTPZ1 `json:"BHHDLMTPZ1,omitempty"`
	NTDUMINFZ3 []*UnitManagerNewNTDUMINFZ3 `json:"NTDUMINFZ3,omitempty"`
	BHHDNWCAZ1 []*UnitManagerNewBHHDNWCAZ1 `json:"BHHDNWCAZ1,omitempty"`
	BHHDNWCCZ1 []*UnitManagerNewBHHDNWCCZ1 `json:"BHHDNWCCZ1,omitempty"`
}

type UnitManagerNewNTDUMINFResponse struct {
	Response struct {
		Body UnitManagerNewNTDUMINFResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewNTDMAWLSBody 6.1设置结算户入账分类条件NTDMAWLS request body.
type UnitManagerNewNTDMAWLSBody struct {
	Ntbusmody  []*UnitManagerNewNtbusmody  `json:"ntbusmody,omitempty"`
	Ntdmawlsx1 []*UnitManagerNewNtdmawlsx1 `json:"ntdmawlsx1,omitempty"`
}

// UnitManagerNewNTDMAWLSResponseBody 6.1设置结算户入账分类条件NTDMAWLS response body.
type UnitManagerNewNTDMAWLSResponseBody struct {
	Ntoprrtnz []*UnitManagerNewNtoprrtnz `json:"ntoprrtnz,omitempty"`
}

type UnitManagerNewNTDMAWLSResponse struct {
	Response struct {
		Body UnitManagerNewNTDMAWLSResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewNTDMAWLDBody 6.2删除结算户入账分类条件NTDMAWLD request body.
type UnitManagerNewNTDMAWLDBody struct {
	Ntbusmody  []*UnitManagerNewNtbusmody  `json:"ntbusmody,omitempty"`
	Ntdmawldx1 []*UnitManagerNewNtdmawldx1 `json:"ntdmawldx1,omitempty"`
	Ntdmawldx2 []*UnitManagerNewNtdmawldx2 `json:"ntdmawldx2,omitempty"`
}

// UnitManagerNewNTDMAWLDResponseBody 6.2删除结算户入账分类条件NTDMAWLD response body.
type UnitManagerNewNTDMAWLDResponseBody struct {
	Ntoprrtnz []*UnitManagerNewNtoprrtnz `json:"ntoprrtnz,omitempty"`
}

type UnitManagerNewNTDMAWLDResponse struct {
	Response struct {
		Body UnitManagerNewNTDMAWLDResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHDZWQCBody 6.3 查询结算户入账分类条件BHHDZWQC request body.
type UnitManagerNewBHHDZWQCBody struct {
	BHHDNWQCY1 []*UnitManagerNewBHHDNWQCY1 `json:"BHHDNWQCY1,omitempty"`
}

// UnitManagerNewBHHDZWQCResponseBody 6.3 查询结算户入账分类条件BHHDZWQC response body.
type UnitManagerNewBHHDZWQCResponseBody struct {
	BHHDNWQCZ1 []*UnitManagerNewBHHDNWQCZ1 `json:"BHHDNWQCZ1,omitempty"`
	BHHDNWQCY1 []*UnitManagerNewBHHDNWQCY1 `json:"BHHDNWQCY1,omitempty"`
}

type UnitManagerNewBHHDZWQCResponse struct {
	Response struct {
		Body UnitManagerNewBHHDZWQCResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHZWMNTBody 6.4修改不满足结算户入账分类条件的入账方式BHHZWMNT request body.
type UnitManagerNewBHHZWMNTBody struct {
	Bhducmody1 []*UnitManagerNewBhducmody1 `json:"bhducmody1,omitempty"`
	Bhhdwmntx1 []*UnitManagerNewBhhdwmntx1 `json:"bhhdwmntx1,omitempty"`
}

// UnitManagerNewBHHZWMNTResponseBody 6.4修改不满足结算户入账分类条件的入账方式BHHZWMNT response body.
type UnitManagerNewBHHZWMNTResponseBody struct {
	Bhhdurtnz1 []*UnitManagerNewBhhdurtnz1 `json:"bhhdurtnz1,omitempty"`
}

type UnitManagerNewBHHZWMNTResponse struct {
	Response struct {
		Body UnitManagerNewBHHZWMNTResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHDNWCABody 7.1新增结算户出账分类条件BHHDNWCA request body.
type UnitManagerNewBHHDNWCABody struct {
	BHDUCMODY1 []*UnitManagerNewBHDUCMODY1 `json:"BHDUCMODY1,omitempty"`
	BHHDNWCAX1 []*UnitManagerNewBHHDNWCAX1 `json:"BHHDNWCAX1,omitempty"`
}

// UnitManagerNewBHHDNWCAResponseBody 7.1新增结算户出账分类条件BHHDNWCA response body.
type UnitManagerNewBHHDNWCAResponseBody struct {
	BHHDURTNZ1 []*UnitManagerNewBHHDURTNZ1 `json:"BHHDURTNZ1,omitempty"`
}

type UnitManagerNewBHHDNWCAResponse struct {
	Response struct {
		Body UnitManagerNewBHHDNWCAResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHDNWCCBody 7.2删除结算户出账分类条件BHHDNWCC request body.
type UnitManagerNewBHHDNWCCBody struct {
	BHDUCMODY1 []*UnitManagerNewBHDUCMODY1 `json:"BHDUCMODY1,omitempty"`
	BHHDNWCCX1 []*UnitManagerNewBHHDNWCCX1 `json:"BHHDNWCCX1,omitempty"`
}

// UnitManagerNewBHHDNWCCResponseBody 7.2删除结算户出账分类条件BHHDNWCC response body.
type UnitManagerNewBHHDNWCCResponseBody struct {
	BHHDURTNZ1 []*UnitManagerNewBHHDURTNZ1 `json:"BHHDURTNZ1,omitempty"`
}

type UnitManagerNewBHHDNWCCResponse struct {
	Response struct {
		Body UnitManagerNewBHHDNWCCResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHDNWCQBody 7.3查询结算户出账分类条件BHHDNWCQ request body.
type UnitManagerNewBHHDNWCQBody struct {
	BHHDNWCQX1 []*UnitManagerNewBHHDNWCQX1 `json:"BHHDNWCQX1,omitempty"`
	BHHDNWCQY1 []*UnitManagerNewBHHDNWCQY1 `json:"BHHDNWCQY1,omitempty"`
}

// UnitManagerNewBHHDNWCQResponseBody 7.3查询结算户出账分类条件BHHDNWCQ response body.
type UnitManagerNewBHHDNWCQResponseBody struct {
	BHHDNWCQY1 []*UnitManagerNewBHHDNWCQY1 `json:"BHHDNWCQY1,omitempty"`
	BHHDNWCQZ1 []*UnitManagerNewBHHDNWCQZ1 `json:"BHHDNWCQZ1,omitempty"`
}

type UnitManagerNewBHHDNWCQResponse struct {
	Response struct {
		Body UnitManagerNewBHHDNWCQResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewNTDMARLTBody 8.1设置记账子单元入账条件NTDMARLT request body.
type UnitManagerNewNTDMARLTBody struct {
	Ntbusmody  []*UnitManagerNewNtbusmody  `json:"ntbusmody,omitempty"`
	Ntdmarltx1 []*UnitManagerNewNtdmarltx1 `json:"ntdmarltx1,omitempty"`
	Ntdmarltx2 []*UnitManagerNewNtdmarltx2 `json:"ntdmarltx2,omitempty"`
	Ntdmarltx3 []*UnitManagerNewNtdmarltx3 `json:"ntdmarltx3,omitempty"`
}

// UnitManagerNewNTDMARLTResponseBody 8.1设置记账子单元入账条件NTDMARLT response body.
type UnitManagerNewNTDMARLTResponseBody struct {
	Ntoprrtnz []*UnitManagerNewNtoprrtnz `json:"ntoprrtnz,omitempty"`
}

type UnitManagerNewNTDMARLTResponse struct {
	Response struct {
		Body UnitManagerNewNTDMARLTResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewNTDMARLDBody 8.2 删除记账子单元入账条件NTDMARLD request body.
type UnitManagerNewNTDMARLDBody struct {
	Ntbusmody  []*UnitManagerNewNtbusmody  `json:"ntbusmody,omitempty"`
	Ntdmarldx1 []*UnitManagerNewNtdmarldx1 `json:"ntdmarldx1,omitempty"`
	Ntdmarldx2 []*UnitManagerNewNtdmarldx2 `json:"ntdmarldx2,omitempty"`
	Ntdmarldx3 []*UnitManagerNewNtdmarldx3 `json:"ntdmarldx3,omitempty"`
}

// UnitManagerNewNTDMARLDResponseBody 8.2 删除记账子单元入账条件NTDMARLD response body.
type UnitManagerNewNTDMARLDResponseBody struct {
	Ntoprrtnz []*UnitManagerNewNtoprrtnz `json:"ntoprrtnz,omitempty"`
}

type UnitManagerNewNTDMARLDResponse struct {
	Response struct {
		Body UnitManagerNewNTDMARLDResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewNTDMARLQBody 8.3查询记账子单元入账条件NTDMARLQ request body.
type UnitManagerNewNTDMARLQBody struct {
	Ntdmarlqy1 []*UnitManagerNewNtdmarlqy1 `json:"ntdmarlqy1,omitempty"`
}

// UnitManagerNewNTDMARLQResponseBody 8.3查询记账子单元入账条件NTDMARLQ response body.
type UnitManagerNewNTDMARLQResponseBody struct {
	Ntdmarlqz1 []*UnitManagerNewNtdmarlqz1 `json:"ntdmarlqz1,omitempty"`
	Ntdmarlqz2 []*UnitManagerNewNtdmarlqz2 `json:"ntdmarlqz2,omitempty"`
	Ntdmarlqz3 []*UnitManagerNewNtdmarlqz3 `json:"ntdmarlqz3,omitempty"`
	Ntdmarlqy1 []*UnitManagerNewNtdmarlqy1 `json:"ntdmarlqy1,omitempty"`
}

type UnitManagerNewNTDMARLQResponse struct {
	Response struct {
		Body UnitManagerNewNTDMARLQResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewNTDMATMNBody 8.4 修改不满足记账子单元入账条件时的入账方式NTDMATMN request body.
type UnitManagerNewNTDMATMNBody struct {
	Ntbusmody  []*UnitManagerNewNtbusmody  `json:"ntbusmody,omitempty"`
	Ntdmatmnx1 []*UnitManagerNewNtdmatmnx1 `json:"ntdmatmnx1,omitempty"`
}

// UnitManagerNewNTDMATMNResponseBody 8.4 修改不满足记账子单元入账条件时的入账方式NTDMATMN response body.
type UnitManagerNewNTDMATMNResponseBody struct {
	Ntoprrtnz []*UnitManagerNewNtoprrtnz `json:"ntoprrtnz,omitempty"`
}

type UnitManagerNewNTDMATMNResponse struct {
	Response struct {
		Body UnitManagerNewNTDMATMNResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHDNDSABody 9.1新增个人账户白名单BHHDNDSA request body.
type UnitManagerNewBHHDNDSABody struct {
	BHDUCMODY1 []*UnitManagerNewBHDUCMODY1 `json:"BHDUCMODY1,omitempty"`
	BHHDNDSAX1 []*UnitManagerNewBHHDNDSAX1 `json:"BHHDNDSAX1,omitempty"`
}

// UnitManagerNewBHHDNDSAResponseBody 9.1新增个人账户白名单BHHDNDSA response body.
type UnitManagerNewBHHDNDSAResponseBody struct {
	BHHDURTNZ1 []*UnitManagerNewBHHDURTNZ1 `json:"BHHDURTNZ1,omitempty"`
}

type UnitManagerNewBHHDNDSAResponse struct {
	Response struct {
		Body UnitManagerNewBHHDNDSAResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHDNDSDBody 9.2删除个人账户白名单BHHDNDSD request body.
type UnitManagerNewBHHDNDSDBody struct {
	BHDUCMODY1 []*UnitManagerNewBHDUCMODY1 `json:"BHDUCMODY1,omitempty"`
	BHHDNDSDX1 []*UnitManagerNewBHHDNDSDX1 `json:"BHHDNDSDX1,omitempty"`
}

// UnitManagerNewBHHDNDSDResponseBody 9.2删除个人账户白名单BHHDNDSD response body.
type UnitManagerNewBHHDNDSDResponseBody struct {
	BHHDURTNZ1 []*UnitManagerNewBHHDURTNZ1 `json:"BHHDURTNZ1,omitempty"`
}

type UnitManagerNewBHHDNDSDResponse struct {
	Response struct {
		Body UnitManagerNewBHHDNDSDResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHDNDSQBody 9.3查询个人账户白名单BHHDNDSQ request body.
type UnitManagerNewBHHDNDSQBody struct {
	BHHDNDSQY1 []*UnitManagerNewBHHDNDSQY1 `json:"BHHDNDSQY1,omitempty"`
}

// UnitManagerNewBHHDNDSQResponseBody 9.3查询个人账户白名单BHHDNDSQ response body.
type UnitManagerNewBHHDNDSQResponseBody struct {
	BHHDNDSQZ1 []*UnitManagerNewBHHDNDSQZ1 `json:"BHHDNDSQZ1,omitempty"`
	BHHDNDSQY1 []*UnitManagerNewBHHDNDSQY1 `json:"BHHDNDSQY1,omitempty"`
}

type UnitManagerNewBHHDNDSQResponse struct {
	Response struct {
		Body UnitManagerNewBHHDNDSQResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewNTDCDADDBody 10.1设置记账子单元关联公司卡NTDCDADD request body.
type UnitManagerNewNTDCDADDBody struct {
	Ntbusmody  []*UnitManagerNewNtbusmody  `json:"ntbusmody,omitempty"`
	Ntdcdaddx1 []*UnitManagerNewNtdcdaddx1 `json:"ntdcdaddx1,omitempty"`
	Ntdcdaddx2 []*UnitManagerNewNtdcdaddx2 `json:"ntdcdaddx2,omitempty"`
}

// UnitManagerNewNTDCDADDResponseBody 10.1设置记账子单元关联公司卡NTDCDADD response body.
type UnitManagerNewNTDCDADDResponseBody struct {
	Ntdcdaddz1 []*UnitManagerNewNtdcdaddz1 `json:"ntdcdaddz1,omitempty"`
	Ntoprrtnz  []*UnitManagerNewNtoprrtnz  `json:"ntoprrtnz,omitempty"`
}

type UnitManagerNewNTDCDADDResponse struct {
	Response struct {
		Body UnitManagerNewNTDCDADDResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewNTDCDDLTBody 10.2解除记账子单元与公司卡关联NTDCDDLT request body.
type UnitManagerNewNTDCDDLTBody struct {
	Ntbusmody  []*UnitManagerNewNtbusmody  `json:"ntbusmody,omitempty"`
	Ntdcddltx1 []*UnitManagerNewNtdcddltx1 `json:"ntdcddltx1,omitempty"`
	Ntdcddltx2 []*UnitManagerNewNtdcddltx2 `json:"ntdcddltx2,omitempty"`
}

// UnitManagerNewNTDCDDLTResponseBody 10.2解除记账子单元与公司卡关联NTDCDDLT response body.
type UnitManagerNewNTDCDDLTResponseBody struct {
	Ntdcdaddz1 []*UnitManagerNewNtdcdaddz1 `json:"ntdcdaddz1,omitempty"`
	Ntoprrtnz  []*UnitManagerNewNtoprrtnz  `json:"ntoprrtnz,omitempty"`
}

type UnitManagerNewNTDCDDLTResponse struct {
	Response struct {
		Body UnitManagerNewNTDCDDLTResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewNTDCDRLTBody 10.3查询记账子单元关联公司卡列表NTDCDRLT request body.
type UnitManagerNewNTDCDRLTBody struct {
	Ntdcdrltx1 []*UnitManagerNewNtdcdrltx1 `json:"ntdcdrltx1,omitempty"`
}

// UnitManagerNewNTDCDRLTResponseBody 10.3查询记账子单元关联公司卡列表NTDCDRLT response body.
type UnitManagerNewNTDCDRLTResponseBody struct {
	Ntdcdrltz1 []*UnitManagerNewNtdcdrltz1 `json:"ntdcdrltz1,omitempty"`
	Ntdcdrltx1 []*UnitManagerNewNtdcdrltx1 `json:"ntdcdrltx1,omitempty"`
}

type UnitManagerNewNTDCDRLTResponse struct {
	Response struct {
		Body UnitManagerNewNTDCDRLTResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewNTDUMCTLBody 11.1设置记账子单元收款/出款限制NTDUMCTL request body.
type UnitManagerNewNTDUMCTLBody struct {
	Ntbusmody  []*UnitManagerNewNtbusmody  `json:"ntbusmody,omitempty"`
	Ntdumctlx1 []*UnitManagerNewNtdumctlx1 `json:"ntdumctlx1,omitempty"`
}

// UnitManagerNewNTDUMCTLResponseBody 11.1设置记账子单元收款/出款限制NTDUMCTL response body.
type UnitManagerNewNTDUMCTLResponseBody struct {
	Ntoprrtnz []*UnitManagerNewNtoprrtnz `json:"ntoprrtnz,omitempty"`
}

type UnitManagerNewNTDUMCTLResponse struct {
	Response struct {
		Body UnitManagerNewNTDUMCTLResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewNTDUMCLEBody 11.2解除记账子单元收款/出款限制NTDUMCLE request body.
type UnitManagerNewNTDUMCLEBody struct {
	Ntbusmody  []*UnitManagerNewNtbusmody  `json:"ntbusmody,omitempty"`
	Ntdumclex1 []*UnitManagerNewNtdumclex1 `json:"ntdumclex1,omitempty"`
}

// UnitManagerNewNTDUMCLEResponseBody 11.2解除记账子单元收款/出款限制NTDUMCLE response body.
type UnitManagerNewNTDUMCLEResponseBody struct {
	Ntoprrtnz []*UnitManagerNewNtoprrtnz `json:"ntoprrtnz,omitempty"`
}

type UnitManagerNewNTDUMCLEResponse struct {
	Response struct {
		Body UnitManagerNewNTDUMCLEResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHJINASBody 12.1 单笔设置子单元利率BHHJINAS request body.
type UnitManagerNewBHHJINASBody struct {
	BHHDINASX1 []*UnitManagerNewBHHDINASX1 `json:"BHHDINASX1,omitempty"`
	BHDUCMODY1 []*UnitManagerNewBHDUCMODY1 `json:"BHDUCMODY1,omitempty"`
}

// UnitManagerNewBHHJINASResponseBody 12.1 单笔设置子单元利率BHHJINAS response body.
type UnitManagerNewBHHJINASResponseBody struct {
	BHHDURTNZ1 []*UnitManagerNewBHHDURTNZ1 `json:"BHHDURTNZ1,omitempty"`
}

type UnitManagerNewBHHJINASResponse struct {
	Response struct {
		Body UnitManagerNewBHHJINASResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHJINCSBody 12.2 单笔删除子单元利率BHHJINCS request body.
type UnitManagerNewBHHJINCSBody struct {
	BHHDINCSX1 []*UnitManagerNewBHHDINCSX1 `json:"BHHDINCSX1,omitempty"`
	BHDUCMODY1 []*UnitManagerNewBHDUCMODY1 `json:"BHDUCMODY1,omitempty"`
}

// UnitManagerNewBHHJINCSResponseBody 12.2 单笔删除子单元利率BHHJINCS response body.
type UnitManagerNewBHHJINCSResponseBody struct {
	BHHDURTNZ1 []*UnitManagerNewBHHDURTNZ1 `json:"BHHDURTNZ1,omitempty"`
}

type UnitManagerNewBHHJINCSResponse struct {
	Response struct {
		Body UnitManagerNewBHHJINCSResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHJINQDBody 12.3 查询子单元利率BHHJINQD request body.
type UnitManagerNewBHHJINQDBody struct {
	BHHDINQDX1 []*UnitManagerNewBHHDINQDX1 `json:"BHHDINQDX1,omitempty"`
}

// UnitManagerNewBHHJINQDResponseBody 12.3 查询子单元利率BHHJINQD response body.
type UnitManagerNewBHHJINQDResponseBody struct {
	BHHDINQDZ1 []*UnitManagerNewBHHDINQDZ1 `json:"BHHDINQDZ1,omitempty"`
	BHHDINQDZ2 []*UnitManagerNewBHHDINQDZ2 `json:"BHHDINQDZ2,omitempty"`
}

type UnitManagerNewBHHJINQDResponse struct {
	Response struct {
		Body UnitManagerNewBHHJINQDResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHDINRPBody 12.4 子单元单笔提前结息BHHDINRP request body.
type UnitManagerNewBHHDINRPBody struct {
	BHHDINRPX1 []*UnitManagerNewBHHDINRPX1 `json:"BHHDINRPX1,omitempty"`
	BHDUCMODY1 []*UnitManagerNewBHDUCMODY1 `json:"BHDUCMODY1,omitempty"`
}

// UnitManagerNewBHHDINRPResponseBody 12.4 子单元单笔提前结息BHHDINRP response body.
type UnitManagerNewBHHDINRPResponseBody struct {
	BHHDURTNZ1 []*UnitManagerNewBHHDURTNZ1 `json:"BHHDURTNZ1,omitempty"`
}

type UnitManagerNewBHHDINRPResponse struct {
	Response struct {
		Body UnitManagerNewBHHDINRPResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHJINRSBody 12.5 子单元单笔手工结息BHHJINRS request body.
type UnitManagerNewBHHJINRSBody struct {
	BHHDINRSX1 []*UnitManagerNewBHHDINRSX1 `json:"BHHDINRSX1,omitempty"`
	BHDUCMODY1 []*UnitManagerNewBHDUCMODY1 `json:"BHDUCMODY1,omitempty"`
}

// UnitManagerNewBHHJINRSResponseBody 12.5 子单元单笔手工结息BHHJINRS response body.
type UnitManagerNewBHHJINRSResponseBody struct {
	BHHDURTNZ1 []*UnitManagerNewBHHDURTNZ1 `json:"BHHDURTNZ1,omitempty"`
}

type UnitManagerNewBHHJINRSResponse struct {
	Response struct {
		Body UnitManagerNewBHHJINRSResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHJINSDBody 12.6 查询子单元结息结果BHHJINSD request body.
type UnitManagerNewBHHJINSDBody struct {
	BHHDINSDX1 []*UnitManagerNewBHHDINSDX1 `json:"BHHDINSDX1,omitempty"`
}

// UnitManagerNewBHHJINSDResponseBody 12.6 查询子单元结息结果BHHJINSD response body.
type UnitManagerNewBHHJINSDResponseBody struct {
	BHHDINSDZ1 []*UnitManagerNewBHHDINSDZ1 `json:"BHHDINSDZ1,omitempty"`
	BHHDINSDZ2 []*UnitManagerNewBHHDINSDZ2 `json:"BHHDINSDZ2,omitempty"`
}

type UnitManagerNewBHHJINSDResponse struct {
	Response struct {
		Body UnitManagerNewBHHJINSDResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHDINDTBody 12.7合并结息明细查询BHHDINDT request body.
type UnitManagerNewBHHDINDTBody struct {
	BHHDINDTX1 []*UnitManagerNewBHHDINDTX1 `json:"BHHDINDTX1,omitempty"`
	BHHDINDTY1 []*UnitManagerNewBHHDINDTY1 `json:"BHHDINDTY1,omitempty"`
}

// UnitManagerNewBHHDINDTResponseBody 12.7合并结息明细查询BHHDINDT response body.
type UnitManagerNewBHHDINDTResponseBody struct {
	BHHDINDTZ1 []*UnitManagerNewBHHDINDTZ1 `json:"BHHDINDTZ1,omitempty"`
	BHHDINDTY1 []*UnitManagerNewBHHDINDTY1 `json:"BHHDINDTY1,omitempty"`
}

type UnitManagerNewBHHDINDTResponse struct {
	Response struct {
		Body UnitManagerNewBHHDINDTResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewNTDMAINTBody 12.8记账子单元利息试算NTDMAINT request body.
type UnitManagerNewNTDMAINTBody struct {
	Ntdmaintx1 []*UnitManagerNewNtdmaintx1 `json:"ntdmaintx1,omitempty"`
}

// UnitManagerNewNTDMAINTResponseBody 12.8记账子单元利息试算NTDMAINT response body.
type UnitManagerNewNTDMAINTResponseBody struct {
	Ntdmaintz1 []*UnitManagerNewNtdmaintz1 `json:"ntdmaintz1,omitempty"`
}

type UnitManagerNewNTDMAINTResponse struct {
	Response struct {
		Body UnitManagerNewNTDMAINTResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewNTDMIQRYBody 12.9记账子单元交易利息试算NTDMIQRY request body.
type UnitManagerNewNTDMIQRYBody struct {
	Ntdmiqryx1 []*UnitManagerNewNtdmiqryx1 `json:"ntdmiqryx1,omitempty"`
}

// UnitManagerNewNTDMIQRYResponseBody 12.9记账子单元交易利息试算NTDMIQRY response body.
type UnitManagerNewNTDMIQRYResponseBody struct {
	Ntdmiqryz1 []*UnitManagerNewNtdmiqryz1 `json:"ntdmiqryz1,omitempty"`
}

type UnitManagerNewNTDMIQRYResponse struct {
	Response struct {
		Body UnitManagerNewNTDMIQRYResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHDINBBBody 12.10按指定金额做利息试算（批量）BHHDINBB request body.
type UnitManagerNewBHHDINBBBody struct {
	BHHDINBBX1 []*UnitManagerNewBHHDINBBX1 `json:"BHHDINBBX1,omitempty"`
	BHHDINBBX2 []*UnitManagerNewBHHDINBBX2 `json:"BHHDINBBX2,omitempty"`
}

// UnitManagerNewBHHDINBBResponseBody 12.10按指定金额做利息试算（批量）BHHDINBB response body.
type UnitManagerNewBHHDINBBResponseBody struct {
	BHHDINBBZ1 []*UnitManagerNewBHHDINBBZ1 `json:"BHHDINBBZ1,omitempty"`
}

type UnitManagerNewBHHDINBBResponse struct {
	Response struct {
		Body UnitManagerNewBHHDINBBResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHDLMADBody 13.1单笔新增支付限额BHHDLMAD request body.
type UnitManagerNewBHHDLMADBody struct {
	BHDUCMODY1 []*UnitManagerNewBHDUCMODY1 `json:"BHDUCMODY1,omitempty"`
	BHHDLMADX1 []*UnitManagerNewBHHDLMADX1 `json:"BHHDLMADX1,omitempty"`
}

// UnitManagerNewBHHDLMADResponseBody 13.1单笔新增支付限额BHHDLMAD response body.
type UnitManagerNewBHHDLMADResponseBody struct {
	BHHDURTNZ1 []*UnitManagerNewBHHDURTNZ1 `json:"BHHDURTNZ1,omitempty"`
}

type UnitManagerNewBHHDLMADResponse struct {
	Response struct {
		Body UnitManagerNewBHHDLMADResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHDLMMDBody 13.2 单笔修改支付限额BHHDLMMD request body.
type UnitManagerNewBHHDLMMDBody struct {
	BHDUCMODY1 []*UnitManagerNewBHDUCMODY1 `json:"BHDUCMODY1,omitempty"`
	BHHDLMMDX1 []*UnitManagerNewBHHDLMMDX1 `json:"BHHDLMMDX1,omitempty"`
}

// UnitManagerNewBHHDLMMDResponseBody 13.2 单笔修改支付限额BHHDLMMD response body.
type UnitManagerNewBHHDLMMDResponseBody struct {
	BHHDURTNZ1 []*UnitManagerNewBHHDURTNZ1 `json:"BHHDURTNZ1,omitempty"`
}

type UnitManagerNewBHHDLMMDResponse struct {
	Response struct {
		Body UnitManagerNewBHHDLMMDResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHDLMCDBody 13.3 单笔关闭支付限额BHHDLMCD request body.
type UnitManagerNewBHHDLMCDBody struct {
	BHDUCMODY1 []*UnitManagerNewBHDUCMODY1 `json:"BHDUCMODY1,omitempty"`
	BHHDLMCDX1 []*UnitManagerNewBHHDLMCDX1 `json:"BHHDLMCDX1,omitempty"`
}

// UnitManagerNewBHHDLMCDResponseBody 13.3 单笔关闭支付限额BHHDLMCD response body.
type UnitManagerNewBHHDLMCDResponseBody struct {
	BHHDURTNZ1 []*UnitManagerNewBHHDURTNZ1 `json:"BHHDURTNZ1,omitempty"`
}

type UnitManagerNewBHHDLMCDResponse struct {
	Response struct {
		Body UnitManagerNewBHHDLMCDResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHDLMTPBody 13.4 单笔设置/关闭临时限额BHHDLMTP request body.
type UnitManagerNewBHHDLMTPBody struct {
	BHDUCMODY1 []*UnitManagerNewBHDUCMODY1 `json:"BHDUCMODY1,omitempty"`
	BHHDLMTPX1 []*UnitManagerNewBHHDLMTPX1 `json:"BHHDLMTPX1,omitempty"`
}

// UnitManagerNewBHHDLMTPResponseBody 13.4 单笔设置/关闭临时限额BHHDLMTP response body.
type UnitManagerNewBHHDLMTPResponseBody struct {
	BHHDURTNZ1 []*UnitManagerNewBHHDURTNZ1 `json:"BHHDURTNZ1,omitempty"`
}

type UnitManagerNewBHHDLMTPResponse struct {
	Response struct {
		Body UnitManagerNewBHHDLMTPResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHDLMQYBody 13.5 查询支付限额BHHDLMQY request body.
type UnitManagerNewBHHDLMQYBody struct {
	BHHDLMQYY1 []*UnitManagerNewBHHDLMQYY1 `json:"BHHDLMQYY1,omitempty"`
}

// UnitManagerNewBHHDLMQYResponseBody 13.5 查询支付限额BHHDLMQY response body.
type UnitManagerNewBHHDLMQYResponseBody struct {
	BHHDLMQYZ1 []*UnitManagerNewBHHDLMQYZ1 `json:"BHHDLMQYZ1,omitempty"`
	BHHDLMQYY1 []*UnitManagerNewBHHDLMQYY1 `json:"BHHDLMQYY1,omitempty"`
}

type UnitManagerNewBHHDLMQYResponse struct {
	Response struct {
		Body UnitManagerNewBHHDLMQYResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHDGPONBody 14.1. 开启多层级架构BHHDGPON request body.
type UnitManagerNewBHHDGPONBody struct {
	BHDUCMODY1 []*UnitManagerNewBHDUCMODY1 `json:"BHDUCMODY1,omitempty"`
	BHHDGPONX1 []*UnitManagerNewBHHDGPONX1 `json:"BHHDGPONX1,omitempty"`
}

// UnitManagerNewBHHDGPONResponseBody 14.1. 开启多层级架构BHHDGPON response body.
type UnitManagerNewBHHDGPONResponseBody struct {
	BHHDURTNZ1 []*UnitManagerNewBHHDURTNZ1 `json:"BHHDURTNZ1,omitempty"`
}

type UnitManagerNewBHHDGPONResponse struct {
	Response struct {
		Body UnitManagerNewBHHDGPONResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHDGPADBody 14.2新增普通分组BHHDGPAD request body.
type UnitManagerNewBHHDGPADBody struct {
	BHDUCMODY1 []*UnitManagerNewBHDUCMODY1 `json:"BHDUCMODY1,omitempty"`
	BHHDGPADX1 []*UnitManagerNewBHHDGPADX1 `json:"BHHDGPADX1,omitempty"`
}

// UnitManagerNewBHHDGPADResponseBody 14.2新增普通分组BHHDGPAD response body.
type UnitManagerNewBHHDGPADResponseBody struct {
	BHHDURTNZ1 []*UnitManagerNewBHHDURTNZ1 `json:"BHHDURTNZ1,omitempty"`
}

type UnitManagerNewBHHDGPADResponse struct {
	Response struct {
		Body UnitManagerNewBHHDGPADResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHDGPMNBody 14.3 修改普通分组BHHDGPMN request body.
type UnitManagerNewBHHDGPMNBody struct {
	BHDUCMODY1 []*UnitManagerNewBHDUCMODY1 `json:"BHDUCMODY1,omitempty"`
	BHHDGPMNX1 []*UnitManagerNewBHHDGPMNX1 `json:"BHHDGPMNX1,omitempty"`
}

// UnitManagerNewBHHDGPMNResponseBody 14.3 修改普通分组BHHDGPMN response body.
type UnitManagerNewBHHDGPMNResponseBody struct {
	BHHDURTNZ1 []*UnitManagerNewBHHDURTNZ1 `json:"BHHDURTNZ1,omitempty"`
}

type UnitManagerNewBHHDGPMNResponse struct {
	Response struct {
		Body UnitManagerNewBHHDGPMNResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHDGPSCBody 14.4关闭普通分组BHHDGPSC request body.
type UnitManagerNewBHHDGPSCBody struct {
	BHDUCMODY1 []*UnitManagerNewBHDUCMODY1 `json:"BHDUCMODY1,omitempty"`
	BHHDGPSCX1 []*UnitManagerNewBHHDGPSCX1 `json:"BHHDGPSCX1,omitempty"`
}

// UnitManagerNewBHHDGPSCResponseBody 14.4关闭普通分组BHHDGPSC response body.
type UnitManagerNewBHHDGPSCResponseBody struct {
	BHHDURTNZ1 []*UnitManagerNewBHHDURTNZ1 `json:"BHHDURTNZ1,omitempty"`
}

type UnitManagerNewBHHDGPSCResponse struct {
	Response struct {
		Body UnitManagerNewBHHDGPSCResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHDGPBLBody 14.5分组余额查询BHHDGPBL request body.
type UnitManagerNewBHHDGPBLBody struct {
	BHHDGPBLX1 []*UnitManagerNewBHHDGPBLX1 `json:"BHHDGPBLX1,omitempty"`
}

// UnitManagerNewBHHDGPBLResponseBody 14.5分组余额查询BHHDGPBL response body.
type UnitManagerNewBHHDGPBLResponseBody struct {
	BHHDGPBLZ1 []*UnitManagerNewBHHDGPBLZ1 `json:"BHHDGPBLZ1,omitempty"`
}

type UnitManagerNewBHHDGPBLResponse struct {
	Response struct {
		Body UnitManagerNewBHHDGPBLResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHDZCTGBody 14.6分组交易查询BHHDZCTG request body.
type UnitManagerNewBHHDZCTGBody struct {
	BHHDZCTGX1 []*UnitManagerNewBHHDZCTGX1 `json:"BHHDZCTGX1,omitempty"`
	BHHDZCTGY1 []*UnitManagerNewBHHDZCTGY1 `json:"BHHDZCTGY1,omitempty"`
}

// UnitManagerNewBHHDZCTGResponseBody 14.6分组交易查询BHHDZCTG response body.
type UnitManagerNewBHHDZCTGResponseBody struct {
	BHHDZCTGZ1 []*UnitManagerNewBHHDZCTGZ1 `json:"BHHDZCTGZ1,omitempty"`
	BHHDZCTGY1 []*UnitManagerNewBHHDZCTGY1 `json:"BHHDZCTGY1,omitempty"`
}

type UnitManagerNewBHHDZCTGResponse struct {
	Response struct {
		Body UnitManagerNewBHHDZCTGResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHDGNADBody 14.7分组下新增下级子单元BHHDGNAD request body.
type UnitManagerNewBHHDGNADBody struct {
	BHDUCMODY1 []*UnitManagerNewBHDUCMODY1 `json:"BHDUCMODY1,omitempty"`
	BHHDGNADX1 []*UnitManagerNewBHHDGNADX1 `json:"BHHDGNADX1,omitempty"`
}

// UnitManagerNewBHHDGNADResponseBody 14.7分组下新增下级子单元BHHDGNAD response body.
type UnitManagerNewBHHDGNADResponseBody struct {
	BHHDURTNZ1 []*UnitManagerNewBHHDURTNZ1 `json:"BHHDURTNZ1,omitempty"`
}

type UnitManagerNewBHHDGNADResponse struct {
	Response struct {
		Body UnitManagerNewBHHDGNADResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHDGNCDBody 14.8分组下删除下级子单元BHHDGNCD request body.
type UnitManagerNewBHHDGNCDBody struct {
	BHDUCMODY1 []*UnitManagerNewBHDUCMODY1 `json:"BHDUCMODY1,omitempty"`
	BHHDGNCDX1 []*UnitManagerNewBHHDGNCDX1 `json:"BHHDGNCDX1,omitempty"`
}

// UnitManagerNewBHHDGNCDResponseBody 14.8分组下删除下级子单元BHHDGNCD response body.
type UnitManagerNewBHHDGNCDResponseBody struct {
	BHHDURTNZ1 []*UnitManagerNewBHHDURTNZ1 `json:"BHHDURTNZ1,omitempty"`
}

type UnitManagerNewBHHDGNCDResponse struct {
	Response struct {
		Body UnitManagerNewBHHDGNCDResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHDGNAGBody 14.9分组下新增下级分组BHHDGNAG request body.
type UnitManagerNewBHHDGNAGBody struct {
	BHDUCMODY1 []*UnitManagerNewBHDUCMODY1 `json:"BHDUCMODY1,omitempty"`
	BHHDGNAGX1 []*UnitManagerNewBHHDGNAGX1 `json:"BHHDGNAGX1,omitempty"`
}

// UnitManagerNewBHHDGNAGResponseBody 14.9分组下新增下级分组BHHDGNAG response body.
type UnitManagerNewBHHDGNAGResponseBody struct {
	BHHDURTNZ1 []*UnitManagerNewBHHDURTNZ1 `json:"BHHDURTNZ1,omitempty"`
}

type UnitManagerNewBHHDGNAGResponse struct {
	Response struct {
		Body UnitManagerNewBHHDGNAGResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHDGNCGBody 14.10分组下删除下级分组BHHDGNCG request body.
type UnitManagerNewBHHDGNCGBody struct {
	BHDUCMODY1 []*UnitManagerNewBHDUCMODY1 `json:"BHDUCMODY1,omitempty"`
	BHHDGNCGX1 []*UnitManagerNewBHHDGNCGX1 `json:"BHHDGNCGX1,omitempty"`
}

// UnitManagerNewBHHDGNCGResponseBody 14.10分组下删除下级分组BHHDGNCG response body.
type UnitManagerNewBHHDGNCGResponseBody struct {
	BHHDURTNZ1 []*UnitManagerNewBHHDURTNZ1 `json:"BHHDURTNZ1,omitempty"`
}

type UnitManagerNewBHHDGNCGResponse struct {
	Response struct {
		Body UnitManagerNewBHHDGNCGResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHDGNGSBody 14.11查询下级节点信息BHHDGNGS request body.
type UnitManagerNewBHHDGNGSBody struct {
	BHHDGNGSX1 []*UnitManagerNewBHHDGNGSX1 `json:"BHHDGNGSX1,omitempty"`
	BHHDGNGSY1 []*UnitManagerNewBHHDGNGSY1 `json:"BHHDGNGSY1,omitempty"`
}

// UnitManagerNewBHHDGNGSResponseBody 14.11查询下级节点信息BHHDGNGS response body.
type UnitManagerNewBHHDGNGSResponseBody struct {
	BHHDGNGSZ1 []*UnitManagerNewBHHDGNGSZ1 `json:"BHHDGNGSZ1,omitempty"`
	BHHDGNGSY1 []*UnitManagerNewBHHDGNGSY1 `json:"BHHDGNGSY1,omitempty"`
}

type UnitManagerNewBHHDGNGSResponse struct {
	Response struct {
		Body UnitManagerNewBHHDGNGSResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// UnitManagerNewBHHDGNGTBody 14.12 查询上级节点信息BHHDGNGT request body.
type UnitManagerNewBHHDGNGTBody struct {
	BHHDGNGTX1 []*UnitManagerNewBHHDGNGTX1 `json:"BHHDGNGTX1,omitempty"`
}

// UnitManagerNewBHHDGNGTResponseBody 14.12 查询上级节点信息BHHDGNGT response body.
type UnitManagerNewBHHDGNGTResponseBody struct {
	BHHDGNGTZ1 []*UnitManagerNewBHHDGNGTZ1 `json:"BHHDGNGTZ1,omitempty"`
}

type UnitManagerNewBHHDGNGTResponse struct {
	Response struct {
		Body UnitManagerNewBHHDGNGTResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}
