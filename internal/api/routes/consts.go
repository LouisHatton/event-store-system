package routes

const (
	ProjectIdParam    = "projectId"
	ProjectUrlParam   = "/{projectId}"
	ProjectPathPrefix = "/projects"
	CreateProjectPath = ProjectPathPrefix
	ProjectIdPath     = ProjectPathPrefix + ProjectUrlParam
)

const (
	ConnectionIdParam     = "connectionId"
	ConnectionUrlParam    = "/{connectionId}"
	ConnectionsPathPrefix = "/connections"
	CreateConnectionsPath = ProjectIdPath + ConnectionsPathPrefix
	ConnectionsListPath   = ProjectIdPath + ConnectionsPathPrefix + "/list"
	ConnectionIdPath      = ProjectIdPath + ConnectionsPathPrefix + ConnectionUrlParam
)
