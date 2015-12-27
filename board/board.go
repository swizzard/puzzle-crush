package board

import (
    "bytes"
    "errors"
    "fmt"
)


type Board struct {
    MaxX, MaxY uint64
    Tiles [][]Tile
}

func (board Board) String() string {
    var buf bytes.Buffer
    for _, row := range board.Tiles {
        fmt.Fprintf(&buf, "%v", row)
    }
    return buf.String()
}

func New (slice []Color, slices ...[]Color) Board {
    maxX := len(slice)
    maxY := len(slices) + 1
    slices = append([][]Color{slice}, slices...)
    var tslices [][]Tile
    for i := 0; i < maxY; i++ {
        tslice := make([]Tile, maxX)
        for j := 0; j < maxX; j++ {
            tslice = append(tslice, NewTile(slices[i][j]))
        }
        tslices = append(tslices, tslice)
    }
    return Board{uint64(maxX), uint64(maxY), tslices}
}

func (board Board) contains (coord Coord) bool {
    return (coord.X <= board.MaxX) && (coord.Y <= board.MaxY)
}

func (board Board) TileAt (coord Coord) (*Tile, error) {
   if board.contains(coord) {
       return &board.Tiles[coord.X][coord.Y], nil
   }
   return nil, errors.New("invalid coordinate")
}

func (board Board) GetSegment (segment Segment) (TileSegment, error) {
    var tsegment TileSegment
    for _, coord := range segment {
        if tile, err := board.TileAt(coord); err != nil {
            return TileSegment{}, err
        } else {
            tsegment = append(tsegment, *tile)
        }
    }
    return tsegment, nil
}

func (board Board) Swap (c1, c2 Coord) {
    if board.contains(c1) && board.contains(c2) {
        board.Tiles[c1.X][c1.Y], board.Tiles[c2.X][c2.Y] = board.Tiles[c2.X][c2.Y], board.Tiles[c1.X][c1.Y]
    }
}
