package config

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

func TestIsParseEmptyIDDiagnostic(t *testing.T) {
	cases := map[string]struct {
		diags []*tfprotov6.Diagnostic
		want  bool
	}{
		"empty endpoint id": {
			diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Error Parsing Endpoint ID",
				Detail:   `Could not parse endpoint ID "" as integer: strconv.Atoi: parsing "": invalid syntax`,
			}},
			want: true,
		},
		"empty cluster id": {
			diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Error Parsing Cluster ID",
				Detail:   `parsing "": invalid syntax`,
			}},
			want: true,
		},
		"populated id parse error is not suppressed": {
			diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Error Parsing Endpoint ID",
				Detail:   `Could not parse endpoint ID "abc" as integer`,
			}},
			want: false,
		},
		"warning is ignored": {
			diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityWarning,
				Summary:  "Error Parsing Endpoint ID",
				Detail:   `parsing ""`,
			}},
			want: false,
		},
		"unrelated error": {
			diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Boom",
				Detail:   "network down",
			}},
			want: false,
		},
		"empty input": {diags: nil, want: false},
		"nil entry": {
			diags: []*tfprotov6.Diagnostic{nil},
			want:  false,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			if got := isParseEmptyIDDiagnostic(tc.diags); got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}
