package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		"enode://43702a98bb7a9e81d78754dd6507eaef3a3107dde80d367cd4d0a2993d274732ab00eeffc917b11cd2f55149891436bb947a2d0bca93128ba0313e58dc4cd126@52.72.214.247:5050",
		"enode://0a2deb8cebf0a79608380aecabebc7ab6493f28f19da51e64b873b6a9c781dee95f241dde5073825504feccaf87f87f92d5d438d814d371fc57a8899e0967944@3.214.69.247:5050",
		"enode://c18703de659e9d3b429e7ff79f414628c1dc805efdfb0bba03069d903583c382ed5cf87b2aa518313c012a9498ebdb1ac3aa25e8ec6f1d6fba96a701742aad39@52.71.129.233:5050",
		"enode://020ecf753f125e8569d29749a42e57e7b540b09dd9e2c93a07573f9472f74d85141df7acabdbf17fe69b71685595d6a199d06f50d78bd252f9f2edf2cd3bf6de@18.233.179.63:5050",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
