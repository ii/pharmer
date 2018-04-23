package fake

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"sync"

	apiv1 "github.com/pharmer/pharmer/apis/v1"
	api "github.com/pharmer/pharmer/apis/v1alpha1"
	"github.com/pharmer/pharmer/store"
	clusterv1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
)

const (
	UID = "fake"
)

func init() {
	store.RegisterProvider(UID, func(ctx context.Context, cfg *api.PharmerConfig) (store.Interface, error) {
		return New(), nil
	})
}

type FakeStore struct {
	credentials  store.CredentialStore
	clusters     store.ClusterStore
	nodeGroups   map[string]store.NodeGroupStore
	machineSet   map[string]store.MachineSetStore
	certificates map[string]store.CertificateStore
	sshKeys      map[string]store.SSHKeyStore

	mux sync.Mutex
}

var _ store.Interface = &FakeStore{}

func New() store.Interface {
	return &FakeStore{
		nodeGroups:   map[string]store.NodeGroupStore{},
		machineSet:   map[string]store.MachineSetStore{},
		certificates: map[string]store.CertificateStore{},
		sshKeys:      map[string]store.SSHKeyStore{},
	}
}

func (s *FakeStore) Credentials() store.CredentialStore {
	s.mux.Lock()
	defer s.mux.Unlock()

	if s.credentials == nil {
		s.credentials = &credentialFileStore{container: map[string]*api.Credential{}}
	}
	return s.credentials
}

func (s *FakeStore) Clusters() store.ClusterStore {
	s.mux.Lock()
	defer s.mux.Unlock()

	if s.clusters == nil {
		s.clusters = &clusterFileStore{container: map[string]*apiv1.Cluster{}}
	}
	return s.clusters
}

func (s *FakeStore) NodeGroups(cluster string) store.NodeGroupStore {
	s.mux.Lock()
	defer s.mux.Unlock()

	if _, found := s.nodeGroups[cluster]; !found {
		s.nodeGroups[cluster] = &nodeGroupFileStore{container: map[string]*api.NodeGroup{}, cluster: cluster}
	}
	return s.nodeGroups[cluster]
}

func (s *FakeStore) MachineSet(cluster string) store.MachineSetStore {
	s.mux.Lock()
	defer s.mux.Unlock()
	if _, found := s.machineSet[cluster]; !found {
		s.machineSet[cluster] = &machineSetFileStore{container: map[string]*clusterv1.MachineSet{}, cluster: cluster}
	}
	return s.machineSet[cluster]
}

func (s *FakeStore) Certificates(cluster string) store.CertificateStore {
	s.mux.Lock()
	defer s.mux.Unlock()

	if _, found := s.certificates[cluster]; !found {
		s.certificates[cluster] = &certificateFileStore{certs: map[string]*x509.Certificate{}, keys: map[string]*rsa.PrivateKey{}, cluster: cluster}
	}
	return s.certificates[cluster]
}

func (s *FakeStore) SSHKeys(cluster string) store.SSHKeyStore {
	s.mux.Lock()
	defer s.mux.Unlock()

	if _, found := s.sshKeys[cluster]; !found {
		s.sshKeys[cluster] = &sshKeyFileStore{container: map[string][]byte{}, cluster: cluster}
	}
	return s.sshKeys[cluster]
}
