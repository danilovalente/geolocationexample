package repository

import (
	"testing"

	"github.com/danilovalente/geolocationexample/gateway/mongodb"
	"github.com/danilovalente/geolocationexample/repository"
)

func TestRepoMap_Add(t *testing.T) {
	type fields struct {
		repositories map[string]repository.Repository
	}
	type args struct {
		repositoryName string
		repository     repository.Repository
	}
	repos := make(map[string]repository.Repository)
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Add complete repository",
			fields: fields{repositories: repos},
			args: args{
				repositoryName: "TransportRepository",
				repository:     mongodb.TransportRepository{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoMap := repository.CreateRepoMap()
			repoMap.Add(tt.args.repositoryName, tt.args.repository)
			if repoMap.Count() == 0 {
				t.Error("Repository not added")
			}
		})
	}

}
