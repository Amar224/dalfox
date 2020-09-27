package scanning

import (
	"net/http"
	"net/url"
	"sync"

	"github.com/hahwul/dalfox/pkg/model"
	"github.com/hahwul/dalfox/pkg/optimization"
)

// SSTIAnalysis is basic check for SSTI
func SSTIAnalysis(target string, options model.Options) {
	// Build-in Grepping payload :: SSTI
	// {444*6664}
	// 2958816
	bpu, _ := url.Parse(target)
	bpd := bpu.Query()
	var wg sync.WaitGroup
	concurrency := options.Concurrence
	reqs := make(chan *http.Request)

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(){
			for req := range reqs {
				SendReq(req, "toGrepping", options)
			}
			wg.Done()
		}()
	}

	for bpk := range bpd {
		for _, ssti := range getSSTIPayload() {
			turl, _ := optimization.MakeRequestQuery(target, bpk, ssti, "toGrepping", options)
			reqs <- turl
		}
	}
	close(reqs)
	wg.Wait()
}

//SqliAnalysis is basic check for SQL Injection
func SqliAnalysis(target string, options model.Options) {
	// sqli payload
	bpu, _ := url.Parse(target)
	bpd := bpu.Query()
	var wg sync.WaitGroup
	concurrency := options.Concurrence
	reqs := make(chan *http.Request)

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(){
			for req := range reqs {
				SendReq(req, "toGrepping", options)
			}
			wg.Done()
		}()
	}

	for bpk := range bpd {
		for _, sqlipayload := range getSQLIPayload() {
			turl, _ := optimization.MakeRequestQuery(target, bpk, sqlipayload, "toGrepping", options)
			reqs <- turl
		}
	}
	close(reqs)
	wg.Wait()
}