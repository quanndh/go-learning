package strategy

type IStrategy interface {
	Execute(int, int) int
}
