package eu
import (
	"github.com/diverse-soles/adidas"
	"net/http"
	"net/url"
//  "time"
  "errors"
  "bytes"
  //"io/ioutil"
 "strconv"
  "github.com/PuerkitoBio/goquery"
 "github.com/sendgrid/sendgrid-go"
)
var (
   UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36"
)
func get_request_url(step,country string)string{
	var url string
   locale := adidas.Locale(country)
	if step=="shipping"{
	url = "https://www.adidas."+adidas.Serverext[country]+ "/on/demandware.store/Sites-adidas-"+country+"-Site/"+locale+"/CODelivery-Start"
}else if step=="payment"{
  url = "https://www.adidas."+adidas.Serverext[country] + "on/demandware.store/Sites-adidas-"+ country+"-Site/"+locale+"/COSummary-Start"
}
return url
}
func get_request_body(step,country,key,shippingisbilling string,profile map[string]string)url.Values{
	data := url.Values{}
  if step=="shipping"{

		data.Add("dwfrm_delivery_deliverymethod","Standard")
data.Add("dwfrm_delivery_shippingOriginalAddress","false")
data.Add("dwfrm_delivery_shippingSuggestedAddress","false")
data.Add("dwfrm_delivery_singleshipping_shippingAddress_isedited","false")
data.Add("dwfrm_delivery_singleshipping_shippingAddress_addressFields_firstName",profile["sfname"])
data.Add("dwfrm_delivery_singleshipping_shippingAddress_addressFields_lastName",profile["slname"])
data.Add("dwfrm_delivery_singleshipping_shippingAddress_addressFields_houseNumber",profile["shousenum"])
data.Add("dwfrm_delivery_singleshipping_shippingAddress_addressFields_address1",profile["saddy1"])
data.Add("dwfrm_delivery_singleshipping_shippingAddress_addressFields_address2",profile["saddy2"])
data.Add("dwfrm_delivery_singleshipping_shippingAddress_addressFields_city",profile["scity"])
data.Add("dwfrm_delivery_singleshipping_shippingAddress_addressFields_zip",profile["szip"])
data.Add("dwfrm_delivery_singleshipping_shippingAddress_addressFields_phone",profile["phone"])
data.Add("dwfrm_delivery_singleshipping_shippingAddress_useAsBillingAddress",shippingisbilling)
//data.Add("dwfrm_delivery_singleshipping_shippingAddress_email_emailAddress","ilikepies@gmail.com")
data.Add("dwfrm_delivery_securekey",key)
data.Add("dwfrm_delivery_billingOriginalAddress","false")
data.Add("dwfrm_delivery_billingSuggestedAddress","false")
data.Add("dwfrm_delivery_billing_billingAddress_isedited","false")
data.Add("dwfrm_delivery_billing_billingAddress_addressFields_firstName",profile["bfname"])
data.Add("dwfrm_delivery_billing_billingAddress_addressFields_lastName",profile["blname"])
data.Add("dwfrm_delivery_billing_billingAddress_addressFields_houseNumber",profile["bhousenum"])
data.Add("dwfrm_delivery_billing_billingAddress_addressFields_address1",profile["baddy1"])
data.Add("dwfrm_delivery_billing_billingAddress_addressFields_address2",profile["baddy2"])
data.Add("dwfrm_delivery_billing_billingAddress_addressFields_city",profile["bcity"])
data.Add("dwfrm_delivery_billing_billingAddress_addressFields_zip",profile["bzip"])
data.Add("dwfrm_delivery_billing_billingAddress_addressFields_country",country)
data.Add("dwfrm_delivery_billing_billingAddress_addressFields_phone",profile["phone"])
data.Add("signup_source","shipping")
data.Add("dwfrm_delivery_singleshipping_shippingAddress_ageConfirmation","true")
data.Add("shipping-group-0","Standard")
data.Add("dwfrm_cart_shippingMethodID_0","Standard")
data.Add("shippingMethodType_0","inline")
data.Add("dwfrm_cart_selectShippingMethod","ShippingMethodID")
data.Add("referer","Cart-Show")
data.Add("dwfrm_delivery_singleshipping_shippingAddress_agreeForSubscription","true")
data.Add("dwfrm_delivery_savedelivery","Review & Pay")
data.Add("format","ajax")

	}
  return data
}
func AqcuireShippingKeys(client *http.Client,country string,retry int) (string,string,error){
if retry>6{
        return "","",errors.New("Error Aqcuireing Shipping Keys")
    }
urlstr := get_request_url("shipping",country)
req, err := http.NewRequest("GET",urlstr,nil)
req.Header.Add("User-Agent",UserAgent)
res, err:= client.Do(req)
if err!=nil{
    url,securekey,err := AqcuireShippingKeys(client,country,retry +1)
    return url,securekey,err
}
if res.StatusCode!=http.StatusOK{
 url,securekey,err := AqcuireShippingKeys(client,country,retry +1)
    return url,securekey,err
}
defer res.Body.Close()
doc, err:= goquery.NewDocumentFromReader(res.Body)
formnode := doc.Find("#dwfrm_delivery")
action,_  := formnode.Attr("action")
key,_ := doc.Find("input[name=dwfrm_delivery_securekey]").Attr("value")

return action,key,nil
  }
  func SubmitShippingDetails(client *http.Client,action,key,shippingisbilling,country string,profile map[string]string,retry int) error{
    if retry>6{
        return errors.New("Error Submitted Shipping")
    }
    data := get_request_body("shipping",country,key,shippingisbilling,profile)
req, err := http.NewRequest("POST", action, bytes.NewBufferString(data.Encode()))
req.Header.Add("User-Agent", UserAgent)
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
    req.Header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
    req.Header.Add("Connection","keep-alive")
res, err := client.Do(req)
if err!=nil{
    err := SubmitShippingDetails(client,action,key,shippingisbilling,country,profile,retry+1)
    return err
}
if res.StatusCode!=http.StatusAccepted && res.StatusCode!=http.StatusOK{
    err := SubmitShippingDetails(client,action,key,shippingisbilling,country,profile,retry+1) 
    return err
}else{
    return nil
}
}
func SendSuccess(shoe,to,email,pass,country string){
  subject := "Diverse Go - "+shoe
  body := "<strong>Shoe : "+shoe+"<br>Country: "+country+"<br>Email: "+email+"<br> Password: "+pass
   sendgridKey := "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"
    sg := sendgrid.NewSendGridClientWithApiKey(sendgridKey)
    message := sendgrid.NewMail()
    message.AddTo(to)
   // message.AddToName("")
    message.SetSubject(subject)
    message.SetHTML(body)
    message.SetFrom("bot@diversego.com")
    if r := sg.Send(message); r == nil {
               // fmt.Println("Email sent!")
        } else {
              //  fmt.Println(r)
        }
}