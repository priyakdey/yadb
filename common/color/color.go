// Contains code for the Color for x-256 term
package color

type Color string

// ref: https://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html
const (
	RESET Color = "\u001b[0m"
	RED         = "\u001b[31m"
	GREEN       = "\u001b[32m"
	BLUE        = "\u001b[34m"
)
