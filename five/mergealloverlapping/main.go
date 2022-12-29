package main

func main() {
	zoo := map[string]struct{}{}

	// We can add members to the set
	// by setting the value of each key to an
	// empty struct
	zoo["elephant"] = struct{}{}
	zoo["tiger"] = struct{}{}
	zoo["owl"] = struct{}{}
	zoo["lion"] = struct{}{}

	// Adding a new member to the set
	zoo["Bear"] = struct{}{}
}
