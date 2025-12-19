package ui

import (
	"image"
	"image/color"

	// gio ui
<<<<<<< HEAD
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
=======
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
>>>>>>> 90f5ff9 (commit init)
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/font"
)

type Terminal struct {
	list			widget.List
	buffer			TerminalBuffer
	bgColor			color.NRGBA
	textColor		color.NRGBA
	fontSize		unit.Sp
	scrollOffset	int
}

type TerminalPanel struct {
	terminal		*Terminal
	clearBtn		widget.Clickable
	stopBtn			widget.Clickable
	runBtn			widget.Clickable
	OnClear			func()
	OnStop			func()
	OnRun			func()
}

type TerminalBuffer interface {
	Lines()				[]string
	Clear()
	Append(text string)
	LineCount()			int
}

func NewTerminal(buffer TerminalBuffer) *Terminal {
	return &Terminal{
		buffer: buffer,
		list: widget.List{
			List: layout.List{
				Axis: layout.Vertical,
			},
		},
<<<<<<< HEAD
		bgColor:		color.NRGBA{R: 30, G: 30, B: 30, A: 255}, // Dark
		textColor:		color.NRGBA{R: 200, G: 200, B: 200, A: 255} // Light
		fontSize:		12,
	}
}

func (t* Terminal) Layout(gtx layout.Context. th *material.Theme) layout.Dimensions {
=======
	bgColor:		color.NRGBA{R: 30, G: 30, B: 30, A: 255}, // Dark
	textColor:		color.NRGBA{R: 200, G: 200, B: 200, A: 255}, // Light
	fontSize:		12,
	}
}

func (t *Terminal) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
>>>>>>> 90f5ff9 (commit init)
	paint.FillShape(gtx.Ops, t.bgColor, clip.Rect{
		Max: gtx.Constraints.Max,
	}.Op())

	lines := t.buffer.Lines()

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx, layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
		return t.list.List.Layout(gtx, len(lines), func(gtx layout.Context, index int) layout.Dimensions {
			return t.layoutLine(gtx, th, lines[index])
		})
	  }),
    )
}

<<<<<<< HEAD
func (t *Terminal) layoutLine(gtx, layout.Context, th *material.Theme, line string) layout.Dimensions {
=======
func (t *Terminal) layoutLine(gtx layout.Context, th *material.Theme, line string) layout.Dimensions {
>>>>>>> 90f5ff9 (commit init)
	return layout.Inset{
		Top:		unit.Dp(2),
		Bottom:		unit.Dp(2),
		Left:		unit.Dp(8),
		Right:		unit.Dp(8),
<<<<<<< HEAD
	}.Layout(gtx, func(gtx layout.Context, th *material.Theme, line string) layout.Dimensions {
=======
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
>>>>>>> 90f5ff9 (commit init)
		// ansi colors later, for now its going to be a basic text rendering
		label := material.Label(th, t.fontSize, line)
		label.Color = t.textColor
		label.Font.Typeface = font.Typeface("monospace")
		return label.Layout(gtx)
	})
}

func (t *Terminal) ScrollToBottom() {
	lineCount := t.buffer.LineCount()
	if lineCount > 0 {
		t.list.List.Position.First = lineCount - 1
		t.list.List.Position.Offset = 0
	}
}

func (t *Terminal) Clear() {
	t.buffer.Clear()
}

func NewTerminalPanel(buffer TerminalBuffer) *TerminalPanel {
	return &TerminalPanel{
		terminal: NewTerminal(buffer),
	}
}

func (tp *TerminalPanel) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
<<<<<<< HEAD
	return layout.Flex{Axis: layout,Vertical}.Layout(gtx, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
=======
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
>>>>>>> 90f5ff9 (commit init)
		return tp.layoutToolbar(gtx, th)
	  }), layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
		return tp.terminal.Layout(gtx, th)
	  }),
	)
}

func (tp *TerminalPanel) layoutToolbar(gtx layout.Context, th *material.Theme) layout.Dimensions {
	if tp.clearBtn.Clicked(gtx) && tp.OnClear != nil {
		tp.OnClear()
	}
	if tp.stopBtn.Clicked(gtx) && tp.OnStop != nil {
		tp.OnStop()
	}
	if tp.runBtn.Clicked(gtx) && tp.OnRun != nil {
		tp.OnRun()
	}

	bg := color.NRGBA{R: 45, G: 45, B: 45, A: 255}
	paint.FillShape(gtx.Ops, bg, clip.Rect{
		Max: image.Pt(gtx.Constraints.Max.X, gtx.Dp(40)),
	}.Op())

	return layout.Inset{
<<<<<<< HEAD
		Top:		unit.Dp(4)
		Bottom:		unit.Dp(4)
		Left:		unit.Dp(8)
		Right:		unit.Dp(8)
=======
		Top:		unit.Dp(4),
		Bottom:		unit.Dp(4),
		Left:		unit.Dp(8),
		Right:		unit.Dp(8),
>>>>>>> 90f5ff9 (commit init)
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween}.Layout(gtx, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				btn := material.Button(th, &tp.runBtn, "▶️ Run")
				btn.Background = color.NRGBA{R: 0, G: 150, B: 0, A: 255}
				return btn.Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Width: unit.Dp(8)}.Layout),
<<<<<<< HEAD
			layout.Rigid(func(gtx, layout.Context) layout.Dimensions {
=======
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
>>>>>>> 90f5ff9 (commit init)
				btn := material.Button(th, &tp.stopBtn, "⏹️ Stop")
				btn.Background = color.NRGBA{R: 200, G: 0, B: 0, A: 255}
				return btn.Layout(gtx)
			}),
		)
	}),
	layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		btn := material.Button(th, &tp.clearBtn, "Clear")
		btn.Background = color.NRGBA{R: 80, G: 80, B: 80, A: 255}
		return btn.Layout(gtx)
	}),
   )
  })
}

func (tp *TerminalPanel) ScrollToBottom() {
	tp.terminal.ScrollToBottom()
}
