package healthcheck

type usecase struct{}

func newUsecase() *usecase {
	return &usecase{}
}

func (u *usecase) Get() string {
	return "pong"
}
