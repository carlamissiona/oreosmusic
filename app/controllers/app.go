package controllers
import(
  
  "github.com/revel/revel"
  "log"
 // "fmt"
//  "html"
 // "net/mail"
 // "errors"
 // "net"
 // _"bytes"
 // "net/smtp"
 // "crypto/tls"
  //_"io"
// "oreos/app"
// "oreos/app/hexarch/ports"
//_ "oreos/app/hexarch/repositories"

)

type App struct {
  * revel.Controller
  //userService ports.UserService
  // userSvc ports.UserService
}
 

func(c App) Index() revel.Result {

  log.Println("Initial DB")
//   log.Println(" ---> %", app.DbInit)
//   log.Println(" ---> %", app.InsDB)

   

  // c.userSvc.NewWS()
  return c.Render()
}

func(c App) Home() revel.Result {
//   log.Println("Initial DB")
//   log.Println(" ---> %", app.DbInit)
//   log.Println(" ---> %", app.InsDB)

  return c.Render()
}
func(c App) Signup() revel.Result {
  return c.Render()
    
}
  
// }
// func(c App) SubmitSignup()revel.Result {

//   firstname := html.EscapeString(c.Params.Form.Get("fname") )
//   lastname := html.EscapeString(c.Params.Form.Get("lname") )
//   email := html.EscapeString(c.Params.Form.Get("email") )
//   password := html.EscapeString(c.Params.Form.Get("password") )
//   country := html.EscapeString(c.Params.Form.Get("address") )

//   c.Validation.Required(firstname).Message("Please enter a firstname")
//   c.Validation.Required(lastname).Message("Please enter a lastname")
//   c.Validation.Required(email).Message("Please enter a email")
//   c.Validation.Required(password).Message("Please enter a password")
//   c.Validation.Required(country).Message("Please enter a country ")


//   newuser, err := c.userSvc.NewUser(firstname, lastname, "username", email, country, password, true , "violet" ) 
//   log.Println("newuser")
//   log.Println(" ---> %",newuser)
//   if err != nil {
//     errors.New("error at new user web service")
//     return c.Redirect(App.Signup)
//   }

  
//   log.Println("THIS IS SignUP >>>>> > ", firstname + " +++ " + c.Params.Form.Get("fname") + " " + lastname + " " +email + " " +  country + password )
//   log.Println("THIS IS SignUP >>>>> > ", firstname + " +++ " + c.Params.Form.Get("fname") + " " + lastname + " " +email + " " +  country + password )

//   if c.Validation.HasErrors() {
// 		c.Validation.Keep()
// 		c.FlashParams()
// 		return c.Redirect(App.Signup)
// 	}

  



//   return c.Redirect(App.Home)
// }

// func(c App) SignupVerify (emailstr string, ) revel.Result {

//   if isEmailValid(emailstr) {

//     c.Flash.Error("The Email Is A Valid Input But We Couldn't Reach Your Email Domain")

//     c.Flash.Success("Success!")
//   }

//   return c.Redirect(App.Home)

// }
// func(c App) EmailPreview(emailstr string) revel.Result {

//   if isEmailValid(emailstr) {

//     c.Flash.Error("The Email Is A Valid Input But We Couldn't Reach Your Email Domain")

//     c.Flash.Success("Success!")
//   }

//   return c.Redirect(App.Home)

// }

// func isEmailValid(emailstr string) bool {
//   log.Println("THIS IS EMAIL ___ > ", emailstr)
//   log.Println("THIS IS EMAIL ___ > ", emailstr)
//   log.Println("THIS IS EMAIL ___ > ", emailstr)
//   from:= mail.Address { "sender", "oreos.music@yandex.com"   }
//   to:= mail.Address { "user","missiona.carla@gmail.com" }

//   subj:= "!!This is the email subject"
//   f:="HI"
//   body:= "<h2>Hello </h2>"+f

//   // Setup headers
//   headers:= make(map[string] string)
//   headers["From"] = from.String()
//   headers["To"] = to.String()
//   headers["Subject"] = subj
//   headers["MIME-Version"] = "1.0"
//   headers["Content-Type"] = "text/html; charset=\"utf-8\" "

// // Setup message

//   // Connect to the SMTP Server
//   servername:= "smtp.yandex.com:465"

//   host, _, _:= net.SplitHostPort(servername)

//   auth:= smtp.PlainAuth("oreos", "oreos.music@yandex.com", "wmhfjvddcpbtiduw", host)


// 	log.Println("===============================================")
// 	log.Println(body)
// 	log.Println("===============================================")
// 	log.Println("===============================================")
// 	log.Println("===============================================")


// 	  message:= ""
//   for k, v:= range headers {

//     message += fmt.Sprintf("%s: %s\r\n", k, v)
//   }
//   message += "\r\n" + body


//   log.Println()
//   log.Println("auth")

//   log.Println(auth)
//   log.Println("auth")
//   log.Println("auth")
//   // TLS config
//   tlsconfig:= & tls.Config {
//     InsecureSkipVerify: true,
//     ServerName: host,
//   }
//   log.Println("TLS")
//   // Here is the key, you need to call tls.Dial instead of smtp.Dial
//   // for smtp servers running on 465 that require an ssl connection
//   // from the very beginning (no starttls)
//   conn, err:= tls.Dial("tcp", servername, tlsconfig)
//   log.Println("TLS")
//   log.Printf("%v", conn)
//   if err != nil {
//     log.Println("TLS Error")
//     log.Println(err)
//     log.Println(err)
//   }


//   c, err:= smtp.NewClient(conn, host)
//   log.Println(c)
//   log.Println("Above is conn")
//   if err != nil {
//     log.Println("NewClient Error")
//     log.Println(err)
//     log.Println(err)
//     log.Println(err)
//   }
//   log.Println("Above is conn")
//   // Auth
//   if err = c.Auth(auth);
//   err != nil {
//     log.Println(err)
//   }
//   log.Println("Above is Auth")
//   log.Println(c)
//   // To && From
//   if err = c.Mail(from.Address);
//   err != nil {
//     log.Println(err)
//   }

//   if err = c.Rcpt(to.Address);
//   err != nil {
//     log.Println(err)
//   }

//   // Data
//   w, err:= c.Data()
//   log.Println("Above is W")
//   log.Println(w)
//   log.Println("Above is W")
//   log.Println("Above is W")
//   if err != nil {
//     log.Println(err)
//   }

//   _, err = w.Write([] byte(message))
//   if err != nil {
//     log.Println(err)
//   }

//   err = w.Close()
//   if err != nil {
//     log.Println(err)
//   }

//   c.Quit()


//   log.Println("Email Sent Successfully!")
//   return true

// }
