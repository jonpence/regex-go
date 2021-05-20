/* Parser tests.
 */

package regex

import (
	"fmt"
	"testing"
)

func TestItos(t *testing.T) {
	testInt := 9876543210
	testStr := "9876543210"

	result := itos(testInt)

	if result != testStr {
		t.Errorf("expected %s, got %s", testStr, result)
	}
}

func TestInitParser(t *testing.T) {
	testParser := initParser()

	if &testParser == nil {
		t.Errorf("expected non-nil value for testParser")
	}

	if testParser.token != DEFAULT {
		t.Errorf("expected testParser.token to be DEFAULT, instead is %s", tokenStr(testParser.token))
	}
}

func TestSetDebug(t *testing.T) {
	testParser := initParser()

	if testParser.debug {
		t.Errorf("expected testParser.debug to be false")
	}

	testParser.setDebug()

	if !testParser.debug {
		t.Errorf("expected testParser.debug to be true")
	}
}

func TestUnsetDebug(t *testing.T) {
	testParser := initParser()
	testParser.debug = true

	testParser.unsetDebug()

	if testParser.debug {
		t.Errorf("expected testParser.debug to be false")
	}
}

func TestRequire(t *testing.T) {
	testParser := initParser()

	if !testParser.require(DEFAULT) {
		t.Errorf("expected testParser.require(DEFAULT) to be true")
	}

	if testParser.token != DEFAULT {
		t.Errorf("expected testParser.token to equal DEFAULT, instead is %s", tokenStr(testParser.token))
	}
}

func TestAccept(t *testing.T) {
	testParser := initParser()
	testParser.lexer.storeInput("()")

	testParser.accept(DEFAULT)

	if !testParser.accept(LPAREN) {
		t.Errorf("expected testParser.require(LPAREN) to be true")
	}

	if testParser.token != RPAREN {
		t.Errorf("expected testParser.token to be RPAREN, instead is %s", tokenStr(testParser.token))
	}
}

func TestOr(t *testing.T) {
	fmt.Println("TESTING OR")
	testParser := initParser()

	testParser.parse("A|B")

	for state := range testParser.nfa.states {
		testParser.nfa.states[state].printNFAState()
	}
	fmt.Println()
}

func TestConcatenate(t *testing.T) {
	fmt.Println("TESTING CONCATENATE")
	testParser := initParser()

	testParser.parse("A.B")

	for state := range testParser.nfa.states {
		testParser.nfa.states[state].printNFAState()
	}
	fmt.Println()
}

func TestKleene(t *testing.T) {
	fmt.Println("TESTING KLEENE")
	testParser := initParser()

	testParser.parse("A*")

	for state := range testParser.nfa.states {
		testParser.nfa.states[state].printNFAState()
	}
	fmt.Println()
}

func TestSymbol(t *testing.T) {
	fmt.Println("TESTING SYMBOL")
	testParser := initParser()

	testParser.parse("A")

	for state := range testParser.nfa.states {
		testParser.nfa.states[state].printNFAState()
	}
	fmt.Println()
}

func TestParse(t *testing.T) {
	testParser := initParser()

	if !testParser.parse("A|B") {
		t.Errorf("expected testParser.parse(\"A|B\") to be true")
	}

	if testParser.parse("A...B") {
		t.Errorf("expected testParser.parse(\"A...B\") to be false")
	}
}
