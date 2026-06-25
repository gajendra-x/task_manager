package seed

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"task_manager/config/db"
	"time"

	"github.com/jackc/pgx/v5"
)

const (
	TotalUsers = 1_000_000
	BatchSize  = 5_000
	Workers    = 10
)

func SeedUsers() {
	start := time.Now()

	var inserted atomic.Int64

	jobs := make(chan int)

	var wg sync.WaitGroup

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			count := inserted.Load()

			fmt.Printf(
				"\rInserted: %d/%d (%.2f%%)",
				count,
				TotalUsers,
				float64(count)*100/float64(TotalUsers),
			)

			if count >= TotalUsers {
				fmt.Println()
				return
			}
		}
	}()

	for i := 0; i < Workers; i++ {
		wg.Add(1)

		go func(workerID int) {
			defer wg.Done()

			for startID := range jobs {

				rows := make([][]any, 0, BatchSize)

				end := min(startID+BatchSize, TotalUsers+1)

				for id := startID; id < end; id++ {
					rows = append(rows, []any{
						fmt.Sprintf("User %d", id),
						fmt.Sprintf("user%d@example.com", id),
					})
				}

				_, err := db.DB.CopyFrom(
					db.CTX,
					pgx.Identifier{"users"},
					[]string{
						"name",
						"email",
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

	for i := 1; i <= TotalUsers; i += BatchSize {
		jobs <- i
	}

	close(jobs)

	wg.Wait()

	elapsed := time.Since(start)

	fmt.Printf("\n\nUsers Seed Complete\n")
	fmt.Printf("Inserted : %d\n", inserted.Load())
	fmt.Printf("Time     : %v\n", elapsed)
	fmt.Printf("Rows/sec : %.0f\n", float64(inserted.Load())/elapsed.Seconds())
}
