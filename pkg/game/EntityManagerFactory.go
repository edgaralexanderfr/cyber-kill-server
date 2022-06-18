package game

type EntityManagerFactory struct{}

func (entityManagerFactory *EntityManagerFactory) NewEntityManager() EntityManagerInterface {
	return nil
}
