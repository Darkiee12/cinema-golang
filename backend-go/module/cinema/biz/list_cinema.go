package cinemabuisness

import (
	"cinema/common"
	cinemamodel "cinema/module/cinema/model"
	"context"
)

type ListCinemaStore interface {
	ListCinemaWithCondition(
		context context.Context,
		filter *cinemamodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]cinemamodel.Cinema, error)
}
type listCinemaBusiness struct {
	store ListCinemaStore
}

func NewListCinemaBusiness(store ListCinemaStore) *listCinemaBusiness {
	return &listCinemaBusiness{store: store}
}
func (business *listCinemaBusiness) ListCinema(
	context context.Context,
	filter *cinemamodel.Filter,
	paging *common.Paging,
) ([]cinemamodel.Cinema, error) {
	res, err := business.store.ListCinemaWithCondition(context, filter, paging, "Owner")
	if err != nil {
		return nil, err
	}
	return res, nil
}
