package engine

type Object struct {
	Name       string
	Type       string
	Mass       float32
	Radius     float32
	Velocity   Vector2D
	Location   coordinates
	isPhysical bool //eğer kapalıysa fizikten etkilenmez
}

func (object *Object) ApplyForce(force *Vector2D, time float32) {
	if object.isPhysical {
		object.Velocity.X += force.X * time / object.Mass
		object.Velocity.Y += force.Y * time / object.Mass
	}
}

func (object *Object) Move(time float32) {
	if object.isPhysical {
		object.Location.X += object.Velocity.X * time
		object.Location.Y += object.Velocity.Y * time
	}
}
func (object *Object) Physical(status bool) {
	object.isPhysical = status
}
