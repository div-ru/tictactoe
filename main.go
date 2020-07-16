package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWIDTH  = 640
	screenHEIGHT = 640
	cellSIZE     = 213
)

var (
	tcrosses       rl.Texture2D
	tnoughts       rl.Texture2D
	startPosV      [2]rl.Vector2
	endPosV        [2]rl.Vector2
	startPosH      [2]rl.Vector2
	endPosH        [2]rl.Vector2
	rectMuteScreen rl.Rectangle
	mousePos       rl.Vector2
)

var (
	gameOver           bool = false
	draw               bool = false
	mouseButtonPressed bool = false
	playersMove        bool = false
	positions          byte = 0
)

type cell struct {
	t      rl.Texture2D
	rect   rl.Rectangle
	chtype byte
	marked bool
}

func main() {
	
	// init
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(screenWIDTH, screenHEIGHT, "Tic Tac Toe")
	rl.SetTargetFPS(60)
	rl.SetMouseScale(1.0, 1.0)

	// loading textures for 'x' and 'o'
	tcrosses = rl.LoadTexture("res/cross.png")
	tnoughts = rl.LoadTexture("res/nought.png")

	// initializing the game board
	cells := [3][3]cell{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			cells[i][j].chtype = ' '
			cells[i][j].marked = false
			cells[i][j].rect.Width = cellSIZE
			cells[i][j].rect.Height = cellSIZE
			cells[i][j].rect.X = float32(i) * cellSIZE
			cells[i][j].rect.Y = float32(j) * cellSIZE
		}
	}

	// drawing a grid on the board
	startPosV = [2]rl.Vector2{rl.NewVector2(screenWIDTH/3, 0), rl.NewVector2(screenWIDTH*2/3, 0)}
	endPosV = [2]rl.Vector2{rl.NewVector2(screenWIDTH/3, screenHEIGHT), rl.NewVector2(screenWIDTH*2/3, screenHEIGHT)}
	startPosH = [2]rl.Vector2{rl.NewVector2(0, screenHEIGHT/3), rl.NewVector2(0, screenHEIGHT*2/3)}
	endPosH = [2]rl.Vector2{rl.NewVector2(screenWIDTH, screenHEIGHT/3), rl.NewVector2(screenWIDTH, screenHEIGHT*2/3)}

	// game over screen
	rectMuteScreen = rl.NewRectangle(0, 0, screenWIDTH, screenHEIGHT)

	for !rl.WindowShouldClose() {
		// game logic
		if !gameOver && !draw {
			mousePos = rl.GetMousePosition()
			// if mouse button is pressed, we check whether an entry is empty or preoccupied by 'x' or 'o'
			// if it is empty, we mark the entry by our pieces: 'x' or 'o'
			if mouseButtonPressed {
				for i := 0; i < 3; i++ {
					for j := 0; j < 3; j++ {
						if mousePos.X > cells[i][j].rect.X &&
							mousePos.Y > cells[i][j].rect.Y &&
							mousePos.X < cells[i][j].rect.X+cellSIZE &&
							mousePos.Y < cells[i][j].rect.Y+cellSIZE && !cells[i][j].marked {

							if playersMove {
								cells[i][j].t = tcrosses
								cells[i][j].chtype = 'x'
							} else {
								cells[i][j].t = tnoughts
								cells[i][j].chtype = 'o'
							}

							cells[i][j].marked = true
							playersMove = !playersMove
							positions++
							break
						}
					}
				}

				mouseButtonPressed = false
			}

			// checking if game is over
			if cells[0][0].marked && cells[0][1].marked && cells[0][2].marked {
				if cells[0][0].chtype == cells[0][1].chtype && cells[0][1].chtype == cells[0][2].chtype {
					gameOver = true
				}
			}

			if cells[1][0].marked && cells[1][1].marked && cells[1][2].marked {
				if cells[1][0].chtype == cells[1][1].chtype && cells[1][1].chtype == cells[1][2].chtype {
					gameOver = true
				}
			}

			if cells[2][0].marked && cells[2][1].marked && cells[2][2].marked {
				if cells[2][0].chtype == cells[2][1].chtype && cells[2][1].chtype == cells[2][2].chtype {
					gameOver = true
				}
			}

			if cells[0][0].marked && cells[1][0].marked && cells[2][0].marked {
				if cells[0][0].chtype == cells[1][0].chtype && cells[1][0].chtype == cells[2][0].chtype {
					gameOver = true
				}
			}

			if cells[0][1].marked && cells[1][1].marked && cells[2][1].marked {
				if cells[0][1].chtype == cells[1][1].chtype && cells[1][1].chtype == cells[2][1].chtype {
					gameOver = true
				}
			}

			if cells[0][2].marked && cells[1][2].marked && cells[2][2].marked {
				if cells[0][2].chtype == cells[1][2].chtype && cells[1][2].chtype == cells[2][2].chtype {
					gameOver = true
				}
			}

			if cells[0][0].marked && cells[1][1].marked && cells[2][2].marked {
				if cells[0][0].chtype == cells[1][1].chtype && cells[1][1].chtype == cells[2][2].chtype {
					gameOver = true
				}
			}

			if cells[2][0].marked && cells[1][1].marked && cells[0][2].marked {
				if cells[2][0].chtype == cells[1][1].chtype && cells[1][1].chtype == cells[0][2].chtype {
					gameOver = true
				}
			}

			// if all positions are occupied, then it is a draw
			if positions >= 9 {
				draw = true
			}
		}

		// process events
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) || rl.IsMouseButtonPressed(rl.MouseRightButton) {
			mouseButtonPressed = true
		}

		// draw to screen
		rl.BeginDrawing()

		rl.ClearBackground(rl.NewColor(0, 179, 255, 255))

		for i := 0; i < 2; i++ {
			rl.DrawLineEx(startPosV[i], endPosV[i], 3, rl.White)
			rl.DrawLineEx(startPosH[i], endPosH[i], 3, rl.White)
		}

		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if cells[i][j].marked {
					rl.DrawTextureRec(cells[i][j].t, cells[i][j].rect, rl.NewVector2(cells[i][j].rect.X, cells[i][j].rect.Y), rl.White)
				}
			}
		}

		if gameOver {
			rl.DrawRectangleRec(rectMuteScreen, rl.NewColor(0, 0, 0, 160))
			rl.DrawText("You won!", screenWIDTH/6, screenHEIGHT*2/5, 100, rl.White)
		} else if draw {
			rl.DrawRectangleRec(rectMuteScreen, rl.NewColor(0, 0, 0, 160))
			rl.DrawText("It's a draw!", screenWIDTH/7, screenHEIGHT*2/5, 80, rl.White)
		}

		rl.EndDrawing()
	}

	rl.UnloadTexture(tnoughts)
	rl.UnloadTexture(tcrosses)
	rl.CloseWindow()
}
