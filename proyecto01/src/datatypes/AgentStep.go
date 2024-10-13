package datatypes

type AgentAction int

const (
	UP AgentAction = iota
	RIGHT
	DOWN
	LEFT
)

// The position of the agent among the previous position
type AgentStep struct {
	Action           AgentAction
	Depth            int
	CurrentPosition  BoardCoordinate
	PreviousPosition BoardCoordinate
}

type ByAction []AgentStep

func (a ByAction) Len() int           { return len(a) }
func (a ByAction) Less(i, j int) bool { return a[i].Action < a[j].Action }
func (a ByAction) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
