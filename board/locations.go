package board

import "errors"

type Coord struct {
    X, Y uint64
}


type Segment []Coord

func MakeRow (y uint64, length uint64) Segment {
    var row Segment
    var i uint64
    for i = 0; i < length; i++ {
        row = append(row, Coord{i, y})
    }
    return row
}

func MakeCol (x uint64, length uint64) Segment {
    var col Segment
    var i uint64
    for i = 0; i < length; i++ {
        col = append(col, Coord{x, i})
    }
    return col
}

func adjacent (c1, c2 Coord) bool {
    if c1 == c2 {
        panic("can't compare identical coordinates")
    } else {
        var diff int
        switch {
            case c1.X == c2.X:
               diff = int(c1.Y) - int(c2.Y)
            case c1.Y == c2.Y:
               diff = int(c1.X) - int(c2.X)
            }
        return diff == 1 || diff == -1
    }
}

func (segment Segment) IsValid () bool {
    for i := 0; i < len(segment) - 1; i++ {
        if adjacent(segment[i], segment[i+1]) {
            continue
        } else {
            return false
        }
    }
    return true
}

func MakeSegment (coords ...Coord) (Segment, error) {
    if Segment(coords).IsValid() {
        return coords, nil
    } else {
        return Segment{}, errors.New("invalid segment")
    }
}

