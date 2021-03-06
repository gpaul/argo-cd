package version

import (
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"

	"github.com/argoproj/argo-cd/util/helm"
	"github.com/argoproj/argo-cd/util/kube"
	"github.com/argoproj/argo-cd/util/kustomize"

	"github.com/argoproj/argo-cd/common"
	"github.com/argoproj/argo-cd/pkg/apiclient/version"
	ksutil "github.com/argoproj/argo-cd/util/ksonnet"
)

type Server struct {
	ksonnetVersion   string
	kustomizeVersion string
	helmVersion      string
	kubectlVersion   string
}

// Version returns the version of the API server
func (s *Server) Version(context.Context, *empty.Empty) (*version.VersionMessage, error) {
	vers := common.GetVersion()
	if s.ksonnetVersion == "" {
		ksonnetVersion, err := ksutil.Version()
		if err != nil {
			return nil, err
		}
		s.ksonnetVersion = ksonnetVersion
	}
	if s.kustomizeVersion == "" {
		kustomizeVersion, err := kustomize.Version()
		if err != nil {
			return nil, err
		}
		s.kustomizeVersion = kustomizeVersion
	}
	if s.helmVersion == "" {
		helmVersion, err := helm.Version()
		if err != nil {
			return nil, err
		}
		s.helmVersion = helmVersion
	}
	if s.kubectlVersion == "" {
		kubectlVersion, err := kube.Version()
		if err != nil {
			return nil, err
		}
		s.kubectlVersion = kubectlVersion
	}
	return &version.VersionMessage{
		Version:          vers.Version,
		BuildDate:        vers.BuildDate,
		GitCommit:        vers.GitCommit,
		GitTag:           vers.GitTag,
		GitTreeState:     vers.GitTreeState,
		GoVersion:        vers.GoVersion,
		Compiler:         vers.Compiler,
		Platform:         vers.Platform,
		KsonnetVersion:   s.ksonnetVersion,
		KustomizeVersion: s.kustomizeVersion,
		HelmVersion:      s.helmVersion,
		KubectlVersion:   s.kubectlVersion,
	}, nil
}

// AuthFuncOverride allows the version to be returned without auth
func (s *Server) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}
