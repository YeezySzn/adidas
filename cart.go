package adidas
import (
	"net/http"
	"errors"
	"strings"
	"net/url"
	"bytes"
	"strconv"
	"io/ioutil"
)
func create_cart_link(country string)string{
 path := "/on/demandware.store/Sites-adidas-" + country+ "-Site/en_" + country + "/Cart-MiniAddProduct"
	url := "http://www.adidas." + Serverext[country] +path 
  return url
}
func ATC(client *http.Client,pid,country string,retry int)error{
	  if retry>4{
        return errors.New("Error Adding: "+pid+" to Cart")
    }
masterpid:= strings.Split(pid,"_")[0]
urlstr := create_cart_link(country)
data := url.Values{}
data.Add("pid",pid)
data.Add("masterPid",masterpid)
data.Add("layer","Add To Bag overlay")
data.Add("Quantity","1")
data.Add("x-PrdRtt","")
data.Add("g-recaptcha-response","")
data.Add("request","ajax")
data.Add("responseformat","json")
data.Add("ajax","true")
req, err := http.NewRequest("POST", urlstr, bytes.NewBufferString(data.Encode()))
req.Header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
    req.Header.Add("Connection","keep-alive")
req.Header.Add("User-Agent", UserAgent)
   req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
res, err := client.Do(req)
if err!=nil{

	err := ATC(client,pid,country,retry+1)
    return err
}
if res.StatusCode!=http.StatusOK{
err := ATC(client,pid,country,retry+1)
return err
}

defer res.Body.Close()
body, _:= ioutil.ReadAll(res.Body)
j := string(body)
added := strings.Contains(j,"SUCCESS")
if added{
return nil
}else{
	err := ATC(client,pid,country,retry+1)
    return err
}
}