package registry

import (
	"posts/adapter/controller"
	"posts/adapter/repository"
	"posts/usecase/usecase"
)

// connect use case with instance of a repo and repo with DB
func (r registry) NewPostController() controller.PostController {
	pu := usecase.NewPostUseCase(
		repository.NewPostRepository(r.db),
	)

	return controller.NewPostController(pu)
}
