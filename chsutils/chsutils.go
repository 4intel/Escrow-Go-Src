/*
 * basic utilies...
 */

package chsutils

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	cr "github.com/coinpool/cpresult"
)

// Ver --
type Ver struct {
	Major uint16
	Miner uint16
	Build uint16
}

// SetVersion --
func (v *Ver) SetVersion(ver uint32) {
	v.Major = uint16(ver / 10000)
	v.Miner = uint16((ver % 10000) / 100)
	v.Build = uint16(ver % 100)
}

// String --
func (v *Ver) String() string {
	return fmt.Sprintf("%d.%02d.%02d", v.Major, v.Miner, v.Build)
}

// Uint32 --
func (v *Ver) Uint32() uint32 {
	return uint32(v.Major)*10000 + uint32(v.Miner)*100 + uint32(v.Build)
}

// NewVersion --
func NewVersion(ver uint32) *Ver {
	v := new(Ver)
	v.SetVersion(ver)
	return v
}

// NewStringVersion --
func NewStringVersion(ver string) *Ver {
	unVer, cpErr := FnGetParamUint32(ver, "Ver")
	if cpErr != nil {
		unVer = 0
	}

	return NewVersion(uint32(unVer))
}

// FnCheckParams --
// 함수 파라메터 검증
func FnCheckParams(err *cr.CPResult, arg string) *cr.CPResult {
	if err != nil {
		return cr.FnGetCPResult(cr.CPErr_InvalidParamValue, "", arg)
	}
	return nil
}

// FnGetParamInt16 --
// 문자열을 int16 로 받는다.
func FnGetParamInt16(arg, argname string) (int16, *cr.CPResult) {
	var nBase = 10
	if strings.HasPrefix(arg, "0x") || strings.HasPrefix(arg, "0X") {
		nBase = 16
		arg = arg[2:]
	}
	nV, err := strconv.ParseInt(arg, nBase, 16)
	if err != nil {
		return 0, cr.FnGetCPResult(cr.CPErr_InvalidParamValue, "", argname)
	}
	return int16(nV), nil
}

// FnGetParamUint16 --
// 문자열을 uint16 로 받는다.
func FnGetParamUint16(arg, argname string) (uint16, *cr.CPResult) {
	var nBase = 10
	if strings.HasPrefix(arg, "0x") || strings.HasPrefix(arg, "0X") {
		nBase = 16
		arg = arg[2:]
	}
	nV, err := strconv.ParseUint(arg, nBase, 16)
	if err != nil {
		return 0, cr.FnGetCPResult(cr.CPErr_InvalidParamValue, "", argname)
	}
	return uint16(nV), nil
}

// FnGetParamInt32 --
// 문자열을 int32 로 받는다.
func FnGetParamInt32(arg, argname string) (int32, *cr.CPResult) {
	var nBase = 10
	if strings.HasPrefix(arg, "0x") || strings.HasPrefix(arg, "0X") {
		nBase = 16
		arg = arg[2:]
	}
	nV, err := strconv.ParseInt(arg, nBase, 32)
	if err != nil {
		return 0, cr.FnGetCPResult(cr.CPErr_InvalidParamValue, "", argname)
	}
	return int32(nV), nil
}

// FnGetParamUint32 --
// 문자열을 uint32 로 받는다.
func FnGetParamUint32(arg, argname string) (uint32, *cr.CPResult) {
	var nBase = 10
	if strings.HasPrefix(arg, "0x") || strings.HasPrefix(arg, "0X") {
		nBase = 16
		arg = arg[2:]
	}
	nV, err := strconv.ParseUint(arg, nBase, 32)
	if err != nil {
		return 0, cr.FnGetCPResult(cr.CPErr_InvalidParamValue, "", argname)
	}
	return uint32(nV), nil
}

// FnGetParamInt64 --
// 문자열을 int64 로 받는다.
func FnGetParamInt64(arg, argname string) (int64, *cr.CPResult) {
	var nBase int = 10
	if strings.HasPrefix(arg, "0x") || strings.HasPrefix(arg, "0X") {
		nBase = 16
		arg = arg[2:]
	}
	nV, err := strconv.ParseInt(arg, nBase, 64)
	if err != nil {
		return 0, cr.FnGetCPResult(cr.CPErr_InvalidParamValue, "", argname)
	}
	return int64(nV), nil
}

// FnGetParamUint64 --
// 문자열을 uint64 로 받는다.
func FnGetParamUint64(arg, argname string) (uint64, *cr.CPResult) {
	var nBase int = 10
	if strings.HasPrefix(arg, "0x") || strings.HasPrefix(arg, "0X") {
		nBase = 16
		arg = arg[2:]
	}
	nV, err := strconv.ParseUint(arg, nBase, 64)
	if err != nil {
		return 0, cr.FnGetCPResult(cr.CPErr_InvalidParamValue, "", argname)
	}
	return uint64(nV), nil
}

// FnGetParamFloat32 --
// 문자열을 float32 로 받는다.
func FnGetParamFloat32(arg, argname string) (float32, *cr.CPResult) {
	fV, err := strconv.ParseFloat(arg, 32)
	if err != nil {
		return 0, cr.FnGetCPResult(cr.CPErr_InvalidParamValue, "", argname)
	}
	return float32(fV), nil
}

// FnGetParamFloat64 --
// 문자열을 float64 로 받는다.
func FnGetParamFloat64(arg, argname string) (float64, *cr.CPResult) {
	fV, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		return 0, cr.FnGetCPResult(cr.CPErr_InvalidParamValue, "", argname)
	}
	return float64(fV), nil
}

// FnGetParamTimestamp --
// 문자열을 uint64 로 받는다.
//  dfmt 형식문자열 : Mon Jan 2 15:04:05 -0700 MST 2006 를 표현한 문자열
//                  ex) "2006-01-02"
//	t, err := time.Parse("2006-01-02", "2011-01-19")
func FnGetParamTimestamp(dfmt, arg, argname string) (int64, *cr.CPResult) {
	t, err := time.Parse(dfmt, arg)
	if err != nil {
		return 0, cr.FnGetCPResult(cr.CPErr_InvalidParamValue, "", argname)
	}
	return int64(t.Unix()), nil
}

// FnGetParamTime --
// 문자열을 uint64 로 받는다.
//  dfmt 형식문자열 : Mon Jan 2 15:04:05 -0700 MST 2006 를 표현한 문자열
//                  ex) "2006-01-02"
//	t, err := time.Parse("2006-01-02", "2011-01-19")
func FnGetParamTime(dfmt, arg, argname string) (time.Time, *cr.CPResult) {
	t, err := time.Parse(dfmt, arg)
	if err != nil {
		return time.Now(), cr.FnGetCPResult(cr.CPErr_InvalidParamValue, "", argname)
	}
	return t, nil
}

// FnCheckPrefixID --
// 파라메터로 전달된 string 형 keyID를 실제 사용되는 keyID로(prefix 포함된) 변경한다.
// 일관성 있는 ID를 만들기위해 prefix를 관리한다.
func FnCheckPrefixID(keyID string, prefixID string) string {
	lenPrefix := len(prefixID)
	for strings.HasPrefix(keyID, prefixID) {
		keyID = keyID[lenPrefix:]
	}
	return prefixID + keyID
}

// FnStripPrefixID --
// 파라메터로 전달된 string 형 keyID에서 prefix 를 제거한다.
// prefix 가 이중으로 붙는 것을 막고 일관성 있는 ID를 만들기위해 이용...
func FnStripPrefixID(keyID string, prefixID string) string {
	lenPrefix := len(prefixID)
	for strings.HasPrefix(keyID, prefixID) {
		keyID = keyID[lenPrefix:]
	}
	return keyID
}

/*
 // Timestamp -> Time
 func FnTimestampFromTime(t *time.Time) *timestamp.Timestamp {
	 s := int64(t.Second())     // from 'int'
	 n := int32(t.Nanosecond()) // from 'int'

	 return &timestamp.Timestamp{Seconds: s, Nanos: n}
 }

 // Time -> Timestamp
 func FnTimeFromTimestamp(ts *timestamp.Timestamp) *time.Time {
	 t := time.Unix(ts.Seconds, int64(ts.Nanos))
	 return &t
 }
*/
// FnAppendStr --
// 문자열 처
// 문자열버퍼에 문자열 붙이기
func FnAppendStr(s, apps string) string {
	n := len(s) + len(apps)
	b := make([]byte, n)
	bp := copy(b, s)
	copy(b[bp:], apps)
	return string(b)
}

// FnAppendRune --
// 문자열버퍼에 문자열 붙이기
func FnAppendRune(buf *[]rune, apps string) {
	for _, ch := range apps {
		*buf = append(*buf, rune(ch))
	}
}

// FnGetLineFromStirng --
// 바이트배열 에서 한라인 읽어들이기
// 리턴 int32는 새로운 offset, newOffet == -1 이면 이미 끝에 도착했음을 의미
func FnGetLineFromStirng(data []byte, offset int32) (line string, newOffset int32) {

	l := int32(len(data))
	if offset < 0 || offset >= l {
		return "", int32(-1)
	}

	nLineEndPos := int32(-1)
	for i := offset; i < l; i++ {
		if data[i] == byte('\n') {
			nLineEndPos = i
		}
	}

	newLineStart := int32(0)
	if nLineEndPos < 0 {
		nLineEndPos = l
		newLineStart = nLineEndPos
	} else {
		newLineStart = nLineEndPos + 1
		if nLineEndPos > 0 && data[nLineEndPos-1] == '\r' {
			nLineEndPos--
		}
	}
	return string(bytes.TrimRight(data[offset:nLineEndPos], " \t")), newLineStart
}

// FnMergeString --
// 문자열 합치기, a[s] ... a[e-1]  까지 모두 붙인다.
// 각 문자열 사이에는 sep 문자열을 삽입한다.
func FnMergeString(a []string, s int16, e int16, sep string) string {
	n := int16(len(a))
	if s < 0 || s >= n {
		return ""
	}
	if e < 0 || e > n {
		e = n
	}
	if e <= s {
		return ""
	}

	sz := int16(len(sep)) * (e - s - 1)
	for i := s; i < e; i++ {
		sz += int16(len(a[i]))
	}

	b := make([]byte, sz)
	bp := copy(b, a[s])
	for _, s := range a[s+1 : e] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return string(b)
}

// FnHexStringFromBytes --
// data[s] ~ data[s+len-1] 까지의 16진수 문자열을 얻는다.
// size < 0 이면 끝까지 이다.
func FnHexStringFromBytes(data []byte, offset int, size int) string {
	var e = offset + size

	if size < 0 {
		e = len(data)
	}

	if offset != 0 || e != len(data) {
		data = data[offset : e-1]
	}
	s := hex.EncodeToString(data)
	return s
	/*
		var result = make([]rune, 0)
		var strCh = ""
		for i := offset; i < e; i++ {
			strCh = fmt.Sprintf("%02X ", data[i])
			FnAppendRune(&result, strCh)
		}

		return string(result)
	*/
}

// FnBytesFromHexString --
func FnBytesFromHexString(sHex string) []byte {

	b, err := hex.DecodeString(sHex)
	if err != nil {
		return nil
	}
	return b
}

// FnWriteStringToBuffer --
// 문자열을 버퍼로 저장한다. 길이 + 데이터 형식
// FnReadStringFromBuffer () 와 Pair 함수이다.
func FnWriteStringToBuffer(buf *bytes.Buffer, s string) {
	b := []byte(s)
	binary.Write(buf, binary.LittleEndian, int16(len(b)))
	buf.Write(b)
}

// FnReadStringFromBuffer --
// 문자열을 버퍼에서 읽는다.  길이 + 데이터 형식
// FnWriteStringToBuffer () 와 Pair 함수이다.
func FnReadStringFromBuffer(buf *bytes.Buffer) (string, error) {
	var c int16 = 0
	var n int = 0
	var err error = nil
	err = binary.Read(buf, binary.LittleEndian, &c)
	if err != nil {
		return "", err
	}
	b := make([]byte, c)
	n, err = buf.Read(b)
	if int16(n) != c {
		fmt.Println("read string invalid length")
		return string(b), err
	}
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// FnRoundF64 --
func FnRoundF64(f64 float64) float64 {
	if f64 < 0 {
		return math.Ceil(f64 - 0.5)
	}
	return math.Floor(f64 + 0.5)
}

// FnBigFromString --
// 10진수로 기록된 문자열을 big.Int 수로 읽는다.
func FnBigFromString(s string) *big.Int {
	ret := new(big.Int)
	ret.SetString(s, 10)
	return ret
}

// FnBigFromHexString --
// 16진수로 기록된 문자열을  big.Int 수로 읽는다.
func FnBigFromHexString(s string) *big.Int {
	ret := new(big.Int)
	ret.SetString(s, 16)
	return ret
}

//------------------------------------------
// time function

// Time_GetUtcUnixTimeStamp function
func Time_GetUtcUnixTimeStamp() int64 {
	loc, _ := time.LoadLocation("UTC")
	utcNow := time.Now().In(loc)
	//return utcNow.Unix()
	return utcNow.Unix()
}

// Time_GetUtcUnixNanoTimeStamp function
func Time_GetUtcUnixNanoTimeStamp() int64 {
	loc, _ := time.LoadLocation("UTC")
	utcNow := time.Now().In(loc)
	return utcNow.UnixNano()
}

// Time_GetUtcUnixMiliTimeStamp function
func Time_GetUtcUnixMiliTimeStamp() int64 {
	loc, _ := time.LoadLocation("UTC")
	utcNow := time.Now().In(loc)
	return utcNow.UnixNano() / int64(time.Millisecond)
}

// Time_GetUtcTime function
func Time_GetUtcTime() time.Time {
	loc, _ := time.LoadLocation("UTC")
	return time.Now().In(loc)
}
