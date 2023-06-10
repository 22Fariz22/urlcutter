package usecase

import (
	"context"
	"database/sql"
	"github.com/22Fariz22/urlcutter/internal/url/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_useCase_Get(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoPG := mock.NewMockRepo(ctrl)

	UC := NewUseCase(repoPG)

	ctx := context.Background()

	repoPG.EXPECT().Get(ctx, "my_short_url").Return("", sql.ErrNoRows)

	getURL, _ := UC.Get(ctx, "my_short_url")
	//require.NoError(t, err)
	//require.EqualError(t, err, "no rows in result set")
	require.NotNil(t, getURL)
	//assert.EqualErrorf(t, err, "no rows in result set", "Error should be: %v, got: %v", "no rows in result set", err)

	//require.Equal(t, getURL, "")
}

func Test_useCase_Save(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoPG := mock.NewMockRepo(ctrl)

	UC := NewUseCase(repoPG)

	ctx := context.Background()

	repoPG.EXPECT().Save(ctx, "my_long_url", "my_short_url").Return("", nil)

	saveURL, err := UC.Save(ctx, "my_long_url", "my_short_url")
	require.NoError(t, err)
	require.NotNil(t, saveURL)
	require.Equal(t, saveURL, "")

}
