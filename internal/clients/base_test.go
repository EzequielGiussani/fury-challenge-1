package clients

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/mercadolibre/fury_go-meli-toolkit-goutils/apierrors"
)

func TestBuildURL(t *testing.T) {
	type args struct {
		params []string
		query  url.Values
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 apierrors.ApiError
	}{
		{
			name: "Test build URL Success",
			args: args{
				params: []string{"/items", "mla-1"},
				query: url.Values{
					"attributes": []string{"id,prices"},
				},
			},
			want:  "/items/mla-1?attributes=id%2Cprices",
			want1: nil,
		},
		{
			name: "Test build URL with no params",
			args: args{
				params: []string{},
				query: url.Values{
					"attributes": []string{"id,prices"},
				},
			},
			want:  "",
			want1: apierrors.NewInternalServerApiError("invalid params URL", nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := BuildURL(tt.args.params, tt.args.query)
			if got != tt.want {
				t.Errorf("BuildURL() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("BuildURL() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
