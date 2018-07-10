// coinpool project coinpool.go
package main

import (
	"encoding/json"
	"fmt"
	"time"

	cu "github.com/coinpool/chsutils"
	cl "github.com/coinpool/coinlib"
	cr "github.com/coinpool/cpresult"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

const (
	CHAIN_Version = "0.9.16"
)

// SimpleAsset implements a simple chaincode to manage an asset
type CoinPool struct {
}

// var cpLogger = logging.MustGetLogger("shim")

//-----------------------------------
// coinpool : 체인코드 이름
// cpchain      : 채널 이름
//-----------------------------------
// 디버깅 환경 서버 실행 예시
// 1. hyperledge 서버(172.30.1.17) ~/work/github/chainpool-debug 에서
//    * 원래 명령 :  docker-compose -f docker-compose-simple.yaml up
//       ./run.sh [start | restart | install ]
//       ./run.sh install    // 기존 최초 설치시 채널만들어줘야 한다.
// 2. vscodeㅣㅣ 에서 이 프로그램 디버깅 환경으로 실행
//    launcher  디버깅 변수가 지정되어 있어야 함
//           "env": {
//	            "CORE_PEER_ADDRESS":"172.30.1.17:7051",
//	            "CORE_CHAINCODE_ID_NAME":"coinpool:1.0"
//----------------------------------------------

// Init is called during chaincode instantiation to initialize any
// data. Note that chaincode upgrade also calls this function to reset
// or to migrate data.
func (t *CoinPool) Init(stub shim.ChaincodeStubInterfaceCp) peer.Response {
	// Get the args from the transaction proposal
	fmt.Println("[]==============================")
	fmt.Println("[] Instantiate coinpool .....")
	fmt.Printf("[] Now : %s\n", time.Now().String())
	fmt.Println("[]==============================")
	//cl.FnDefaultInstance(stub)
	//cl.FnLoadManagerKey()
	//return initCoinPool(stub)
	return shim.Success(nil)
}

// Invoke is called per transaction on the chaincode. Each transaction is
// either a 'get' or a 'set' on the asset created by Init function. The Set
// method may create a new asset by specifying a new key-value pair.
func (t *CoinPool) Invoke(stub shim.ChaincodeStubInterfaceCp) peer.Response {
	var cpRes *cr.CPResult
	var result []byte

	// Extract the function and args from the transaction proposal
	fn, args := stub.GetFunctionAndParameters()
	sArgs, err1 := json.Marshal(args)
	fmt.Printf("Invoke : %s %s\n", fn, string(sArgs))
	if err1 != nil {
		return shim.Error(err1.Error())
	}

	if len(args) < 1 {
		return shim.Error(err1.Error())
	}

	for _, arg := range args {
		if len(arg) > 256 {
			cpRes = cr.FnGetCPResult(cr.CPErr_InvalidParamValue, "Arg length so long", "")
			result, _ = json.Marshal(cpRes)
			return shim.Error(string(result))
		}
	}

	sPid := args[0]
	args = args[1:]
	//==============================================
	// Coin
	// CPKey_RootPublicKey = "owner_pub_key4567890123456789012"
	// CPKey_BaseCoinId    = "coinpool_basecoin"
	switch fn {
	/* -------------------------------------------------- */
	case "version":
		cpRes = cr.FnGetCPResult(cr.CPErr_OK, "{ \"ver\":\""+CHAIN_Version+"\" }", fn)
	/* -------------------------------------------------- */
	case "createCoin":
		// "10000", "coinpool_basecoin", "놀이방 코인", "0x000001", "GPIN", "GPON", "1000", "20", "1.0", "0", "10000", "2017-01-01", "2999-12-31", "1000000000000", "memo", "noncenumber", "signautrekey", "owner_pub_key4567890123456789012"
		fallthrough
	case "stateCoin": // "10000","CoinID"
		fallthrough
	case "setCoinFee": // "10000","CoinID",rate,max,fix,feewid,nonce,sig,man100-pubkey
		fallthrough
	case "setCoinExchRate": // "10000","CoinID",exchRate,nonce,sig,manager(root)-pubkey
		fallthrough
	case "setCoinState": // "10000","CoinID",u32State,nonce,sig,manager(root)-pubkey
		fallthrough
	case "setCoinDate": // "10000","CoinID","start date","end date",nonce,sig,manager(root)-pubkey
		// date format ex "2017-05-12"
		fallthrough

	case "queryCoin": // "10000", "CoinID","CoinID2",nonce,sig,manager-pubkey
		cpRes = cl.FnMainCoin(stub, fn, args)

	/* -------------------------------------------------- */
	case "createWallet":
		//"10000", "coinpool_basecoin", "내지갑", "0x03", "1000000000", "Memo",	"noncenumber", "signautrekey", "owner_pub_key4567890123456789012"
		//"10000", "coinpool_basecoin", "내지갑", "0x03", "1000000000", "Memo",	"noncenumber", "signautrekey", "owner_pub_key4567890123456789012"
		fallthrough
	case "setWalletName": // "10000", "WID" , "지갑이름","지갑 메모"  "man300 or owner pub_key -- 지갑 등록정보 변경
		fallthrough
	case "setWalletDayLimit": // "10000", "WID" , "일일전송한도","지갑 메모"  "man300 or owner pub_key -- 지갑 등록정보 일일전송한도 변경
		fallthrough
	case "stateWallet": // "10000", "WAL1" 지갑 등록정보 확인
		fallthrough
	case "setWalletState": // "10000", "WID" , "상태값", noncenumber", "signautrekey", "man 200_pub_key -- 지갑 상태 정보 변경
		fallthrough
	case "attachWalletTag": // args : "ver","wid","tag",nonce,sig,pubk -- 지갑 보조정보 등록
		fallthrough
	case "stateWalletTag": // args : "ver","wid",,nonce,sig,pubk -- 지갑 보조 정보 조회
		fallthrough
	case "queryWallet": // "10000", "WAL1","WAL2"  , "noncenumber", "signautrekey", "owner_pub_key4567890123456789012"
		cpRes = cl.FnMainWallet(stub, fn, args)

	/* -------------------------------------------------- */
	case "createTrans":
		//  John(owner_xxx), Nora(rec...) Worker(sende...)
		// John(owner_xxx)이 Nora(rec...)이 운영하는 Cafe에서 3000 CashCoin Green-tea 구매함
		// "10000", "WAreceiver_wallet_id90123456789013", "R","3000000", "Cafe  Green-tea", "John", "Nora","realWID","noncenumber", "signautrekey", "owner_pub_key4567890123456789012"
		fallthrough
	case "escrowTrans": // same as createTrans except recive wallet is escrow wallet
		fallthrough
	case "infoTrans":
		//-----------------------------------------------------
		// 거래을  하기위해 확인을 받기위한 정보를 일괄 요청 하는 함수
		// access : man300
		//            ver ,   SenderWID,    ReceiverWID,  금액기준(R|S),   금액 ,
		// args =  [ "10000", "SenderWID",    "ReceiverWID",   "S",     "3000000", "noncenumber", "signautrekey", "man300 or sender PubK" ]
		// return [ ver ,   SenderWID,    ReceiverWID,  금액기준(R|S),   금액 ,
		//           거래요금,보낼금액, 받을금액
		//           보내는 코인ID,보내는 코인NAME,보내는 코인 거래단위,보내는 코인 코인단위, 보내는 코인환율,보내는 지갑 이름
		//           받는 코인ID,받는 코인NAME,받는 코인 거래단위,받는 코인 코인단위, 받는 코인환율, 받는 지갑 이름]
		fallthrough
	case "initBaseTrans":
		//                      ver ,       ReceiverWID,                    초기금액,    거래메시지,    자기지갑표시,
		//"10000", 베이스지갑ID, "90000000000", "초기 코인발행", "COIN_BASE",	"noncenumber", "signautrekey", rootkey}
		fallthrough
	case "stateTrans": // 관리자가 사용자 지갑 잔고 확인 "ver","wid","nonce","sig","man300_pubk"
		fallthrough
	case "walletTrans": // 사용자가 자기 지갑 잔고 확인 "ver","nonce","sig","user_pubk"
		fallthrough
	case "attachTransTag": // 거래관련 평가자료 등록 , args : "ver","txid","estmate","nonce","sig","user_pubk"
		// Estimate version for Escrow-service
		fallthrough
	case "historyTrans": // 관리자가 사용자 지갑 거래내역 확인
		fallthrough
	case "historyTransE": // Estimate version for Escro-service
		// 프로토콜버전, 조회 할 지갑ID,"시작일","종료일","최신순(Y/N)","레코드시작번호","최대레코드갯수","난수문자열","서명","관리자공개키"
		// "10000","WAowner_pub_key4567890123456789012","","","Y","0","200","noncenumber", "signautrekey", "owner_pub_key4567890123456789012"
		fallthrough
	case "historyNTrans": // 관리자가 사용자 지갑 최근 거래내역 확인
		fallthrough
	case "historyNTransE": // Estimate version for Escro-service
		// 프로토콜버전, 조회 할 지갑ID,"최대레코드갯수","난수문자열","서명","관리자공개키"
		// "10000","WAowner_pub_key4567890123456789012","20","noncenumber", "signautrekey", "owner_pub_key4567890123456789012"
		fallthrough
	case "userTrans": // 소유자가 자기 지갑 거래내역 확인
		fallthrough
	case "userTransE": // Estimate version for Escro-service
		// 프로토콜버전,"시작일","종료일","최신순(Y/N)","레코드시작번호","최대레코드갯수","난수문자열","서명","관리자공개키"
		// "10000","","","Y","0","10","noncenumber", "signautrekey", "owner_pub_key4567890123456789012"
		fallthrough
	case "userNTrans": // 소유자가 자기 지갑 최근 거래내역 확인
		fallthrough
	case "userNTransE": // Estimate version for Escro-service
		// 프로토콜버전,"최대레코드갯수","난수문자열","서명","관리자공개키"
		// "10000","3","noncenumber", "signautrekey", "owner_pub_key4567890123456789012"
		cpRes = cl.FnMainTrans(stub, fn, args)
	/* -------------------------------------------------- */
	case "getNonce":
		fallthrough
	case "setManagerKey":
		fallthrough
	case "reloadManagerKeys":
		cpRes = cl.FnMainConfig(stub, fn, args)
	/* -------------------------------------------------- */
	case "onReady":
		if fnOnReady(stub) { // 처음 실행되며 호출됨
			cpRes = cr.FnGetCPResult(cr.CPErr_OK, "", "")
		} else {
			cpRes = cr.FnGetCPResult(cr.CPErr_NotAllow, "", "")
		}
	case "deleteState": // 강제로 임의 키를 지운다. (ROOT 권한  + Reg Nonce)
		cpRes = fnDeleteState(stub, args)
		if cpRes == nil {
			cpRes = cr.FnGetCPResult(cr.CPErr_OK, "", "")
		}
	case "initChain":
		cl.FnDefaultInstance(stub)
		cpRes = cr.FnGetCPResult(cr.CPErr_OK, "", "")
	default:
		cpRes = cr.FnGetCPResult(cr.CPErr_UnknownCommand, "", "")
	}

	cpRes.Pid = sPid
	result, _ = json.Marshal(cpRes)
	if cpRes.Code == 0 {
		// Return the result as success payload
		return shim.Success(result)
	}
	return shim.Error(string(result))
}

var s_bIsOnReady = false

func fnOnReady(stub shim.ChaincodeStubInterfaceCp) bool {
	if s_bIsOnReady {
		fmt.Println("[fail] onReady already ...")
		return false
	}
	s_bIsOnReady = true
	fmt.Println("[] load manager key ...")
	cpErr := cl.FnLoadManagerKey(stub)
	if cpErr != nil {
		return false
	}

	return true
}

func fnDeleteState(stub shim.ChaincodeStubInterfaceCp, args []string) *cr.CPResult {
	var retErr *cr.CPResult
	if len(args) != 5 { //  ["ver","key","nonce","sig","root_pubk"]
		return cr.FnGetCPResult(cr.CPErr_ParamMismatch, "", "")
	}
	stVer := cu.NewStringVersion(args[0])
	if stVer.Major < 1 {
		retErr = cr.FnGetCPResult(cr.CPErr_InvalidParamValue, "", "ver")
		return retErr
	}

	if !cl.FnCheckBaseSignature(args, cl.CPAcc_Root, stVer, true) {
		retErr = cr.FnGetCPResult(cr.CPErr_AccessDeny, "", "")
	} else {
		err := stub.DelState(args[1])
		if err != nil {
			retErr = cr.FnGetCPResult(cr.CPErr_FailDelState, "", err.Error())
		}

	}
	return retErr
}

// main function starts up the chaincode in the container during instantiate
func main() {
	fmt.Println("[]==============================")
	fmt.Println("[] start coinpool ...")

	cl.FnInitCoin()
	cl.FnInitWallet()
	cl.FnInitTrans()
	cl.FnInitConfig()

	go cl.FnStartNonceManager()

	fmt.Println("[]==============================")

	if err := shim.Start(new(CoinPool)); err != nil {
		fmt.Printf("Error starting CoinPool chaincode: %s", err)
	}
	fmt.Println("[]==============================")
	fmt.Println("[] exit coinpool ...")

	cl.FnExitNonceManager()
	fmt.Println("[]==============================")
}
