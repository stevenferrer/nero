package repository

import (
	"github.com/sf9v/nero/sort"
)

type SortFunc func(*sort.Sorts)

func Asc(col Column) SortFunc {
	return func(s *sort.Sorts) {
		s.Add(&sort.Sort{
			Col:       col.String(),
			Direction: sort.Asc,
		})
	}
}

func Desc(col Column) SortFunc {
	return func(s *sort.Sorts) {
		s.Add(&sort.Sort{
			Col:       col.String(),
			Direction: sort.Desc,
		})
	}
}
