package models

// MainAccountPayResultResponse
//  @Description:  支付结果
//  @Author  ahKevinXy
//  @Date2023-04-10 15:12:08
type MainAccountPayResultResponse struct {
	Response struct {
		Body struct {
			Bb1Qrybtz1 []struct {
				BthNbr string `json:"bthNbr,omitempty"`
				BusCod string `json:"busCod,omitempty"`
				BusMod string `json:"busMod,omitempty"`
				DtlAmt string `json:"dtlAmt,omitempty"`
				DtlNum string `json:"dtlNum,omitempty"`
				ReqSts string `json:"reqSts,omitempty"`
				Rsv30Z string `json:"rsv30z,omitempty"`
				RtnFlg string `json:"rtnFlg,omitempty"`
				SucAmt string `json:"sucAmt,omitempty"`
				SucNum string `json:"sucNum,omitempty"`
				TrsDat string `json:"trsDat,omitempty"`
				TrsTim string `json:"trsTim,omitempty"`
			} `json:"bb1qrybtz1,omitempty"`
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

// QueryAccountPaymentTransInfoResponse
//  @Description: 企银批量支付批次查询 返回结果
//  @Author  ahKevinXy
//  @Date2023-04-13 16:45:22
type QueryAccountPaymentTransInfoResponse struct {
	Response struct {
		Body struct {
			Bb1Qrybtz1 []struct {
				BthNbr string `json:"bthNbr,omitempty"`
				BusCod string `json:"busCod,omitempty"`
				BusMod string `json:"busMod,omitempty"`
				DtlAmt string `json:"dtlAmt,omitempty"`
				DtlNum string `json:"dtlNum,omitempty"`
				ReqSts string `json:"reqSts,omitempty"`
				Rsv30Z string `json:"rsv30z,omitempty"`
				RtnFlg string `json:"rtnFlg,omitempty"`
				SucAmt string `json:"sucAmt,omitempty"`
				SucNum string `json:"sucNum,omitempty"`
				TrsDat string `json:"trsDat,omitempty"`
				TrsTim string `json:"trsTim,omitempty"`
			} `json:"bb1qrybtz1,omitempty"`
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

// QueryAccountPaymentDetailResponse
//  @Description:  明细数据
//  @Author  ahKevinXy
//  @Date2023-04-13 16:52:47
type QueryAccountPaymentDetailResponse struct {
	Response struct {
		Body struct {
			Bb1Qrybdy1 []struct {
				BthNbr string `json:"bthNbr,omitempty"`
			} `json:"bb1qrybdy1,omitempty"`
			Bb1Qrybdz1 []struct {
				BnkFlg string `json:"bnkFlg,omitempty"`
				BrdNbr string `json:"brdNbr,omitempty"`
				BthNbr string `json:"bthNbr,omitempty"`
				CcyNbr string `json:"ccyNbr,omitempty"`
				CnvNbr string `json:"cnvNbr,omitempty"`
				CopNbr string `json:"copNbr,omitempty"`
				CouCod string `json:"couCod,omitempty"`
				CrtAcc string `json:"crtAcc,omitempty"`
				CrtBbk string `json:"crtBbk,omitempty"`
				CrtNam string `json:"crtNam,omitempty"`
				CrtSqn string `json:"crtSqn,omitempty"`
				CtyCod string `json:"ctyCod,omitempty"`
				DbtAcc string `json:"dbtAcc,omitempty"`
				DbtBbk string `json:"dbtBbk,omitempty"`
				DmaNbr string `json:"dmaNbr,omitempty"`
				DrpFlg string `json:"drpFlg,omitempty"`
				EptDat string `json:"eptDat,omitempty"`
				EptTim string `json:"eptTim,omitempty"`
				ErrTxt string `json:"errTxt,omitempty"`
				InpTel string `json:"inpTel,omitempty"`
				IssRef string `json:"issRef,omitempty"`
				KjtAcc string `json:"kjtAcc,omitempty"`
				MsgTxt string `json:"msgTxt,omitempty"`
				NpsTyp string `json:"npsTyp,omitempty"`
				NtfCh1 string `json:"ntfCh1,omitempty"`
				NtfCh2 string `json:"ntfCh2,omitempty"`
				NusAge string `json:"nusAge,omitempty"`
				PasNbr string `json:"pasNbr,omitempty"`
				PayTyp string `json:"payTyp,omitempty"`
				RcvChk string `json:"rcvChk,omitempty"`
				RemNbr string `json:"remNbr,omitempty"`
				ReqSts string `json:"reqSts,omitempty"`
				RsvAmt string `json:"rsvAmt,omitempty"`
				RsvNa1 string `json:"rsvNa1,omitempty"`
				RsvNa2 string `json:"rsvNa2,omitempty"`
				RsvNb1 string `json:"rsvNb1,omitempty"`
				RsvNb2 string `json:"rsvNb2,omitempty"`
				RtnFlg string `json:"rtnFlg,omitempty"`
				SplC80 string `json:"splC80,omitempty"`
				StlChn string `json:"stlChn,omitempty"`
				TrsAmt string `json:"trsAmt,omitempty"`
				TrsCat string `json:"trsCat,omitempty"`
				TrsTyp string `json:"trsTyp,omitempty"`
				TrxAmt string `json:"trxAmt,omitempty"`
				TrxCod string `json:"trxCod,omitempty"`
				YurRef string `json:"yurRef,omitempty"`
			} `json:"bb1qrybdz1,omitempty"`
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
