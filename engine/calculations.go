package engine

import "math"

func CalculateGravityForce(mass1 float32, mass2 float32, distance float32) float32 {
	return (G * mass1 * mass2) / (distance * distance * pixel * pixel)
}

func CalculateDistance(object1 coordinates, object2 coordinates) float32 {

	return float32(math.Sqrt(math.Pow((float64(object1.X)-float64(object2.X)), 2) + math.Pow((float64(object1.Y)-float64(object2.Y)), 2)))

}
