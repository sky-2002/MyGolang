package bpm

type PageID int

const pageSize = 5

type Page struct {
	id       PageID         // id of the page, unique to each page
	pinCount int            // tracks the number of concurrent accesses
	isDirty  bool           // checks if the page has been modified after it was read
	data     [pageSize]byte // data stored in the form of bytes
}

// PinCount retunds the pin count
func (p *Page) PinCount() int {
	return p.pinCount
}

// ID retunds the page id
func (p *Page) ID() PageID {
	return p.id
}

// DecPinCount decrements pin count
func (p *Page) DecPinCount() {
	if p.pinCount > 0 {
		p.pinCount--
	}
}
