package repl

import (
	"bufio"
	"fmt"
	"io"
	"bat-go/lexer"
	"bat-go/parser"
)

const PROMPT = ">> "

const BAT_FACE = `
                      _..-'(                       )'-.._
                   ./'. '||\\.       (\_/)       .//|' .'\.
                ./'.|'.'||||\\|..    )O O(    ..|//|||'.'|.'\.
             ./'..|'.|| |||||\'''''' '"'" ''''''/||||| ||'|..'\.
           ./'.||'.|||| ||||||||||||.     .|||||||||||| |||||'||.'\.
          /'|||'.|||||| ||||||||||||{     }|||||||||||| ||||||'|||'\
         '.|||'.||||||| ||||||||||||{     }|||||||||||| |||||||'|||.'
        '.||| ||||||||| |/'   ''\||''     ''||/''   '\| ||||||||| |||.'
        |/' \./'     '\./         \!|\   /|!/         \./'     '\./ '\|
        V    V         V          }' '\ /' '{          V         V    V
        '    '         '               V               '         '    '
`

func Start(in io.Reader, out io.Writer) {
	Scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := Scanner.Scan()

		if !scanned {
			return
		}

		line := Scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, BAT_FACE)
	io.WriteString(out, "Eek! Something went sideways in the bat cave.\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
