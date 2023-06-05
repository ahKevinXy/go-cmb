package models

type QueryPayrollOldTransCodeResponse struct {
	Response struct {
		Body struct {
			Ntagtls2Z []struct {
				Accnbr  string `json:"accnbr,omitempty"`
				CTrstyp string `json:"c_trstyp,omitempty"`
				Ccynbr  string `json:"ccynbr,omitempty"`
				Cnvnbr  string `json:"cnvnbr,omitempty"`
				Eftdat  string `json:"eftdat,omitempty"`
				Expdat  string `json:"expdat,omitempty"`
				Sgndat  string `json:"sgndat,omitempty"`
				Stscod  string `json:"stscod,omitempty"`
				Trstyp  string `json:"trstyp,omitempty"`
			} `json:"ntagtls2z,omitempty"`
		} `json:"body,omitempty"`
		Head struct {
			Bizcode    string `json:"bizcode,omitempty"`
			Funcode    string `json:"funcode,omitempty"`
			Reqid      string `json:"reqid,omitempty"`
			Resultcode string `json:"resultcode,omitempty"`
			Resultmsg  string `json:"resultmsg,omitempty"`
			Rspid      string `json:"rspid,omitempty"`
			Userid     string `json:"userid,omitempty"`
		} `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

type PayrollOldCreditOtherBySupResponse struct {
	Response struct {
		Body struct {
			Ntagcagcz1 []struct {
				Reqnbr string `json:"reqnbr,omitempty"`
				Reqsta string `json:"reqsta,omitempty"`
				Rsv50Z string `json:"rsv50z,omitempty"`
			} `json:"ntagcagcz1,omitempty"`
		} `json:"body,omitempty"`
		Head struct {
			Bizcode    string `json:"bizcode,omitempty"`
			Funcode    string `json:"funcode,omitempty"`
			Reqid      string `json:"reqid,omitempty"`
			Resultcode string `json:"resultcode,omitempty"`
			Resultmsg  string `json:"resultmsg,omitempty"`
			Rspid      string `json:"rspid,omitempty"`
			Userid     string `json:"userid,omitempty"`
		} `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

type GetPayrollPdfResponse struct {
	Response struct {
		Body struct {
			Begidx  string `json:"begidx,omitempty"`
			Printid string `json:"printid,omitempty"`
			Total   string `json:"total,omitempty"`
		} `json:"body,omitempty"`
		Head struct {
			Bizcode    string `json:"bizcode,omitempty"`
			Funcode    string `json:"funcode,omitempty"`
			Reqid      string `json:"reqid,omitempty"`
			Resultcode string `json:"resultcode,omitempty"`
			Resultmsg  string `json:"resultmsg,omitempty"`
			Rspid      string `json:"rspid,omitempty"`
			Userid     string `json:"userid,omitempty"`
		} `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// QueryOldPayRollOrderResponse
//  @Description:   代发概要查询
//  @Author  ahKevinXy
//  @Date  2023-06-05 17:50:15
type QueryOldPayRollOrderResponse struct {
	Response struct {
		Body struct {
			Ntagcinny1 []struct {
				Bgndat string `json:"bgndat,omitempty"`
				Buscod string `json:"buscod,omitempty"`
				Busmod string `json:"busmod,omitempty"`
				Datflg string `json:"datflg,omitempty"`
				Enddat string `json:"enddat,omitempty"`
				Ctnkey string `json:"ctnkey,omitempty"`
			} `json:"ntagcinny1,omitempty"`
			Ntagcinqz []struct {
				Accnbr string `json:"accnbr,omitempty"` //账号
				Accnam string `json:"accnam,omitempty"` //账号名称
				Agctyp string `json:"agctyp,omitempty"` // 是否超网
				Athflg string `json:"athflg,omitempty"` // 附件标志
				Bbknbr string `json:"bbknbr,omitempty"` // 分行号
				Buscod string `json:"buscod,omitempty"` // 业务类型
				Busmod string `json:"busmod,omitempty"` // 业务模式
				Ccynbr string `json:"ccynbr,omitempty"` // 币种
				Eptdat string `json:"eptdat,omitempty"` // 期望日期
				Epttim string `json:"epttim,omitempty"` // 期望时间
				Nusage string `json:"nusage,omitempty"` // 备注 用途
				Oprals string `json:"oprals,omitempty"` // 操作别名
				Oprdat string `json:"oprdat,omitempty"` // 操作时间
				Oprsqn string `json:"oprsqn,omitempty"` // 待操作序列
				Oprstp string `json:"oprstp,omitempty"` // 待操作步骤
				Rtnflg string `json:"rtnflg,omitempty"` // 业务请求结果 S 成功 F 失败 R 否决 D 过期 C 撤销
				Reqnbr string `json:"reqnbr,omitempty"` // 流程实例号
				Reqsta string `json:"reqsta,omitempty"` //业务请求状态  AUT 等待处理  NTE 终审完毕 BNK 银行处理  FIN 完成 OPR 数据接受中 APW 银行人工审批 WRF 可疑
				Seqcod string `json:"seqcod,omitempty"` //处理批次
				Stscod string `json:"stscod,omitempty"` // 记录状态
				Sucamt string `json:"sucamt,omitempty"` // 成功金额
				Sucnum string `json:"sucnum,omitempty"` // 成功数目
				Totamt string `json:"totamt,omitempty"` // 总金额
				Dmanbr string `json:"dmanbr,omitempty"` //  子单元
				Trsnum string `json:"trsnum,omitempty"` // 交易笔数
				Trstyp string `json:"trstyp,omitempty"` //  交易类型
				Yurref string `json:"yurref,omitempty"` // 业务参考号
			} `json:"ntagcinqz,omitempty"`
		} `json:"body,omitempty"`
		Head struct {
			Bizcode    string `json:"bizcode,omitempty"`
			Funcode    string `json:"funcode,omitempty"`
			Reqid      string `json:"reqid,omitempty"`
			Resultcode string `json:"resultcode,omitempty"`
			Resultmsg  string `json:"resultmsg,omitempty"`
			Rspid      string `json:"rspid,omitempty"`
			Userid     string `json:"userid,omitempty"`
		} `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// QueryOldPayRollOrderDetailResponse
//  @Description:  概要查询
//  @Author  ahKevinXy
//  @Date  2023-06-05 17:51:22
type QueryOldPayRollOrderDetailResponse struct {
	Response struct {
		Body struct {
			Ntagcdtly1 []struct {
				Accnam string `json:"accnam,omitempty"` //户名
				Accnbr string `json:"accnbr,omitempty"` // 账号
				Bnkflg string `json:"bnkflg,omitempty"` // 系统内标志
				Lgramt string `json:"lgramt,omitempty"` //
				Stscod string `json:"stscod,omitempty"` // 记录状态  A  待处理 E 失败 S 成功
				Trsamt string `json:"trsamt,omitempty"` //  金额
				Trsdsp string `json:"trsdsp,omitempty"` // 注释
				Trxseq string `json:"trxseq,omitempty"` // 交易序号
			} `json:"ntagcdtly1,omitempty"`
			Ntagdinfy1 []Ntagdinfy1 `json:"ntagdinfy1,omitempty"`
		} `json:"body,omitempty"`
		Head struct {
			Bizcode    string `json:"bizcode,omitempty"`
			Funcode    string `json:"funcode,omitempty"`
			Reqid      string `json:"reqid,omitempty"`
			Resultcode string `json:"resultcode,omitempty"`
			Resultmsg  string `json:"resultmsg,omitempty"`
			Rspid      string `json:"rspid,omitempty"`
			Userid     string `json:"userid,omitempty"`
		} `json:"head,omitempty"`
	} `json:"response,omitempty"`
}
