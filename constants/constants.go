package constants

const (
	DefaultFolderPath      = "."
	DefaultMgitObjectsPath = ".mgit/objects/"
	DefaultMgitPath        = ".mgit/"
	DefaultRefsFolder      = "refs/tags/"
)

var IgnoredPaths = []string{".git", ".hg", ".svn", ".mgit", ".idea", "mgit.exe"}
