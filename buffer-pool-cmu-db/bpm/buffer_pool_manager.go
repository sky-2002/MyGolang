package bpm

//MaxPoolSize is the size of the pool
const MaxPoolSize = 4

//BufferPoolManager represents the buffer pool manager
type BufferPoolManager struct {
	diskManager DiskManager
	pages       [MaxPoolSize]*Page
	replacer    *ClockReplacer
	freeList    []FrameID
	pageTable   map[PageID]FrameID
}
