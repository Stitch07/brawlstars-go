package brawlstars

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
)

// BaseURL is the base API URL all requests go to
const BaseURL = "https://api.brawlapi.cf/v1/"

// New creates a new instance of a Brawl Stars client with an empty http.Client{}
func New(token string) *Client {
	return &Client{
		HTTP:     http.DefaultClient,
		Token:    token,
		Compress: true,

		ratelimit: Ratelimit{
			Remaining: new(int32),
		},
	}
}

// Client is an API client. To avoid accepting GZIP encoded response bodies, set Compress to false
type Client struct {
	HTTP     *http.Client // the HTTP client used to create all requests
	Token    string       // the authentication token used to authorize requests
	Compress bool         // whether the client should accept gzip encoding from the API

	ratelimit Ratelimit
}

// GetPlayer returns a Player instance from the provided player tag
func (c *Client) GetPlayer(tag string) (p *Player, err error) {
	tag, ok := ValidateTag(tag)
	if !ok {
		err = errors.New("invalid tag provided")
		return
	}
	err = c.doJSON(http.MethodGet, "player?tag="+tag, &p)
	return
}

// SearchPlayer returns a slice of Players matching the provided query
func (c *Client) SearchPlayer(query string) (p []*PlayerSearch, err error) {
	err = c.doJSON(http.MethodGet, "player/search?name="+query, &p)
	return
}

// GetClub returns a Club instance from the provided club tag
func (c *Client) GetClub(tag string) (club *Club, err error) {
	tag, ok := ValidateTag(tag)
	if !ok {
		err = errors.New("invalid tag provided")
		return
	}
	err = c.doJSON(http.MethodGet, "club?tag="+tag, &club)
	return
}

// SearchClub returns a slice of Clubs matching the provided query
func (c *Client) SearchClub(query string) (clubs []*ClubSearch, err error) {
	err = c.doJSON(http.MethodGet, "club/search?name="+query, &clubs)
	return
}

// GetEvents retrieves the current and upcoming events
// parameters:
//   - type: The type of events to return (current/upcoming/all)
func (c *Client) GetEvents(evtType string) (events *Events, err error) {
	err = c.doJSON(http.MethodGet, "events?type="+evtType, &events)
	return
}

// TopClubs returns the top <count> clubs in the leaderboard
func (c *Client) TopClubs(count uint) (clubs []*TopClub, err error) {
	err = c.doJSON(http.MethodGet, fmt.Sprint("leaderboards/clubs?count=", count), &clubs)
	return
}

// TopPlayers returns the top <count> players in the leaderboard
func (c *Client) TopPlayers(count uint) (players []*TopPlayer, err error) {
	err = c.doJSON(http.MethodGet, fmt.Sprint("leaderboards/players?count=", count), &players)
	return
}

// TopPlayersByBrawler returns the top <count> players in the <brawler> leaderboard
func (c *Client) TopPlayersByBrawler(brawler string, count uint) (players []*TopPlayer, err error) {
	err = c.doJSON(http.MethodGet, fmt.Sprint("leaderboards/players?count=", count, "&brawler=", brawler), &players)
	return
}

// doJSON creates an http request to the Brawl API and decodes the response body into a JSON object
func (c *Client) doJSON(method, path string, respBody interface{}) error {
	// if we can't do any more requests, wait until we can
	remaining := atomic.LoadInt32(c.ratelimit.Remaining)
	if remaining <= 0 { // this shouldn't be less than 0, but possible edge case
		wait := time.Until(c.ratelimit.Reset)
		time.Sleep(wait)
	}

	req, _ := http.NewRequest(method, BaseURL+path, nil)
	req.Header.Set("Authorization", c.Token)
	req.Header.Set("User-Agent", "brawlstars-go (https://github.com/Soumil07/brawlstars-go)")
	if c.Compress {
		req.Header.Set("Accept-Encoding", "gzip")
	}
	response, err := c.HTTP.Do(req)
	if err != nil {
		return fmt.Errorf("http request error: %v", err)
	}
	defer response.Body.Close()

	// handle status codes
	switch {
	case response.StatusCode >= 200 && response.StatusCode <= 300:
		// decompress gzip body if set to true
		var body = response.Body
		if c.Compress {
			body, err = gzip.NewReader(body)
			if err != nil {
				return fmt.Errorf("gzip compress error: %v", err)
			}
			defer body.Close()
		}
		// update ratelimit headers
		remaining, _ := strconv.Atoi(req.Header.Get("X-Ratelimit-Remaining"))
		reset, _ := strconv.ParseInt(req.Header.Get("X-Ratelimit-Reset"), 10, 64)
		c.ratelimit.Reset = time.Unix(reset, 0)
		atomic.StoreInt32(c.ratelimit.Remaining, int32(remaining))
		// decode JSON body
		return json.NewDecoder(body).Decode(respBody)

	default:
		return fmt.Errorf("unexpected status code: %s", response.Status)
	}
}
