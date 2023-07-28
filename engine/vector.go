package engine

type Math interface {
	AddVectors2D()
}

type Vector2D struct {
	X float32
	Y float32
}

func AddVectors2D(vector1 *Vector2D, vector2 *Vector2D) Vector2D {
	return Vector2D{vector1.X + vector2.X, vector1.Y + vector2.Y}
}
