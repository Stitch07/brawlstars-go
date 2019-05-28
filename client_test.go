package brawlstars

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var client = New(os.Getenv("TOKEN"))

func TestClient_GetPlayer(t *testing.T) {
	player, err := client.GetPlayer("#Y2QPGG")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "Lex_YouTube", player.Name)
}

func TestRatelimit(t *testing.T) {
	t.Skip()
	// do 10 requests to check the lib properly handles ratelimits
	for i := 0; i <= 10; i++ {
		player, err := client.GetPlayer("#Y2QPGG")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "Lex_YouTube", player.Name)
	}
}

func TestClient_GetClub(t *testing.T) {
	club, err := client.GetClub("QCGUUYJ")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "Boom Bandits", club.Name)
}

func TestClient_GetEvents(t *testing.T) {
	_, err := client.GetEvents("all")
	assert.Equal(t, nil, err)
}
