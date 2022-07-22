package bpm

//DiskMaxNumPages sets the disk capacity
const DiskMaxNumPages = 15

type DiskManager interface {
	Readpage(PageID) (*Page, error)
	WritePage(*Page) error
	AllocatePage() *PageID
	DeallocatePage(PageID)
}
