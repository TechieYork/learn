package Marsrover

var (
    DirectionNorth = "north"
    DirectionSouth = "south"
    DirectionEast = "east"
    DirectionWest = "west"
)

type Location struct {
    X int
    Y int
    Direction string
}

func (loc *Location) GetX() int {
    return loc.X
}

func (loc *Location) SetX(x int) {
   loc.X = x
}

func (loc *Location) GetY() int {
    return loc.Y
}

func (loc *Location) SetY(y int) {
    loc.Y = y
}

func (loc *Location) GetDirection() string {
    return loc.Direction
}

func (loc *Location) SetDirection(direction string) {
    loc.Direction = direction
}

type Marsrover struct {
    Loc *Location
}

func (rover *Marsrover) GetLocation() *Location {
    return rover.Loc
}

func (rover *Marsrover) GoForward() {
    rover.Loc.SetX(rover.Loc.GetX() + 1)
}
