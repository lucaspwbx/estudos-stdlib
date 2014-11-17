package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type earthMass float64
type au float64

type Planet struct {
	name     string
	mass     earthMass
	distance au
}

type By func(p1, p2 *Planet) bool

func (by By) Sort(planets []Planet) {
	ps := &planetSorter{
		planets: planets,
		by:      by,
	}
	sort.Sort(ps)
}

type planetSorter struct {
	planets []Planet
	by      func(p1, p2 *Planet) bool
}

func (s *planetSorter) Len() int {
	return len(s.planets)
}

func (s *planetSorter) Swap(i, j int) {
	s.planets[i], s.planets[j] = s.planets[j], s.planets[i]
}

func (s *planetSorter) Less(i, j int) bool {
	return s.by(&s.planets[i], &s.planets[j])
}

var planets = []Planet{
	{"Mercury", 0.055, 0.4},
	{"Venus", 0.815, 0.7},
	{"Earth", 1.0, 1.0},
	{"Mars", 0.107, 1.5},
}

func main() {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

	name := func(p1, p2 *Planet) bool {
		return p1.name < p2.name
	}

	mass := func(p1, p2 *Planet) bool {
		return p1.mass < p2.mass
	}

	distance := func(p1, p2 *Planet) bool {
		return p1.distance < p2.distance
	}

	decreasingDistance := func(p1, p2 *Planet) bool {
		return !distance(p1, p2)
	}

	By(name).Sort(planets)
	fmt.Println(planets)

	By(mass).Sort(planets)
	fmt.Println(planets)

	By(distance).Sort(planets)
	fmt.Println(planets)

	By(decreasingDistance).Sort(planets)
	fmt.Println(planets)
}
