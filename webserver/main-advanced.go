package main

import (
    "fmt"
    "log"
    "net/http"
    "context"
    "time"
    "os"
    "strconv"
    "math/rand"
    // Alternative library for generating distributions:
    // "gonum.org/v1/gonum/stat/distuv"

    "golang.org/x/sync/semaphore"
)

var sem *semaphore.Weighted

func Health(w http.ResponseWriter, r *http.Request) {

    aquired := sem.TryAcquire(1)

    // No slots available. 10% chance of failure. 90% chance of just waiting.
    if !aquired{
        if rand.Intn(10) > 1{
            sem.Acquire(context.Background(), 1)
        }else{
            http.Error(w, "Congested", http.StatusInternalServerError)
            return
        }
    }

    sleepRandTime()
    
    fmt.Fprintf(w, "I'm OK!")
    sem.Release(1)
}


func main() {
    // Get optional configurations
    listen := ":8080"
    if os.Getenv("LISTEN") != ""{
        listen = os.Getenv("LISTEN")
    }

    var err error
    semaphoreSlots := 200
    if os.Getenv("SEMAPHORE_SLOTS") != "" {
        if semaphoreSlots, err = strconv.Atoi(os.Getenv("SEMAPHORE_SLOTS")); err != nil{
            log.Fatalf("Could not convert %v to integer: %v", os.Getenv("SEMAPHORE_SLOTS"), err)
        }
    }

    // Initialize random seed
    rand.Seed(time.Now().UnixNano())

    // Initialize semaphore
    fmt.Printf("Creating semaphore with %v slots\n", semaphoreSlots)
    sem = semaphore.NewWeighted(int64(semaphoreSlots))

    // Configure and start webserver
    http.HandleFunc("/health", Health)
    fmt.Printf("Starting webserver on %v\n", listen)
    log.Fatal(http.ListenAndServe(listen, nil))
}

// Sleeps for a random time determined by a normal distribution
func sleepRandTime() {
    n := rand.NormFloat64() * 50 + 50

    if n < 0 {
        n = n * -1
    }

    time.Sleep(time.Duration(n)*time.Millisecond)
}