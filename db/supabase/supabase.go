package supabase

import (
	"github.com/supabase-community/supabase-go"
)

func SetupDB(ApiUrl string, ApiKey string) (*supabase.Client, error) {
	client, err := supabase.NewClient(ApiUrl, ApiKey, &supabase.ClientOptions{})
	if err != nil {
		return nil, err
	}
	return client, nil
}
