package hmap

const seed uint32 = 131 // 31 131 1313 13131 131313 etc..
const SHARD_COUNT = 10

//bkd encryption
func bkdrHash(str string) uint32 {
	var h uint32

	for _, c := range str {
		h = h*seed + uint32(c)
	}

	return h
}

// hash table structure designed.
type ShardMap struct {
	shardCount uint8 //hash table number
	shards     []*SafeMap
}

func Shard(args ...uint8) *ShardMap {
	shm := &ShardMap{}
	if len(args) >= 1 && args[0] > 0 {
		shm.shardCount = args[0]
	} else {
		shm.shardCount = SHARD_COUNT
	}

	shm.shards = make([]*SafeMap, shm.shardCount)
	return shm
}

// locate the current safemap.
func (shm *ShardMap) Find(key string) *SafeMap {
	return shm.shards[shm.IndexBy(key)]
}

func (shm *ShardMap) IndexBy(key string) uint8 {
	return uint8(bkdrHash(key) & uint32(shm.shardCount-1))
}

//Alias of find
func (shm *ShardMap) Locate(key string) *SafeMap {
	return shm.Find(key)
}

func (shm *ShardMap) Set(key string, value interface{}) error {
	sm := shm.Find(key)

	if sm == nil {
		sm = Make()
		shm.shards[shm.IndexBy(key)] = sm
	}
	sm.Set(key, value)
	return nil
}

func (shm *ShardMap) Get(key string) interface{} {
	ret, err := shm.Find(key).Get(key)

	if err != nil {
		return nil
	}
	return ret
}

var (
	// Alias of Shard
	New = Shard
)
