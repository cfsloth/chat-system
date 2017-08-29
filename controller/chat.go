package controller

import(
	"fmt"
	"net/http"
	"strings"
)

func GetIP(w http.ResponseWriter ,r *http.Request){
	ip := strings.Split(r.RemoteAddr,":")[0]
	fmt.Println(ip)
}	
