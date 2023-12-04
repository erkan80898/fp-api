package main

//sm -> 612515
//fol -> 570935
//sns -> 47537
//ab -> 48226
import (
	Lib "flx/lib"
	Mod "flx/model"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
)

const POOLLIMIT = 40
const WORKERS = POOLLIMIT

type Tokens struct {
	sources  []string
	channels []string
}

func gatherTokens() Tokens {
	return Tokens{
		sources: []string{"AB", "SS", "SM"},
	}
}

func main() {
	CountSourceVariant()
}

func CountSourceVariant() {
	tokens := gatherTokens()
	sourceTokens := []string{os.Getenv("FLX_AB_TOKEN"), os.Getenv("FLX_SS_TOKEN"), os.Getenv("FLX_SM_TOKEN")}
	for i := 0; i < len(sourceTokens); i++ {
		var wg sync.WaitGroup
		var ops atomic.Uint64

		wg.Add(WORKERS)

		for j := 1; j <= WORKERS; j++ {
			go ConcurrentCount(j, &wg, &ops, sourceTokens[i])
		}

		wg.Wait()
		fmt.Printf("%s count: %d\n", tokens.sources[i], ops.Load())
	}
}

func ConcurrentCount(step int, wg *sync.WaitGroup, ops *atomic.Uint64, token string) {
	count := 0
	query := Mod.GetInventoryVariant{Page: step, PageSize: 100}

	for count = len(Lib.GetDataList(Mod.GET_INVENTORY_VARIANTS_PATH+Mod.QueryUrl(query), token)); count != 0; count = len(Lib.GetDataList(Mod.GET_INVENTORY_VARIANTS_PATH+Mod.QueryUrl(query), token)) {
		ops.Add(uint64(count))
		query.Page += WORKERS
	}
	wg.Done()
}
