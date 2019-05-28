package repository

//Repository is the base interface for all Repositories
type Repository interface{}

//RepoMap is the type defining a map of Repositories
type RepoMap struct {
	repositories map[string]Repository
}

//Repos keeps all repositories available, initialized in the startup
var Repos RepoMap

//Add a repository By Name
func (repoMap *RepoMap) Add(name string, repository Repository) {
	repoMap.repositories[name] = repository
}

//Get a repository By Name
func (repoMap *RepoMap) Get(repositoryName string) Repository {
	return repoMap.repositories[repositoryName]
}

//Count returns the count of repositories registered
func (repoMap *RepoMap) Count() int {
	return len(repoMap.repositories)
}

//CreateRepoMap creates a new RepoMap instance
func CreateRepoMap() RepoMap {
	return RepoMap{repositories: make(map[string]Repository)}
}

func init() {
	Repos = CreateRepoMap()
}
