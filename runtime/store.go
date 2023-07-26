package runtime

import (
	"regexp"
	"sync"

	"github.com/therealgauravr/pocketbase-clone-spa/types"
)

type GlobalStore struct {
	mu   sync.Mutex
	data map[string]interface{}
}

var gs *GlobalStore

func RetrieveGlobalStore() *GlobalStore {
	if gs == nil {
		gs = new(GlobalStore)
	}
	return gs
}

func (gs *GlobalStore) Init() {
	gs.mu.Lock()
	gs.data = make(map[string]interface{}, 2000)
	gs.data["api_key"] = ""
	gs.data["users"] = make(map[string]types.User, 3)
	gs.data["apiProtectedURLs"] = []*regexp.Regexp{
		regexp.MustCompile("^/api"),
	}
	gs.data["basicProtectedURLs"] = []*regexp.Regexp{
		regexp.MustCompile("^/session"),
	}
	gs.mu.Unlock()
}

func (gs *GlobalStore) Set(key string, value interface{}) {
	gs.mu.Lock()
	gs.data[key] = value
	gs.mu.Unlock()

}

func (gs *GlobalStore) Get(key string) interface{} {
	if _, ok := gs.data[key]; ok {
		return gs.data[key]
	}
	return nil
}
