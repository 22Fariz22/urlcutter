package grpc

import (
	"context"
	"github.com/22Fariz22/urlcutter/internal/config"
	"github.com/22Fariz22/urlcutter/internal/url/mock"
	urlcutter "github.com/22Fariz22/urlcutter/proto"
	"github.com/golang/mock/gomock"
	gonanoid "github.com/matoous/go-nanoid"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_service_Post(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	UC := mock.NewMockUseCase(ctrl)
	cfg := &config.Config{}
	serverGRPC := NewServerGRPC(nil, cfg, UC)

	//reqValue := urlcutter.ShortURL{ShortURL: "mock_short"}

	t.Run("Post", func(t *testing.T) {
		t.Parallel()
		//respValue := urlcutter.ShortURL{ShortURL: "mock_short"}

		ctx := context.Background()

		id, err := gonanoid.Nanoid(10)
		require.NoError(t, err)

		UC.EXPECT().Save(ctx, gomock.Any(), gomock.Any()).Return(cfg.BaseURL+"/"+id, nil)

		response, err := serverGRPC.Post(ctx, &urlcutter.LongURL{LongURL: "mock_long"})
		require.NoError(t, err)
		require.NotNil(t, response)
	})

}

func Test_service_Get(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	UC := mock.NewMockUseCase(ctrl)
	cfg := &config.Config{}
	serverGRPC := NewServerGRPC(nil, cfg, UC)

	t.Run("Get", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		UC.EXPECT().Get(ctx, gomock.Any()).Return("", nil)

		response, err := serverGRPC.Get(ctx, &urlcutter.ShortURL{ShortURL: "mock_short_url"})
		require.NoError(t, err)
		require.NotNil(t, response)
	})
}
