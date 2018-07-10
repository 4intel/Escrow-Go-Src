package chsutils

import (
	"fmt"
	"strings"
	"testing"
)

// VTag --
type VTag struct {
}

// IsValid --
// key 무결성 확인
func (vt *VTag) IsValid(tagKey string) bool {

	if len(tagKey) <= 4 {
		return false
	}
	if rune(tagKey[3]) != rune('_') {
		return false
	}

	if !strings.ContainsRune("sr", rune(tagKey[0])) {
		return false
	}

	if !strings.ContainsRune("aox", rune(tagKey[1])) {
		return false
	}
	if !strings.ContainsRune("ox", rune(tagKey[2])) {
	}
	return true
}

// IsRewriteable --
// key의 값이 재기록 가능한지 여부
func (vt *VTag) IsRewriteable(tagKey string) bool {
	if len(tagKey) < 2 {
		return true
	}
	return rune(tagKey[1]) != rune('x')
}

func TestVTag(t *testing.T) {
	sTag := "sao_data:1; sax_data:1; soo_data:1; sox_data:1; sxo_data:1; sxx_data:1; " +
		"rao_data:1; rax_data:1; roo_data:1; rox_data:1; rxo_data:1; rxx_data:1;"

	fmt.Println("[0]", sTag)

	var itag = IMaskTag(&VTag{})
	mapTag := FnGetMapFromTag(sTag, itag)

	fmt.Println("[1]", FnGetTagFromMap(mapTag))
	mapSqX, _ := FnFilterAllowMask(mapTag, "s?x")
	mapRqX, _ := FnFilterDenyMask(mapTag, "s?x")
	fmt.Println("[2]", FnGetTagFromMap(mapSqX))
	fmt.Println("[3]", FnGetTagFromMap(mapRqX))

	sTag2 := "sao_data:1; rao_data:1; rao_data2:3;"
	mapTag2 := FnGetMapFromTag(sTag2, itag)
	FnRemoveDuplicateKey(mapTag, mapTag2)
	fmt.Println("[4]", FnGetTagFromMap(mapTag2))

	sTag3 := "sao_data:2; sax_data:2; soo_data:2; sox_data:2; sxo_data:2; sxx_data:2; " +
		"rao_data:2; rax_data:2; roo_data:2; rox_data:2; rxo_data:2; rxx_data:2;"
	mapTag3 := FnGetMapFromTag(sTag3, itag)
	fmt.Println("[5]", FnGetTagFromMap(mapTag3))
	FnMergeTagMap(mapTag, mapTag3, itag)
	fmt.Println("[6]", FnGetTagFromMap(mapTag))

	mapTag = FnGetMapFromTag(sTag, itag)

}
