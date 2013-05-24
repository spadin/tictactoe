package game

import (
	"github.com/remogatto/prettytest"
	"testing"
)

type testTictactoeRulesSuite struct {
	prettytest.Suite
}

func rulesSuiteSetup() (board Board, rules Rules) {
	board = NewTictactoeBoard()
	rules = NewTictactoeRules(board)
	return
}

func TestTictactoeRulesRunner(t *testing.T) {
	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(testTictactoeRulesSuite),
	)
}

func (t *testTictactoeRulesSuite) TestNewTictactoeRulesHasNoWinner() {
	_, rules := rulesSuiteSetup()
	t.False(rules.HasWinner())
}

func (t *testTictactoeRulesSuite) TestXWinningTictactoeRulesHasWinner() {
	board, rules := rulesSuiteSetup()
	PlayXWinningGame(board.(*TictactoeBoard))
	t.True(rules.HasWinner(), "X winning rules should have winner")
}

func (t *testTictactoeRulesSuite) TestTictactoeRulesoverFalseWithNewRules() {
	_, rules := rulesSuiteSetup()
	t.False(rules.IsGameover(), "New rules is not rulesover")
}

func (t *testTictactoeRulesSuite) TestTictactoeRulesoverTrueWithWinningRules() {
	board, rules := rulesSuiteSetup()
	PlayXWinningGame(board.(*TictactoeBoard))
	t.True(rules.IsGameover(), "TictactoeRules is over when there is a winner")
}

func (t *testTictactoeRulesSuite) TestTictactoeRulesoverTrueWithFullBoard() {
	board, rules := rulesSuiteSetup()
	PlayTiedGame(board.(*TictactoeBoard))
	t.True(rules.IsGameover(), "TictactoeRules is over when board full")
}

func (t *testTictactoeRulesSuite) TestTictactoeWinnerTiedRules() {
	board, rules := rulesSuiteSetup()
	PlayTiedGame(board.(*TictactoeBoard))
	t.Nil(rules.Winner(), "Winner is null for tied rules")
}

func (t *testTictactoeRulesSuite) TestTictactoeWinnerXWinner() {
	board, rules := rulesSuiteSetup()
	PlayXWinningGame(board.(*TictactoeBoard))
	t.Equal(X, rules.Winner())
}
