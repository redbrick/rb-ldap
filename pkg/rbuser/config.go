package rbuser

const timeLayout = "2006-01-02 15:04:05"

// Shell defaults
const expiredShell = "/usr/local/shells/expired"
const noLoginShell = "/usr/local/shells/disusered"
const defaultShell = "/usr/local/shells/shell"

// user groups
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
