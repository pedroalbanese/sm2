// Parameters for the SM2 Elliptic curve
package sm2p256v1

import (
	"crypto/elliptic"
	"math/big"
	"sync"
)

var initonce sync.Once
var sm2 *elliptic.CurveParams

func initP256() {
	sm2 = new(elliptic.CurveParams)
	sm2.P, _ = new(big.Int).SetString("fffffffeffffffffffffffffffffffffffffffff00000000ffffffffffffffff", 16)
	sm2.N, _ = new(big.Int).SetString("fffffffeffffffffffffffffffffffff7203df6b21c6052b53bbf40939d54123", 16)
	sm2.B, _ = new(big.Int).SetString("28e9fa9e9d9f5e344d5a9e4bcf6509a7f39789f515ab8f92ddbcbd414d940e93", 16)
	sm2.Gx, _ = new(big.Int).SetString("32c4ae2c1f1981195f9904466a39c9948fe30bbff2660be1715a4589334c74c7", 16)
	sm2.Gy, _ = new(big.Int).SetString("bc3736a2f4f6779c59bdcee36b692153d0a9877cc62a474002df32e52139f0a0", 16)
	sm2.BitSize = 256
}

func P256() elliptic.Curve {
	initonce.Do(initP256)
	return sm2
}
