package huifupay

func NewConfig(systemID, productID, rsaHuifuPublicKey, rsaMerchantPrivateKey string, verifySign bool) Config {
	return Config{
		SystemID:              systemID,
		ProductId:             productID,
		RsaHuifuPublicKey:     rsaHuifuPublicKey,
		RsaMerchantPrivateKey: rsaMerchantPrivateKey,
		VerifySign:            verifySign,
	}
}

type Config struct {
	SystemID  string `json:"system_id"`
	ProductId string `json:"product_id"`

	// 汇付公钥
	RsaHuifuPublicKey string `json:"rsa_huifu_public_key"`

	// 商户私钥
	RsaMerchantPrivateKey string `json:"rsa_merchant_private_key"`

	// 是否对请求返回的结果进行验签
	VerifySign bool `json:"verify_sign"`
}
