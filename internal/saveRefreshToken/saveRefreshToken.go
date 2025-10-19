package saverefreshtoken

import (
	"fmt"
	"os"

	"golang.org/x/oauth2"
)

func Save(token *oauth2.Token) error {
	f, err := os.Create("token.txt")
	if err != nil {
		return fmt.Errorf("Failed to create file to save token %w", err)
	}
	defer f.Close()
	f.WriteString(token.RefreshToken)
	return nil
}
