package config

// GetChannelID returns the channel ID
func (c *Config) GetChannelID() int64 {
	return c.ChannelID
}

// GetTelegramGroupURL returns the Telegram group URL
func (c *Config) GetTelegramGroupURL() string {
	return c.TelegramGroupURL
}

// GetMiniAppURL returns the URL for the specified game type
func (c *Config) GetMiniAppURL(gameType string) string {
	url, ok := c.MiniAppURLs[gameType]
	if !ok {
		// Default to bank game if game type is not found
		return c.MiniAppURLs["bank"]
	}
	return url
}

// GetDefaultAppText returns the default app text
func (c *Config) GetDefaultAppText() string {
	return c.DefaultAppText
}

// GetDefaultAppURL returns the default app URL
func (c *Config) GetDefaultAppURL() string {
	return c.DefaultAppURL
}

// GetVoucherButtonText returns the voucher button text
func (c *Config) GetVoucherButtonText() string {
	return c.VoucherButtonText
}

// GetVoucherTypeParam returns the voucher type parameter
func (c *Config) GetVoucherTypeParam() string {
	return c.VoucherTypeParam
}

// GetKYCButtonText returns the KYC button text
func (c *Config) GetKYCButtonText() string {
	return c.KYCButtonText
}

// GetKYCTypeParam returns the KYC type parameter
func (c *Config) GetKYCTypeParam() string {
	return c.KYCTypeParam
}
