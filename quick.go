package main

import (
    "fmt"
    "math/rand"
    "time"
)

const DATASET_LEN = 100000000
const DATASET_RANGE = 1000000

func gen_dataset() []int {
    arr := [DATASET_LEN]int{}
    for i := 0; i < DATASET_LEN; i++ {
        arr[i] = rand.Intn(DATASET_RANGE)
    }
    return arr[:]
}

func quick(dataset []int) chan struct{} {
    done := make(chan struct{})

    if len(dataset) <= 1 {
        close(done)
        return done
    }

    go func() {
        defer close(done)
        p1 := 0
        p2 := len(dataset) - 1
        for p1 != p2 {
            if (dataset[p1] > dataset[p2] && p1 < p2) || (dataset[p1] < dataset[p2] && p1 > p2) {
                dataset[p1], dataset[p2] = dataset[p2], dataset[p1]
                p1, p2 = p2, p1
            }
            if p1 < p2 {
                p1 += 1
            } else {
                p1 -= 1
            }
        }
        doneleft := quick(dataset[0:p1])
        <-quick(dataset[p1+1:])
        <-doneleft
    }()

    return done
}

func main() {
    dataset := gen_dataset()

    //fmt.Printf("DATASET: %v\n\n", dataset)

    startt := time.Now()
    <-quick(dataset)
    endt := time.Now()

    //fmt.Printf("SORTED DATASET: %v\n\n", dataset)

    fmt.Printf("Sorted %d values ranging from 0 to %d in %s\n", DATASET_LEN, DATASET_RANGE, endt.Sub(startt).String())
}

func init() {
    rand.Seed(time.Now().UnixNano())
}
