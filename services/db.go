package services

import (
	"fmt"
	"os"

	"github.com/supabase-community/supabase-go"
)

type IDbClient[T any] interface {
	Connect() (T, error)
}

type SupabaseClient struct {
}

func NewSupabaseClient() *SupabaseClient {
	return &SupabaseClient{}
}

func (s *SupabaseClient) Connect() (*supabase.Client, error) {
	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("SUPABASE_PUB_KEY")

	client, err := supabase.NewClient(url, key, &supabase.ClientOptions{})

	if err != nil {
		return nil, fmt.Errorf("failed to initialize db client: %w", err)
	}

	return client, nil
}
