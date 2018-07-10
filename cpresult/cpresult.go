package cpresult

import (
	"strings"
)

//-----------------------------------------
// 상수
const (
	CPErr_OK = iota
	CPErr_UnknownFunction
	CPErr_InvalidVersion
	CPErr_ParamMismatch
	CPErr_InvalidParamValue
	CPErr_InvalidCoin
	CPErr_InvalidCoinID
	CPErr_NotFoundCoinID
	CPErr_NotAllowedCoin
	CPErr_ExistCoinAlready
	CPErr_InvalidWallet
	CPErr_InvalidWalletID
	CPErr_NotFoundWalletID
	CPErr_ExistWalletAlready
	CPErr_DisableWallet
	CPErr_InvalidSignature
	CPErr_InvalidTrans
	CPErr_InvalidTransID
	CPErr_NotFoundTransID
	CPErr_InvalidConfigID
	CPErr_NotFoundConfigID
	CPErr_SelfTrans
	CPErr_NotAllow
	CPErr_NotSupported
	CPErr_NotSufficient
	CPErr_MinTransValue
	CPErr_MaxTransValue
	CPErr_PreTransNotFound
	CPErr_OverDayTrans
	CPErr_FailTransReceive
	CPErr_HistoryForKey
	CPErr_AccessDeny
	CPErr_StateByRange
	CPErr_UnderTransTime
	CPErr_FailPutState
	CPErr_FailDelState
	CPErr_FailParse
	CPErr_UnknownCommand
	CPErr_Unknown
	CPErr_TagNotChanged
	CPErr_AccessExpired
	CPErr_InvalidSignaturePublicKey
	CPErr_End
)

// CPResult -- RPC 함수 리턴형
type CPResult struct {
	Code   int    `json:"ec"`
	Pid    string `json:"pid"`
	Result string `json:"value"`
	Ref    string `json:"ref"`
}

var mapErrMessage = map[int]string{
	CPErr_OK:                        "OK",
	CPErr_UnknownFunction:           "Unknown function",
	CPErr_InvalidVersion:            "Version mismatch",
	CPErr_ParamMismatch:             "Parameter mismatch",
	CPErr_InvalidParamValue:         "Invalid parameter value",
	CPErr_InvalidCoin:               "Invalid coin",
	CPErr_InvalidCoinID:             "Invalid coin ID",
	CPErr_NotFoundCoinID:            "Not found coin ID",
	CPErr_NotAllowedCoin:            "Not allowed coin",
	CPErr_ExistCoinAlready:          "Exist coin already",
	CPErr_InvalidWallet:             "Invalid wallet",
	CPErr_InvalidWalletID:           "Invalid wallet ID",
	CPErr_NotFoundWalletID:          "Not found Wallet ID",
	CPErr_ExistWalletAlready:        "Exist wallet already",
	CPErr_DisableWallet:             "Disabled wallet",
	CPErr_InvalidSignature:          "Invalid signature",
	CPErr_InvalidTrans:              "Invalid transaction data",
	CPErr_InvalidTransID:            "Invalid transaction ID",
	CPErr_NotFoundTransID:           "Not found trans ID",
	CPErr_InvalidConfigID:           "Invalid configuration ID",
	CPErr_NotFoundConfigID:          "Configuration not found",
	CPErr_SelfTrans:                 "Not allow the self transaction",
	CPErr_NotAllow:                  "Not allow",
	CPErr_NotSupported:              "Not supported",
	CPErr_NotSufficient:             "Enough not coin",
	CPErr_MinTransValue:             "So minvalue for transaction",
	CPErr_MaxTransValue:             "So maxvalue for transaction",
	CPErr_PreTransNotFound:          "Previous-trans is not found",
	CPErr_OverDayTrans:              "Over day trans limit",
	CPErr_FailTransReceive:          "Fail Transaction for receiver wallet",
	CPErr_HistoryForKey:             "Invalid history iterator for key",
	CPErr_StateByRange:              "Invalid range iterator for key",
	CPErr_AccessDeny:                "Access deny",
	CPErr_UnderTransTime:            "Not past min transastion time",
	CPErr_FailPutState:              "Fail PutState",
	CPErr_FailDelState:              "Fail DelState",
	CPErr_FailParse:                 "Fail Parse",
	CPErr_UnknownCommand:            "Unknown command",
	CPErr_Unknown:                   "Unknown",
	CPErr_TagNotChanged:             "There is no tag to change",
	CPErr_AccessExpired:             "Access expired",
	CPErr_InvalidSignaturePublicKey: "Invalid signature publickey",
	CPErr_End:                       "Last code",
}

func getCPErrorMsg(code int) string {
	if code < CPErr_OK || code >= CPErr_End {
		code = CPErr_Unknown
	}

	return mapErrMessage[code]
}

// FnGetCPResult -- return result
func FnGetCPResult(code int, result string, ref string) *CPResult {
	cpResult := CPResult{}
	cpResult.Code = code
	cpResult.Ref = getCPErrorMsg(code)
	if ref != "" {
		cpResult.Ref += "-" + ref
	}
	result = strings.Trim(result, " \t\r\n")
	if result == "" {
		cpResult.Result = "{}"
	} else {
		cpResult.Result = result
	}
	return &cpResult

}

/*
// FnGetCPError -- error string
func FnGetCPError(en int, etc string) error {
	var eMsg = ""

	if etc != "" {
		eMsg = fmt.Sprintf("[%d]%s, %s", en, getCPErrorMsg(en), etc)
	} else {
		eMsg = fmt.Sprintf("[%d]%s", en, getCPErrorMsg(en))
	}
	return errors.New(eMsg)
}
*/
