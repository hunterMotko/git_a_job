package fancy

import (
	"bytes"
	"text/template"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

const content = `
# By: 
  {{ .By }}
# Title: 
  {{ .JsonTitle }}
# Score: 
  {{ .Score }}
# Time: 
  {{ .Time }}
# Type: 
  {{ .Type }}
# Url: 
  {{ .Url }}
# Text: 
  {{ .Text }}
`

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render

type jobView struct {
	viewport viewport.Model
}

func newJobView(i Item) (*jobView, error) {
	const width = 78
	vp := viewport.New(width, 20)
	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(2)

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(width),
	)
	if err != nil {
		return nil, err
	}

	tmpl, err := template.New(content).Parse(content)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, i)
	if err != nil {
		panic(err)
	}

	str, err := renderer.Render(buf.String())
	if err != nil {
		return nil, err
	}

	vp.SetContent(str)
	return &jobView{
		viewport: vp,
	}, nil
}

func (e jobView) Init() tea.Cmd {
	return nil
}

func (e jobView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return e, tea.Quit
		default:
			var cmd tea.Cmd
			e.viewport, cmd = e.viewport.Update(msg)
			return e, cmd
		}
	default:
		return e, nil
	}
}
func (e jobView) View() string {
	return e.viewport.View() + e.helpView()
}
func (e jobView) helpView() string {
	return helpStyle("\n  ↑/↓: Navigate • q: Quit\n")
}

//func main() {
//	model, err := newJobView()
//	if err != nil {
//		fmt.Println("Could not initialize Bubble Tea model:", err)
//		os.Exit(1)
//	}
//	if _, err := tea.NewProgram(model).Run(); err != nil {
//		fmt.Println("Bummer, there's been an error:", err)
//		os.Exit(1)
//	}
//}
