package models

// AddUnitAccountV1Response
//
//	@Description:  添加记账单元
//	@Author  ahKevinXy
//	@Date2023-04-13 17:43:53
type AddUnitAccountV1Response struct {
	Response struct {
		Body struct {
			Ntdumaddz1 []struct {
				Dyanbr string `json:"dyanbr,omitempty"`
				Errcod string `json:"errcod,omitempty"`
				Inbacc string `json:"inbacc,omitempty"`
				Reqnbr string `json:"reqnbr,omitempty"`
				Reqsts string `json:"reqsts,omitempty"`
			} `json:"ntdumaddz1,omitempty"`
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

// CloseUnitAccountResponse
//
//	@Description:  关闭记账单元
//	@Author  ahKevinXy
//	@Date2023-04-13 17:44:38
type CloseUnitAccountResponse struct {
	Response struct {
		Body struct {
			Ntdumdltz1 []struct {
				Dyanbr string `json:"dyanbr,omitempty"`
				Errcod string `json:"errcod,omitempty"`
				Inbacc string `json:"inbacc,omitempty"`
				Reqnbr string `json:"reqnbr,omitempty"`
				Reqsts string `json:"reqsts,omitempty"`
			} `json:"ntdumdltz1,omitempty"`
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

// AccountUnitInfoResponse
//
//	@Description:  获取余额信息
//	@Author  ahKevinXy
//	@Date  2023-04-13 17:50:47
type AccountUnitInfoResponse struct {
	Response struct {
		Body struct {
			Ntdmabalz []struct {
				Accnbr string `json:"accnbr"`
				Actbal string `json:"actbal"`
				Dmaccy string `json:"dmaccy"`
				Dmanam string `json:"dmanam"`
				Dmanbr string `json:"dmanbr"`
				Eftdat string `json:"eftdat"`
				Stscod string `json:"stscod"`
			} `json:"ntdmabalz"`
		} `json:"body"`
		Head struct {
			Bizcode    string `json:"bizcode"`
			Funcode    string `json:"funcode"`
			Reqid      string `json:"reqid"`
			Resultcode string `json:"resultcode"`
			Resultmsg  string `json:"resultmsg"`
			Rspid      string `json:"rspid"`
			Userid     string `json:"userid"`
		} `json:"head"`
	} `json:"response"`
}

// QueryUnitTransByBusNoResponse
//
//	@Description:  获取交易信息
//	@Author  ahKevinXy
//	@Date  2023-04-13 18:04:30
type QueryUnitTransByBusNoResponse struct {
	Response struct {
		Body struct {
			Ntdmaqryz1 []struct {
				Bbknbr string `json:"bbknbr,omitempty"`
				Buscod string `json:"buscod,omitempty"`
				Busmod string `json:"busmod,omitempty"`
				Inbacc string `json:"inbacc,omitempty"`
				Oprdat string `json:"oprdat,omitempty"`
				Oprtyp string `json:"oprtyp,omitempty"`
				Reqnbr string `json:"reqnbr,omitempty"`
				Reqsta string `json:"reqsta,omitempty"`
				Trsamt string `json:"trsamt,omitempty"`
				Yurref string `json:"yurref,omitempty"`
			} `json:"ntdmaqryz1,omitempty"`
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

type QueryUnitAccountBalanceHistoryResponse struct {
	Response struct {
		Body struct {
			Ntdmahbdz1 []struct {
				Accnbr string `json:"accnbr,omitempty"`
				Ccynbr string `json:"ccynbr,omitempty"`
				Dmabal string `json:"dmabal,omitempty"`
				Dmanam string `json:"dmanam,omitempty"`
				Dmanbr string `json:"dmanbr,omitempty"`
				Psedat string `json:"psedat,omitempty"`
			} `json:"ntdmahbdz1,omitempty"`
			Ntdmahbdz2 []struct {
				Ctnflg string `json:"ctnflg,omitempty"`
			} `json:"ntdmahbdz2,omitempty"`
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

type QueryUnitAccountSingleBalanceHistoryResponse struct {
	Response struct {
		Body struct {
			Ntdmahadz1 []struct {
				Accnbr string `json:"accnbr,omitempty"`
				Ccynbr string `json:"ccynbr,omitempty"`
				Dmabal string `json:"dmabal,omitempty"`
				Dmanam string `json:"dmanam,omitempty"`
				Dmanbr string `json:"dmanbr,omitempty"`
				Psedat string `json:"psedat,omitempty"`
			} `json:"ntdmahadz1,omitempty"`
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

// UpdateUnitAccountV1Response
//
//	@Description:  更新记账单元
//	@Author  ahKevinXy
//	@Date  2023-05-23 18:38:00
type UpdateUnitAccountV1Response struct {
	Response struct {
		Body struct {
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

// QueryUnitAccountInfoV2Response
//
//	@Description:  查询记账单元返回
//	@Author  ahKevinXy
//	@Date  2023-05-24 10:05:18
type QueryUnitAccountInfoV2Response struct {
	Response struct {
		Body struct {
			Ntdumqryz1 []struct {
				Dyaccy string `json:"dyaccy,omitempty"` // 币种
				Dyanam string `json:"dyanam,omitempty"` // 单元名称
				Dyanbr string `json:"dyanbr,omitempty"` // 记账单元编号
				Eftdat string `json:"eftdat,omitempty"` // 生效时间
				Enddat string `json:"enddat,omitempty"` // 终止时间
				Inbacc string `json:"inbacc,omitempty"` // 账号
				Ovrctl string `json:"ovrctl,omitempty"` // 是否允许透支
				Pstbal string `json:"pstbal,omitempty"` // 上日余额
				Pstdat string `json:"pstdat,omitempty"` // 上日日期
				Rcvlck string `json:"rcvlck,omitempty"` // 是否入账控制
				Stscod string `json:"stscod,omitempty"` // 状态
				Uptbal string `json:"uptbal,omitempty"` // 实时更新余额
			} `json:"ntdumqryz1,omitempty"`
			Ntdumqryz2 []struct {
				Avilmt string `json:"avilmt,omitempty"` // 余额
				Ballmt string `json:"ballmt,omitempty"` // 余额上线
				Dyanbr string `json:"dyanbr,omitempty"` // 记账单元编号
				Lmtflg string `json:"lmtflg,omitempty"` // 额度标志
			} `json:"ntdumqryz2,omitempty"`
			Ntdumqryy1 []*Ntdumqryy1 `json:"ntdumqryy1,omitempty"`
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

// AccountSetWhitePayResponse
//
//	@Description:  设置白名单返回
//	@Author  colornote
//	@Date  2023-06-27 10:05:18
type AccountSetWhitePayResponse struct {
	Response struct {
		Body struct {
			Ntoprrtnz []struct {
				Errcod string `json:"errcod,omitempty"` // 错误码
				Reqnbr string `json:"reqnbr,omitempty"` // 请求流水号
				Reqsts string `json:"reqsts,omitempty"` // 请求状态
			} `json:"ntoprrtnz,omitempty"`
		} `json:"body,omitempty"`
		Head struct {
			Bizcode    string `json:"bizcode,omitempty"`    // 业务码
			Funcode    string `json:"funcode,omitempty"`    // 功能码
			Reqid      string `json:"reqid,omitempty"`      // 请求流水号
			Resultcode string `json:"resultcode,omitempty"` // 返回码
			Resultmsg  string `json:"resultmsg,omitempty"`  // 返回信息
			Rspid      string `json:"rspid,omitempty"`      // 响应流水号
			Userid     string `json:"userid,omitempty"`     // 用户ID
		} `json:"head,omitempty"`
	} `json:"response,omitempty"`
}
