package main

import (
    "testing"

    . "github.com/smartystreets/goconvey/convey"

    "github.com/TechieYork/learn/golang/libraries/test_tdd/Marsrover"
)

// Marsrover test

func TestMarsrover(t *testing.T) {
    // Test convey
    Convey("Test Marsrover", t, func() {
            Convey("Test Single Command", func() {
                    Convey("Test FORWARD", func() {
                        marsrover := &Marsrover.Marsrover{
                            Loc: &Marsrover.Location{
                                X:         0,
                                Y:         0,
                                Direction: Marsrover.DirectionNorth,
                            },
                        }

                        originalLoc := marsrover.GetLocation()
                        marsrover.GoForward()
                        currentLoc := marsrover.GetLocation()

                        ShouldEqual(originalLoc.GetX() + 1, currentLoc.GetX())
                        ShouldEqual(originalLoc.GetDirection(), currentLoc.GetDirection())
                    })
                },
            )
        },
    )
}