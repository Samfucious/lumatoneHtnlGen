package lumatone

import (
	"encoding/hex"
	"fmt"
	"strings"
)

type Key struct {
	id      int
	Pitch   int
	Color   string
	Channel int
}

const (
	BoardKeyCount = 56
)

func (key *Key) SetColor(r byte, g byte, b byte) {
	bytes := []byte{r, g, b}
	key.Color = hex.EncodeToString(bytes)
}

func (key *Key) SetChannel(channel int) {
	key.Channel = channel
}

func (key *Key) GetId() int {
	return key.id
}

func (key *Key) WriteLtnt(sb *strings.Builder) {
	fmt.Fprintf(sb, "Key_%d=%d\n", key.id, key.Pitch)
	fmt.Fprintf(sb, "Col_%d=%s\n", key.id, key.Color)
	fmt.Fprintf(sb, "Chan_%d=%d\n", key.id, key.Channel)
}

type Board struct {
	id   int
	keys []*Key
}

func (b *Board) GetKey(id int) *Key {
	return b.keys[id]
}

func (b *Board) GetKeys() []*Key {
	return b.keys
}

func NewBoard(id int, channel int) *Board {
	board := &Board{
		id: id,
	}
	for i := 0; i < BoardKeyCount; i++ {
		board.keys = append(board.keys, &Key{
			id:      i,
			Color:   "A5A500",
			Channel: channel,
		})
	}
	return board
}

func (b *Board) WriteLtnt(sb *strings.Builder) {
	sb.WriteString(fmt.Sprintf("[Board%d]\n", b.id))
	for _, k := range b.keys {
		k.WriteLtnt(sb)
	}
}

type Lumatone struct {
	sections []*Board
}

func NewLumatone() *Lumatone {
	lt := &Lumatone{
		sections: make([]*Board, 5),
	}
	for i := 0; i < 5; i++ {
		lt.sections[i] = NewBoard(i, 1)
	}
	return lt
}

func (lt *Lumatone) GetSection(sectionId int) *Board {
	if sectionId < 0 || sectionId > 4 {
		return nil
	}
	return lt.sections[sectionId]
}

func (lt *Lumatone) WriteLtnt(sb *strings.Builder) {
	for _, s := range lt.sections {
		s.WriteLtnt(sb)
	}
}
