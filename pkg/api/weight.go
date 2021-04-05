package api

// WeightService contains the methods of the weight service
type WeightService interface{}

// WeightRepository is what lets our service do db operations without knowing anything about the implementation
type WeightRepository interface{}

type weightService struct {
	storage WeightRepository
}

func NewWeightService(weightRepo WeightRepository) WeightService {
	return &weightService{
		storage: weightRepo,
	}
}
