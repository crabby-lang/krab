package terminal

<<<<<<< HEAD
=======
import "fmt"

type Color string

const (
	Reset  Color = "\033[0m"
	Red    Color = "\033[31m"
	Green  Color = "\033[32m"
	Yellow Color = "\033[33m"
	Blue   Color = "\033[34m"
)

func Colorize(text string, color Color) string {
	return fmt.Sprintf("%s%s%s", color, text, Reset)
}
>>>>>>> 90f5ff9 (commit init)

