package constant

const (
	//It is not recommended to hard-code keys like this. Better to store it in a keystore and retrieve it.
	WalletDemoKey = "435042b0fac31048adf0fa1cb71fbf0d5989d30c6bea2c91a22dd297187cfd2a"
)

// This simulates getting a key from a keystore
func GetWalletKey() string {
	return WalletDemoKey
}
