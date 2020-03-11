package feed

type UseCase interface {
	UpdatePopularPublications()
	GenerateUsersFeed()
}