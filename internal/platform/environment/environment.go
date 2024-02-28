package environment

type Environment int

const (
	Development Environment = iota
	Production
)

func GetFromString(s string) Environment {
	switch s {
	case "production":
		return Production
	default:
		return Development
	}
}
