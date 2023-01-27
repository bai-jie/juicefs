package odrvcookie

import (
	"net/http"
	"strings"
)

//type SpCookie struct {
//	Cookie string
//	expire time.Time
//}
//
//func (sp SpCookie) IsExpire() bool {
//	return time.Now().After(sp.expire)
//}
//
//var cookiesMap = struct {
//	sync.Mutex
//	m map[string]*SpCookie
//}{m: make(map[string]*SpCookie)}

/*func GetCookie(username, password, siteUrl string) (string, error) {
	//cookiesMap.Lock()
	//defer cookiesMap.Unlock()
	//spCookie, ok := cookiesMap.m[username]
	//if ok {
	//	if !spCookie.IsExpire() {
	//		log.Debugln("sp use old cookie.")
	//		return spCookie.Cookie, nil
	//	}
	//}
	//log.Debugln("fetch new cookie")
	ca := New(username, password, siteUrl)
	tokenConf, err := ca.Cookies()
	if err != nil {
		return "", err
	}
	return cookieToString([]*http.Cookie{&tokenConf.RtFa, &tokenConf.FedAuth}), nil
	//spCookie = &SpCookie{
	//	Cookie: cookie.ToString([]*http.Cookie{&tokenConf.RtFa, &tokenConf.FedAuth}),
	//	expire: time.Now().Add(time.Hour * 12),
	//}
	//cookiesMap.m[username] = spCookie
	//return spCookie.Cookie, nil
}*/

func (ca *CookieAuth) CookieString() (string, error) {
	tokenConf, err := ca.Cookies()
	if err != nil {
		return "", err
	}
	return cookieToString([]*http.Cookie{&tokenConf.RtFa, &tokenConf.FedAuth}), nil
}

func cookieToString(cookies []*http.Cookie) string {
	if cookies == nil {
		return ""
	}
	cookieStrings := make([]string, len(cookies))
	for i, cookie := range cookies {
		cookieStrings[i] = cookie.String()
	}
	return strings.Join(cookieStrings, ";")
}
