/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package tbls

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math"
	"net"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/IBM/TSS/mpc/bls"
	comm "github.com/IBM/TSS/net"
	"github.com/IBM/TSS/testutil/tlsgen"
	. "github.com/IBM/TSS/threshold"
	. "github.com/IBM/TSS/types"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestThresholdBLS(t *testing.T) {
	var commParties []*comm.Party
	var signers []*tlsgen.CertKeyPair
	var loggers []*commLogger
	var listeners []net.Listener
	var stopFuncs []func()

	n := 3

	members, certPool, loggers, signers, listeners, commParties, membershipFunc, parties, kgf := setup(t, n, loggers, signers, listeners, commParties)

	for id := 1; id <= n; id++ {
		stop, s := createParty(id, kgf, signers[id-1], n, certPool, listeners, loggers, commParties, membershipFunc)
		parties = append(parties, s)
		stopFuncs = append(stopFuncs, stop)
	}

	defer func() {
		for _, stop := range stopFuncs {
			stop()
		}
	}()

	shares, start := keygen(t, parties, n)

	elapsed := time.Since(start)
	t.Log("DKG elapsed", elapsed)

	// Create the threshold signers.
	thresholdSigners := make([]*bls.TBLS, n)
	for id := 1; id <= n; id++ {
		thresholdSigners[id-1] = &bls.TBLS{
			Logger: logger(id, t.Name()),
			Party:  uint16(id),
		}
	}

	// Initialize them with a nil send function
	for i, signer := range thresholdSigners {
		signer.Init(members, n-1, nil)
		signer.SetShareData(shares[i])
	}

	var signatures [][]byte

	// Sign a message
	msg := []byte("Three can keep a secret, if two of them are dead.")
	digest := sha256Digest(msg)
	for _, signer := range thresholdSigners {
		sig, err := signer.Sign(nil, digest)
		assert.NoError(t, err)
		signatures = append(signatures, sig)
	}

	// Lastly ensure the message verifies with a threshold public key obtained from each signer
	for _, signer := range thresholdSigners {
		pk, err := signer.ThresholdPK()
		assert.NoError(t, err)

		var v bls.Verifier
		err = v.Init(pk)
		assert.NoError(t, err)

		// Iterate over all combinations of signatures and public keys and verify each aggregated signature

		sig, err := v.AggregateSignatures([][]byte{signatures[0], signatures[1]}, []uint16{1, 2})
		assert.NoError(t, err)

		err = v.Verify(digest, sig)
		assert.NoError(t, err)

		sig, err = v.AggregateSignatures([][]byte{signatures[0], signatures[2]}, []uint16{1, 3})
		assert.NoError(t, err)

		err = v.Verify(digest, sig)
		assert.NoError(t, err)

		sig, err = v.AggregateSignatures([][]byte{signatures[1], signatures[2]}, []uint16{2, 3})
		assert.NoError(t, err)

		err = v.Verify(digest, sig)
		assert.NoError(t, err)
	}
}

func TestBenchmark(t *testing.T) {
	var commParties []*comm.Party
	var signers []*tlsgen.CertKeyPair
	var loggers []*commLogger
	var listeners []net.Listener
	var stopFuncs []func()

	nEnv := os.Getenv("N")
	n, err := strconv.Atoi(nEnv)
	if err != nil || n <= 0 {
		n = 3
	}

	threshold := int(math.Floor(float64(n/2)) + 1)

	t.Logf(">>>> n = %d, t = %d", n, threshold)

	members, certPool, loggers, signers, listeners, commParties, membershipFunc, parties, kgf := setup(t, n, loggers, signers, listeners, commParties)

	for _, l := range loggers {
		l.mute()
	}

	// Create Party
	for id := 1; id <= n; id++ {
		stop, s := createParty(id, kgf, signers[id-1], n, certPool, listeners, loggers, commParties, membershipFunc)
		parties = append(parties, s)
		stopFuncs = append(stopFuncs, stop)
	}

	defer func() {
		for _, stop := range stopFuncs {
			stop()
		}
	}()

	// DKG
	shares, start := keygen(t, parties, n)

	fmt.Println(">>>> shares", len(shares))

	elapsed := time.Since(start)
	t.Log(">>>> DKG took", elapsed)

	loop := 100

	totalSignerInitTime := time.Duration(0)
	totalSignTime := time.Duration(0)
	totalVerifierCreationTime := time.Duration(0)
	totalAggregateTime := time.Duration(0)

	for i := 0; i < loop; i++ {
		_, _, _, signerInitEndTime, signTime, verifierCreationEndTime, aggregateTime := signAndAggregate(t, threshold, members, shares)

		totalSignerInitTime += signerInitEndTime
		totalSignTime += signTime
		totalVerifierCreationTime += verifierCreationEndTime
		totalAggregateTime += aggregateTime
	}

	t.Log(">>>> Average signer init time took", totalSignerInitTime/time.Duration(loop))
	t.Log(">>>> Average sign took", totalSignTime/time.Duration(loop))
	t.Log(">>>> Average verifier creation time took", totalVerifierCreationTime/time.Duration(loop))
	t.Log(">>>> Average aggregate took", totalAggregateTime/time.Duration(loop))
}

// example usage: go test -bench BenchmarkParallelInvocation -run=^$ -cpu=1,2,4,8,16,32,64
func BenchmarkParallelInvocation(b *testing.B) {

	var commParties []*comm.Party
	var signers []*tlsgen.CertKeyPair
	var loggers []*commLogger
	var listeners []net.Listener
	var stopFuncs []func()

	n := 3

	members, certPool, loggers, signers, listeners, commParties, membershipFunc, parties, kgf := setup(b, n, loggers, signers, listeners, commParties)

	for _, l := range loggers {
		l.mute()
	}

	for id := 1; id <= n; id++ {
		stop, s := createParty(id, kgf, signers[id-1], n, certPool, listeners, loggers, commParties, membershipFunc)
		parties = append(parties, s)
		stopFuncs = append(stopFuncs, stop)
	}

	defer func() {
		for _, stop := range stopFuncs {
			stop()
		}
	}()

	shares, start := keygen(b, parties, n)

	elapsed := time.Since(start)
	b.Log("DKG elapsed", elapsed)

	// Create the threshold signers.
	thresholdSigners := make([]*bls.TBLS, n)
	for id := 1; id <= n; id++ {
		thresholdSigners[id-1] = &bls.TBLS{
			Logger: logger(id, b.Name()),
			Party:  uint16(id),
		}
	}

	// Initialize them with a nil send function
	for i, signer := range thresholdSigners {
		signer.Init(members, n-1, nil)
		err := signer.SetShareData(shares[i])
		if err != nil {
			return
		}
	}

	// Sign a message
	msg := []byte("Three can keep a secret, if two of them are dead.")
	digest := sha256Digest(msg)

	var signatures [][]byte
	//var signatureCount uint32

	for _, signer := range thresholdSigners {
		sig, err := signer.Sign(nil, digest)
		assert.NoError(b, err)
		signatures = append(signatures, sig)
	}

	pk, err := thresholdSigners[0].ThresholdPK()
	assert.NoError(b, err)

	var v bls.Verifier
	err = v.Init(pk)
	assert.NoError(b, err)

	sig1, err := v.AggregateSignatures([][]byte{signatures[0], signatures[1]}, []uint16{1, 2})
	assert.NoError(b, err)

	sig2, err := v.AggregateSignatures([][]byte{signatures[0], signatures[2]}, []uint16{1, 3})
	assert.NoError(b, err)

	sig3, err := v.AggregateSignatures([][]byte{signatures[1], signatures[2]}, []uint16{2, 3})
	assert.NoError(b, err)

	tSigs := [][]byte{sig1, sig2, sig3}

	parallelism := 1

	b.Run(fmt.Sprintf("sign-p%d", parallelism), func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			signer := thresholdSigners[parallelism%len(thresholdSigners)]
			var sig []byte
			var err error
			for pb.Next() {
				sig, err = signer.Sign(nil, digest)
			}
			// store results to prevent compiler optimizations
			gsig = sig
			gerr = err
		})
		b.StopTimer()
	})

	b.Run(fmt.Sprintf("verify-p%d", parallelism), func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			sig := tSigs[parallelism%len(tSigs)]
			var err error
			for pb.Next() {
				err = v.Verify(digest, sig)
			}
			// store results to prevent compiler optimizations
			gerr = err
		})
		b.StopTimer()
	})
}

var gsig []byte
var gerr error

func keygen(t TestingT, parties []MpcParty, n int) ([][]byte, time.Time) {
	var wg sync.WaitGroup
	wg.Add(len(parties))

	shares := make([][]byte, n)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	start := time.Now()
	for i, p := range parties {
		go func(i int, p MpcParty) {
			defer wg.Done()
			secretShareData, err := p.KeyGen(ctx, len(parties), int(math.Floor(float64(n/2)))+1)
			shares[i] = secretShareData
			assert.NoError(t, err)
			assert.NotNil(t, secretShareData)
		}(i, p)
	}

	wg.Wait()
	cancel()
	return shares, start
}

func setup(t TestingT, n int, loggers []*commLogger, signers []*tlsgen.CertKeyPair, listeners []net.Listener, commParties []*comm.Party) ([]uint16, *x509.CertPool, []*commLogger, []*tlsgen.CertKeyPair, []net.Listener, []*comm.Party, func() map[UniversalID]PartyID, []MpcParty, func(id uint16) KeyGenerator) {
	var members []uint16
	for i := 1; i <= n; i++ {
		members = append(members, uint16(i))
	}

	ca, err := tlsgen.NewCA()
	assert.NoError(t, err)

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(ca.CertBytes())

	tlsCert, err := ca.NewServerCertKeyPair("127.0.0.1")
	assert.NoError(t, err)

	membership := make(map[UniversalID]PartyID)

	for id := 1; id <= n; id++ {
		l := logger(id, t.Name())
		loggers = append(loggers, l)

		s := newSigner(ca, t)
		signers = append(signers, s)

		lsnr := comm.Listen("127.0.0.1:0", tlsCert.Cert, tlsCert.Key)
		listeners = append(listeners, lsnr)

		commParties = append(commParties, &comm.Party{
			Logger:   l,
			Address:  lsnr.Addr().String(),
			Identity: s.Cert,
		})

		membership[UniversalID(id)] = PartyID(id)
	}

	membershipFunc := func() map[UniversalID]PartyID {
		return membership
	}

	var parties []MpcParty

	kgf := func(id uint16) KeyGenerator {
		return &bls.TBLS{
			Logger: logger(int(id), fmt.Sprintf(t.Name())),
			Party:  id,
		}
	}
	return members, certPool, loggers, signers, listeners, commParties, membershipFunc, parties, kgf
}

func createParty(id int, kgf func(id uint16) KeyGenerator, signer *tlsgen.CertKeyPair, n int, certPool *x509.CertPool, listeners []net.Listener, loggers []*commLogger, commParties []*comm.Party, membershipFunc func() map[UniversalID]PartyID) (func(), MpcParty) {
	remoteParties := make(comm.SocketRemoteParties)

	auth := func(tlsContext []byte) comm.Handshake {
		h := comm.Handshake{
			TLSBinding: tlsContext,
			Identity:   signer.Cert,
			Timestamp:  time.Now().Unix(),
		}

		sig, err := signer.Sign(rand.Reader, sha256Digest(h.Bytes()), nil)
		if err != nil {
			panic("failed signing")
		}

		h.Signature = sig

		return h
	}

	for i := 0; i < n; i++ {
		if i+1 == id {
			continue
		}

		remoteParties[i+1] = comm.NewSocketRemoteParty(comm.PartyConnectionConfig{
			AuthFunc: auth,
			TlsCAs:   certPool,
			Id:       i + 1,
			Endpoint: listeners[i].Addr().String(),
		}, loggers[id-1])

	}

	commParties[id-1].SendMessage = remoteParties.Send

	p2id := make(map[string]uint16)
	for i, p := range commParties {
		p2id[hex.EncodeToString(sha256Digest(p.Identity))] = uint16(i + 1)
	}

	in, stop := comm.ServiceConnections(listeners[id-1], p2id, loggers[id-1])
	commParties[id-1].InMessages = in

	s := LoudScheme(uint16(id), loggers[id-1], kgf, nil, len(commParties)-1, remoteParties.Send, membershipFunc)

	go func(in <-chan comm.InMsg) {
		for msg := range in {
			inMsg := &IncMessage{
				MsgType: msg.Type,
				Data:    msg.Data,
				Topic:   msg.Topic,
				Source:  msg.From,
			}

			s.HandleMessage(inMsg)
		}
	}(in)
	return stop, s
}

func logger(id int, testName string) *commLogger {
	logConfig := zap.NewDevelopmentConfig()
	baseLogger, _ := logConfig.Build()
	logger := baseLogger.With(zap.String("t", testName)).With(zap.String("id", fmt.Sprintf("%d", id)))
	return &commLogger{Logger: &loggerWithDebug{SugaredLogger: logger.Sugar()}, conf: &logConfig}
}

type loggerWithDebug struct {
	*zap.SugaredLogger
}

func (lwd *loggerWithDebug) DebugEnabled() bool {
	return false
}

type commLogger struct {
	conf *zap.Config
	Logger
}

func (l *commLogger) mute() {
	l.conf.Level.SetLevel(zapcore.WarnLevel)
}

func (*commLogger) DebugEnabled() bool {
	return false
}

func sha256Digest(b ...[]byte) []byte {
	hash := sha256.New()
	for _, bytes := range b {
		hash.Write(bytes)
	}
	return hash.Sum(nil)
}

func newSigner(ca tlsgen.CA, t TestingT) *tlsgen.CertKeyPair {
	clientPair, err := ca.NewClientCertKeyPair()
	assert.NoError(t, err, "failed to create client key pair")

	return clientPair
}

// TestingT is an interface wrapper around *testing.T
type TestingT interface {
	Errorf(format string, args ...interface{})
	Name() string
}

func combinations(n uint16, t uint16) [][]uint16 {
	var ans [][]uint16
	if t == 1 {
		for i := uint16(1); i <= n; i++ {
			ans = append(ans, []uint16{i})
		}
		return ans
	}
	for i := n; i >= t; i-- {
		tem := combinations(i-1, t-1)
		for _, v := range tem {
			v = append(v, i)
			ans = append(ans, v)
		}
	}
	return ans
}

func signAndAggregate(t *testing.T, threshold int, members []uint16, shares [][]byte) ([]byte, bls.Verifier, []byte, time.Duration, time.Duration, time.Duration, time.Duration) {

	signerInitStartTime := time.Now()
	thresholdSigners := make([]*bls.TBLS, threshold)
	for id := 1; id <= threshold; id++ {
		thresholdSigners[id-1] = &bls.TBLS{
			Logger: logger(id, t.Name()),
			Party:  uint16(id),
		}
	}

	// Initialize them with a nil send function
	for i, signer := range thresholdSigners {
		signer.Init(members, threshold, nil)
		err := signer.SetShareData(shares[i])
		if err != nil {
			return nil, bls.Verifier{}, nil, 0, 0, 0, 0
		}
	}

	signerInitTime := time.Since(signerInitStartTime)

	// random message
	msg, err := generateRandomBytes(50)
	assert.NoError(t, err)
	digest := sha256Digest(msg)

	signatures := make([][]byte, threshold)
	signerIndices := make([]uint16, threshold)

	var wg sync.WaitGroup

	wg.Add(threshold)

	signStartTime := time.Now()
	for worker := 0; worker < threshold; worker++ {
		go func(worker int) {
			signer := thresholdSigners[worker]
			defer wg.Done()

			sig, err := signer.Sign(nil, digest)
			assert.NoError(t, err)

			signatures[worker] = sig
			signerIndices[worker] = uint16(worker + 1)
		}(worker)
	}

	wg.Wait()

	pk, err := thresholdSigners[0].ThresholdPK()
	assert.NoError(t, err)
	signTime := time.Since(signStartTime)

	verifierCreationStartTime := time.Now()
	var v bls.Verifier
	err = v.Init(pk)
	assert.NoError(t, err)
	verifierCreationTime := time.Since(verifierCreationStartTime)

	aggregateStartTime := time.Now()
	sig, err := v.AggregateSignatures(signatures, signerIndices)
	assert.NoError(t, err)
	aggregateTime := time.Since(aggregateStartTime)

	return sig, v, digest, signerInitTime, signTime, verifierCreationTime, aggregateTime
}

func generateRandomBytes(length int) ([]byte, error) {
	randomBytes := make([]byte, length)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	return randomBytes, nil
}
