예약 키
-------------------------------------------------------------------------
* cpchannel : 체인코드 이름
* coinpool : 채널 이름
* 아래는 키는 설명을 위한 예시 키이고 , 실제는 별도록 관리된다.
* "man_root_key34567890123456789012"  // MAN_ROOT   , root key
* "man_100_pub_key7890123456789100"  // MAN_100 , 관리자 키
* "man_200_pub_key7890123456789200"  // MAN_200
* "man_300_pub_key7890123456789300"  // MAN_300
키 권한 레벨
-------------------------------------------------------------------------
* API에는 권한 레벨속성이 필요하다.  상위 레벨은 하위 레벨의 권한을 행사 할수 있다.
* MAN_ROOT   : 수퍼유저로 모든 모든 권한을 갖고 있음(단 MAN_OWNER 속성 API는 제한됨)
* MAN_100    : 높은 보안이 요구되는 API
* MAN_200    : 보통 보안이 요구되는 API
* MAN_300    : 낮은 보안이 요구되는 API
--------------------------------------------------------------------------
* MAN_OWNER  : 소유권을 갖은 자만 허락 ( (ex)사용자 지갑 관련  명령 경우 ) 
* MAN_FREE   : 권한을 요구하지 않는 API 

체인코드 Install
-------------------------------------------------------------------------
<code>
* docker exec -it cli bash
* peer chaincode install -p chaincodedev/chaincode/sacc -n coinpool -v 0
* peer chaincode instantiate -n coinpool -v 0 -c '{"Args":[]}' -C cpchannel
</code>

간편 디버깅용 시스템 지갑 생성
-------------------------------------------------------------------------
* 코인(COIN_00000001)을 만들고, 지갑을 만들고, 지갑에 초기 코인(COIN_00000001)을 넣어준다. WAman_root_key34567890123456789012 지갑이 만들어짐 , 개발 이후는 없어 질 것임
<code>
* peer chaincode invoke -n coinpool -c '{"Args":["debugSet","pid00","10000"]}' -C cpchannel
</code>

코인 만들기
-------------------------------------------------------------------------
* API 권한 : MAN_ROOT
* 코인 단위      : 코인 기본 단위
* 거래 단위      : 거래에 사용되는 단위
* 단위비         : 코인단위/거래단위 비
* 최소 재거래시간 : 지갑 연속 보내기 제한 시간
* 환율           : usercoin/기준코인
* 코인ID        :  1 이상인 숫자로, 10진 8자리 이하로 표기 가능한 숫자를 이용한다. 내부적으로 "COIN_" 선행문자열이 사용되고 숫자부분은 8자리로 전환되어 이용된다. 
*                 따라서, "1","COIN_1","COIN_0001" , "001","0001" 등은 모두 "COIN_00000001" 과 같은 것으로 처리된다. 
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["createCoin","pid00","10000", "코인ID", "코인이름", "코인속성",  "코인단위","거래단위", "코인단위/거래단위비",	"최소재거래시간(초)", "환율",최대거래한도,최소거래한도,거래유효기간시작,거래유효기간종료,발행총코인,메모, noncenumber", "signautrekey", "man_root_key34567890123456789012"]}' -C cpchannel
* peer chaincode invoke -n coinpool -c '{"Args":["createCoin","pid00","10000", "COIN_00000001","homecoin", "0x000002","HOM", "HAM", "1000", "20", "1.0","0","10000","2017-01-01","2999-12-31","1000000000000","memo" ,"noncenumber", "signautrekey", "man_root_key34567890123456789012"]}' -C cpchannel
</code></pre-wrap>

코인 환율 설정 
-------------------------------------------------------------------------
* API 권한 : MAN_ROOT
* 코인의 환율 을 지정한다. ( COIN/기준코인) 비이다. 
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["setExchRate","pid00","10000","코인ID", 환율, nonce, sig, manager-pubkey]}' -C cpchannel
* peer chaincode invoke -n coinpool -c '{"Args":["setCosetExchRateinFee","pid00","10000","COIN_00000001","1.00","noce", "sig", "man_root_key34567890123456789012"] }' -C cpchannel 
</code></pre-wrap>

코인 거래 요금 정책 설정 
-------------------------------------------------------------------------
* API 권한 : MAN_100
* 요금율 방식 : 요금율, 최대요금은 쌍으로 같이 설정되며, 고정요금과는 배타적이다.
* 고정요금 방식 : 고정요금 설정(코인단위), 요금율,최대요금 = 0으로 설정 
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["setCoinFee","pid00","10000","코인ID", 요금율,최대요금, 고정요금,요금수급지갑, nonce, sig, manager-pubkey]}' -C cpchannel
* peer chaincode invoke -n coinpool -c '{"Args":["setCoinFee","pid00","10000","COIN_00000001","0","0","2000", "WA_man_root_key34567890123456789012", "noce", "sig", "man_root_key34567890123456789012"] }' -C cpchannel 
</code></pre-wrap>

COINBASE 지갑 초기 코인 설정
-------------------------------------------------------------------------
* API 권한 : MAN_ROOT
* 지갑ID는 "WA_" + publickey로 만들어진다, 지갑ID부분에 publickey만 사용해도 내부처리에서 "WA_"를 붙여서 사용된다.
<code><pre-wrap>
* peer chaincode invoke -n coinpool -c '{"Args":["initBaseTrans","pid00","10000", "베이스지갑ID", "금액", "메모","통장메모","noncenumber", "signautrekey", "root public key"] }' -C cpchannel
* peer chaincode invoke -n coinpool -c '{"Args":["initBaseTrans","pid00","10000", "WA_owner_pub_key4567890123456789013", "9000000000000",  "Initial coin", "COIN_BASE","noncenumber", "signautrekey","man_root_key34567890123456789012"] }' -C cpchannel
</pre-wrap></code>

코인의 상태를 본다(등록 내용, 초기 발행 금액 등)
-------------------------------------------------------------------------
* API 권한 : MAN_FREE
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":[명령어, 쿼리ID, 프로토콜버전, 코인ID]}' -C cpchannel
* peer chaincode invoke -n coinpool -c '{"Args":["stateCoin","pid00","10000","COIN_1"]}' -C cpchannel
</code></pre-wrap>

등록된 코인 리스트를 확인한다.
-------------------------------------------------------------------------
* API 권한 : MAN_100
* 시작코인ID : "" = "00000000"을 의미
* 끝코인ID  : "" = "99999999"을 의미  
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":[명령어, 쿼리ID, 프로토콜버전,시작코인ID,끝코인ID,난수문자열, 서명, 관리자공개키]}' -C cpchannel
* peer chaincode invoke -n coinpool -c '{"Args":["queryCoin","pid00","10000","","","noncenumber", "signautrekey", "man_root_key34567890123456789012"]}' -C cpchannel
</code></pre-wrap>

지갑 잔고확인(관리자용) 
-------------------------------------------------------------------------
* API 권한 : MAN_300
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":[명령어, 쿼리ID, 프로토콜버전, 지갑ID,난수문자열, 서명, 지갑공개키]}' -C cpchannel
* peer chaincode invoke -n coinpool -c '{"Args":["stateTrans","pid00","10000", "WA_man_root_key34567890123456789012","noncenumber", "signautre", "man_300_key"]}' -C cpchannel
* 아래 지갑 만들기 수행후 쿼리
* peer chaincode invoke -n coinpool -c '{"Args":["stateTrans","pid00","10000", "users_pub_key4567890123456789020","noncenumber", "signautre", "owner_pub_key4567890123456789020"]]}' -C cpchannel
</code></pre-wrap>

사용자 자기 지갑 잔고확인 
-------------------------------------------------------------------------
* API 권한 : MAN_OWNER
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":[명령어, 쿼리ID, 프로토콜버전, 난수문자열, 서명, 지갑공개키 ]}' -C cpchannel
* peer chaincode invoke -n coinpool -c '{"Args":["walletTrans","pid00","10000", "noncenumber", "signautre","man_root_key"]}' -C cpchannel
* 아래 지갑 만들기 수행후 쿼리
* peer chaincode invoke -n coinpool -c '{"Args":["walletTrans","pid00","10000", "noncenumber", "signautre",man200_key"]}' -C cpchannel
</code></pre-wrap>

지갑 만들기
-------------------------------------------------------------------------
* API 권한 : MAN_200 or MAN_ROOT(base_wallet)
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":[명령어, 쿼리ID, 프로토콜버전, 지갑ID,코인ID, 지갑이름, 지갑상태(3:사용가능), 일거래한도액, 지갑설명, 난수문자열, 서명, 지갑공개키]}' -C cpchannel
* peer chaincode invoke -n coinpool -c '{"Args":["createWallet","pid00","10000","WA_users_pub_key4567890123456789022","COIN_00000001", "Madam Wang", "0x03", "1000000000", "Memo",	"noncenumber", "signautre", "users_pub_key4567890123456789020"]}' -C cpchannel
* peer chaincode invoke -n coinpool -c '{"Args":["createWallet","pid00","10000","WA_users_pub_key4567890123456789020","COIN_00000001", "Wang's main husband", "0x03", "1000000000", "Memo",	"noncenumber", "signautre", "users_pub_key4567890123456789022"]}' -C cpchannel
</code></pre-wrap>

전체 지갑 목록 조회(관리지용)
-------------------------------------------------------------------------
* API 권한 : MAN_300
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":[명령어,  쿼리ID, 프로토콜버전,시작지갑코드ID,마지막지갑코드, 난수문자열, 서명, 지갑공개키]}' -C cpchannel
* peer chaincode invoke -n coinpool -c '{"Args":["queryWallet","pid00","10000","","","noncenumber", "signautrekey", "users_pub_key4567890123456789020"]}' -C cpchannel
</code></pre-wrap>

지갑 등록 정보 조회(잔고확인이 아니다) 
-------------------------------------------------------------------------
* API 권한 : MAN_FREE
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":[명령어, 쿼리ID, 프로토콜버전, 지갑ID]}' -C cpchannel
* peer chaincode invoke -n coinpool -c '{"Args":["stateWallet","pid00","10000","WA_man_root_key34567890123456789012"]}' -C cpchannel
* peer chaincode invoke -n coinpool -c '{"Args":["stateWallet","pid00","10000","WA_users_pub_key4567890123456789020"]}' -C cpchannel
</code></pre-wrap>

거래 등록(송금)
-------------------------------------------------------------------------
* API 권한 : MAN_OWNER
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":[명령어, 쿼리ID, 프로토콜버전, 받는사람 지갑ID, 기준코인종류(S/R, S:보내는 코인 기준), 보내는 금액, 거래내용, 보내는지갑메모, 받는지갑메모, 랜덤문자열, 서명값, 보내는지갑공개키]}' -C cpchannel
* peer chaincode invoke -n coinpool -c '{"Args":["createTrans","pid00","10000", "WA_users_pub_key4567890123456789020", "S","3000000", "Rose tea-shop Miss Lee", "Her husband", "Madam Wang","noncenumber", "signautrekey", "man_root_key34567890123456789012"]}' -C cpchannel
</code></pre-wrap>

지갑 거래 내역 조회
-------------------------------------------------------------------------
* API 권한 : MAN_300
* 일자      : "YYYY-mm-dd" , ""는 오늘 
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":[historyTrans,         쿼리ID, 프로토콜버전,조회할 지갑ID,"시작일","종료일","최신순(Y/N)","레코드시작번호","최대레코드갯수","난수문자열","서명","관리자공개키"]}' -C cpchannel
* peer chaincode invoke -n coinpool -c '{"Args":["historyTrans","pid00","10000","WA_man_root_key34567890123456789012","","","Y","0","200","noncenumber", "signautrekey", "man_root_key34567890123456789012"]}' -C cpchannel
</code></pre-wrap>

자기 지갑 거래 내역 조회
-------------------------------------------------------------------------
* API 권한 : MAN_OWNER
* 일반권한 : 6개월 이하조회 , 최대레코드수 100 이하조회)
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":[userTrans,         쿼리ID, 프로토콜버전,"시작일","종료일","최신순(Y/N)","레코드시작번호","최대레코드갯수","난수문자열","서명","관리자공개키"]}' -C cpchannel
* peer chaincode invoke -n coinpool -c '{"Args":["userTrans","pid00","10000","","","Y","0","200","noncenumber", "signautrekey", "man_root_key34567890123456789012"]}' -C cpchannel
</code></pre-wrap>


테스트 쿼리 시나리오 (DB 초기화 상태에서 시작)
=========================================================================

COINBASE 코인 만들기
-------------------------------------------------------------------------
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["createCoin","pid00","10000", "1","homecoin", "0x000003","JUN", "WON", "1000", "20", "1.0","0","10000","2017-01-01","2999-12-31","1000000000000","memo" ,"noncenumber", "signautrekey", "man_root_key34567890123456789012"]}' -C cpchannel
</code></pre-wrap>

COINBASE BASE 지갑 만들기
-------------------------------------------------------------------------
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["createWallet","pid00","10000","WA_man_root_key34567890123456789012","1", "BaseWallet", "0xA0003", "1000000000", "Memo",	"noncenumber", "signautre", "man_root_key34567890123456789012"]}' -C cpchannel
</code></pre-wrap>

코인 거래 요금 설정 
-------------------------------------------------------------------------
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["setCoinFee","pid00","10000","COIN_00000001","0","0","2000",  "noce", "sig", "man_root_key34567890123456789012"] }' -C cpchannel 
</code></pre-wrap>

코인 정보 조회
-------------------------------------------------------------------------
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["stateCoin","pid00","10000","COIN_00000001"]}' -C cpchannel
</code></pre-wrap>

등록된 코인 리스트를 확인한다.
-------------------------------------------------------------------------
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["queryCoin","pid00","10000","","","noncenumber", "signautrekey", "man_root_key34567890123456789012"]}' -C cpchannel
</code></pre-wrap>

COINBASE BASE 지갑 초기 코인 설정
-------------------------------------------------------------------------
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["initBaseTrans","pid00","10000", "WA_man_root_key34567890123456789012", "9000000000000", "Initial issue", "COIN_BASE","noncenumber", "signautrekey", "man_root_key34567890123456789012"] }' -C cpchannel
</code></pre-wrap>


지갑 잔고확인 (관리자용)
-------------------------------------------------------------------------
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["stateTrans","pid00","10000", "WA_man_root_key34567890123456789012","noncenumber", "signautrekey", "man_root_key34567890123456789012"]}' -C cpchannel
</code></pre-wrap>

왕마담,기둥서방 지갑 만듦 
-------------------------------------------------------------------------
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["createWallet","pid00","10000","WA_users_pub_key4567890123456789020","COIN_00000001", "Madam Wang", "0x03", "1000000000", "Memo",	"noncenumber", "signautre", "man_root_key34567890123456789012"]}' -C cpchannel
* peer chaincode invoke -n coinpool -c '{"Args":["createWallet","pid00","10000","WA_users_pub_key4567890123456789022","COIN_00000001", "Main guy of Wang", "0x03", "1000000000", "Memo","noncenumber", "signautre", "man_root_key34567890123456789012"]}' -C cpchannel
</code></pre-wrap>

전체 지갑 리스트 조회
-------------------------------------------------------------------------
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["queryWallet","pid00","10000","","","noncenumber", "signautrekey", "man_200_pub_key7890123456789200"]}' -C cpchannel
</code></pre-wrap>

베이스지갑 에서 왕마담에게 코인 보냄 
-------------------------------------------------------------------------
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["createTrans","pid00","10000", "WA_users_pub_key4567890123456789020", "S","20000000", "Rose coffee shop Miss Lee", "Base Wallet", "Madam Wang","noncenumber", "signautrekey", "man_root_key34567890123456789012"]}' -C cpchannel
</code></pre-wrap>

왕마담 거래 내역 
-------------------------------------------------------------------------
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["historyTrans","pid00","10000","WA_users_pub_key4567890123456789020","","","Y","0","30","noncenumber", "signautrekey", "man_root_key34567890123456789012"]}' -C cpchannel
</code></pre-wrap>

왕마담 -> 기둥서방 코인 보냄 
-------------------------------------------------------------------------
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["createTrans","pid00","10000", "WA_users_pub_key4567890123456789022", "S","1500000", "Pin money", "Madam Wang","Main husband","noncenumber", "signautrekey", "users_pub_key4567890123456789020"]}' -C cpchannel
</code></pre-wrap>

왕마담 지갑 잔고확인 
-------------------------------------------------------------------------
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["stateTrans","pid00","10000", "WA_users_pub_key4567890123456789020","noncenumber", "signautrekey", "man_root_key34567890123456789012"]}' -C cpchannel
</code></pre-wrap>

기둥서방 지갑 잔고확인 
-------------------------------------------------------------------------
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["stateTrans","pid00","10000", "WA_users_pub_key4567890123456789022","noncenumber", "signautrekey", "man_root_key34567890123456789012"]}' -C cpchannel
</code></pre-wrap>

base 지갑 거래내역 확인 , 요금 받은것도 이어야...
-------------------------------------------------------------------------
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["historyTrans","pid00","10000","WA_man_root_key34567890123456789012","","","Y","0","30","noncenumber", "signautrekey", "man_root_key34567890123456789012"]}' -C cpchannel
</code></pre-wrap>

COINBASE2 ,코인 만들기 
-------------------------------------------------------------------------
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["createCoin","pid00","10000", "COIN_00000002","homecoin2", "0x000003","CEN", "CRA", "1000", "20", "1.0","0","10000","2017-01-01","2999-12-31","1000000000000","memo" ,"noncenumber", "signautrekey", "man_root_key34567890123456789012"]}' -C cpchannel
</code></pre-wrap>

COINBASE2 BASE지갑 만들기
-------------------------------------------------------------------------
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["createWallet","pid00","10000","WA_owner_pub_key4567890123456789020","COIN_00000002", "BaseWallet", "0xA0003", "1000000000", "Memo",	"noncenumber", "signautre", "owner_pub_key4567890123456789020"]}' -C cpchannel
</code></pre-wrap>

COINBASE2 코인 거래 요금 설정 
-------------------------------------------------------------------------
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["setCoinFee","pid00","10000","COIN_00000002","0.03","200000","0", "WA_owner_pub_key4567890123456789020", "noce", "sig", "man_root_key34567890123456789012"] }' -C cpchannel 
</code></pre-wrap>

COINBASE2 코인 정보 조회 
-------------------------------------------------------------------------
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["stateCoin","pid00","10000","COIN_00000002"]}' -C cpchannel
</code></pre-wrap>

COINBASE2 등록된 코인 리스트를 확인한다.
-------------------------------------------------------------------------
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["queryCoin","pid00","10000","","","noncenumber", "signautrekey", "man_root_key34567890123456789012"]}' -C cpchannel
</code></pre-wrap>

COINBASE2 BASE 지갑 초기 코인 설정
-------------------------------------------------------------------------
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["initBaseTrans","pid00","10000", "WA_owner_pub_key4567890123456789020", "9000000000000","Initial issue", "COIN_BASE","noncenumber", "signautrekey","man_root_key34567890123456789012"] }' -C cpchannel
</code></pre-wrap>

COINBASE2 BASE 지갑 잔고확인 
-------------------------------------------------------------------------
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["stateTrans","pid00","10000", "WA_owner_pub_key4567890123456789020","noncenumber", "signautrekey", "man_root_key34567890123456789012"]}' -C cpchannel
</code></pre-wrap>

COINBASE2 사용자 지갑 만들기
-------------------------------------------------------------------------
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["createWallet","pid00","10000","users_pub_key4567890123456789030","COIN_00000002", "coin2 manager30", "0x03", "1000000000", "manager30",	"noncenumber", "signautre", "man_root_key34567890123456789012"]}' -C cpchannel
* peer chaincode invoke -n coinpool -c '{"Args":["createWallet","pid00","10000","users_pub_key4567890123456789031","COIN_00000002", "coin2 manager31", "0x03", "1000000000", "manager31",	"noncenumber", "signautre", "man_root_key34567890123456789012"]}' -C cpchannel
</code></pre-wrap>

COINBASE2 base지갑 에서 manager30에게 코인 보냄 
-------------------------------------------------------------------------
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["createTrans","pid00","10000", "WA_users_pub_key4567890123456789030", "S","8000000", "Blue moon", "Base Wallet", "Madam Wang","noncenumber", "signautrekey", "owner_pub_key4567890123456789020"]}' -C cpchannel
</code></pre-wrap>

잔고확인 
-------------------------------------------------------------------------
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["stateTrans","pid00","10000", "WA_owner_pub_key4567890123456789020","noncenumber", "signautrekey", "man_root_key34567890123456789012"]}' -C cpchannel
* peer chaincode invoke -n coinpool -c '{"Args":["stateTrans","pid00","10000", "WA_users_pub_key4567890123456789030","noncenumber", "signautrekey", "man_root_key34567890123456789012"]}' -C cpchannel
</code></pre-wrap>

COINBASE2 manager30지갑 에서 COINBASE 왕마담 에게 코인 보냄 
-------------------------------------------------------------------------
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["createTrans","pid00","10000", "WA_users_pub_key4567890123456789020", "S","1600000", "Blue moon", "Base Wallet", "Madam Wang","noncenumber", "signautrekey", "users_pub_key4567890123456789030"]}' -C cpchannel
</code></pre-wrap>

왕마담 잔고
-------------------------------------------------------------------------
<pre-wrap><code>
* peer chaincode invoke -n coinpool -c '{"Args":["stateTrans","pid00","10000", "WA_users_pub_key4567890123456789020","noncenumber", "signautrekey", "man_root_key34567890123456789012"]}' -C cpchannel
* COINBASE2 manager30지갑 잔고 확인 
* peer chaincode invoke -n coinpool -c '{"Args":["stateTrans","pid00","10000", "WA_users_pub_key4567890123456789030","noncenumber", "signautrekey", "man_root_key34567890123456789012"]}' -C cpchannel
</code></pre-wrap>

