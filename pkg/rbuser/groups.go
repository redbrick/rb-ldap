package rbuser

func gidToGroup(gid int) string {
	switch gid {
	case 100:
		return "committe"
	case 101:
		return "society"
	case 102:
		return "club"
	case 105:
		return "founder"
	case 107:
		return "associat"
	case 109:
		return "staff"
	case 1016:
		return "intersoc"
	case 1017:
		return "redbrick"
	case 1014:
		return "projects"
	case 31382:
		return "dcu"
	default:
		return "member"
	}
}

func groupToGID(group string) int {
	switch group {
	case "committe":
		return 100
	case "society":
		return 101
	case "club":
		return 102
	case "founder":
		return 105
	case "associat":
		return 107
	case "staff":
		return 109
	case "intersoc":
		return 1016
	case "redbrick":
		return 1017
	case "projects":
		return 1014
	case "dcu":
		return 31382
	case "member":
		return 103
	default:
		return 103
	}
}
