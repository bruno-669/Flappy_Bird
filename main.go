package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

func PrintMatrix(matrix [][]string) {
	for i := range matrix {
		for j := 0; j < len(matrix[i]); j++ {
			print(matrix[i][j])
		}
		println()
	}
}

func AggregateMatrix(tmpmatrix [][]string, symbol string) {

	for i := range tmpmatrix {
		for j := range tmpmatrix[i] {
			tmpmatrix[i][j] = symbol
		}
	}
}

func CreateMatrix(rows, cols int) [][]string {
	matrix := make([][]string, rows)
	for i := range matrix {
		matrix[i] = make([]string, cols)
	}
	return matrix
}

func PaintBoardMatrix(matrix [][]string) {

	for i := range matrix {
		matrix[i][0] = "║"
		matrix[i][len(matrix[i])-1] = "║"
	}
	for i := range matrix[0] {
		matrix[0][i] = "═"
		matrix[len(matrix)-1][i] = "═"
	}
	matrix[len(matrix)-1][0] = "╚"
	matrix[0][0] = "╔"
	matrix[0][len(matrix[0])-1] = "╗"
	matrix[len(matrix)-1][len(matrix[0])-1] = "╝"
}

func CreatePipe(rows int) [][]string {
	pipe := CreateMatrix(rows, 8)
	max := (len(pipe) / 2) - 6
	min := (len(pipe) / 2) + 6
	polarity_random := rand.Intn(2)
	random := rand.Intn(10)
	if polarity_random == 1 {
		random *= -1
	}
	max += random
	min += random
	AggregateMatrix(pipe, "░")
	for i := max; i < min; i++ {
		for j := range pipe[0] {
			pipe[i][j] = " "
		}
	}
	for i := range pipe {
		if i < max || i > min {
			pipe[i][0] = "║"
			pipe[i][len(pipe[i])-1] = "║"
		}
	}
	for j := range pipe[0] {
		pipe[max][j] = "═"
		pipe[min][j] = "═"
	}
	pipe[max][0] = "╚"
	pipe[max][len(pipe[0])-1] = "╝"
	pipe[min][0] = "╔"
	pipe[min][len(pipe[0])-1] = "╗"
	return pipe
}

func AddPipe(position int, matrix, pipe [][]string) {
	if position+len(pipe[0]) <= 0 {
		return
	}
	if position >= len(matrix[0]) {
		return
	}
	for i := range matrix {
		for j := range matrix[i] {
			pipeJ := j - position
			if pipeJ >= 0 && pipeJ < len(pipe[0]) {
				matrix[i][j] = pipe[i][pipeJ]
			}
		}
	}
}

func MovementPipes(frame, rows, cols int, matrix [][]string, pipe1, pipe2, pipe3, pipe4 *[][]string) {
	AggregateMatrix(matrix, " ")
	AddRoad(matrix)
	pipeWidth := len((*pipe1)[0])
	screensize := cols + pipeWidth

	position := screensize - ((frame + screensize) % (screensize)) - 8

	if position == -7 {
		tmp := CreatePipe(rows)
		*pipe1 = tmp
	}
	if (position+screensize+40)%screensize-8 == -7 {
		tmp := CreatePipe(rows)
		*pipe2 = tmp
	}
	if (position+screensize+72)%screensize-8 == -7 {
		tmp := CreatePipe(rows)
		*pipe3 = tmp
	}
	if (position+screensize+104)%screensize-8 == -7 {
		tmp := CreatePipe(rows)
		*pipe4 = tmp
	}
	AddPipe(position, matrix, *pipe1)
	if frame > 33 {
		AddPipe((position+screensize+40)%screensize-8, matrix, *pipe2)
	}
	if frame > 65 {
		AddPipe((position+screensize+72)%screensize-8, matrix, *pipe3)
	}
	if frame > 97 {
		AddPipe((position+screensize+104)%screensize-8, matrix, *pipe4)
	}

}

func CreateBurd() [][]string {
	burd := CreateMatrix(3, 5)
	AggregateMatrix(burd, " ")
	burd[2][2] = "╘"
	burd[2][1] = "╘"
	burd[1][4] = "►"
	burd[1][3] = "◙"
	burd[1][2] = "█"
	burd[1][1] = "█"
	burd[0][1] = "▲"
	return burd
}

func AddBurd(position int, matrix, burd [][]string, Game *bool) {
	centerY := len(matrix)/2 + position
	centerX := len(matrix[0]) / 2
	height := len(burd)
	width := len(burd[0])
	top := centerY - height/2
	left := centerX - width/2
	obstacles := "║╔╗╚╝═░▒"

	if top < 0 || top+height >= len(matrix) {
		*Game = false
		return
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if burd[i][j] == " " {
				continue
			}

			y := top + i
			x := left + j

			if strings.ContainsAny(matrix[y][x], obstacles) {
				*Game = false
				return
			}
		}
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			y := top + i
			x := left + j

			if burd[i][j] != " " {
				matrix[y][x] = burd[i][j]
			}
		}
	}
}

func WingReplacement(frame int, burd [][]string) {
	switch_wird := frame % 2
	if switch_wird == 0 {
		burd[0][1] = "▲"
		burd[1][0] = " "
	} else {
		burd[0][1] = " "
		burd[1][0] = "◄"
	}
}

func PrintMatrixBuffered(matrix [][]string) string {
	var buffer strings.Builder
	for i := range matrix {
		for j := 0; j < len(matrix[i]); j++ {
			buffer.WriteString(matrix[i][j])
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}

func waitForInput() {
	var input string
	fmt.Scanln(&input)
	if input == "w" || input == "" {

	}
	fmt.Printf("Вы ввели: %s\n", input)
}

func ScoreAdded(frame int, score *int) {
	if frame == 64 {
		*score += 1
	}
	if frame > 64 {
		if frame%32 == 0 {
			*score += 1
		}
	}
}

func AddedScore(matrix [][]string, score, pos_y, pos_x int) {
	text := "score :"

	for index, char := range text {
		matrix[pos_y][index+pos_x] = string(char)

	}
	if score >= 100 {
		matrix[pos_y][pos_x-1+8] = strconv.Itoa(score / 100)
	}
	if score >= 10 {
		matrix[pos_y][pos_x-1+9] = strconv.Itoa((score / 10) % 10)
	}
	matrix[pos_y][pos_x-1+10] = strconv.Itoa(score % 10)

}

func KeyboardReader(Up, Game *bool) {
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			break
		}

		if key == keyboard.KeyArrowUp {
			*Up = true
		} else if key == keyboard.KeySpace {
			*Up = true
		} else if key == keyboard.KeyEsc || char == 'q' {
			*Game = false
		}

	}
}

func BurdPositionCorrect(Frame int, Up *bool, position_burd *int) {
	if Frame%2 == 0 {
		*position_burd++
	}
	if *Up {
		*position_burd -= 5
		*Up = false
	}
}

func DrawGameOverFancy(matrix [][]string) {
	gameOver := []string{
		"   _____                         ____                  ",
		"  / ____|                       / __ \\                 ",
		" | |  __  __ _ _ __ ___   ___  | |  | |_   _____ _ __  ",
		" | | |_ |/ _` | '_ ` _ \\ / _ \\ | |  | \\ \\ / / _ \\ '__| ",
		" | |__| | (_| | | | | | |  __/ | |__| |\\ V /  __/ |    ",
		"  \\_____|\\__,_|_| |_| |_|\\___|  \\____/  \\_/ \\___|_|    ",
	}

	centerY := len(matrix) / 2
	centerX := len(matrix[0]) / 2
	startY := centerY - len(gameOver)/2
	startX := centerX - len(gameOver[0])/2

	for i, line := range gameOver {
		y := startY + i
		if y >= 0 && y < len(matrix) {
			for j, ch := range line {
				x := startX + j
				if x >= 0 && x < len(matrix[0]) {
					matrix[y][x] = string(ch)
				}
			}
		}
	}
}

func AddRoad(matrix [][]string) {
	for i := 37; i < len(matrix); i++ {
		for j := range matrix[i] {
			matrix[i][j] = "/"
		}
	}
}

func main() {
	rows := 40
	cols := 120
	score := 0
	Game := true
	Up := false
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()
	go KeyboardReader(&Up, &Game)
	matrix := CreateMatrix(rows, cols)
	pipe1 := CreatePipe(rows)
	pipe2 := CreatePipe(rows)
	pipe3 := CreatePipe(rows)
	pipe4 := CreatePipe(rows)
	burd := CreateBurd()
	position_burd := 0

	fmt.Print("\033[?25l")
	defer fmt.Print("\033[?25h")

	for i := 0; i < 100000 && Game; i++ {

		var output strings.Builder
		ScoreAdded(i, &score)
		WingReplacement(i, burd)
		MovementPipes(i, rows, cols, matrix, &pipe1, &pipe2, &pipe3, &pipe4)
		AddBurd(position_burd, matrix, burd, &Game)
		AddedScore(matrix, score, 1, 1)
		if !Game {
			for i := 14; i < 27; i++ {
				for j := 30; j < 90; j++ {
					matrix[i][j] = " "
					if i == 14 || i == 26 || j == 30 || j == 89 || j == 31 || j == 88 {
						matrix[i][j] = "▓"
					}

				}
			}
			DrawGameOverFancy(matrix)
			AddedScore(matrix, score, 24, 54)
		}
		PaintBoardMatrix(matrix)

		output.WriteString("\033[H")
		output.WriteString(PrintMatrixBuffered(matrix))
		fmt.Print(output.String())
		BurdPositionCorrect(i, &Up, &position_burd)
		time.Sleep(150 * time.Millisecond)
	}
}
