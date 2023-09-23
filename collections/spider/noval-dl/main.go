// 小说下载器后端版,内网均可访问,附win成品及go源码
// https://www.52pojie.cn/thread-1836863-1-1.html
// (出处: 吾爱破解论坛)
//go build -ldflags "-s -w"

package main

import (
	"crypto/tls"
	"io"
	"net/http"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

var (
	sdyd = map[string]string{

		`Referer`: `https://m.sdyd.org/`,

		`Content-Type`: `application/x-www-form-urlencoded`,

		`Cookie`: `8av__v=1; 8av_t=5; 8av_cdn=https%3A%2F%2F23.225.160.18%2Fxcdn8%2F; 8avlatest=6t67; 8avauth=%012169526422943f9IjQ2NTYyOVx0NmE0ODJjZGZcdGNjZGRielx0MTMwMzA1MDYyMTBcdDM1NTg4MzU1OTVAcXEuY29tXHQxNjk3ODU2MjI5XHQi`,

		`User-Agent`: `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36`}

	uri = `https://m.sdyd.org`

	r1 = regexp.MustCompile(`<div class="sc-title">.*?<span><a href="/html/(\w+).html" target="_blank">(.*?)</a>`)

	r2 = regexp.MustCompile(`<li>⑵ <a href="(.*?)"`)
)

func main() {

	if runtime.GOOS == `windows` {

		exec.Command(`cmd`, `/c`, `start`, `http://localhost/`).Start()

	} else if runtime.GOOS == `linux` {

		exec.Command(`bash`, `-c`, `x-www-browser`, `http://localhost/`).Start()

	} else if runtime.GOOS == `darwin` {

		exec.Command(`open`, `http://localhost/`).Start()
	}

	http.HandleFunc(`/`, Search)

	http.HandleFunc(`/bk`, Book)

	http.ListenAndServe(`:80`, nil)
}

func Search(w http.ResponseWriter, r *http.Request) {

	if r.Method == `GET` {

		io.WriteString(w, `<!DOCTYPE html><html><head><title>Search</title><style>
 
ul,ol,li{list-style-position:inside;text-align:center}
 
@font-face{font-family:t3;src:url(http://23.225.160.18/xcdn8/ft/y3.woff) format('woff2')}.y3{font-family:t3}
 
body{background-image:url([img]http://img1.imgtp.com/2023/09/21/NtHOWm9r.jpg[/img]);background-attachment:fixed;background-repeat:no-repeat;background-size:cover}</style></head><body><div id=div align=center><input type=search name=s id=s><input id=ss type=button value=Search><ol id=ol></ol></div><script>ss.onclick=()=>{
 
let xhr=new XMLHttpRequest()
 
xhr.open('POST',location.href)
 
xhr.setRequestHeader('Content-Type','application/x-www-form-urlencoded')
 
xhr.send('s='+s.value)
 
xhr.onload=()=>{
 
ol.innerHTML=''
 
let re=xhr.response.split('||')
 
for(let i=0;i<re.length-1;i+=2)ol.insertAdjacentHTML('beforeend','<li><a href=bk?s='+re[i]+'>'+re[i+1]+'</a></li>')}}</script></body></html>`)

	} else {
		r.ParseForm()

		c := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}

		q, _ := http.NewRequest(`POST`, uri+`/search/`, strings.NewReader(`keyword=`+r.Form[`s`][0]+`&submit=`))

		for k, v := range sdyd {
			q.Header.Set(k, v)
		}

		p, _ := c.Do(q)

		defer p.Body.Close()

		b, _ := io.ReadAll(p.Body)

		s1 := r1.FindAllStringSubmatch(string(b), -1)

		for i := 0; i < len(s1); i++ {
			io.WriteString(w, s1[i][1]+`||`+s1[i][2]+`||`)
		}
	}
}

func Book(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	c := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}

	q, _ := http.NewRequest(`GET`, uri+`/downpage/`+r.Form[`s`][0]+`/?sure=2`, nil)

	for k, v := range sdyd {
		q.Header.Set(k, v)
	}

	p, _ := c.Do(q)

	b, _ := io.ReadAll(p.Body)

	defer p.Body.Close()

	s1 := r2.FindSubmatch(b)

	http.Redirect(w, r, dl(uri+string(s1[1])), 302)
}

func dl(u string) string {

	c := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}, CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse }}

	q, _ := http.NewRequest(`GET`, u, nil)

	for k, v := range sdyd {
		q.Header.Set(k, v)
	}

	p, _ := c.Do(q)

	defer p.Body.Close()

	return p.Header.Get(`Location`)
}
