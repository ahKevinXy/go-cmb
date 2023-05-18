package models

// QueryAccountCallbackAsyncResponse
//  @Description:   异步回单
//  @Author  ahKevinXy
//  @Date2023-04-10 15:07:07
type QueryAccountCallbackAsyncResponse struct {
	Response struct {
		Head struct {
			Funcode    string `json:"funcode,omitempty"`
			Reqid      string `json:"reqid,omitempty"`
			Resultcode string `json:"resultcode,omitempty"`
			Resultmsg  string `json:"resultmsg,omitempty"`
			Rspid      string `json:"rspid,omitempty"`
			Userid     string `json:"userid,omitempty"`
		} `json:"head,omitempty"`
		Body struct {
			Asycalhdz1 struct {
				Rtncod string `json:"rtncod,omitempty"` // 返回码
				Rtnmsg string `json:"rtnmsg,omitempty"` // 返回信息
				Rtndat string `json:"rtndat,omitempty"` // 打印任务编号
			} `json:"asycalhdz1,omitempty"`
			Ctnkeyz2 struct {
				Begamt string `json:"begamt,omitempty"`
				Begdat string `json:"begdat,omitempty"`
				Daltag string `json:"daltag,omitempty"`
				Eacnbr string `json:"eacnbr,omitempty"`
				Endamt string `json:"endamt,omitempty"`
				Enddat string `json:"enddat,omitempty"`
				Nxtdat string `json:"nxtdat,omitempty"`
				Nxtnbr string `json:"nxtnbr,omitempty"`
				Nxttim string `json:"nxttim,omitempty"`
				Oprtyp string `json:"oprtyp,omitempty"`
				Pagcnt string `json:"pagcnt,omitempty"`
				Pattyp string `json:"pattyp,omitempty"`
				Predat string `json:"predat,omitempty"`
				Prenbr string `json:"prenbr,omitempty"`
				Pretim string `json:"pretim,omitempty"`
				Rrccod string `json:"rrccod,omitempty"`
				Rrcflg string `json:"rrcflg,omitempty"`
				Spc100 string `json:"spc100,omitempty"`
			} `json:"ctnkeyz2,omitempty"`
		} `json:"body,omitempty"`
	} `json:"response,omitempty"`
}

// SingleCallBackPdfResponse
//  @Description:   获取单笔交易回单
//  @Author  ahKevinXy
//  @Date2023-04-10 15:10:52
type SingleCallBackPdfResponse struct {
	Response struct {
		Body struct {
			Checod string `json:"checod,omitempty"`
			Fildat string `json:"fildat,omitempty"`
			Istnbr string `json:"istnbr,omitempty"`
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

// QueryAccountCallbackDownloadPdfResponse
//  @Description:   获取回单文件
//  @Author  ahKevinXy
//  @Date2023-04-10 15:10:42
type QueryAccountCallbackDownloadPdfResponse struct {
	Response struct {
		Body struct {
			Fileurl string `json:"fileurl,omitempty"`
			Fintim  string `json:"fintim,omitempty"`
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

type BatchStatementQueryResponse struct {
	Response struct {
		Body struct {
			Dctrspdfz []struct {
				Printid string `json:"printid,omitempty"`
				Taskid  string `json:"taskid,omitempty"`
			} `json:"dctrspdfz,omitempty"`
			Ntqacctny []struct {
				Ctndta string `json:"ctndta,omitempty"`
				Pagnbr string `json:"pagnbr,omitempty"`
			} `json:"ntqacctny,omitempty"`
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

type QueryPayrollStatementResponse struct {
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
