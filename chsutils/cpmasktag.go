package chsutils

import (
	"errors"
	"fmt"
	"strings"
)

// IMaskTag --
type IMaskTag interface {
	IsValid(tagKey string) bool       // key 무결성 확인
	IsRewriteable(tagKey string) bool // key의 값이 재기록 가능한지 여부
}

// FnGetMapFromTag --
// key1=val1; key2=var2; .... 형식의 문자열을 map[string]string 로 변환화여 리턴한다.
func FnGetMapFromTag(tag string, imt IMaskTag) map[string]string {
	var key, val string
	var delimiterPos int

	mapTag := make(map[string]string)

	aryItem := strings.Split(tag, ";")
	for _, item := range aryItem {
		delimiterPos = strings.Index(item, ":")
		if delimiterPos < 0 {
			continue // no key or non delimiter
		}
		key = item[:delimiterPos]
		val = item[delimiterPos+1:]
		key = strings.TrimSpace(key)
		val = strings.TrimSpace(val)
		if !imt.IsValid(key) {
			continue // key is invalid
		}
		mapTag[key] = val
	}
	return mapTag
}

// FnMergeTagMap --
// 0. mapAdd의 소유자 검증과 , key prefix의 합법성 검증은 이미 되었다고 가정함
//    - FnFilterDenyMask(), isValidXXXTagKey()
// 1. mapBase 과 mapAdd 합친다. mapBase에도 있고 mapAdd에도 것은 mapAdd것을 취한다.
// 2. 단, 수정불가 태그는 mapBase에 없는 경우만 합쳐진다
//  return 합쳐진 mapBase 개체를 리턴한다( mapBase 는 변경된다.)
//         bool 리턴은 리턴되는 mapBase에 수정된 사항이 있는 지(true) 없는 지(false)를 나타낸다.
func FnMergeTagMap(mapBase, mapAdd map[string]string, imt IMaskTag) (map[string]string, bool) {
	mapMerge := make(map[string]string)
	for key, valAdd := range mapAdd {
		if valBase, bExist := mapBase[key]; bExist {
			if imt.IsRewriteable(key) {
				// remove key that has no value
				if len(valAdd) != 0 && valBase != valAdd {
					mapMerge[key] = valAdd // valAdd 가 계산식일 수도 있겠다. 하지만 지금은 단순 텍스트로
				}
			}
		} else {
			mapMerge[key] = valAdd
		}
	}

	if len(mapMerge) == 0 {
		return mapMerge, false
	}
	appendMap(mapBase, mapMerge)
	return mapBase, true
}

// FnGetTagFromMap --
// map[string]string 를 key1=val1; key2=var2; .... 형식의 문자열로 리턴한다.
func FnGetTagFromMap(mapTag map[string]string) string {

	var tag = ""
	for key, val := range mapTag {
		tag += fmt.Sprintf("%s:%s; ", key, val)
	}
	return tag
}

// FnRemoveDuplicateKey --
// mapCheck의 키중에서 mapBase 존재하는 것은 mapCheck에서 제거된다.
// 즉, mapCheck에는 mapBase에 없는 것만 남게 된다.
// return : 중복 제거된 mapCheck를 리턴한다.
func FnRemoveDuplicateKey(mapBase, mapCheck map[string]string) map[string]string {
	var bExist bool

	var sDeleteKeys []string
	nDelSum := int(0)
	for key := range mapCheck {
		if _, bExist = mapBase[key]; bExist {
			sDeleteKeys = append(sDeleteKeys, key)
			nDelSum++
		}
	}
	for _, key := range sDeleteKeys {
		delete(mapCheck, key)
	}
	return mapCheck
}

// FnFilterAllowMask --
// sMask 해당하는 태그만 모은 새로운 map를 리턴한다.
// 마스크에서 ?  고려치 않음, 각 항목은 and 조건임 , '?'는 true로 판단된다, 즉 모아진다.
func FnFilterAllowMask(mapCheck map[string]string, sMask string) (map[string]string, error) {

	mapFiltered := make(map[string]string)

	nMskLen := len(sMask)
	if nMskLen < 1 {
		return nil, errors.New("invalid tagmask")
	}

	var chMask rune
	var bAllow bool
	for key, val := range mapCheck {
		if len(key) < nMskLen+2 {
			continue
		}
		if rune(key[nMskLen]) != rune('_') {
			continue
		} // format mismatch

		bAllow = true
		for i := 0; i < nMskLen; i++ {
			chMask = rune(sMask[i])
			if chMask != rune('?') {
				if rune(key[i]) != chMask {
					bAllow = false
					break
				}
			}
		}

		if bAllow { // 필터 조건 만족, 포함 한다.
			mapFiltered[key] = val
		}
	}
	return mapFiltered, nil
}

// FnFilterDenyMask --
// sMask 해당하는 태그를 제거한 새로운 map를 리턴한다.
// 마스크에서 ?  고려치 않음, 각 항목은 and 조건임 , '?'는 true로 판단된다, 즉 제거한다.
func FnFilterDenyMask(mapCheck map[string]string, sMask string) (map[string]string, error) {

	mapFiltered := make(map[string]string)

	nMskLen := len(sMask)
	if nMskLen < 1 {
		return nil, errors.New("invalid tagmask")
	}

	var chMask rune
	var bAllow bool
	for key, val := range mapCheck {
		if len(key) < nMskLen+2 {
			continue
		}
		if rune(key[nMskLen]) != rune('_') {
			continue
		} // format mismatch

		bAllow = true
		for i := 0; i < nMskLen; i++ {
			chMask = rune(sMask[i])
			if chMask != rune('?') {
				if rune(key[i]) != chMask {
					bAllow = false
					break
				}
			}
		}

		if bAllow { // 필터 조건 만족, 포함하지 않는다..
			continue
		}
		mapFiltered[key] = val
	}
	return mapFiltered, nil

}

// mapBase에 mapAdd를 추가 한다.
// return : 합쳐진  mapBase를 리턴한다.
func appendMap(mapBase, mapAdd map[string]string) map[string]string {
	for key, val := range mapAdd {
		mapBase[key] = val
	}
	return mapBase
}
