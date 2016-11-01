package lockutils

import "fmt"

const SPACE_SEP = '/'

type lockspace struct {
	_map *lmap
}

func NewSpace(_map *lmap) *lockspace {
	space := &lockspace{_map:_map}
	return space
}

func splitPath(path string) []string {
	spaces := make([]string, 0)

	for i, ch := range path {
		if ch == SPACE_SEP && i != 0 {
			space := path[:i]
			spaces = append(spaces, space)
		}
	}

	return spaces
}

func (ls *lockspace) RLock(path string) {
	spaces := splitPath(path)

	for space := range spaces {
		//ls._map.RLock(space)
		fmt.Println(space)
	}
}