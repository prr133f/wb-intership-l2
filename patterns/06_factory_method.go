package patterns

import "errors"

type IVehicle interface {
	setType(typeOfVehicle string)
	getType() string
}

type Vehicle struct {
	typeOfVehicle string
}

func (v *Vehicle) setType(mark string) {
	v.typeOfVehicle = mark
}

func (v *Vehicle) getType() string {
	return v.typeOfVehicle
}

type Car struct {
	Vehicle
}

func NewCar() IVehicle {
	return &Car{
		Vehicle: Vehicle{
			typeOfVehicle: "Car",
		},
	}
}

type Ship struct {
	Vehicle
}

func NewShip() IVehicle {
	return &Ship{
		Vehicle: Vehicle{
			typeOfVehicle: "Ship",
		},
	}
}

type Plane struct {
	Vehicle
}

func NewPlane() IVehicle {
	return &Plane{
		Vehicle: Vehicle{
			typeOfVehicle: "Plane",
		},
	}
}

func VechileFactory(typeOfVehicle string) (IVehicle, error) {
	switch typeOfVehicle {
	case "Car":
		return NewCar(), nil
	case "Ship":
		return NewShip(), nil
	case "Plane":
		return NewPlane(), nil
	default:
		return nil, errors.New("Unknown type of vehicle")
	}
}
