package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/time/rate"
)

type limitersByIP struct {
	limiters sync.Map
	rate     rate.Limit
}

func NewLimitersByIP(r rate.Limit, b int, allowedIPs []string) *limitersByIP {
	concurrentMap := sync.Map{}
	for _, ip := range allowedIPs {
		concurrentMap.Store(ip, rate.NewLimiter(r, b))
	}
	return &limitersByIP{
		limiters: concurrentMap,
		rate:     r,
	}
}

func limitedHandler(lmts *limitersByIP, nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var clmt *rate.Limiter

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err.Error())
			w.Write([]byte(err.Error()))

		}
		ip = strings.TrimSpace(ip)

		lmt, ok := lmts.limiters.Load(ip)
		if !ok || lmt == nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Error when getting limiter from map.")
			return
		}
		clmt = lmt.(*rate.Limiter)

		w.Header().Add("X-Rate-Limit-Limit", fmt.Sprintf("%.2f", lmts.rate))
		w.Header().Add("X-Rate-Limit-Duration", "1")
		w.Header().Add("X-Rate-Limit-Request-Forwarded-For", r.Header.Get("X-Forwarded-For"))
		w.Header().Add("X-Rate-Limit-Request-Remote-Addr", r.RemoteAddr)

		canDoRequest := clmt.Allow()
		if canDoRequest {
			nextFunc.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusTooManyRequests)
		}
	}
}

func main() {

	lbyip := NewLimitersByIP(10, 5, []string{"127.0.0.1"})

	handlerWithLimiter := limitedHandler(lbyip, handler)
	http.Handle("/", handlerWithLimiter)

	log.Println("Starting Server on Port 9000.")
	http.ListenAndServe("localhost:9000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello!")
}
