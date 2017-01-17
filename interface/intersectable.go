package src

import "../util"

type Intersectable interface {
	Intersect(*util.IntersectInfo)
}