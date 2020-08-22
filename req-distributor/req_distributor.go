/*
3. Implement an HTTP service which accepts incoming requests and assigns the request to 1 of 3 buckets (A, B or C) based on the given distribution.
Write a unit test to verify your solution.
*/

package main

import (
	"fmt"
	"net/http"
)

type httpreq struct {
	r *http.Request
}

func (hr *httpreq) process() {
	// process request
}

func (hr *httpreq) close() {
	// cleanup
}

type bucketInfo struct {
	load  int
	reqCh chan *http.Request
}

func (b *bucketInfo) init() {
	b.load = 0                              // counter for load
	b.reqCh = make(chan *http.Request, 200) // each bucket is configured a buffered channel
	go handleReq(b.reqCh)                   // start req handler on each of the bucket
}

// distriute among three buckets, round robbin
func selectBucket(bInfo [3]*bucketInfo) int {
	var minLoad int
	var minIdx int

	// find the least loaded bucket
	for i, b := range bInfo {
		if i == 0 {
			minLoad = b.load
			minIdx = 0
		} else if b.load < minLoad {
			minLoad = b.load
			minIdx = i
		}
	}
	fmt.Printf("request comes up, distribute to bucket %d (%d,%d,%d)\n", minIdx,
		bInfo[0].load, bInfo[1].load, bInfo[2].load)
	bInfo[minIdx].load++
	return minIdx
}

func main() {
	fmt.Println("")

	var bInfo [3]*bucketInfo
	for i := 0; i < 3; i++ {
		bInfo[i] = &bucketInfo{}
		bInfo[i].init()
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Passed")

		minIdx := selectBucket(bInfo)
		bInfo[minIdx].reqCh <- r // send the request to the least loaded bucket
	})

	http.ListenAndServe(":12345", mux)
}

// request handler
func handleReq(reqCh chan *http.Request) {
	for {
		select {
		case r := <-reqCh:
			req := &httpreq{r}
			req.process()
			req.close()
		}
	}
}
