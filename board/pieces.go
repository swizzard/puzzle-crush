package board

import "fmt"


type Color uint8

var colors = []string{"black", "blue", "green", "red", "yellow", "purple", "white"}

func (c Color) String() string {
    return colors[c]
}

func (c1 Color) Matches(c2 Color) bool {
    if int(c1) == 7 || int(c2) == 7 {
        return true
    } else {
        return c1 == c2
    }
}


type Tile struct {
    Color
    Status string
}

func NewTile (color Color) Tile {
    return Tile{color, ""}
}

func (tile *Tile) SetStatus (newStatus string) {
    tile.Status = newStatus
}

func (tile *Tile) RemoveStatus() {
    tile.Status = ""
}

func (tile Tile) IsRegular() bool {
    return tile.Status == ""
}

func (tile Tile) String() string {
    return fmt.Sprintf("<Tile: %v (%s)>", tile.Color, tile.Status)
}

func (tile Tile) Matches(t2 Tile) bool {
    return tile.Color.Matches(t2.Color)
}


type TileSegment []Tile

func (tsegment TileSegment) IsMatch () bool {
    var start = tsegment[0]
    for _, tile := range tsegment[1:] {
        if !start.Matches(tile) {
            return false
        }
    }
    return true
}

