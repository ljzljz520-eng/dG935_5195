package domain

func CanTransition(from, to string) bool {
	if from == to {
		return true
	}
	if from == "open" && (to == "review" || to == "archived") {
		return true
	}
	if from == "review" && to == "archived" {
		return true
	}
	return false
}
func NormalizeStatus(status string) string {
	switch status {
	case "open", "review", "archived":
		return status
	default:
		return "open"
	}
}
func ReviewDecisionValid(d string) bool {
	return d == "approved" || d == "changes_requested" || d == "rejected"
}
