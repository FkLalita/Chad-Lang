package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/FkLalita/Chad-Lang/lexer"
	"github.com/FkLalita/Chad-Lang/token"
)

const PROMPT = ">>"

func START(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)

		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("*+v\n", tok)
		}

	}
}
