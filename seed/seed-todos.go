package seed

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"task_manager/config/db"
	"time"

	"github.com/jackc/pgx/v5"
)

const TotalTodos = 1_000_000

func SeedTodos() {
	start := time.Now()

	var inserted atomic.Int64

	jobs := make(chan int)
	var wg sync.WaitGroup

	statuses := []string{
		"pending",
		"in_progress",
		"completed",
	}

	// Reuse the same description for every row.
	description := strings.Repeat("A", 50)

	// Progress Logger
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			count := inserted.Load()

			fmt.Printf(
				"\rInserted: %d/%d (%.2f%%)",
				count,
				TotalTodos,
				float64(count)*100/float64(TotalTodos),
			)

			if count >= TotalTodos {
				fmt.Println()
				return
			}
		}
	}()

	// Workers
	for i := 0; i < Workers; i++ {
		wg.Add(1)

		go func(workerID int) {
			defer wg.Done()

			// Each worker gets its own RNG.
			r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(workerID)))

			for startID := range jobs {

				baseTime := time.Now()
				rows := make([][]any, 0, BatchSize)

				end := min(startID+BatchSize, TotalTodos+1)

				for id := startID; id < end; id++ {

					days := r.Intn(211) - 90
					dueDate := baseTime.AddDate(0, 0, days)

					rows = append(rows, []any{
						"Todo 1M " + strconv.Itoa(id), // title
						r.Int63n(TotalUsers) + 1,      // user_id
						description,                   // description
						statuses[r.Intn(3)],           // status
						dueDate,                       // due_date
					})
				}

				_, err := db.DB.CopyFrom(
					db.CTX,
					pgx.Identifier{"todos"},
					[]string{
						"title",
						"user_id",
						"description",
						"status",
						"due_date",
					},
					pgx.CopyFromRows(rows),
				)

				if err != nil {
					log.Fatal(err)
				}

				inserted.Add(int64(len(rows)))
			}

		}(i)
	}

	// Queue jobs
	for i := 1; i <= TotalTodos; i += BatchSize {
		jobs <- i
	}

	close(jobs)

	wg.Wait()

	elapsed := time.Since(start)

	fmt.Printf("\n\nTodos Seed Complete\n")
	fmt.Printf("Inserted : %d\n", inserted.Load())
	fmt.Printf("Time     : %v\n", elapsed)
	fmt.Printf("Rows/sec : %.0f\n", float64(inserted.Load())/elapsed.Seconds())
}
