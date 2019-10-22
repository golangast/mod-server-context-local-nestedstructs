package Context

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"unsafe"

	"github.com/gorilla/securecookie"
	. "github.com/logrusorgru/aurora"
)

//function that actually gets the data
func AddContext(ctx context.Context, next http.Handler) http.Handler {
	var err error
	//wrapped
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
			log.Fatal(err)
		}

		log.Println(r.Method, "-", r.RequestURI)
		//everything after this point is used to grab data
		//#Response DATA
		bodyBuffer, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// Error occurred while parsing request body
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		//#Time
		start := time.Now()
		duration := time.Now().Sub(start)

		//#Cookie
		var hashKey = []byte("very-secret")
		var s = securecookie.New(hashKey, nil)
		encoded, err := s.Encode("cookie-name", "cookie-value")
		if err == nil {
			cookie := &http.Cookie{
				Name:  "cookie-name",
				Value: encoded,
				Path:  "/",
			}
			http.SetCookie(w, cookie)
			if cookie != nil {
				//Add data to context
				ctx := context.WithValue(r.Context(), cookie.Name, cookie.Value)
				next.ServeHTTP(w, r.WithContext(ctx))

			} else {
				next.ServeHTTP(w, r)
			}

			r.ParseForm()

			// 			//~~~~~~~~~database starting
			// 			db, err := sql.Open("mysql", "zendrulat:@/c9")
			// 			if err != nil {
			// 				log.Fatal(err, "didnt hit querymap")
			// 			}
			// 			defer db.Close()
			// 			err = db.Ping()
			// 			if err != nil {
			// 				log.Fatal(err)
			// 			}

			//~~~~~~~~~~~~~~~~~~db ended

			//#Logging
			fmt.Println(Blue("/ʕ◔ϖ◔ʔ/```````"))
			fmt.Printf("Host:%s - Addr:%s - Method:%v - URL:%s - PROTO:%s - Status:%s - Dur:%02d-00:00 - CookieName:%s - FormValue:%s - BodySize:%d - RequestBody:%d - Context:%s  \r\n",
				Cyan(r.Host),                          //url
				Magenta(r.RemoteAddr),                 //ip
				Brown(r.Method),                       //method request
				Red(r.RequestURI),                     //second url segment
				Green(r.Proto),                        //protocal
				Red(r.Header.Get("X-Forwarded-Port")), //status code

				Red(duration),                   //duration of request
				Brown(cookie.Name),              //cookie
				Magenta(r.Form),                 //form data
				Cyan(unsafe.Sizeof(bodyBuffer)), //size of content body
				Brown(bodyBuffer),               //request body
				//Red(db.Stats()),                 //database stats

				Red(r.WithContext(ctx)), //logging context gives really big logs https://golang.org/pkg/net/http
				//Magenta(r.Header.Get("User-Agent")),

			)

		}

	})
}
