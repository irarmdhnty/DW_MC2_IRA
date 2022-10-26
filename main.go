package main

import (
	"fmt"
)

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func main() {
	profile := MakeProfile("Naruto")
	fmt.Println("Name :", profile.Name)
	fmt.Println("Health :", profile.Health)
	fmt.Println("Power :", profile.Power)
	fmt.Println("Exp :", profile.Exp)
	fmt.Println("===HEROES POWER UP===")
	powerUp := PowerUp(profile, 4)
	fmt.Println("Name :", powerUp.Name)
	fmt.Println("Health :", powerUp.Health)
	fmt.Println("Power :", powerUp.Power)
	fmt.Println("Exp :", powerUp.Exp)

}

func MakeProfile(inputName string) Profile {
	var ProfileDetail Profile
	ProfileDetail.Name = inputName
	if inputName == "Sasuke" {
		ProfileDetail.Health = 200
		ProfileDetail.Power = 100
		ProfileDetail.Exp = 0
	} else if inputName == "Goku" {
		ProfileDetail.Health = 400
		ProfileDetail.Power = 300
		ProfileDetail.Exp = 100
	} else if inputName == "Naruto" {
		ProfileDetail.Health = 150
		ProfileDetail.Power = 200
		ProfileDetail.Exp = 50
	}

	return ProfileDetail
}

func PowerUp(profile Profile, multiplier int) Profile {
	var Multiplier Profile
	Multiplier.Name = profile.Name
	if profile.Name == "Sasuke" {
		Multiplier.Health = 200 + (200 * multiplier)
		Multiplier.Power = 100 + (100 * multiplier)
		Multiplier.Exp = 0 + (0 * multiplier)
	} else if profile.Name == "Goku" {
		Multiplier.Health = 400 + (400 * multiplier)
		Multiplier.Power = 300 + (300 * multiplier)
		Multiplier.Exp = 100 + (100 * multiplier)
	} else if profile.Name == "Naruto" {
		Multiplier.Health = 150 + (150 * multiplier)
		Multiplier.Power = 200 + (200 * multiplier)
		Multiplier.Exp = 50 + (50 * multiplier)
	}

	return Multiplier
}