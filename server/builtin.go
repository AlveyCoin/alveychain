package server

import (
	"github.com/alveycoin/alveychain/consensus"
	consensusDev "github.com/alveycoin/alveychain/consensus/dev"
	consensusDummy "github.com/alveycoin/alveychain/consensus/dummy"
	consensusIBFT "github.com/alveycoin/alveychain/consensus/ibft"
	"github.com/alveycoin/alveychain/secrets"
	"github.com/alveycoin/alveychain/secrets/awsssm"
	"github.com/alveycoin/alveychain/secrets/gcpssm"
	"github.com/alveycoin/alveychain/secrets/hashicorpvault"
	"github.com/alveycoin/alveychain/secrets/local"
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
