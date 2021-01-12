package controllers

import (
    "net/http"
)

// GetEmptyPage returns predicted response if route not found
func GetEmptyPage(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(GetErrorJSON("Page not found")))
}
