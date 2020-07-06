package toolkit

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

func Sign(params map[string]string, appSecret string) string {
	var b strings.Builder

	// params key list
	keyList := make([]string, 0, len(params))
	for k := range params {
		keyList = append(keyList, k)
	}
	// sort
	sort.Strings(keyList)

	for _, key := range keyList {
		if val, ok := params[key]; ok {
			b.WriteString(fmt.Sprintf("%s%s", key, val))
		}
	}

	text := fmt.Sprintf("%s%s%s", appSecret, b.String(), appSecret)

	hash := md5.New()
	hash.Write([]byte(text))
	return strings.ToUpper(hex.EncodeToString(hash.Sum(nil)))
}
