package main

import (
	"SellerApp/UserHandler"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kataras/golog"
)

func main() {
	fmt.Println("main")

	http.HandleFunc("/", f_defaultPage) //need to change to defaultPage to login

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources/"))))

	err := http.ListenAndServe(":9002", nil)
	if err != nil {
		log.Fatal("Listen and Serve:", err)
		golog.Error("Listen and Serve:", err)
	}

}

func f_defaultPage(w http.ResponseWriter, r *http.Request) {

	user := UserHandler.GetUserHandlerInstance()
	err := r.ParseForm()
	if err != nil {
		golog.Error("E", "Form parse error")
		return
	}
	fmt.Println("Form     :=", r.Form)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	fmt.Println("Body     :=", string(body))
	fmt.Println("path     :=  ", r.URL.Path)

	switch r.URL.Path {
	case "/transformuserdata":
		if r.Method == "POST" {
			profile := new(UserHandler.Profile)
			err := json.Unmarshal(body, &profile)
			if err != nil {
				golog.Println("body Unmarshal failed ", err)
				fmt.Fprintf(w, err.Error())
			} else {
				user.TransformUserData(*profile)
			}
		} else {
			fmt.Fprintf(w, "%s", http.StatusMethodNotAllowed)
		}
	case "/getuserdata":
		if r.Method == "POST" {
			profile := new(UserHandler.Profile)
			err := json.Unmarshal(body, &profile)
			if err != nil {
				golog.Println("body Unmarshal failed ", err)
				fmt.Fprintf(w, err.Error())
			} else {
				data, err := user.GetUserData(*profile)
				if err != nil {
					fmt.Fprintf(w, err.Error())
				} else {
					dataByte, err := json.Marshal(data)
					if err != nil {
						fmt.Fprintf(w, err.Error())
					} else {
						fmt.Fprintf(w, string(dataByte))
					}
				}
			}
		} else {
			fmt.Fprintf(w, "%s", http.StatusMethodNotAllowed)
		}

	default:
		fmt.Fprintf(w, "StatusMethodNotAllowed")
	}

}
