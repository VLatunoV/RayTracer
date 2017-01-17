package src

import "../util"

type Transformable interface {
	Transform(util.Transform)
	Reverse(util.Transform)
}