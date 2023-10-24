package tetris

import (
	"encoding/json"
	"fmt"

	"github.com/gdamore/tcell"
)

// loadBoards loads the internal boards
func loadBoards() error {
	var boardsJSON []BoardsJSON
	err := json.Unmarshal(boardsInternal, &boardsJSON)
	if err != nil {
		return fmt.Errorf("unmarshal error: %v", err)
	}

	numInternalBoards = len(boardsJSON)
	boards = make([]Boards, numInternalBoards)

	for boardNum, boardLoad := range boardsJSON {
		aBoards := Boards{name: boardLoad.Name}
		aBoards.colors = make([][]tcell.Color, len(boardLoad.Mino))
		aBoards.rotation = boardLoad.Rotation

		for i := 0; i < len(boardLoad.Mino); i++ {
			aBoards.colors[i] = make([]tcell.Color, len(boardLoad.Mino[i]))
			for j := 0; j < len(boardLoad.Mino[i]); j++ {
				switch boardLoad.Mino[i][j] {
				case "b":
					aBoards.colors[i][j] = colorBlank
				case "i":
					aBoards.colors[i][j] = colorCyan
				case "j":
					aBoards.colors[i][j] = colorBlue
				case "l":
					aBoards.colors[i][j] = colorWhite
				case "o":
					aBoards.colors[i][j] = colorYellow
				case "s":
					aBoards.colors[i][j] = colorGreen
				case "t":
					aBoards.colors[i][j] = colorMagenta
				case "z":
					aBoards.colors[i][j] = colorRed
				}
			}
		}
		boards[boardNum] = aBoards
	}

	boardsJSON = nil

	return nil
}

var boardsInternal = []byte(`
[
	{
		"name":"10 x 20 blank",
		"mino":[
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"]
		],
		"rotation":[
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
		]
	},
	{
		"name":"10 x 20 checkerboard double",
		"mino":[
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b"],
			["b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b"],
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b"],
			["b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b"],
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"]
		],
		"rotation":[
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
		]
	},
	{
		"name":"10 x 20 checkerboard single",
		"mino":[
			["b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"],
			["b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"],
			["b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"],
			["b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"],
			["b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"]
		],
		"rotation":[
			[0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
		]
	},
	{
		"name":"20 x 20 blank",
		"mino":[
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"]
		],
		"rotation":[
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
		]
	},
	{
		"name":"20 x 20 heart",
		"mino":[
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","l","l","l","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","z","z","t","l","o","o","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","z","z","t","t","t","o","o","j","b","b","b","b","b"],
			["b","b","b","b","b","b","b","l","l","l","i","i","i","i","j","b","b","b","b","b"],
			["b","b","b","b","b","b","b","s","s","l","l","l","z","z","j","j","b","b","b","b"],
			["b","b","b","b","b","b","b","b","s","s","l","z","z","i","i","i","i","b","b","b"],
			["b","b","b","b","b","b","b","b","j","j","l","s","s","t","t","t","z","z","b","b"],
			["b","b","b","b","b","b","b","b","s","j","o","o","s","s","t","z","z","b","b","b"],
			["b","b","b","b","b","b","b","s","s","j","o","o","o","o","l","l","b","b","b","b"],
			["b","b","b","b","b","b","b","s","z","z","s","s","o","o","l","b","b","b","b","b"],
			["b","b","b","b","b","b","b","z","z","t","j","s","s","j","l","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","t","t","j","j","j","j","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","t","j","j","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"]
		],
		"rotation":[
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,1,1,1,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,1,1,3,1,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,1,1,3,3,3,0,0,2,0,0,0,0,0],
			[0,0,0,0,0,0,0,1,1,1,3,3,3,3,2,0,0,0,0,0],
			[0,0,0,0,0,0,0,1,1,1,2,2,1,1,2,2,0,0,0,0],
			[0,0,0,0,0,0,0,0,1,1,2,1,1,1,1,1,1,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,2,1,1,1,1,1,1,1,0,0],
			[0,0,0,0,0,0,0,0,0,0,1,1,1,1,1,1,1,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,1,1,3,3,2,2,0,0,0,0],
			[0,0,0,0,0,0,0,0,1,1,1,1,3,3,2,0,0,0,0,0],
			[0,0,0,0,0,0,0,1,1,0,2,1,1,1,2,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,2,3,3,3,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,2,2,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
		]
	},
	{
		"name":"20 x 20 checkerboard double",
		"mino":[
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b"],
			["b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b"],
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b"],
			["b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b"],
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b"],
			["b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b"],
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b"],
			["b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b"],
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b"],
			["b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b"]
		],
		"rotation":[
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
		]
	},
	{
		"name":"20 x 20 checkerboard single",
		"mino":[
			["b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"],
			["b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"],
			["b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"],
			["b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"],
			["b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"],
			["b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"],
			["b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"],
			["b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"],
			["b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"],
			["b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"]
		],
		"rotation":[
			[0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
		]
	},
	{
		"name":"26 x 28 peace symbol",
		"Mino":[
			["b","b","b","b","b","b","b","b","b","b","b","b","j","i","i","i","i","o","o","l","l","j","j","j","z","j","j","i"],
			["b","b","b","b","b","b","b","b","j","l","j","j","j","s","j","j","j","o","o","l","z","j","l","l","z","z","j","i"],
			["b","b","b","b","b","b","b","b","j","l","l","l","s","s","j","b","b","b","b","l","z","z","l","o","o","z","j","i"],
			["b","b","b","b","b","b","b","t","j","j","z","z","s","b","b","b","j","l","b","b","b","z","l","o","o","j","z","i"],
			["b","b","b","b","b","b","t","t","t","z","z","b","b","b","j","j","j","l","l","l","b","b","b","j","j","j","z","z"],
			["b","b","b","b","b","b","i","i","i","i","b","b","j","z","s","s","t","t","t","z","z","s","b","b","t","t","t","z"],
			["b","b","b","b","b","b","t","t","t","b","b","s","j","z","z","s","s","t","z","z","s","s","b","b","b","t","l","l"],
			["b","b","b","b","b","b","b","t","b","b","s","s","j","j","z","j","j","j","o","o","s","b","b","s","b","b","l","i"],
			["b","b","b","b","b","b","z","z","b","l","s","i","i","i","i","j","s","s","o","o","b","b","s","s","j","b","l","i"],
			["b","b","b","b","b","z","z","b","b","l","o","o","l","l","t","t","t","s","s","b","b","z","s","b","j","b","b","i"],
			["b","b","b","b","b","l","l","b","l","l","o","o","l","o","o","t","z","z","b","b","t","z","z","l","j","j","b","i"],
			["b","b","b","b","b","l","b","b","i","i","i","i","l","o","o","z","z","b","b","t","t","t","z","l","l","l","b","b"],
			["b","b","b","b","b","l","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","j","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","j","b","b","o","o","l","l","t","t","t","z","z","b","b","i","i","i","i","j","o","o","b","b"],
			["b","b","b","b","b","j","j","b","o","o","l","s","s","t","z","z","s","s","b","b","s","s","b","j","o","o","b","l"],
			["b","b","b","b","b","b","z","b","b","z","l","i","s","s","l","l","l","s","s","b","b","s","s","j","j","b","b","l"],
			["b","b","b","b","b","b","z","z","b","z","z","i","j","j","j","i","l","j","j","j","b","b","t","t","t","b","l","l"],
			["b","b","b","b","b","b","l","z","b","b","z","i","j","o","o","i","s","j","o","o","z","b","b","t","b","b","j","j"],
			["b","b","b","b","b","b","l","l","l","b","b","i","l","o","o","i","s","s","o","o","z","z","b","b","b","o","o","j"],
			["b","b","b","b","b","b","i","i","i","i","b","b","l","l","l","i","j","s","t","t","t","z","b","b","t","o","o","j"],
			["b","b","b","b","b","b","b","b","t","s","s","b","b","b","j","j","j","z","z","t","b","b","b","t","t","t","z","z"],
			["b","b","b","b","b","b","b","b","t","t","s","s","l","b","b","b","z","z","b","b","b","i","i","i","i","z","z","i"],
			["b","b","b","b","b","b","b","b","t","o","o","z","l","l","l","b","b","b","b","j","j","j","t","t","t","j","j","i"],
			["b","b","b","b","b","b","b","b","b","o","o","z","z","j","j","j","l","l","l","j","t","s","s","t","o","o","j","i"],
			["b","b","b","b","b","b","b","b","i","i","i","i","z","j","i","i","i","i","l","t","t","t","s","s","o","o","j","i"]
		],
		"Rotation":[
			[0,0,0,0,0,0,0,0,0,0,0,0,3,3,3,3,3,0,0,2,2,1,1,1,0,0,0,2],
			[0,0,0,0,0,0,0,0,2,3,3,3,3,0,1,1,1,0,0,2,0,1,2,2,0,0,0,2],
			[0,0,0,0,0,0,0,0,2,3,3,3,0,0,1,0,0,0,0,2,0,0,2,1,1,0,0,2],
			[0,0,0,0,0,0,0,3,2,2,1,1,0,0,0,0,3,3,0,0,0,0,2,1,1,3,0,2],
			[0,0,0,0,0,0,3,3,3,1,1,0,0,0,3,3,3,3,3,3,0,0,0,3,3,3,0,0],
			[0,0,0,0,0,0,1,1,1,1,0,0,2,0,3,3,1,1,1,3,3,0,0,0,1,1,1,0],
			[0,0,0,0,0,0,1,1,1,0,0,2,2,0,0,3,3,1,3,3,0,0,0,0,0,1,2,2],
			[0,0,0,0,0,0,0,1,0,0,2,2,2,2,0,1,1,1,2,2,0,0,0,2,0,0,2,0],
			[0,0,0,0,0,0,3,3,0,0,2,3,3,3,3,1,1,1,2,2,0,0,2,2,2,0,2,0],
			[0,0,0,0,0,3,3,0,0,0,1,1,2,2,1,1,1,1,1,0,0,0,2,0,2,0,0,0],
			[0,0,0,0,0,2,2,0,0,0,1,1,2,3,3,1,3,3,0,0,3,0,0,3,2,2,0,0],
			[0,0,0,0,0,2,0,0,1,1,1,1,2,3,3,3,3,0,0,3,3,3,0,3,3,3,0,0],
			[0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,2,0,0,0,0,2,2,1,1,1,3,3,0,0,1,1,1,1,2,1,1,0,0],
			[0,0,0,0,0,2,2,0,0,0,2,1,1,1,3,3,1,1,0,0,3,3,0,2,1,1,0,0],
			[0,0,0,0,0,0,2,0,0,0,2,2,1,1,1,1,1,1,1,0,0,3,3,2,2,0,0,0],
			[0,0,0,0,0,0,2,2,0,0,0,2,1,1,1,0,1,1,1,1,0,0,1,1,1,0,0,0],
			[0,0,0,0,0,0,3,2,0,0,0,2,1,3,3,0,0,1,0,0,2,0,0,1,0,0,0,0],
			[0,0,0,0,0,0,3,3,3,0,0,2,3,3,3,0,0,0,0,0,2,2,0,0,0,2,2,0],
			[0,0,0,0,0,0,1,1,1,1,0,0,3,3,3,0,3,0,1,1,1,2,0,0,3,2,2,0],
			[0,0,0,0,0,0,0,0,2,3,3,0,0,0,3,3,3,3,3,1,0,0,0,3,3,3,1,1],
			[0,0,0,0,0,0,0,0,2,2,3,3,3,0,0,0,3,3,0,0,0,3,3,3,3,1,1,2],
			[0,0,0,0,0,0,0,0,2,1,1,2,3,3,3,0,0,0,0,1,1,1,1,1,1,0,0,2],
			[0,0,0,0,0,0,0,0,0,1,1,2,2,1,1,1,1,1,1,1,3,3,3,1,3,3,0,2],
			[0,0,0,0,0,0,0,0,3,3,3,3,2,1,3,3,3,3,1,3,3,3,3,3,3,3,0,2]
		]
	},
	{
		"name":"30 x 30 blank",
		"mino":[
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"],
			["b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b"]
		],
		"rotation":[
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
		]
	},
	{
		"name":"30 x 30 checkerboard double",
		"mino":[
			["b","b","b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b"],
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b"],
			["b","b","b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b"],
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b"],
			["b","b","b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b"],
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b"],
			["b","b","b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b"],
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b"],
			["b","b","b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b"],
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b"],
			["b","b","b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b"],
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b"],
			["b","b","b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b"],
			["b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b"],
			["b","b","b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"],
			["b","b","b","b","b","b","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z","b","b","z","z"]
		],
		"rotation":[
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
		]
	},
	{
		"name":"30 x 30 checkerboard single",
		"mino":[
			["b","b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"],
			["b","b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"],
			["b","b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"],
			["b","b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"],
			["b","b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"],
			["b","b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"],
			["b","b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"],
			["b","b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"],
			["b","b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"],
			["b","b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"],
			["b","b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"],
			["b","b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"],
			["b","b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"],
			["b","b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"],
			["b","b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z"],
			["b","b","b","b","b","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b","z","b"]
		],
		"rotation":[
			[0,0,0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
			[0,0,0,0,0,0,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2,0,2],
			[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
		]
	}
]
`)
