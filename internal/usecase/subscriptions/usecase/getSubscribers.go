package usecase

import (
	"context"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/google/uuid"
)

func (u *Usecase) GetSubscribers(ctx context.Context, userID uuid.UUID) ([]models.User, error) {

	//Тут когда в будущем у юзеров появятся настройки приватности, желательно делать проверку, что мы можем смотреть чужие
	//Подписки и подписчиков. То есть будет отдельная табличка и с настройками юзера, мы будем по айдишке таргет юзера
	//Ее доставать и проверять там поле. Пока что смотреть могут все, но потом нужно дописать

	subscribers, err := u.subsRepo.GetSubscribers(ctx, userID)
	if err != nil {
		return nil, u.handleError(err)
	}

	return subscribers, err
}
