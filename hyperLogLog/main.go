package main

import (
	"bufio"
	"fmt"
	"github.com/dchest/siphash"
	"math/rand"
	"os"
	"time"
)

func main() {

	//dataGen()
	//return

	fmt.Println("Test with 1000000 random uint64 numbers")
	s, err := SizeByP(16)
	if err != nil {
		panic(err)
	}
	h := make(HLL, s)
	fi, err := os.Open("dataset56.txt")
	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		h.Add(siphash.Hash(2, 57, []byte(scanner.Text())))
	}
	est := h.EstimateCardinality()

	pres := float64(1000000 - est) / 1000000 / 100

	fmt.Printf("count by hll algo = %d\n%% of precision = %f", est, pres)
	fmt.Println("%")
	_ = fi.Close()

	fmt.Println("Test with 61766393 random uint64 numbers")
	s, err = SizeByP(16)
	if err != nil {
		panic(err)
	}
	h = make(HLL, s)
	fi, err = os.Open("dataset35.txt")
	scanner = bufio.NewScanner(fi)
	for scanner.Scan() {
		h.Add(siphash.Hash(2, 57, []byte(scanner.Text())))
	}
	est = h.EstimateCardinality()

	pres = float64(61766393 - est) / 61766393 / 100

	fmt.Printf("count by hll algo = %d\n%% of precision = %f", est, pres)
	fmt.Println("%")
	_ = fi.Close()
}

func dataGen () {
	fo, _ := os.Create(fmt.Sprintf("dataset%d.txt", time.Now().Second()))
	for i := 0; i < 10000000; i++ {
		_,_ = fo.WriteString(fmt.Sprintf("%d\n", rand.Uint64()))
	}
	_ = fo.Close()
}