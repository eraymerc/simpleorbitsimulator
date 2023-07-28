package engine

import (
	"fmt"
	"math"
)

//cisimler arası etkileşmeleri denetler

func Run(list *[]*Object, time float32) {

	for e := 0; e < len(*list); e++ {

		for b := 0; b < len(*list); b++ {

			dist := CalculateDistance((*list)[e].Location, (*list)[b].Location)
			if dist == 0 {
				continue
			}

			if (*list)[e].isPhysical {
				forceVal := CalculateGravityForce((*list)[e].Mass, (*list)[b].Mass, dist)

				force_x := forceVal * FindCosElementOfForce(&(*list)[e].Location, &(*list)[b].Location)
				force_y := forceVal * FindSinElementOfForce(&(*list)[e].Location, &(*list)[b].Location)
				force := Vector2D{force_x, force_y}
				fmt.Println(force)
				(*list)[e].ApplyForce(&force, time)
				(*list)[e].Move(time)
			}
		}
	}
}

func PythagoreanTheorem(targetObject *coordinates, object2 *coordinates) float32 {
	x := float64(object2.X - targetObject.X)
	y := float64(object2.Y - targetObject.Y)

	return float32(math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2)))
}

func FindCosElementOfForce(targetObject *coordinates, object2 *coordinates) float32 {
	y := object2.Y - targetObject.Y
	if y == 0 {
		return 1
	}

	return (object2.X - targetObject.X) / PythagoreanTheorem(targetObject, object2)
}
func FindSinElementOfForce(targetObject *coordinates, object2 *coordinates) float32 {

	x := object2.X - targetObject.X
	if x == 0 {
		return 1
	}

	return (object2.Y - targetObject.Y) / PythagoreanTheorem(targetObject, object2)
}
