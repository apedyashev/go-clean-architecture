package registry

import (
	"posts/adapter/controller"
	"posts/infrastructure/datastore"
)

type Registry interface {
	NewAppController() controller.AppController
}

type registry struct {
	db *datastore.InMemoryDB
}

func NewRegistry(db *datastore.InMemoryDB) registry {
	return registry{db}
}

// Creates the AppControllwer that exposes all the controllers belonging to this domain
// Each controller will connect useses with ports(repositories), and repository with datastore
func (r registry) NewAppController() controller.AppController {
	return controller.AppController{
		Post: r.NewPostController(),
	}
}
