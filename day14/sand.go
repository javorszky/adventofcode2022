package day14

import (
	"github.com/pkg/errors"
)

// I hate sand, it gets everywhere

const (
	moveRow int = 0b0000000000100000000000
	moveCol     = 0b0000000000000000000001
)

var (
	errAbyss   = errors.New("went into the abyss")
	errBlocked = errors.New("can't move anywhere else")
)

type sand struct {
	coord int
	world map[int]material
}

func newSand(world map[int]material) *sand {
	return &sand{
		coord: xyToBinary(0, 500),
		world: world,
	}
}

func moveDown(prev int) int {
	return prev + moveRow
}

func moveDownLeft(prev int) int {
	return prev + moveRow - moveCol
}

func moveDownRight(prev int) int {
	return prev + moveRow + moveCol
}

func (s *sand) findRestingPlace() (int, error) {
	var err error
	for {
		err = s.move()
		if errors.Is(err, errAbyss) {
			return 0, err
		}

		if err != nil {
			break
		}
	}

	return s.coord, nil
}

func (s *sand) move() error {
	down, ok := s.world[moveDown(s.coord)]
	if !ok {
		return errAbyss
	}

	if down == matAir {
		s.coord = moveDown(s.coord)
		return nil
	}

	downLeft, ok := s.world[moveDownLeft(s.coord)]
	if !ok {
		return errAbyss
	}

	if downLeft == matAir {
		s.coord = moveDownLeft(s.coord)
		return nil
	}

	downRight, ok := s.world[moveDownRight(s.coord)]
	if !ok {
		return errAbyss
	}

	if downRight == matAir {
		s.coord = moveDownRight(s.coord)
		return nil
	}

	return errBlocked
}
