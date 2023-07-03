package htnl // Package htnl - Harmonic Table Note Layout

import (
	"encoding/json"
	"fmt"
	"htnlgen/pkg/lumatone"
	"image/color"
)

type Layout struct {
	Board        int    `json:"Board"`
	Key0Pitch    int    `json:"Key0Pitch"`
	Channel      int    `json:"Channel"`
	ColorDefault string `json:"ColorDefault"`
	ColorEdge    string `json:"ColorEdge"`
	ColorC       string `json:"ColorC"`
	ColorMiddleC string `json:"ColorMiddleC"`
}

type LayoutEntryList []*Layout

type colors struct {
	ColorDefault color.RGBA
	ColorEdge    color.RGBA
	ColorC       color.RGBA
	ColorMiddleC color.RGBA
}

func (layout *Layout) SetPitchesForSection(board *lumatone.Board) {
	keyIdOffsets := []int{0, 2, 7, 13, 19, 25, 31, 37, 43, 49, 54}                 // hardware layout - the keys
	keyCountInRow := []int{2, 5, 6, 6, 6, 6, 6, 6, 6, 5, 2}                        // hardware layout - the rows
	pitchValueOffsets := []int{0, -3, -10, -13, -20, -23, -30, -33, -40, -39, -34} // The harmonic table note layout diff per row.

	for i := 0; i < len(keyIdOffsets); i++ {
		for j := 0; j < keyCountInRow[i]; j++ {
			idx := keyIdOffsets[i] + j
			key := board.GetKey(idx)
			key.Pitch = layout.Key0Pitch + (pitchValueOffsets[i] + (j * 4))
		}
	}
}

func (layout *Layout) SetChannelForSection(board *lumatone.Board) {
	for _, k := range board.GetKeys() {
		k.Channel = layout.Channel
	}
}

func (layout *Layout) SetColorsForSection(board *lumatone.Board) {
	c := &colors{
		ColorDefault: hexToColor(layout.ColorDefault),
		ColorEdge:    hexToColor(layout.ColorEdge),
		ColorC:       hexToColor(layout.ColorC),
		ColorMiddleC: hexToColor(layout.ColorMiddleC),
	}

	for i := 0; i < lumatone.BoardKeyCount; i++ {
		key := board.GetKey(i)
		p := key.Pitch

		var color color.RGBA
		if p%12 == 0 {
			if p == 60 {
				color = c.ColorMiddleC
			} else {
				color = c.ColorC
			}
		} else if p%4 == 0 || p%3 == 0 {
			color = c.ColorEdge
		} else {
			color = c.ColorDefault
		}
		key.SetColor(color.R, color.G, color.B)
	}
}

func hexToColor(hex string) color.RGBA {
	var color color.RGBA
	_, err := fmt.Sscanf(hex, "%02x%02x%02x", &color.R, &color.G, &color.B)
	if err != nil {
		panic(err)
	}
	return color
}

func UnmarshalJSON(data []byte) (*LayoutEntryList, error) {
	list := &LayoutEntryList{}

	if err := json.Unmarshal(data, list); err != nil {
		return nil, err
	}

	return list, nil
}
