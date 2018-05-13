package rbuser

type group struct {
	name string
	id   int
}

var associatGroup = group{"associat", 107}
var clubGroup = group{"club", 102}
var committeeGroup = group{"committe", 100}
var dcuGroup = group{"dcu", 31382}
var founderGroup = group{"founder", 105}
var intersocGroup = group{"intersoc", 1016}
var memberGroup = group{"member", 103}
var projectsGroup = group{"projects", 1014}
var redbrickGroup = group{"redbrick", 1017}
var societyGroup = group{"society", 101}
var staffGroup = group{"staff", 109}

var groups = []group{
	associatGroup,
	clubGroup,
	committeeGroup,
	dcuGroup,
	founderGroup,
	intersocGroup,
	memberGroup,
	projectsGroup,
	redbrickGroup,
	societyGroup,
	staffGroup,
}

var userGroups = []group{
	associatGroup,
	memberGroup,
	staffGroup,
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
