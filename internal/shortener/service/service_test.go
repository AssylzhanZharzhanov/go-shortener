package service

import (
	"context"
	"testing"

	"github.com/AssylzhanZharzhanov/go-shortener/internal/mocks"
	"github.com/AssylzhanZharzhanov/go-shortener/internal/shortener/domain"
	"github.com/go-test/deep"

	"github.com/golang/mock/gomock"
)

func TestService_CreateShortenURL(t *testing.T) {
	var (
		ctx       = context.Background()
		validLink = &domain.Link{
			OriginalURL: "test_original_url",
			Hash:        "test",
		}
	)

	// Setup mocks.
	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mocks repository.
	repoStub := mocks.NewMockShortenerRepository(stubCtrl)
	redisStub := mocks.NewMockShortenerRedisRepository(stubCtrl)
	service := NewShortenerService(repoStub, redisStub)

	repoStub.EXPECT().Create(validLink).
		Return(&validLink, nil).AnyTimes()

	type arguments struct {
		link *domain.Link
	}

	type result struct {
		expected *domain.Link
	}

	tests := []struct {
		name        string
		arguments   arguments
		result      result
		expectError bool
	}{
		{
			name: "Success: shorten url created",
			arguments: arguments{
				link: validLink,
			},
			result: result{
				expected: validLink,
			},
			expectError: false,
		},
		{
			name: "Fail: empty original url",
			arguments: arguments{
				link: &domain.Link{
					OriginalURL: "",
					Hash:        "test_hash",
				},
			},
			result:      result{},
			expectError: true,
		},
		{
			name: "Fail: empty hash",
			arguments: arguments{
				link: &domain.Link{
					OriginalURL: "test_original_url",
					Hash:        "",
				},
			},
			result:      result{},
			expectError: true,
		},
		{
			name: "Fail: nil link",
			arguments: arguments{
				link: nil,
			},
			result:      result{},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.arguments
			expected := test.result

			link, err := service.CreateShortenURL(ctx, args.link)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error: %s", err)
				}
				actual := result{
					expected: link,
				}
				if diff := deep.Equal(expected, actual); diff != nil {
					t.Error(diff)
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}

func TestService_Get(t *testing.T) {
	var (
		ctx             = context.Background()
		validShortenURL = "test_shorten_url"
		validLink       = &domain.Link{
			OriginalURL: "test_original_url",
			Hash:        "test",
		}
	)

	// Setup mocks.
	stubCtrl := gomock.NewController(t)
	defer stubCtrl.Finish()

	// Mocks repository.
	repoStub := mocks.NewMockShortenerRepository(stubCtrl)
	redisStub := mocks.NewMockShortenerRedisRepository(stubCtrl)

	service := NewShortenerService(repoStub, redisStub)

	repoStub.EXPECT().Get(validLink).
		Return(&validLink, nil).AnyTimes()

	type arguments struct {
		shortenURL string
	}

	type result struct {
		expected *domain.Link
	}

	tests := []struct {
		name        string
		arguments   arguments
		result      result
		expectError bool
	}{
		{
			name: "Success: ",
			arguments: arguments{
				shortenURL: validShortenURL,
			},
			result: result{
				expected: validLink,
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.arguments
			expected := test.result

			link, err := service.Get(ctx, args.shortenURL)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error: %s", err)
				}
				actual := result{
					expected: link,
				}
				if diff := deep.Equal(expected, actual); diff != nil {
					t.Error(diff)
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}
