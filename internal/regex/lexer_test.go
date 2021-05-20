/* Lexer tests.
 */

package regex

import (
	"testing"
)

func TestTokenStr(t *testing.T) {
	if tokenStr(SYMBOL) != "SYMBOL" {
		t.Errorf("expected SYMBOL but got %s", tokenStr(SYMBOL))
	}

	if tokenStr(STAR) != "STAR" {
		t.Errorf("expected STAR but got %s", tokenStr(STAR))
	}

	if tokenStr(BAR) != "BAR" {
		t.Errorf("expected BAR but got %s", tokenStr(BAR))
	}

	if tokenStr(DOT) != "DOT" {
		t.Errorf("expected DOT but got %s", tokenStr(DOT))
	}

	if tokenStr(LPAREN) != "LPAREN" {
		t.Errorf("expected LPAREN but got %s", tokenStr(DOT))
	}

	if tokenStr(RPAREN) != "RPAREN" {
		t.Errorf("expected RPAREN but got %s", tokenStr(RPAREN))
	}

	if tokenStr(END) != "END" {
		t.Errorf("expected END but got %s", tokenStr(END))
	}

	if tokenStr(DEFAULT) != "DEFAULT" {
		t.Errorf("expected DEFAULT but got %s", tokenStr(DEFAULT))
	}
}

func TestInitLexer(t *testing.T) {
	testLexer := initLexer()

	if &testLexer == nil {
		t.Errorf("expected non-nil value for &testLexer")
	}
}

func TestStoreInput(t *testing.T) {
	testString := "test"
	testLexer  := initLexer()

	testLexer.storeInput(testString)

	if testLexer.input != testString {
		t.Errorf("expected %s but got %s", testString, testLexer.input)
	}
}

func TestNextToken(t *testing.T) {
	testString := "()*|.A"
	testLexer  := initLexer()

	testLexer.storeInput(testString)

	token := testLexer.nextToken()
	if token != LPAREN {
		t.Errorf("expected LPAREN but got %s", tokenStr(token))
	}

	token = testLexer.nextToken()
	if token != RPAREN {
		t.Errorf("expected RPAREN but got %s", tokenStr(token))
	}

	token = testLexer.nextToken()
	if token != STAR {
		t.Errorf("expected STAR but got %s", tokenStr(token))
	}

	token = testLexer.nextToken()
	if token != BAR {
		t.Errorf("expected BAR but got %s", tokenStr(token))
	}

	token = testLexer.nextToken()
	if token != DOT {
		t.Errorf("expected DOT but got %s", tokenStr(token))
	}

	token = testLexer.nextToken()
	if token != SYMBOL {
		t.Errorf("expected SYMBOL but got %s", tokenStr(token))
	}

	token = testLexer.nextToken()
	if token != END {
		t.Errorf("expected END but got %s", tokenStr(token))
	}
}
