package main

type Soldier struct {
	health int
	damage int
}

func SimpleFight(first *Soldier, second *Soldier) {

	for first.health > 0 && second.health > 0 {
		Attack(first, second)
		Attack(second, first)
	}
}

func Attack(attacker *Soldier, victim *Soldier) {
	victim.health = victim.health - attacker.damage
}
