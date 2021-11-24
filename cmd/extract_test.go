package cmd

import (
	"reflect"
	"testing"
)

func Test_extractMDImages(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []*MDImage
	}{
		{
			args: args{`foo ![imag\]e](image "2.p'ng "hel\"l'o wo\)rld")  ![](image2.png)barGroup world.`},
			want: []*MDImage{
				{
					altText: "imag]e",
					path:    "image \"2.p'ng",
					title:   "hel\"l'o wo\\)rld",
				},
				{
					altText: "",
					path:    "image2.png",
					title:   "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractMDImages(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractMDImages() = %v, want %v", got, tt.want)
			}
		})
	}
}
