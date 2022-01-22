package main

type mushket struct {
	gun
}

func newMushket() iGun {
	return &mushket{
		gun: gun{
			name:  "Mushket",
			power: 2,
		},
	}
}
