package utils

import "github.com/sathishkumar64/adv_golang/type_exp/old"

// IsMajor return whether a given person is engough
func IsMajor(p old.Person) bool {

	return p.Age > 18
}
