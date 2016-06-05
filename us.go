package adidas
import (
	"net/http"
    "io/ioutil"
    "github.com/PuerkitoBio/goquery"
    "bytes"
    "net/url"
    "strconv"
    "strings"
    "encoding/json"
    "errors"
)
func ATC(client *http.Client,pid string,retry int)error{
    if retry>9{
        return errors.New("Error Adding: "+pid+" to Cart")
    }
masterpid:= strings.Split(pid,"_")[0]
urlstr := "http://www.adidas.com/on/demandware.store/Sites-adidas-US-Site/en_US/Cart-MiniAddProduct"
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
req.Header.Add("User-Agent", UserAgent)
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
res, err := client.Do(req)
if err!=nil{
	err := ATC(client,pid,retry+1)
    return err
}
if res.StatusCode!=http.StatusOK{
err := ATC(client,pid,retry+1)
return err
}

defer res.Body.Close()
body, _:= ioutil.ReadAll(res.Body)
j := string(body)
added := strings.Contains(j,"SUCCESS")
if added{
return nil
}else{
	err := ATC(client,pid,retry+1)
    return err
}
}
func AqcuireShippingKeys(client *http.Client,retry int) (string,string,error){
    if retry>9{
        return "","",errors.New("Error Aqcuireing Shipping Keys")
    }
urlstr := "https://www.adidas.com/us/delivery-start";
req, err := http.NewRequest("GET",urlstr,nil)
req.Header.Add("User-Agent",UserAgent)
res, err:= client.Do(req)
if err!=nil{
    url,securekey,err := AqcuireShippingKeys(client,retry +1)
    return url,securekey,err
}
if res.StatusCode!=http.StatusOK{
 url,securekey,err := AqcuireShippingKeys(client,retry +1)
    return url,securekey,err
}
defer res.Body.Close()
doc, err:= goquery.NewDocumentFromReader(res.Body)
formnode := doc.Find("#dwfrm_delivery")
action,_  := formnode.Attr("action")
key,_ := doc.Find("input[name=dwfrm_delivery_securekey]").Attr("value")
return action,key,nil

}
func SubmitShippingDetails(client *http.Client,action,key,shippingisbilling string,profile map[string]string,retry int) error{
    if retry>9{
        return errors.New("Error Submitted Shipping")
    }
data := url.Values{}
data.Add("dwfrm_delivery_shippingOriginalAddress","false")
data.Add("dwfrm_delivery_shippingSuggestedAddress","false")
data.Add("dwfrm_delivery_singleshipping_shippingAddress_isedited","false")
data.Add("dwfrm_delivery_singleshipping_shippingAddress_addressFields_firstName",profile["sfname"])
data.Add("dwfrm_delivery_singleshipping_shippingAddress_addressFields_lastName",profile["slname"])
data.Add("dwfrm_delivery_singleshipping_shippingAddress_addressFields_address1",profile["saddy1"])
data.Add("dwfrm_delivery_singleshipping_shippingAddress_addressFields_address2",profile["saddy2"])
data.Add("dwfrm_delivery_singleshipping_shippingAddress_addressFields_city",profile["scity"])
data.Add("dwfrm_delivery_singleshipping_shippingAddress_addressFields_countyProvince",profile["sstate"])
data.Add("state","")
data.Add("dwfrm_delivery_singleshipping_shippingAddress_addressFields_zip",profile["szip"])
data.Add("dwfrm_delivery_singleshipping_shippingAddress_addressFields_phone",profile["phone"])
 data.Add("dwfrm_delivery_securekey",key)
data.Add("dwfrm_delivery_singleshipping_shippingAddress_useAsBillingAddress",shippingisbilling)
data.Add("dwfrm_delivery_billingOriginalAddress","false")
data.Add("dwfrm_delivery_billingSuggestedAddress","false")
data.Add("dwfrm_delivery_billing_billingAddress_isedited","false")
data.Add("dwfrm_delivery_billing_billingAddress_addressFields_country","US")
data.Add("dwfrm_delivery_billing_billingAddress_addressFields_firstName",profile["bfname"])
data.Add("dwfrm_delivery_billing_billingAddress_addressFields_lastName",profile["blname"])
data.Add("dwfrm_delivery_billing_billingAddress_addressFields_address1",profile["baddy1"])
data.Add("dwfrm_delivery_billing_billingAddress_addressFields_address2",profile["baddy2"])
data.Add("dwfrm_delivery_billing_billingAddress_addressFields_city",profile["bcity"])
data.Add("dwfrm_delivery_billing_billingAddress_addressFields_countyProvince",profile["bstate"])
data.Add("state","")
data.Add("dwfrm_delivery_billing_billingAddress_addressFields_zip",profile["bzip"])
data.Add("dwfrm_delivery_billing_billingAddress_addressFields_phone",profile["phone"])
data.Add("dwfrm_delivery_singleshipping_shippingAddress_email_emailAddress",profile["email"])
data.Add("signup_source","shipping")
data.Add("dwfrm_delivery_singleshipping_shippingAddress_ageConfirmation","true")
data.Add("dwfrm_delivery_singleshipping_shippingAddress_agreeForSubscription","true")
data.Add("shipping-group-0","Standard")
data.Add("dwfrm_cart_shippingMethodID_0","Standard")
data.Add("shippingMethodType_0","inline")
data.Add("dwfrm_cart_selectShippingMethod","ShippingMethodID")
data.Add("referer","Cart-Show")
data.Add("dwfrm_delivery_savedelivery","Review and Pay")
data.Add("format","ajax")
req, err := http.NewRequest("POST", action, bytes.NewBufferString(data.Encode()))
req.Header.Add("User-Agent", UserAgent)
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
    req.Header.Add("Origin","https://www.adidas.com")
    req.Header.Add("Referer","https://www.adidas.com/us/delivery-start")
    
res, err := client.Do(req)
if err!=nil{
    err := SubmitShippingDetails(client,action,key,shippingisbilling,profile,retry+1)
    return err
}
if res.StatusCode!=http.StatusOK{
err := SubmitShippingDetails(client,action,key,shippingisbilling,profile,retry+1)
    return err
}else{
    return nil
}

}
func AqcuirePaymentKeys(client *http.Client,retry int) (string,string,error){
    if retry>9{
        return "","",errors.New("Error Submitted Shipping")
    }
    urlstr := "https://www.adidas.com/on/demandware.store/Sites-adidas-US-Site/en_US/COSummary-Start"
    req, err := http.NewRequest("GET",urlstr,nil)
req.Header.Add("User-Agent",UserAgent)
res, err:= client.Do(req)
if err!=nil{
    action,paymentkey,err := AqcuirePaymentKeys(client,retry+1)
    return action,paymentkey,err
}
if res.StatusCode!=http.StatusOK{
 action,paymentkey,err := AqcuirePaymentKeys(client,retry+1)
    return action,paymentkey,err
}
defer res.Body.Close()
doc, err:= goquery.NewDocumentFromReader(res.Body)
formnode := doc.Find("#dwfrm_delivery_billing")
action,_  := formnode.Attr("action")
paymentkey,_ := doc.Find("input[name=dwfrm_payment_securekey]").Attr("value")
return action,paymentkey,nil

}
func SubmitPayDetails(client *http.Client,action, paymentkey string,profile map[string]string,retry int) (map[string]interface{},error){
    if retry>9{
        return nil,errors.New("Error Submitted Shipping")
    }
    cardint:= "001"
    data := url.Values{}
    data.Add("dwfrm_payment_creditCard_type",cardint)
    data.Add("dwfrm_payment_creditCard_owner",profile["bfname"]+profile["blname"])
    data.Add("dwfrm_payment_creditCard_month",profile["expmonth"])
    data.Add("dwfrm_payment_creditCard_year",profile["expyear"])
    data.Add("dwfrm_payment_securekey",paymentkey)
    data.Add("dwfrm_payment_signcreditcardfields","sign")
    req, err := http.NewRequest("POST", action, bytes.NewBufferString(data.Encode()))
req.Header.Add("User-Agent", UserAgent)
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
    res, err:= client.Do(req)
    if err!=nil{
        r,err := SubmitPayDetails(client,action,paymentkey,profile,retry+1)
        return r,err
    }
  if res.StatusCode!=http.StatusOK{
 r,err := SubmitPayDetails(client,action,paymentkey,profile,retry+1)
        return r,err
}
    defer res.Body.Close()
    body,err := ioutil.ReadAll(res.Body)
    cyberobj := CyberSourceResponse{}
   json.Unmarshal([]byte(string(body)),&cyberobj)
fields := cyberobj.Fields
return fields,nil

    

}
func CyberSourceSubmit(client *http.Client,profile map[string]string,fields map[string]interface{},retry int)error{
    if retry>9{
        return errors.New("Error Submitted Shipping")
    }
    urlstr := "https://secureacceptance.cybersource.com/silent/pay"
    data := url.Values{}
    for key,val := range fields{
       switch val.(type){
       case int:
            s := strconv.Itoa(val.(int))
            data.Add(key,s)
       case float64:
         s := strconv.FormatFloat(val.(float64), 'f', 2, 64)
         nodecimals := strings.Contains(s,".00")
         if nodecimals{
            s = strings.Split(s,".")[0]
         }
            data.Add(key,s)
        
        case string: 
        data.Add(key,val.(string))

       }
        
}
    data.Add("card_cvn",profile["cvv"])
    data.Add("card_number",profile["cardnum"])
    req,err := http.NewRequest("POST",urlstr,bytes.NewBufferString(data.Encode()))
    req.Header.Add("User-Agent",UserAgent)
    req.Header.Add("Origin","https://www.adidas.com")
    req.Header.Add("Referer","https://www.adidas.com/on/demandware.store/Sites-adidas-US-Site/en_US/COSummary-Start")
    req.Header.Add("Content-Type","application/x-www-form-urlencoded")
    req.Header.Add("Content-Length",strconv.Itoa(len(data.Encode())))
    res,err := client.Do(req)
    if err!=nil{
        err := CyberSourceSubmit(client,profile,fields,retry+1)
        return err
    }
    if res.StatusCode!=http.StatusOK{
    err := CyberSourceSubmit(client,profile,fields,retry+1)
        return err
}
    defer res.Body.Close()
    return nil
   
}