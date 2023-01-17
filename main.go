package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	
)

func main (){
	http.HandleFunc("/home",func(w http.ResponseWriter, r *http.Request) {
	    t,err := template.ParseFiles("Home.html")
		if err != nil{
			log.Fatalln(err)
		}
		t.Execute(w,nil)
	})
	http.HandleFunc("/home/create/",func(w http.ResponseWriter, r *http.Request) {
	   t,err :=template.ParseFiles("create.html")
	   if err != nil{
		log.Fatalln(err)
	   }
	   t.Execute(w,nil)
	})

	http.HandleFunc("/home/UserList",func(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("user.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    type User struct{
		List []string
	}
	ul := User{}
    s := bufio.NewScanner(f)
    for s.Scan() {
        ul.List = append(ul.List,s.Text())
    }
	t,err := template.ParseFiles("UserList.html")
	if err != nil{
	 log.Fatalln(err)
	}
	t.Execute(w,ul)

	 })
	http.HandleFunc("/home/edit/",func(w http.ResponseWriter, r *http.Request) {
	    t,err := template.ParseFiles("edit.html")
		if err != nil{
			log.Fatalln(err)
		}
		t.Execute(w,nil)

	})
	

	http.HandleFunc("/home/create/store",func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm();err!=nil{
			log.Fatalln()
		}
		name := r.FormValue("name")
		ID := r.FormValue("ID")
		sub1 := r.FormValue("sub1")
		Mark1 := r.FormValue("mark1")
		sub2 := r.FormValue("sub2")
		Mark2 := r.FormValue("mark2")
		if name == "" || ID == "" ||sub1 == "" || Mark1 == "" || sub2 == "" || Mark2=="" {
			type ErrorMessage struct{
				Message string
			}
			t,err :=template.ParseFiles("create.html")
	        if err != nil{
		       log.Fatalln(err)
	        }
	        t .Execute(w,ErrorMessage{Message: "All filesds need to be filled"})
			return
		}else{
			type ErrorMessage struct{
				Message string
			}
			t,err :=template.ParseFiles("create.html")
	        if err != nil{
		       log.Fatalln(err)
	        }
	        t .Execute(w,ErrorMessage{Message: "Successfull !"})

		}
		f,err := os.OpenFile("user.txt",os.O_APPEND|os.O_WRONLY,0644)
		if err != nil{
			log.Fatalln(err)
		}
		defer f.Close()

		fmt.Fprintln(f,"Name : ",name)
		fmt.Fprintln(f,"ID : ",ID)
		fmt.Fprintln(f,"Subject : ",sub1)
		fmt.Fprintln(f,"Mark of subject 1 : ",Mark1)
		value1,err := strconv.Atoi(Mark1)
		if err != nil{
			log.Fatalln(err)
		}
		if value1 <=100 && value1 >= 80 {
			fmt.Fprintln(f,"Grade : A+")
		}else if value1<=79 && value1 >= 70{
			fmt.Fprintln(f,"Grade : A")
		}else if value1<=69 && value1>=60 {
			fmt.Fprintln(f,"Grade : A-")
		}else if value1<=59 && value1>=50 {
			fmt.Fprintln(f,"Grade : B")
		}else if value1<=49 && value1>=40 {
			fmt.Fprintln(f,"Grade : C")
		}else{
			fmt.Fprintln(f,"Grade : F")
		}
	    fmt.Fprintln(f,"Subject : ",sub2)
		fmt.Fprintln(f,"Mark of subject 2 : ",Mark2)
		value2,err := strconv.Atoi(Mark2)
		if err != nil{
			log.Fatalln(err)
		}
		if value2 <=100 && value2 >= 80 {
			fmt.Fprintln(f,"Grade : A+")
		}else if value2<=79 && value2 >= 70{
			fmt.Fprintln(f,"Grade : A")
		}else if value2<=69 && value2>=60 {
			fmt.Fprintln(f,"Grade : A-")
		}else if value2<=59 && value2>=50 {
			fmt.Fprintln(f,"Grade : B")
		}else if value2<=49 && value2>=40 {
			fmt.Fprintln(f,"Grade : C")
		}else{
			fmt.Fprintln(f,"Grade : F")
		}
		fmt.Fprintln(f,"------------------------------")
	
		http.Redirect(w,r,"/home/create/",http.StatusPermanentRedirect)
	 })
	
	if err := http.ListenAndServe(":3030",nil);err != nil{
		log.Fatalln(err)
	}
}