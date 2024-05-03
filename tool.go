package hardcrash

import "net/http"

type Tool interface { //previously "xyz", this time of event driven queue manager can be used for any tool
	Reset()
	//Next() int64
	//Previous() int64
	GetChannels() []chan any
	GetGui() []byte
	GetGuiFunctions() map[string]func(w http.ResponseWriter, r *http.Request)
}

/*type XYZStruct struct {
	x    int64 `default:"0"`
	y    int64 `default:"0"`
	z    int64 `default:"0"`
	maxX int64
	maxY int64
	maxZ int64
}*/
