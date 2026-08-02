package mcp

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"mime"
	"net"
	"net/http"
	"net/url"
	"strings"
)

// HTTPHandler serves one stateless Streamable HTTP endpoint. This server uses
// single application/json replies because its tools do not emit progress.
type HTTPHandler struct {
	server *Server
}

func NewHTTPHandler(server *Server) *HTTPHandler {
	return &HTTPHandler{server: server}
}

func (rcvr *HTTPHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	if err := validateServer(rcvr.server); err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	if request.Method != http.MethodPost {
		response.Header().Set("Allow", http.MethodPost)
		writeHTTPError(response, nil, http.StatusMethodNotAllowed, -32600, "MCP endpoint accepts POST only", nil)
		return
	}
	if !validOrigin(request.Header.Get("Origin")) {
		writeHTTPError(response, nil, http.StatusForbidden, -32600, "Origin is not allowed", nil)
		return
	}
	if !acceptsMCPResponse(request.Header.Values("Accept")) {
		writeHTTPError(response, nil, http.StatusNotAcceptable, -32600, "Accept must include application/json and text/event-stream", nil)
		return
	}
	mediaType, _, err := mime.ParseMediaType(request.Header.Get("Content-Type"))
	if err != nil || mediaType != "application/json" {
		writeHTTPError(response, nil, http.StatusUnsupportedMediaType, -32600, "Content-Type must be application/json", nil)
		return
	}
	request.Body = http.MaxBytesReader(response, request.Body, maxMessageBytes)
	payload, err := io.ReadAll(request.Body)
	if err != nil {
		writeHTTPError(response, nil, http.StatusRequestEntityTooLarge, -32600, "Request exceeds the MCP message size limit", nil)
		return
	}
	message, err := decodeMessage(payload)
	if err != nil {
		writeHTTPBody(response, rcvr.server.Handle(request.Context(), payload))
		return
	}
	if result := validateHTTPHeaders(request, message); result != nil {
		writeHTTPBody(response, *result)
		return
	}
	writeHTTPBody(response, rcvr.server.Handle(request.Context(), payload))
}

func validateHTTPHeaders(request *http.Request, message rpcMessage) *HandleResult {
	versionHeader := request.Header.Get("MCP-Protocol-Version")
	methodHeader := request.Header.Get("Mcp-Method")
	nameHeader, nameErr := decodeMirroredHeader(request.Header.Get("Mcp-Name"))
	versionBody, _, metadataErr := requestProtocolVersion(message.Params)
	bodyName := decodeRequestName(message.Params)

	if versionHeader == "" || methodHeader == "" {
		result := HandleResult{Body: protocolErrorBody(message.ID, -32020, "Required MCP request header is missing", nil), HTTPStatus: http.StatusBadRequest}
		return &result
	}
	if metadataErr != nil || versionHeader != versionBody || methodHeader != message.Method {
		result := HandleResult{Body: protocolErrorBody(message.ID, -32020, "MCP request headers do not match the JSON-RPC body", nil), HTTPStatus: http.StatusBadRequest}
		return &result
	}
	if message.Method == "tools/call" || message.Method == "resources/read" || message.Method == "prompts/get" {
		if request.Header.Get("Mcp-Name") == "" || nameErr != nil || nameHeader != bodyName {
			result := HandleResult{Body: protocolErrorBody(message.ID, -32020, "Mcp-Name header does not match the JSON-RPC body", nil), HTTPStatus: http.StatusBadRequest}
			return &result
		}
	}
	return nil
}

func validOrigin(raw string) bool {
	if raw == "" {
		return true
	}
	origin, err := url.Parse(raw)
	if err != nil || (origin.Scheme != "http" && origin.Scheme != "https") || origin.User != nil {
		return false
	}
	host := strings.TrimSuffix(strings.ToLower(origin.Hostname()), ".")
	if host == "localhost" {
		return true
	}
	ip := net.ParseIP(host)
	return ip != nil && ip.IsLoopback()
}

func acceptsMCPResponse(values []string) bool {
	jsonAccepted, eventStreamAccepted := false, false
	for _, value := range values {
		for _, part := range strings.Split(value, ",") {
			mediaType, params, err := mime.ParseMediaType(strings.TrimSpace(part))
			if err != nil || params["q"] == "0" {
				continue
			}
			switch mediaType {
			case "application/json", "*/*":
				jsonAccepted = true
			case "text/event-stream":
				eventStreamAccepted = true
			}
		}
	}
	return jsonAccepted && eventStreamAccepted
}

func decodeMirroredHeader(value string) (string, error) {
	if strings.HasPrefix(value, "=?base64?") && strings.HasSuffix(value, "?=") {
		encoded := strings.TrimSuffix(strings.TrimPrefix(value, "=?base64?"), "?=")
		decoded, err := base64.StdEncoding.DecodeString(encoded)
		if err != nil {
			return "", err
		}
		return string(decoded), nil
	}
	if strings.TrimSpace(value) != value {
		return "", errors.New("unsafe mirrored header")
	}
	for _, character := range []byte(value) {
		if character < 0x20 || character > 0x7e {
			return "", errors.New("unsafe mirrored header")
		}
	}
	return value, nil
}

func writeHTTPBody(response http.ResponseWriter, result HandleResult) {
	if result.Notification {
		response.WriteHeader(http.StatusAccepted)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("Cache-Control", "no-store")
	response.WriteHeader(result.HTTPStatus)
	_, _ = response.Write(result.Body)
}

func writeHTTPError(response http.ResponseWriter, id json.RawMessage, status, code int, message string, data any) {
	writeHTTPBody(response, HandleResult{Body: protocolErrorBody(id, code, message, data), HTTPStatus: status})
}
