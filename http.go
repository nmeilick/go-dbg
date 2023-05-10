package dbg

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/kballard/go-shellquote"
	"github.com/motemen/go-loghttp"
)

func LoggingHTTPClient(t http.RoundTripper) *http.Client {
	t = &loghttp.Transport{
		Transport: t,
		LogRequest: func(req *http.Request) {
			dump, err := httputil.DumpRequestOut(req, true)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Fatal: failed to dump outgoing request: %s %s: %s\n", req.Method, req.URL, err)
				os.Exit(1)
			}
			fmt.Println(dump)
		},
		LogResponse: func(resp *http.Response) {
			dump, err := httputil.DumpResponse(resp, true)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Fatal: failed to dump response: %s\n", err)
				os.Exit(1)
			}
			fmt.Println(dump)
		},
	}

	return &http.Client{Transport: t}
}

func curlCommandFromRequest(req *http.Request) (string, error) {
	if req == nil {
		return "", fmt.Errorf("request is nil")
	}

	args := []string{
		"-X", req.Method,
	}

	for key, values := range req.Header {
		for _, value := range values {
			args = append(args, "-H", fmt.Sprintf("%s: %s", key, value))
		}
	}

	var dataString string
	if req.Body != nil {
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			return "", fmt.Errorf("failed to read request body: %w", err)
		}
		dataString = string(bodyBytes)
		args = append(args, "--data-binary", "@-") // Use stdin for data-binary
	}

	args = append(args, req.URL.String())
	curlCmd := "curl " + shellquote.Join(args...)

	return fmt.Sprintf("cat <<EOT | %s\n%s\nEOT", curlCmd, dataString), nil
}
