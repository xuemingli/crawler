package model

type SearchResult struct {
	Hits      int64
	Start     int
	Query     string
	PrevFrom  int
	NextFrom  int
	PageTh    int
	PageTotal int64
	Items     []interface{}
}
