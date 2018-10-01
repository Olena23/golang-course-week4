package main

import (

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"

	"fmt"
	"io/ioutil"
	"time"
	"strconv"
)

func main() {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello! This is an echo server"))
	})

	r.Route("/", func(r chi.Router) {
		r.Get("/ip", ipInfo)
		r.Get("/useragent", userAgent)
		r.Get("/headers", headerDict)
		r.Get("/get", getInfo)
		r.Post("/post", postInfo)
		r.Put("/put", putInfo)
		r.Delete("/delete", deleteInfo)
		r.Get("/response-headers", respHeaders)
		r.Get("/code/{code}", returnCode)
		r.Get("/stream/{lines}", streamLines)
		r.Get("/delay/{sec}", delayResponse)
		r.Get("/cookies/set/{name}/{value}", setCookie)
		r.Get("/cookies", getCookie)
	})
	http.ListenAndServe(":3333", r)
}


func ipInfo (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("USER IP: " + r.RemoteAddr))
	return
}

func userAgent (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("USER AGENT: " + r.UserAgent()))
	return
}

func headerDict (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header field %q, Value %q\n", k, v)
	}
	return
}

func getInfo (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	response := r.URL.RawQuery
	w.Write([]byte("Get parameters: " + response))
	return
}

func postInfo (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	r.ParseForm()
	for k, _ := range r.Form {
		fmt.Fprintf(w, "Key-value pairs: %q", k)
		return
	}
}

func putInfo (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	r.ParseForm()
	for k, _ := range r.Form {
		fmt.Fprintf(w, "Key-value pairs: %q", k)
		}
	return
}


func deleteInfo (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	defer r.Body.Close()
	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, "response Headers : %q, response Body: %q", r.Header, string(rBody))
	return
}

func returnCode (w http.ResponseWriter, r *http.Request) {
	receivedCode := chi.URLParam(r, "code")
	code, err := strconv.Atoi(receivedCode)
	if err == nil {
		w.WriteHeader(code)
		fmt.Fprintf(w, "Code: %d", code)
	} else {
		fmt.Fprintf(w, "Only numbers accepted")
	}
}

func respHeaders (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header field %q, Value %q\n", k, v)
	}
	response := r.URL.RawQuery
	w.Write([]byte("Get parameters: " + response))
	return
}

func setCookie (w http.ResponseWriter, r *http.Request) {
	cookieName := chi.URLParam(r, "name")
	fmt.Println(cookieName)
	cookieValue := chi.URLParam(r, "value")
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: cookieName, Value: cookieValue, Expires: expiration, Path: "/"}
	http.SetCookie(w, &cookie)
	for _, cookie := range r.Cookies() {
		fmt.Fprint(w, "Name: " + cookie.Name + ", value: " + cookie.Value + "\n")
	}
	return
}

func getCookie (w http.ResponseWriter, r *http.Request) {
	for _, cookie := range r.Cookies() {
		fmt.Fprint(w, "Name: " + cookie.Name + ", value: " + cookie.Value + "\n")
	}
	return
}

func streamLines (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	numberOfLines := chi.URLParam(r, "lines")
	numberToPrint, err := strconv.Atoi(numberOfLines)
	if err == nil {
		if numberToPrint <= 10 {
		for i := 0; i < numberToPrint; i++ {
		for k, v := range r.Header {
		fmt.Fprintf(w, "Header field %q, Value %q", k, v)
	}
		fmt.Fprintf(w, "\n \n")
	}} else {
		fmt.Fprintf(w, "To much lines")
	}
	} else {
		fmt.Fprintf(w, "Only numbers are accepted")
	}
	return
}

func delayResponse (w http.ResponseWriter, r *http.Request) {
	numberOfSeconds := chi.URLParam(r, "sec")
	waitTime, err := strconv.Atoi(numberOfSeconds)
	if err == nil {
		if waitTime <= 10 {
			time.Sleep(time.Duration(waitTime) * time.Second)
			w.WriteHeader(200)
			for k, v := range r.Header {
				fmt.Fprintf(w, "Header field %q, Value %q\n", k, v)
			}
		} else {
			fmt.Fprintf(w, "Only numbers below 10 are accepted")
		}
		}
	return
}
