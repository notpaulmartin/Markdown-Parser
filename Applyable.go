package main

import "mdParser/Parse"

type Applyable interface {
	apply() (bool, Parse.Tag)
}
