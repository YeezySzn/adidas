static void getkeys(CookieContainer cookies)
        {
            var watch = Stopwatch.StartNew();


            string url = "https://www.adidas.com/us/delivery-start";
            HttpWebRequest request = (HttpWebRequest)WebRequest.Create(url);
            request.CookieContainer = cookies;
            request.UserAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.86 Safari/537.36";
            HttpWebResponse response = (HttpWebResponse)request.GetResponse();
            Stream stream;
            stream = response.GetResponseStream();
            StreamReader sr = new StreamReader(stream);

            string read = sr.ReadToEnd();

            HtmlAgilityPack.HtmlDocument document = new HtmlAgilityPack.HtmlDocument();

            document.LoadHtml(read);
            HtmlNode formtag = document.GetElementbyId("dwfrm_delivery");

            string action = formtag.Attributes["action"].Value;
            HtmlNode rawkey = document.DocumentNode.SelectSingleNode("//input[@name='dwfrm_delivery_securekey']");
            string key = rawkey.Attributes["value"].Value;
            var elapsedMs = watch.ElapsedMilliseconds;
            string time = "Elapsed Time(ms) - " + elapsedMs;

            Console.WriteLine("Thread ID " + Thread.CurrentThread.ManagedThreadId + " got shipping keys... "+time);
           

            sr.Close();
            submitship(cookies, action, key);
        }
        static void submitship(CookieContainer cookies, string action, string key)
        {
            var watch = Stopwatch.StartNew();
            ASCIIEncoding encoding = new ASCIIEncoding();
            string pdata = "";

            pdata += "dwfrm_delivery_shippingOriginalAddress=false";
            pdata += "&dwfrm_delivery_shippingSuggestedAddress=false";
            pdata += "&dwfrm_delivery_singleshipping_shippingAddress_isedited=false";
            pdata += "&dwfrm_delivery_singleshipping_shippingAddress_addressFields_firstName=ALEX";
            pdata += "&dwfrm_delivery_singleshipping_shippingAddress_addressFields_lastName=AMIH";
            pdata += "&dwfrm_delivery_singleshipping_shippingAddress_addressFields_address1=5+china+ave";
            pdata += "&dwfrm_delivery_singleshipping_shippingAddress_addressFields_address2=";
            pdata += "&dwfrm_delivery_singleshipping_shippingAddress_addressFields_city=china";
            pdata += "&dwfrm_delivery_singleshipping_shippingAddress_addressFields_countyProvince=NY";
            pdata += "&dwfrm_delivery_singleshipping_shippingAddress_addressFields_zip=11111";
            pdata += "&dwfrm_delivery_singleshipping_shippingAddress_addressFields_phone=8282020020";
            pdata += "&dwfrm_delivery_singleshipping_shippingAddress_useAsBillingAddress=true";
            pdata += "&dwfrm_delivery_securekey=" + key;
            pdata += "&dwfrm_delivery_billingOriginalAddress=false";
            pdata += "&dwfrm_delivery_billingSuggestedAddress=false";
            pdata += "&dwfrm_delivery_billing_billingAddress_isedited=false";
            pdata += "&dwfrm_delivery_billing_billingAddress_addressFields_country=US";
            pdata += "&dwfrm_delivery_billing_billingAddress_addressFields_firstName=ALEX";
            pdata += "&dwfrm_delivery_billing_billingAddress_addressFields_lastName=AMIH";
            pdata += "&dwfrm_delivery_billing_billingAddress_addressFields_address1=5+china+ave";
            pdata += "&dwfrm_delivery_billing_billingAddress_addressFields_address2=";
            pdata += "&dwfrm_delivery_billing_billingAddress_addressFields_city=china";
            pdata += "&dwfrm_delivery_billing_billingAddress_addressFields_countyProvince=NY";
            pdata += "&dwfrm_delivery_billing_billingAddress_addressFields_zip=11111";
            pdata += "&dwfrm_delivery_billing_billingAddress_addressFields_phone=8282020020";
            pdata += "&dwfrm_delivery_singleshipping_shippingAddress_email_emailAddress=fvfkv%40skk.com";
            pdata += "&signup_source=shipping";
            pdata += "&dwfrm_delivery_singleshipping_shippingAddress_ageConfirmation=true";
            pdata += "&dwfrm_delivery_singleshipping_shippingAddress_agreeForSubscription=true";
            pdata += "&shipping-group-0=2ndDay";
            pdata += "&dwfrm_cart_shippingMethodID_0=2ndDay";
            pdata += "&shippingMethodType_0=inline";
            pdata += "&dwfrm_cart_selectShippingMethod=ShippingMethodID";
            pdata += "&referer=Cart-Show";
            pdata += "&dwfrm_delivery_savedelivery=Review%20and%20Pay";
            pdata += "&format=ajax";
            HttpWebRequest request = (HttpWebRequest)WebRequest.Create(action);
            request.Method = "POST";
            // request.KeepAlive = true;
            request.CookieContainer = cookies;
            request.ContentType = "application/x-www-form-urlencoded";


            byte[] data = encoding.GetBytes(pdata);

            Stream stream;
            stream = request.GetRequestStream();
            stream.Write(data, 0, data.Length);
            stream.Close();

            HttpWebResponse response = (HttpWebResponse)request.GetResponse();
            stream = response.GetResponseStream();
            StreamReader sr = new StreamReader(stream);
            sr.Close();
            place1(cookies, action);

        }
        static void place1(CookieContainer cookies, string refer)
        {
            var watch = Stopwatch.StartNew();
            string url = "https://www.adidas.com/on/demandware.store/Sites-adidas-US-Site/en_US/COSummary-Start";
            HttpWebRequest request = (HttpWebRequest)WebRequest.Create(url);
            request.CookieContainer = cookies;
            request.UserAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.86 Safari/537.36";
            HttpWebResponse response = (HttpWebResponse)request.GetResponse();
            Stream stream;
            stream = response.GetResponseStream();
            StreamReader sr = new StreamReader(stream);
            string read = sr.ReadToEnd();
            HtmlAgilityPack.HtmlDocument document = new HtmlAgilityPack.HtmlDocument();
            document.LoadHtml(read);
            HtmlNode form = document.GetElementbyId("dwfrm_delivery_billing");
            string action = form.Attributes["action"].Value;
            HtmlNode rawkey = document.DocumentNode.SelectSingleNode("//input[@name='dwfrm_payment_securekey']");
            string key = rawkey.Attributes["value"].Value;
            var elapsedMs = watch.ElapsedMilliseconds;
            string time = "Elapsed Time(ms) - " + elapsedMs;

            Console.WriteLine("Thread ID " + Thread.CurrentThread.ManagedThreadId + " got billing key and url request... "+time);
            

            sr.Close();
            place2(cookies, key, action, refer);

        }
        static void place2(CookieContainer cookies, string key, string action, string refer)
        {
            var watch = Stopwatch.StartNew();

            ASCIIEncoding encoding = new ASCIIEncoding();
            string pdata = "";
            pdata += "dwfrm_payment_creditCard_type=001";
            pdata += "&dwfrm_payment_creditCard_owner=john+smith";
            pdata += "&dwfrm_payment_creditCard_month=03";
            pdata += "&dwfrm_payment_creditCard_year=2017&";
            pdata += "dwfrm_payment_securekey=" + key;
            pdata += "&dwfrm_payment_signcreditcardfields=sign";




            HttpWebRequest request = (HttpWebRequest)WebRequest.Create(action);
            request.Method = "POST";
            // request.KeepAlive = true;
            request.CookieContainer = cookies;
            request.ContentType = "application/x-www-form-urlencoded";
            request.UserAgent = "Mozilla / 5.0(Windows NT 6.1; WOW64) AppleWebKit / 537.36(KHTML, like Gecko) Chrome / 46.0.2490.86 Safari / 537.36";
            byte[] data = encoding.GetBytes(pdata);
            Stream stream;
            stream = request.GetRequestStream();
            stream.Write(data, 0, data.Length);
            stream.Close();
            HttpWebResponse response = (HttpWebResponse)request.GetResponse();
            stream = response.GetResponseStream();
            StreamReader sr = new StreamReader(stream);
            string json = sr.ReadToEnd();
            var elapsedMs = watch.ElapsedMilliseconds;
            string time = "Elapsed Time(ms) - " + elapsedMs;

            Console.WriteLine("Thread ID " + Thread.CurrentThread.ManagedThreadId + " Got Response Fields to submit... "+time);
            sr.Close();
            lock (locker)
            {
                if (status == 0)
                {
                    status += 1;
                   
                    place3(cookies, json);

                }
                if (status != 0)
                {
                    return;
                }


            }


        }

        static void place3(CookieContainer cookies, string json)
        {
            var watch = Stopwatch.StartNew();
            ASCIIEncoding encoding = new ASCIIEncoding();
            string pdata = "";
            JObject obj = JObject.Parse(json);
            string fields = obj["fieldsToSubmit"].ToString();
            obj = JObject.Parse(fields);

            foreach (var x in obj)
            {
                string name = Uri.EscapeDataString(x.Key);
                JToken value = Uri.EscapeDataString(x.Value.ToString());
                if ((pdata) == "")
                {
                    pdata += name + "=" + value;
                }

                else
                {
                    if (x.Key == "signed_date_time")
                    {
                        string v = x.Value.ToString().Replace(" ", "T");
                        v += "Z";
                        pdata += "&" + name + "=" + Uri.EscapeDataString(v);

                    }
                    else
                    {
                        pdata += "&" + name + "=" + value;
                    }

                }
            }
            pdata += "&card_cvn=123";
            pdata += "&card_number=4111111111111111";


            string url = "https://secureacceptance.cybersource.com/silent/pay";
            HttpWebRequest request = (HttpWebRequest)WebRequest.Create(url);
            request.Method = "POST";
            request.KeepAlive = true;
            request.Headers["Origin"] = "https://www.adidas.com";
            request.Headers["Accept-Language"] = "en-US,en;q=0.8";
            request.Accept = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8";
            request.Referer = "https://www.adidas.com/on/demandware.store/Sites-adidas-US-Site/en_US/COSummary-Start?fromdelivery=true";
            request.UserAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.86 Safari/537.36";

            // request.KeepAlive = true;
            request.CookieContainer = cookies;
            request.ContentType = "application/x-www-form-urlencoded";
            byte[] data = encoding.GetBytes(pdata);
            Stream stream;
            stream = request.GetRequestStream();
            stream.Write(data, 0, data.Length);
            stream.Close();
            HttpWebResponse response = (HttpWebResponse)request.GetResponse();
            stream = response.GetResponseStream();
            StreamReader sr = new StreamReader(stream);
            var elapsedMs = watch.ElapsedMilliseconds;
            string time = "Elapsed Time(ms) - " + elapsedMs;

            Console.WriteLine("Thread ID " + Thread.CurrentThread.ManagedThreadId + " Submited to cyber Source... "+time);

            final(cookies, sr.ReadToEnd());
            sr.Close();





        }
        static void final(CookieContainer cookies, string src)

        {
            var watch = Stopwatch.StartNew();

            string pdata = "";
            HtmlAgilityPack.HtmlDocument document = new HtmlAgilityPack.HtmlDocument();

            document.LoadHtml(src);
            HtmlNode nodea = document.GetElementbyId("custom_redirect");

            string action = nodea.Attributes["action"].Value;
            foreach (HtmlNode node in document.DocumentNode.SelectNodes("//input"))
            {
                string val = Uri.EscapeDataString(node.Attributes["value"].Value);
                string name = Uri.EscapeDataString(node.Attributes["name"].Value);
                if (pdata == "")
                {
                    pdata += name + "=" + val;
                }
                else
                {
                    pdata += "&" + name + "=" + val;
                }
            }
            ASCIIEncoding encoding = new ASCIIEncoding();
            HttpWebRequest request = (HttpWebRequest)WebRequest.Create(action);
            request.Method = "POST";
            request.KeepAlive = true;
            request.Headers["Origin"] = "https://www.adidas.com";

            request.UserAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.86 Safari/537.36";

            // request.KeepAlive = true;
            request.CookieContainer = cookies;
            request.ContentType = "application/x-www-form-urlencoded";
            byte[] data = encoding.GetBytes(pdata);
            Stream stream;
            stream = request.GetRequestStream();
            stream.Write(data, 0, data.Length);
            stream.Close();
            HttpWebResponse response = (HttpWebResponse)request.GetResponse();
            stream = response.GetResponseStream();
            StreamReader sr = new StreamReader(stream);
            var elapsedMs = watch.ElapsedMilliseconds;
            string time = "Elapsed Time(ms) - " + elapsedMs;

            Console.WriteLine("Thread ID " + Thread.CurrentThread.ManagedThreadId + " Order Submited to Adidas..."+time);
            return;



        }


        static void Main(string[] args)
        {

            

            for (int x = 0; x < 5; x++)
            {
                
                Thread thread = new Thread(add);
                thread.Start();
                Thread.Sleep(400);
                
                
            }

            

        }

    }
}