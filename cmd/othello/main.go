package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
)

var (
	mplusNormalFont font.Face
)

// init initializes the application by parsing the goregular.TTF font file and creating a new face with specified options.
// It sets the mplusNormalFont variable to the newly created face.
// If any error occurs during the initialization process, it logs the error and terminates the program.
func init() {
	tt, err := opentype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    16,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

const (
	Empty = iota
	Black
	White
)

type Board [8][8]int

type Game struct {
	SquareNum     int
	SquareMergin  int
	SquareSize    int
	SquareColor   color.NRGBA
	BoardSize     int
	BoardColor    color.Color
	ScreenColor   color.RGBA
	ScreenSize    int
	Board         *Board
	CurrentPlayer int
}

func NewGame() *Game {
	g := &Game{
		SquareNum:     8,
		SquareMergin:  4,
		SquareSize:    100,
		SquareColor:   color.NRGBA{0x00, 0x8b, 0x45, 0xff},
		BoardColor:    color.Black,
		ScreenColor:   color.RGBA{0xfa, 0xf8, 0xef, 0xff},
		Board:         NewBoard(),
		CurrentPlayer: Black,
	}
	g.BoardSize = (g.SquareSize+g.SquareMergin)*g.SquareNum + g.SquareMergin

	screenMergin := 40
	g.ScreenSize = g.BoardSize + screenMergin

	return g
}

func NewBoard() *Board {
	b := new(Board)

	b[3][3] = White
	b[4][4] = White
	b[3][4] = Black
	b[4][3] = Black
	return b
}

func (g *Game) Update() error {
	// ユーザーの入力を取得
	x, y := ebiten.CursorPosition()

	// マウスの左ボタンが押されていない場合は何もしない
	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		return nil
	}

	// クリックされたセルの座標を計算
	i := x / (g.SquareSize + g.SquareMergin)
	j := y / (g.SquareSize + g.SquareMergin)

	// クリックされたセルがボード内であることを確認
	if i < 0 || i >= g.SquareNum || j < 0 || j >= g.SquareNum {
		return nil
	}

	// クリックされたセルが空であることを確認
	if g.Board[i][j] != Empty {
		return nil
	}

	// 合法手であることを確認
	if !g.IsValidMove(i, j, g.CurrentPlayer) {
		return nil
	}

	// 石を置く
	g.Board[i][j] = g.CurrentPlayer

	// 敵の石を自分の石に変える
	g.FlipDisks(i, j)

	// 手番を交代
	if g.CurrentPlayer == Black {
		g.CurrentPlayer = White
	} else {
		g.CurrentPlayer = Black
	}

	return nil
}

func (g *Game) IsValidMove(x, y, player int) bool {
	if g.Board[x][y] != Empty {
		return false
	}

	dx := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	dy := []int{-1, -1, -1, 0, 0, 1, 1, 1}

	for i := 0; i < 8; i++ {
		nx, ny := x+dx[i], y+dy[i]
		if nx >= 0 && nx < 8 && ny >= 0 && ny < 8 && g.Board[nx][ny] == 3-player {
			for {
				nx += dx[i]
				ny += dy[i]
				if nx < 0 || nx >= 8 || ny < 0 || ny >= 8 {
					break
				}
				if g.Board[nx][ny] == Empty {
					break
				}
				if g.Board[nx][ny] == player {
					return true
				}
			}
		}
	}

	return false
}

func (g *Game) FlipDisks(x, y int) {
	dx := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	dy := []int{-1, -1, -1, 0, 0, 1, 1, 1}

	for i := 0; i < 8; i++ {
		nx, ny := x+dx[i], y+dy[i]
		if nx >= 0 && nx < 8 && ny >= 0 && ny < 8 && g.Board[nx][ny] == 3-g.CurrentPlayer {
			for {
				nx += dx[i]
				ny += dy[i]
				if nx < 0 || nx >= 8 || ny < 0 || ny >= 8 || g.Board[nx][ny] == Empty {
					break
				}
				if g.Board[nx][ny] == g.CurrentPlayer {
					for nx, ny = nx-dx[i], ny-dy[i]; g.Board[nx][ny] == 3-g.CurrentPlayer; nx, ny = nx-dx[i], ny-dy[i] {
						g.Board[nx][ny] = g.CurrentPlayer
					}
					break
				}
			}
		}
	}
}

func (g *Game) IsBoardFull() bool {
	for i := 0; i < g.SquareNum; i++ {
		for j := 0; j < g.SquareNum; j++ {
			if g.Board[i][j] == Empty {
				// Found an empty cell, so the board is not full
				return false
			}
		}
	}
	// No empty cells were found, so the board is full
	return true
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.ScreenColor)
	board := ebiten.NewImage(g.BoardSize, g.BoardSize)
	board.Fill(g.BoardColor)
	for j := 0; j < g.SquareNum; j++ {
		for i := 0; i < g.SquareNum; i++ {
			square := ebiten.NewImage(g.SquareSize, g.SquareSize)
			square.Fill(g.SquareColor)
			op := &ebiten.DrawImageOptions{}
			x := i*g.SquareSize + (i+1)*g.SquareMergin
			y := j*g.SquareSize + (j+1)*g.SquareMergin
			op.GeoM.Translate(float64(x), float64(y))
			board.DrawImage(square, op)

			// Draw the disk if there is one
			if g.Board[i][j] == Black {
				g.SetDisk(board, i, j, color.Black)
			} else if g.Board[i][j] == White {
				g.SetDisk(board, i, j, color.White)
			}
		}
	}

	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
	bw, bh := board.Bounds().Dx(), board.Bounds().Dy()
	x := (sw - bw) / 2
	y := (sh - bh) / 2
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(board, op)

	// Display the current player's turn or Game Set
	if g.IsGameSet() {
		winner := g.DetermineWinner()
		if winner == Black {
			text.Draw(screen, "Game Set! Winner: Black", mplusNormalFont, 10, 15, color.RGBA{0xff, 0x00, 0x00, 0xff})
		} else if winner == White {
			text.Draw(screen, "Game Set! Winner: White", mplusNormalFont, 10, 15, color.RGBA{0xff, 0x00, 0x00, 0xff})
		} else {
			text.Draw(screen, "Game Set! Draw", mplusNormalFont, 10, 15, color.RGBA{0xff, 0x00, 0x00, 0xff})
		}
	} else if g.CurrentPlayer == Black {
		text.Draw(screen, "Current turn: Black", mplusNormalFont, 10, 15, color.Black)
	} else {
		text.Draw(screen, "Current turn: White", mplusNormalFont, 10, 15, color.Black)
	}

	return
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ScreenSize, g.ScreenSize
}

func (g *Game) CanPlaceDisk(player int) bool {
	for i := 0; i < g.SquareNum; i++ {
		for j := 0; j < g.SquareNum; j++ {
			if g.Board[i][j] == Empty && g.IsValidMove(i, j, player) {
				return true
			}
		}
	}
	return false
}

func (g *Game) ChangeTurn() {
	if g.CanPlaceDisk(g.CurrentPlayer) {
		return
	}
	g.CurrentPlayer = g.OtherPlayer()
}

func (g *Game) OtherPlayer() int {
	if g.CurrentPlayer == Black {
		return White
	}
	return Black
}

func (g *Game) IsGameSet() bool {
	return !g.CanPlaceDisk(Black) && !g.CanPlaceDisk(White)
}

func (g *Game) CountDisks(player int) int {
	count := 0
	for i := 0; i < g.SquareNum; i++ {
		for j := 0; j < g.SquareNum; j++ {
			if g.Board[i][j] == player {
				count++
			}
		}
	}
	return count
}

func (g *Game) DetermineWinner() int {
	blackDisks := g.CountDisks(Black)
	whiteDisks := g.CountDisks(White)

	switch {
	case blackDisks > whiteDisks:
		return Black
	case blackDisks < whiteDisks:
		return White
	default:
		return Empty // draw
	}
}

func (g *Game) SetDisk(board *ebiten.Image, x int, y int, color color.Color) {
	vector.DrawFilledCircle(board, float32(g.SquareMergin+g.SquareSize)*(float32(x)+0.5), float32(g.SquareMergin+g.SquareSize)*(float32(y)+0.5), float32(g.SquareSize)*0.8/2, color, true)
}

func main() {
	game := NewGame()
	ebiten.SetWindowSize(game.ScreenSize, game.ScreenSize)
	ebiten.SetWindowTitle("Othello")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
