package cse

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"

	query "github.com/google/go-querystring/query"
)

// Agent is return
type Agent struct {
	token     string
	tokenUpts int64

	cx   string
	lang string

	client *http.Client
	mux    sync.Mutex
}

// Request is cse api query params
type Request struct {
	Start int64 `json:"start" url:"start"`
	Num   int64 `json:"num" url:"num"`

	Hl     string `json:"hl" url:"hl"`
	Key    string `json:"key" url:"key"`
	Cx     string `json:"cx" url:"cx"`
	CSETok string `json:"cse_tok" url:"cse_tok"`
	Q      string `json:"q" url:"q"`
}

// New return a search instance
func New(cx string, lang string) (agent Agent) {
	agent = Agent{cx: cx, lang: lang}

	// @todo bind to local address
	tr := &http.Transport{
		// Proxy: http.ProxyFromEnvironment,
		// DialContext: (&net.Dialer{
		// 	LocalAddr: &localTCPAddr,
		// 	Timeout:   30 * time.Second,
		// 	KeepAlive: 30 * time.Second,
		// 	DualStack: true,
		// }).DialContext,
	}
	agent.client = &http.Client{Transport: tr}
	return
}

// Query is do search and return ret list
func (a *Agent) Query(q string, page int64, pageSize int64) (ret *SearchRet, err error) {
	if time.Now().Unix()-a.tokenUpts > 300 {
		a.refreshToken()
	}

	request := newRequest(a.cx, a.token, a.lang, q, page, pageSize)
	params, err := query.Values(request)

	resp, err := http.Get("https://www.googleapis.com/customsearch/v1element?" + params.Encode())
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("response status : " + resp.Status + " Url: " + resp.Request.URL.String())
	}

	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return
	}

	return
}

func (a *Agent) refreshToken() {
	a.mux.Lock()
	defer a.mux.Unlock()

	token, err := a.parseToken()

	if err != nil {
		return
	}

	a.token = token
	a.tokenUpts = time.Now().Unix()
}

func (a *Agent) parseToken() (token string, err error) {
	res, err := http.Get("https://cse.google.com/cse.js?cx=" + a.cx)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	html := string(body)
	tokenIndex := strings.Index(html, "cse_token")
	if tokenIndex < 0 {
		err = errors.New("token parse faild")
		return
	}

	return strings.Split(html[tokenIndex+len(`cse_token": "`):], `"`)[0], nil
}

func newRequest(cx string, token string, lang string, q string, page int64, pageSize int64) Request {
	if page < 1 {
		page = 1
	}
	return Request{
		Start:  pageSize * (page - 1),
		Num:    pageSize,
		Hl:     lang,
		Key:    "AIzaSyCVAXiUzRYsML1Pv6RwSG1gunmMikTzQqY",
		Cx:     cx,
		CSETok: token,
		Q:      q,
	}
}
