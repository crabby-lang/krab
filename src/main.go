package main

import (
	"log"
	"os"
	
	// gio ui
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"

	"internal/terminal"
	"internal/ui"
	"internal/runner"
)

type IDE struct {
	terminalBuffer			*terminal.Buffer
	terminalExecutor		*terminal.Executor
	terminalPanel		 	*ui.TerminalPanel
	crabRunner				*runner.CrabRunner
	currentFile				string
}

func NewIDE() *IDE {
	workingDir, err := os.Getwd()
	if err != nil {
		workingDir = "."
	}

	termBuffer 		:= terminal.NewBuffer(1000)
	termExecutor	:= terminal.NewExecutor(termBuffer, workingDir)
	termPanel		:= ui.NewTerminalPanel(termBuffer)
	crabRunner		:= runner.NewCrabRunner("crabby", workingDir)

	ide := &IDE{
		terminalBuffer:		termBuffer,
		terminalExecutor:   termExecutor,
		terminalPanel:      termPanel,
		crabRunner:         crabRunner,
		currentFile:		"main.crab"
	}

	termPanel.OnClear = func() {
		termBuffer.Clear()
	}

	termPanel.OnStop = func() {
		termExecutor.Stop()
	}

	termPanel.OnRun = func() {
		go func() {
			if err := ide.crabRunner.Run(termExecutor, ide.currentFile); err != nil {
				termBuffer.Append(terminal.Colorize("Error: "+err.Error(), terminal.Red))
			}
		}()
	}

	return ide
}

func (ide *IDE) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx, layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
		return ide.terminalPanel.Layout(gtx, th)
	}),
  )
}

func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Title("Krab IDE"))
		w.Option(app.Size(1200, 800))

		ide := NewIDE()
		th := material.NewTheme()

		var ops op.Ops
		for {
			switch e := w.Event().(type) {
			case app.DestroyEvent:
				os.Exit(0)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				ide.Layout(gtx, th)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
