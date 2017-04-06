// Code generated by zanzibar
// @generated

package googlenow

import (
	"context"

	"github.com/uber/zanzibar/examples/example-gateway/build/clients"
	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"
)

// HandleCheckCredentialsRequest handles "/googlenow/check-credentials".
func HandleCheckCredentialsRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
	clients *clients.Clients,
) {
	if !req.CheckHeaders([]string{"x-uuid", "x-token"}) {
		return
	}

	headers := map[string]string{}

	workflow := CheckCredentialsEndpoint{
		Clients: clients,
		Logger:  req.Logger,
		Request: req,
	}

	_, err := workflow.Handle(ctx, headers)
	if err != nil {
		req.Logger.Warn("Workflow for endpoint returned error",
			zap.String("error", err.Error()),
		)
		res.SendErrorString(500, "Unexpected server error")
		return
	}

	res.WriteJSONBytes(202, nil)
}

// CheckCredentialsEndpoint calls thrift client GoogleNow.CheckCredentials
type CheckCredentialsEndpoint struct {
	Clients *clients.Clients
	Logger  *zap.Logger
	Request *zanzibar.ServerHTTPRequest
}

// Handle calls thrift client.
func (w CheckCredentialsEndpoint) Handle(
	ctx context.Context,
	headers map[string]string,
) (map[string]string, error) {

	_, err := w.Clients.GoogleNow.CheckCredentials(ctx, nil)
	if err != nil {
		w.Logger.Warn("Could not make client request",
			zap.String("error", err.Error()),
		)
		return nil, err
	}

	return nil, nil
}
