package constants

// 账务查询
const (
	CmbAccountCanPayMod                   = "DCLISMOD"     //可经办业务模式查询 DCLISMOD
	CmbAccountUserList                    = "DCLISACC"     // 查询可经办的账户列表 DCLISACC todo
	CmbAccountInfo                        = "NTQACINF"     //账户详细信息查询 NTQACINF
	CmbAccountHistoryBalance              = "NTQABINF"     //查询账户历史余额 NTQABINF
	CmbAccountQueryBankLinkNo             = "NTACCBBK"     //查询分行号信息
	CmbAccountBatchQueryBalance           = "NTQADINF"     //批量查询余额
	CmbAccountQueryTransInfo              = "DCTRSINF"     //账户交易信息查询DCTRSINF
	CmbAccountGetTransInfo                = "GetTransInfo" //账户交易信息查询GetTransInfo
	CmbAccountQueryTransDetail            = "DCTRSPDF"     //账户交易明细对账单查询 DCTRSPDF
	CmbAccountQuerySingleStatement        = "DCSIGREC"     //单笔回单查询DCSIGREC
	CmbAccountQueryAsyncStatement         = "ASYCALHD"     //电子回单异步查询ASYCALHD
	CmbAccountQueryAsyncDownloadStatement = "DCTASKID"     //异步打印结果查询DCTASKID
)

// 支付转账
const (
	CmbAccountSinglePay           = "BB1PAYOP" // 企银支付单笔经办BB1PAYOP
	CmbAccountPayModQuery         = "BB1PAYQR" // 企银支付业务查询BB1PAYQR
	CmbAccountBatchPayQuery       = "BB1PAYBH" // 企银支付批量经办BB1PAYBH
	CmbAccountBatchPayQueryInfo   = "BB1QRYBT" // 企银批量支付批次查询BB1QRYBT
	CmbAccountBatchPayQueryDetail = "BB1QRYBD" // 企银批量支付明细查询BB1QRYBD
	CmbAccountBatchPayRefund      = "BB1PAYQB" // 支付退票明细查询BB1PAYQB

)

//代发代扣 old

// 代发
const (
	CmbPayroll               = "BB6BTHHL" //代发经办 BB6BTHHL
	CmbPayrollQuery          = "BB6BPDQY" //批次与明细查询 BB6BPDQY
	CmbPayrollQueryAll       = "BB6INTQY" //代发代扣综合查询 BB6INTQY
	CmbPayrollQueryBatchList = "BB6BTHQY" //代发批次查询 BB6BTHQY
	CmbPayrollQueryDetail    = "BB6DTLQY" //代发明细查询 BB6DTLQY
	CmbPayrollRefund         = "BB6RFDQY" //代发退票查询 BB6RFDQY
	CmbPayrollQueryType      = "BB6AGTQY" //代发类型查询 BB6AGTQY
	CmbWithholde             = "BB6ACLAK" //代扣经办 BB6ACLAK
	CmbWithholdeQuery        = "BB6CDCQY" //代扣批次查询 BB6CDCQY
	CmbWithholdeQueryDetail  = "BB6ACLIF" //代扣明细查询 BB6ACLIF
)

// 管家 管理
const (
	CmbUnitManageAddAccount                 = "NTDMAADD"    //新增记账子单元 NTDMAADD
	CmbUnitManageCloseAccount               = "NTDUMDLT"    //关闭记账子单元NTDUMDLT
	CmbUnitManageAccountQuery               = "NTDUMQRY"    //记账子单元查询NTDUMQRY
	CmbUnitManageAccountTransQueryByBusNo   = "NTDUMRED"    //按业务参考号查询结果NTDUMRED todo
	CmbUnitManageAccountTransQueryDetail    = "NTDUMINF"    //详情查询NTDUMINF todo
	CmbUnitManageAccountSetWhitePay         = "NTDUMRLA"    //设置记账子单元付方白名单信息NTDUMRLA
	CmbUnitManageAccountDelWhitePay         = "NTDUMRLD"    //删除记账子单元付方白名单信息NTDUMRLD
	CmbUnitManageAccountQueryWhitePay       = "NTDUMRLQ"    //查询记账子单元付方白名单信息NTDUMRLQ
	CmbUnitManageAccountSetCtl              = "NTDUMCTL"    //设置记账子单元入账控制NTDUMCTL
	CmbUnitManageAccountUnSetCtl            = "NTDUMCLE"    //解除记账子单元入账控制NTDUMCLE
	CmbUnitManageAccountTransDaily          = "NTDMTQRD"    //记账子单元当天交易查询NTDMTQRD
	CmbUnitManageAccountTransHistory        = "NTDMTQRY"    //记账子单元历史交易查询NTDMTQRY
	CmbUnitManageAccountTransStatement      = "NTDMTQRYPDF" //记账子单元交易明细对账单获取NTDMTQRYPDF
	CmbUnitManageAccountTransQueryStatement = "DCTASKID"    //记账子单元交易明细对账单处理结果查询请求DCTASKI
	CmbUnitManageAccountPay                 = "NTOPRDMP"    //记账子单元付款经办NTOPRDMP
	CmbUnitManageAccountPayRefund           = "NTOPRDMR"    //记账子单元退款经办NTOPRDMR
	CmbUnitManageAccountPayIn               = "NTDMITRX"    //记账子单元内部转账NTDMITRX

)

// 提醒
const (
	CmbNoticeAccountBalance      = "YQN01010" //账务变动通知
	CmbNoticeSatement            = "YQF01010" //回单结果通知
	CmbNoticePayResult           = "YQN02030" //支付结果通知
	CmbNoticePayrollResult       = "YQN03010" // 代发结果通知
	CmbNoticePayrollDetailResult = "YQF03010" // 代发明细对账单结果通知
	CmbNoticeWithHoldResult      = "YQN03030" // 代扣结果通知
	CmbNoticeSystemMaintenance   = "YQSYSTEM" // 维护通知
)
