package server

import (
	"github.com/ksandr84on/nera/consensus"
	consensusDev "github.com/ksandr84on/nera/consensus/dev"
	consensusDummy "github.com/ksandr84on/nera/consensus/dummy"
	consensusIBFT "github.com/ksandr84on/nera/consensus/ibft"
	"github.com/ksandr84on/nera/secrets"
	"github.com/ksandr84on/nera/secrets/awsssm"
	"github.com/ksandr84on/nera/secrets/gcpssm"
	"github.com/ksandr84on/nera/secrets/hashicorpvault"
	"github.com/ksandr84on/nera/secrets/local"
)

type ConsensusType string

const (
	DevConsensus   ConsensusType = "dev"
	IBFTConsensus  ConsensusType = "ibft"
	DummyConsensus ConsensusType = "dummy"
)

var consensusBackends = map[ConsensusType]consensus.Factory{
	DevConsensus:   consensusDev.Factory,
	IBFTConsensus:  consensusIBFT.Factory,
	DummyConsensus: consensusDummy.Factory,
}

// secretsManagerBackends defines the SecretManager factories for different
// secret management solutions
var secretsManagerBackends = map[secrets.SecretsManagerType]secrets.SecretsManagerFactory{
	secrets.Local:          local.SecretsManagerFactory,
	secrets.HashicorpVault: hashicorpvault.SecretsManagerFactory,
	secrets.AWSSSM:         awsssm.SecretsManagerFactory,
	secrets.GCPSSM:         gcpssm.SecretsManagerFactory,
}

func ConsensusSupported(value string) bool {
	_, ok := consensusBackends[ConsensusType(value)]

	return ok
}
