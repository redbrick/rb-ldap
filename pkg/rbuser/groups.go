package rbuser

type group struct {
	name string
	id   int
}

func gidToGroup(gid int) string {
	for _, group := range groups {
		if group.id == gid {
			return group.name
		}
	}
	return ""
}

func groupToGID(groupName string) int {
	for _, group := range groups {
		if group.name == groupName {
			return group.id
		}
	}
	return 0
}
