package adidas
import (
  "errors"
  "net/url"
  "bytes"
  "strconv"
  "net/http"
  "github.com/PuerkitoBio/goquery"
  "time"
  "strings"
  //"fmt"
)
func get_server_url(country string) Server{
  locale := Locale(country)
  return Server{
    sign_in_page: "https://cp.adidas." + Serverext[country] + "/web/eCom/" + locale + "/loadsignin?target=account",
    start_sso_session: "https://cp.adidas." + Serverext[country] + "/idp/startSSO.ping",
    create_sso_cookie: "https://cp.adidas." + Serverext[country] + "/web/ssoCookieCreate?resume",
    create_sso_domain_cookie: "https://cp.adidasspecialtysports." + Serverext[country] + "/web/createSSODomainCookie?domain=.adidasspecialtysports.com&ssoiniturl=https://cp.adidas.com",
    cp_resume: "https://cp.adidas." + Serverext[country],
    cp_saml: "https://cp.adidas." + Serverext[country] + "/sp/ACS.saml2",
    resume_login: "https://www.adidas." + Serverext[country] + "/on/demandware.store/Sites-adidas-" + country + "-Site/" + locale + "/MyAccount-ResumeLogin",
    target_resource: "https://www.adidas." + Serverext[country] + "/on/demandware.store/Sites-adidas-" + country + "-Site/" + locale +"/MyAccount-ResumeLogin?target=account&target=account",
    my_account: "https://www.adidas." + Serverext[country] + "/on/demandware.store/Sites-adidas-" + country + "-Site/" + locale + "/MyAccount-Show?fromlogin=true",
    relay_state: "https://www.adidas." + Serverext[country] + "/on/demandware.store/Sites-adidas-" + country + "-Site/" + locale + "/MyAccount-ResumeLogin?target=account&target=account"}
 }
  func LoadSignInPage(client *http.Client,urlstr string,retry int) (s string,err error){
    if retry>6{
      return "",errors.New("Error Loading Sign In Page")
    }
req, err := http.NewRequest("GET",urlstr,nil)
req.Header.Add("Pragma","no-cache")
  //req.Header.Add("Origin","https://cp.adidas.com")
  req.Header.Add("Accept-Language","en-US,en;q=0.8")
  req.Header.Add("Upgrade-Insecure-Requests","1")
  req.Header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
  req.Header.Add("Cache-Control","no-cache")
 // req.Header.Add("Referer","https://cp.adidas.com/web/eCom/en_US/loadsignin?target=account")
  req.Header.Add("Connection","keep-alive")
req.Header.Add("User-Agent",UserAgent)
res, err:= client.Do(req)
if err!=nil{
  time.Sleep(time.Millisecond * 200)
  csrf,err := LoadSignInPage(client,urlstr,retry+1)
    return csrf,err
  }
defer res.Body.Close()
doc, err:= goquery.NewDocumentFromReader(res.Body)
csrf,_ := doc.Find("[name=CSRFToken]").Attr("value")

return csrf,nil

  }
func start_sso_session(client *http.Client,urlstr string, data url.Values,retry int)(string,error){
  if retry>7{
      return "",errors.New("Error Creating SSO Session")
    }
req, err := http.NewRequest("POST", urlstr, bytes.NewBufferString(data.Encode()))
req.Header.Add("Pragma","no-cache")
  //req.Header.Add("Origin","https://cp.adidas.com")
  req.Header.Add("Accept-Language","en-US,en;q=0.8")
  req.Header.Add("Upgrade-Insecure-Requests","1")
  req.Header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
  req.Header.Add("Cache-Control","no-cache")
  //req.Header.Add("Referer","https://cp.adidas.com/web/eCom/en_US/loadsignin?target=account")
  req.Header.Add("Connection","keep-alive")
  req.Header.Add("User-Agent", UserAgent)
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
res, err := client.Do(req)
if err!=nil{
  if !strings.Contains(err.Error(),"302"){
  time.Sleep(time.Millisecond * 200)
 location,err := start_sso_session(client, urlstr,data,retry +1)
  return location,err
}
}
if res.StatusCode==302{
  location := res.Header.Get("Location")
  return location,nil
}else{
  time.Sleep(time.Millisecond * 200)
  location,err := start_sso_session(client, urlstr,data,retry +1)
  return location,err
}

}
func get_request_body(username,password,country,csrf string) url.Values{
   var (
    locale = Locale(country)
    signinSubmit= "Sign in"
    IdpAdapterId= "adidasIdP10"
    SpSessionAuthnAdapterId= "https://cp.adidas." + Serverext[country] + "/web/"
    PartnerSpId= "sp:demandware"
    validator_id= "adieComDWus"
    TargetResource= "https://www.adidas." + Serverext[country] + "/on/demandware.store/Sites-adidas-" + country + "-Site/" + locale + "/MyAccount-ResumeLogin?target=account&target=account"
    InErrorResource= "https://www.adidas." + Serverext[country] + "/on/demandware.store/Sites-adidas-" + country + "-Site/" + locale + "/null"
    loginUrl= "https://cp.adidas." + Serverext[country] + "/web/eCom/"+locale+"/loadsignin"
    cd= "eCom|" + locale + "|cp.adidas." + Serverext[country] + "|null"
    app= "eCom"
    domain= "cp.adidas." + Serverext[country]
    email= ""
    pfRedirectBaseURL_test= "https://cp.adidas." + Serverext[country]
    pfStartSSOURL_test= "https://cp.adidas." + Serverext[country] + "/idp/startSSO.ping"
    resumeURL_test= ""
    FromFinishRegistraion= ""
    CSRFToken= csrf
    )
  data := url.Values{}
  data.Add("username",username)
  data.Add("password",password)
  data.Add("signinSubmit",signinSubmit)
  data.Add("IdpAdapterId",IdpAdapterId)
  data.Add("SpSessionAuthnAdapterId",SpSessionAuthnAdapterId)
  data.Add("PartnerSpId",PartnerSpId)
  data.Add("validator_id",validator_id)
  data.Add("TargetResource",TargetResource)
  data.Add("InErrorResource",InErrorResource)
  data.Add("loginUrl",loginUrl)
  data.Add("cd",cd)
  data.Add("app",app)
  data.Add("locale",locale)
  data.Add("domain",domain)
  data.Add("email",email)
  data.Add("pfRedirectBaseURL_test",pfRedirectBaseURL_test)
  data.Add("pfStartSSOURL_test",pfStartSSOURL_test)
  data.Add("resumeURL_test",resumeURL_test)
  data.Add("FromFinishRegistraion",FromFinishRegistraion)
  data.Add("CSRFToken",CSRFToken)

return data
}
func FollowSSORedirect(client *http.Client, redirect_url string,retry int)error{
  if retry>7{
      return errors.New("Error Following Redirect")
    }
  req, err := http.NewRequest("GET",redirect_url,nil)
  req.Header.Add("Pragma","no-cache")
  //req.Header.Add("Origin","https://cp.adidas.com")
  req.Header.Add("Accept-Language","en-US,en;q=0.8")
  req.Header.Add("Upgrade-Insecure-Requests","1")
  req.Header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
  req.Header.Add("Cache-Control","no-cache")
 // req.Header.Add("Referer","https://cp.adidas.com/web/eCom/en_US/loadsignin?target=account")
  req.Header.Add("Connection","keep-alive")
  req.Header.Add("User-Agent",UserAgent)
res, err:= client.Do(req)
if err!=nil{
  time.Sleep(time.Millisecond * 200)
  followed := FollowSSORedirect(client,redirect_url,retry+1)
  return followed
}
defer res.Body.Close()
  return nil

}
func CreateSSODomainCookie(client *http.Client, urlstr string,retry int)error{
  if retry>7{
      return errors.New("Error Creating Cookie")
    }
req, err := http.NewRequest("GET",urlstr,nil)
req.Header.Add("Pragma","no-cache")
 // req.Header.Add("Origin","https://cp.adidas.com")
  req.Header.Add("Accept-Language","en-US,en;q=0.8")
  req.Header.Add("Upgrade-Insecure-Requests","1")
  req.Header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
  req.Header.Add("Cache-Control","no-cache")
 // req.Header.Add("Referer","https://cp.adidas.com/web/eCom/en_US/loadsignin?target=account")
  req.Header.Add("Connection","keep-alive")
  req.Header.Add("User-Agent",UserAgent)
res, err:= client.Do(req)
if err!=nil{
  time.Sleep(time.Millisecond * 200)
  err := CreateSSODomainCookie(client,urlstr,retry +1)
  return err
}
defer res.Body.Close()
  return nil


}
func ResumeCP(client *http.Client,urlstr string,retry int)(string,error){
   if retry>7{
      return "",errors.New("Error Resuming CP")
    }
  req, err := http.NewRequest("GET",urlstr,nil)
req.Header.Add("Pragma","no-cache")
 // req.Header.Add("Origin","https://cp.adidas.com")
  req.Header.Add("Accept-Language","en-US,en;q=0.8")
  req.Header.Add("Upgrade-Insecure-Requests","1")
  req.Header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
  req.Header.Add("Cache-Control","no-cache")
  //req.Header.Add("Referer","https://cp.adidas.com/web/eCom/en_US/loadsignin?target=account")
  req.Header.Add("Connection","keep-alive")
  req.Header.Add("User-Agent",UserAgent)
res, err:= client.Do(req)
if err!=nil{
  time.Sleep(time.Millisecond * 200)
  saml,err := ResumeCP(client,urlstr,retry +1)
  return saml,err
}
defer res.Body.Close()
doc, err:= goquery.NewDocumentFromReader(res.Body)
saml,_ := doc.Find("[name=SAMLResponse]").Attr("value")

return saml,nil

}
func PostSaml(client *http.Client,urlstr,SAML,relay_state string,retry int) (string,error){
   if retry>7{
      return "", errors.New("Error Posting Saml")
    }
   data := url.Values{}
  data.Add("SAMLResponse",SAML)
  data.Add("RelayState",relay_state)
  data.Add("submit","Resume")
  req, err := http.NewRequest("POST", urlstr, bytes.NewBufferString(data.Encode()))
  req.Header.Add("Pragma","no-cache")
 // req.Header.Add("Origin","https://cp.adidas.com")
  req.Header.Add("Accept-Language","en-US,en;q=0.8")
  req.Header.Add("Upgrade-Insecure-Requests","1")
  req.Header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
  req.Header.Add("Cache-Control","no-cache")
  //req.Header.Add("Referer","https://cp.adidas.com/web/eCom/en_US/loadsignin?target=account")
  req.Header.Add("Connection","keep-alive")
  req.Header.Add("User-Agent", UserAgent)
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
res, err := client.Do(req)
if err!=nil{
  time.Sleep(time.Millisecond * 200)
  ref,err := PostSaml(client,urlstr,SAML,relay_state,retry+1)
  return ref,err
}
defer res.Body.Close()
doc, err:= goquery.NewDocumentFromReader(res.Body)
ref,_ := doc.Find("[name=REF]").Attr("value")
return ref,nil

}
func PostRef(client *http.Client,urlstr,REF,target string,retry int)error{
   if retry>7{
      return errors.New("Error Posting Ref")
    }
   data := url.Values{}
  data.Add("TargetResource",target)
  data.Add("REF",REF)
  req, err := http.NewRequest("POST", urlstr, bytes.NewBufferString(data.Encode()))
  req.Header.Add("Pragma","no-cache")
 // req.Header.Add("Origin","https://cp.adidas.com")
  req.Header.Add("Accept-Language","en-US,en;q=0.8")
  req.Header.Add("Upgrade-Insecure-Requests","1")
  req.Header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
  req.Header.Add("Cache-Control","no-cache")
  //req.Header.Add("Referer","https://cp.adidas.com/web/eCom/en_US/loadsignin?target=account")
  req.Header.Add("Connection","keep-alive")
  req.Header.Add("User-Agent", UserAgent)
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
res, err := client.Do(req)
if err!=nil{
  time.Sleep(time.Millisecond * 200)
   err:= PostRef(client,urlstr,REF,target,retry+1)
   return err
}
defer res.Body.Close()
return nil
}
func MakeResumeURL(location,cp_resume string) string{
  u,_:= url.Parse(location)
m, _ := url.ParseQuery(u.RawQuery)
resume := m.Get("resume")
urlstr := cp_resume + resume
return urlstr
}
func Login(client *http.Client, username,password,country string) bool{
server := get_server_url(country)
csrf,err := LoadSignInPage(client,server.sign_in_page,0)
data := get_request_body(username,password,country,csrf)
location,err := start_sso_session(client,server.start_sso_session,data,0)
err = FollowSSORedirect(client,location,0)
err = CreateSSODomainCookie(client,server.create_sso_domain_cookie,0)
resume_url:= MakeResumeURL(location,server.cp_resume)
saml,err := ResumeCP(client,resume_url,0)
ref,err := PostSaml(client,server.cp_saml,saml,server.relay_state,0)
PostRef(client,server.resume_login,ref,server.target_resource,0)
if err!=nil{
  return false
}
return true

}