package main

import (

    "log"
    "net/http"
    "encoding/json"
    "./model"
    "io/ioutil"
)

func main() {
    http.HandleFunc("/ips/", func(w http.ResponseWriter, r *http.Request) {
        if(r.Method=="GET"){
            listIp := model.Ips{
            model.Ip{Destination: "google.fr"},
            model.Ip{Destination: "youtube.fr"},
            model.Ip{Destination: "bing.fr"},
            model.Ip{Destination: "szszdzdadafdff.fr"},
            }
        json.NewEncoder(w).Encode(listIp)           
        }
        if(r.Method=="POST"){
            newDestination := model.Ip{}
            body, err := ioutil.ReadAll(r.Body)
            r.Body.Close()
            if(err!=nil){
                log.Fatal(err)
            } 
            err = json.Unmarshal(body, &newDestination)
            if(err!=nil){
                log.Fatal(err)
            }
            listIp := model.Ips{
            model.Ip{Destination: "google.fr"},
            model.Ip{Destination: "youtube.fr"},
            model.Ip{Destination: "bing.fr"},
            model.Ip{Destination: "szszdzdadafdff.fr"},
            model.Ip{Destination: newDestination.Destination},
            }
        json.NewEncoder(w).Encode(listIp) 
        }

    })
        http.HandleFunc("/ips/results", func(w http.ResponseWriter, r *http.Request) {
        if(r.Method=="GET"){      
        }
        if(r.Method=="POST"){
            newDestination := model.Ip{}
            body, err := ioutil.ReadAll(r.Body)
            r.Body.Close()
            if(err!=nil){
                log.Fatal(err)
            } 
            err = json.Unmarshal(body, &newDestination)
            if(err!=nil){
                log.Fatal(err)
            }
            listIp := model.Ips{
            model.Ip{Destination: newDestination.Destination},
            }
        json.NewEncoder(w).Encode(listIp) 
        }
    })
    log.Fatal(http.ListenAndServe(":8000", nil))
}
