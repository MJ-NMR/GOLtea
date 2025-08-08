package main

import (
	"fmt"
	"os"

	"github.com/MJ-NMR/GOL"
	tea "github.com/charmbracelet/bubbletea"
)

func main()  {
	pro := tea.NewProgram(initialModel())
	if _, err := pro.Run(); err != nil {
		fmt.Printf("there an error : %v", err)
		os.Exit(1)
	}
}

type model struct {
    frame  GOL.State
	courser struct{y,x int}
}

func (m model) Init() tea.Cmd {
    // Just return `nil`, which means "no I/O right now, please."
    return nil
}

func initialModel() model {
	return model{frame: GOL.CreateState(10,10)}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
            return m, tea.Quit

		case "up", "k":
			if m.courser.y > 0 {
				m.courser.y -= 1
			}

		case "down", "l":
			if m.courser.y < len(m.frame)-1 {
				m.courser.y += 1
			}

		case "right", ";":
			if m.courser.x < len(m.frame[0])-1 {
				m.courser.x += 1
			}

		case "left", "j":
			if m.courser.x > 0 {
				m.courser.x -= 1
			}
		case " ":
			if m.frame[m.courser.y][m.courser.x] {
				m.frame[m.courser.y][m.courser.x] = false
			} else {
				m.frame[m.courser.y][m.courser.x] = true
			}

		case "enter":
			m.frame = GOL.PlayRound(m.frame)
		}
	}
    return m, nil
}

func (m model) View() string {
    // The header
    s := "Game Of life\n\n"

    // Iterate over our choices
    for y, row := range m.frame {
		for x := range row {
			if m.courser.y == y && m.courser.x == x {
				s += ">"
			} else {
				s += " "
			}
			if m.frame[y][x] {
				s += "\033[32m#\033[0m"
			} else {
				s += "\033[31m*\033[0m"
			}
		}
		s += "\n"
    }

    // The footer
	s += "\nPress \033[32mq\033[0m: quit, \033[32mEnter\033[0m: next round, \033[32mSpace\033[0m: toggele cell.\n"

    // Send the UI for rendering
    return s
}
