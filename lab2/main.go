package lab2

import (
	"pad/PADLabs/lab2/node"
	"pad/PADLabs/lab2/common"
)

func main()  {

}

func runNodes() {
	// 14400
	nodes := []int{14401, 14403}
	go node.Start([]common.Character {
		{
			"Undertaker",
			"Lorencia",
			96,
			30,
		},
	}, nodes, "14400")
	// 14401
	nodes = []int{14400, 14402}
	go node.Start([]common.Character {
		{
			"Warmashine",
			"Davias",
			51,
			211,
		},
	}, nodes, "14401")
	// 14402
	nodes = []int{14401, 14403}
	go node.Start([]common.Character {
		{
			"Mortarion",
			"Dungeon",
			17,
			98,
		},
	}, nodes, "14402")
	// 14403
	nodes = []int{14400, 14402}
	go node.Start([]common.Character {
		{
			"Boreus",
			"Davias",
			20,
			29,
		},
	}, nodes, "14403")
	// 14404
	nodes = []int{14405}
	go node.Start([]common.Character {
		{
			"Shooter",
			"Tarkan",
			178,
			17,
		},
	}, nodes, "14404")
	// 14405
	nodes = []int{14404}
	go node.Start([]common.Character {
		{
			"21Savage",
			"Atlanta",
			50,
			60,
		},
	}, nodes, "14405")
}
