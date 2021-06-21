package rest_test

import (
	"log"
	"net/http"
	"testing"
	"time"
)

func BenchmarkRestCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		http.Get("http://localhost:8080/pricing/eth/")
	}
}

func TestTimeFunc(t *testing.T) {
	start := time.Now()

	http.Get("http://localhost:8080/pricing/eth/")

	elapsed := time.Since(start)
	log.Println(elapsed)
}
