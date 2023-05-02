package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/toMaTo0502/monkey-interpreter-sample/lexer"
	"github.com/toMaTo0502/monkey-interpreter-sample/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
