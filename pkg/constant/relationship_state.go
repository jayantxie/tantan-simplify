package constant

type State int8

const (
	Liked State = iota + 1
	Disliked
	Matched
	UnknownState
)

const RelationshipStr = "relationship"

const UserStr = "user"

func ParseState(state string) State {
	switch state {
	case "liked":
		return Liked
	case "disliked":
		return Disliked
	case "matched":
		return Matched
	default:
		return UnknownState
	}
}

func (s State) String() string {
	switch s {
	case Liked:
		return "liked"
	case Disliked:
		return "disliked"
	case Matched:
		return "matched"
	default:
		return "unknown"
	}
}

func StateFromInt(state int8) State {
	if state >= int8(UnknownState) && state < int8(Liked) {
		return UnknownState
	}
	return State(state)
}
