# implements the coin library...

type CPCoin struct {
	Ver         uint32  `json:"ver"` // 10진수 표기법으로 (v/10000) . (v % 10000)/100 . (v % 100)
	CoinID      string  `json:"cid"` // COIN ID  under 16 char (ansi)
	Name        string  `json:"nm"`  // under 32 char(utf8) , 코인 이름
	State       uint32  `json:"st"`  // 코인 상태 bit mask
	CoinUnit    string  `json:"cou"` // 코인단위 under 8 char(utf8) , 표현단위
	CashUnit    string  `json:"cau"` // 사용코인단위 under 8 char(utf8) , 표현단위
	CoinPerCash uint32  `json:"cpc"` // 코인단위 비
	ExchRate    float32 `json:"exr"` // POOL 코인과의 환율 (usercoin/basecoin)

	MinTransTime int16  `json:"mit"` //		 최소 재거래 시간 초
	MaxVPerDay   uint32 `json:"mvd"` // 회당 최대 거래 한도 , MaxVPerDay
	MinVPerTX    uint32 `json:"mvt"` // 회당 최소 거래 한도 , MinVPerTX

	DateBegine int64 `json:"dtb"` // 거래 유효기간 시작 ( time.Time.sec )
	DateEnd    int64 `json:"dte"` // 거래 유효기간 종료 ( time.Time.sec )

	Amount uint64 `json:"amt"` // 발행 총코인
	Memo   string `json:"mno"` // 메모
}
