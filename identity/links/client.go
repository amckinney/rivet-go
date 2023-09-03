// This file was auto-generated by Fern from our API Definition.

package links

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	fmt "fmt"
	rivetgo "github.com/rivet-gg/rivet-go"
	core "github.com/rivet-gg/rivet-go/core"
	identity "github.com/rivet-gg/rivet-go/identity"
	io "io"
	http "net/http"
	url "net/url"
)

type Client struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
	}
}

// Begins the process for linking an identity with the Rivet Hub.
// # Importance of Linking Identities
// When an identity is created via `rivet.api.identity#SetupIdentity`, the identity is temporary
// and is not shared with other games the user plays.
// In order to make the identity permanent and synchronize the identity with
// other games, the identity must be linked with the hub.
// # Linking Process
// The linking process works by opening `identity_link_url` in a browser then polling
// `rivet.api.identity#GetGameLink` to wait for it to complete.
// This is designed to be as flexible as possible so `identity_link_url` can be opened
// on any device. For example, when playing a console game, the user can scan a
// QR code for `identity_link_url` to authenticate on their phone.
func (c *Client) Prepare(ctx context.Context) (*identity.PrepareGameLinkResponse, error) {
	baseURL := "https://identity.api.rivet.gg/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "game-links"

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 500:
			value := new(rivetgo.InternalError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 429:
			value := new(rivetgo.RateLimitError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 403:
			value := new(rivetgo.ForbiddenError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 408:
			value := new(rivetgo.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(rivetgo.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 400:
			value := new(rivetgo.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *identity.PrepareGameLinkResponse
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPost,
		nil,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Returns the current status of a linking process. Once `status` is `complete`, the identity's profile should be fetched again since they may have switched accounts.
func (c *Client) Get(ctx context.Context, request *identity.GetGameLinkRequest) (*identity.GetGameLinkResponse, error) {
	baseURL := "https://identity.api.rivet.gg/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "game-links"

	queryParams := make(url.Values)
	queryParams.Add("identity_link_token", fmt.Sprintf("%v", request.IdentityLinkToken))
	queryParams.Add("watch_index", fmt.Sprintf("%v", request.WatchIndex))
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 500:
			value := new(rivetgo.InternalError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 429:
			value := new(rivetgo.RateLimitError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 403:
			value := new(rivetgo.ForbiddenError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 408:
			value := new(rivetgo.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(rivetgo.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 400:
			value := new(rivetgo.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *identity.GetGameLinkResponse
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		request,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}
