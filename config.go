package adidas
import (

	"labix.org/v2/mgo/bson"
	"time"
)
var (
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36"
  Serverext = mkserverext()
)
func mkserverext()map[string]string{
  s := make(map[string]string)
  s["US"]="com"
  s["GB"] = "co.uk"
  s["AU"]="com.au"
  s["CA"]="ca"
  s["DE"]="de"
  s["ES"]="es"
  s["BE"]="be"
  s["CZ"]="cz"
  s["DK"]="dk"
  s["FR"]="fr"
  s["IT"]="it"
  s["IE"]="ie"
  s["NL"]="nl"
  s["SE"]="se"
  return s
}
type CyberSourceResponse struct{
	Fields map[string]interface{} `json:"fieldsToSubmit"`
}
type AdidasTask struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Type string
	Country string
	Starttime time.Time
	PID string
	Profile string 						 				 
	Email string					
	Started string
	Done string
	Notes string
	Timestamp time.Time
}
type Server struct{
  sign_in_page string
  start_sso_session string
  create_sso_cookie string
  create_sso_domain_cookie string
  cp_resume string
  cp_saml string
  resume_login string
  target_resource string
  my_account string
  relay_state string
}
