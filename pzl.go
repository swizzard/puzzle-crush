package main

import (
    // "rand"
    "github.com/swizzard/util"
    "github.com/swizzard/puzzle-crush/board"
)




func main() {
    c1 := board.Coord{0,0}
    c2 := board.Coord{0,1}
    c3 := board.Coord{0,2}
    // c4 := board.Coord{1,1}
    // c5 := board.Coord{1,2}

    theBoard := board.New([]board.Color{0,1,2}, []board.Color{0,2,3}, []board.Color{0,1,3})
    tile, _ := theBoard.TileAt(c1)
    tile.SetStatus("foo")
    tile2, _ := theBoard.TileAt(c1)
    util.Printfn("%v == %v: %t", tile, tile2, tile == tile2)
    util.Printfn("tile at 1, 1 is %v", tile)
    seg := board.Segment{c1, c2, c3}
    tsg, _ := theBoard.GetSegment(seg)
    util.Printfn("tsg: %v, match: %t", tsg, tsg.IsMatch())
    theBoard.Swap(c2, c1)
    util.Printfn("after swap: %v", theBoard)
}

