package glog

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"cloud.google.com/go/compute/metadata"
)

type Entry struct {
	Message  string `json:"message"`
	Severity string `json:"severity,omitempty"`
	Trace    string `json:"logging.googleapis.com/trace,omitempty"`
	SpanID   string `json:"logging.googleapis.com/spanId,omitempty"`

	HttpRequest *HttpRequest `json:"httpRequest,omitempty"`
}

type HttpRequest struct {
	RequestMethod string `json:"requestMethod,omitempty"`
	RequestUrl    string `json:"requestUrl,omitempty"`
	UserAgent     string `json:"userAgent,omitempty"`
	RemoteIp      string `json:"remoteIp,omitempty"`
	Referer       string `json:"referer,omitempty"`
	Protocol      string `json:"protocol,omitempty"`
}

var projectID string

func init() {
	projectID = os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		projectID, _ = metadata.ProjectID()
	}
	if projectID == "" {
		Println("Could not determine Google Cloud Project. Running without log correlation. For local use set the GOOGLE_CLOUD_PROJECT environment variable.")
	}
}

type Severity int32

const (
	// The log entry has no assigned severity level.
	DEFAULT Severity = iota * 100

	// Debug or trace information.
	DEBUG

	// Routine information, such as ongoing status or performance.
	INFO

	// Normal but significant events, such as start up, shut down, or a configuration change.
	NOTICE

	// Warning events might cause problems.
	WARNING

	// Error events are likely to cause problems.
	ERROR

	// Critical events cause more severe problems or outages.
	CRITICAL

	// A person must take an action immediately.
	ALERT

	// One or more systems are unusable.
	EMERGENCY
)

func (s Severity) String() string {
	switch s {
	default:
		return ""
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case NOTICE:
		return "NOTICE"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case CRITICAL:
		return "CRITICAL"
	case ALERT:
		return "ALERT"
	case EMERGENCY:
		return "EMERGENCY"
	}
}

func (s Severity) File() *os.File {
	if s >= ERROR {
		return os.Stderr
	} else {
		return os.Stdout
	}
}

func printGlog(s Severity, l Logger, msg string) {
	entry := Entry{
		Message:     msg,
		Severity:    s.String(),
		Trace:       l.trace,
		SpanID:      l.spanID,
		HttpRequest: l.request,
	}

	json.NewEncoder(s.File()).Encode(entry)
}

func print(s Severity, l Logger, v ...any) {
	printGlog(s, l, fmt.Sprint(v...))
}

func println(s Severity, l Logger, v ...any) {
	printGlog(s, l, fmt.Sprintln(v...))
}

func printf(s Severity, l Logger, format string, v ...any) {
	printGlog(s, l, fmt.Sprintf(format, v...))
}

// WithRequest returns a Logger with trace and executionID set from the request.
func WithRequest(r *http.Request) (l Logger) {
	if projectID != "" {
		traceHeader := r.Header.Get("X-Cloud-Trace-Context")
		traceParts := strings.Split(traceHeader, "/")
		if len(traceParts) > 0 && len(traceParts[0]) > 0 {
			l.trace = fmt.Sprintf("projects/%s/traces/%s", projectID, traceParts[0])
		}

		l.executionID = r.Header.Get("Function-Execution-Id")
	}

	l.request = &HttpRequest{
		RequestMethod: r.Method,
		RequestUrl:    r.RequestURI,
		UserAgent:     r.UserAgent(),
		RemoteIp:      r.RemoteAddr,
		Referer:       r.Referer(),
		Protocol:      r.Proto,
	}

	return l
}
