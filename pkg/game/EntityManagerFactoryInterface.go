package game

type EntityManagerFactoryInterface interface {
	NewEntityManager() EntityManagerInterface
}
