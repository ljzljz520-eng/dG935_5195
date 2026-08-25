package review

func NeedsFollowup(decision string) bool { return decision == "changes_requested" }
func Label(decision string) string {
	switch decision {
	case "approved":
		return "通过"
	case "rejected":
		return "驳回"
	case "changes_requested":
		return "需修改"
	default:
		return "待处理"
	}
}
func AllowedReviewer(name string) bool { return len(name) > 1 }
