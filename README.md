You can fetch exchange rates for cryptocurrencies from several free APIs. Here are some popular options:

1. CoinGecko API
CoinGecko offers a free API that provides comprehensive data on cryptocurrencies. It's widely used and has generous rate limits.

API Endpoint for Simple Price: https://api.coingecko.com/api/v3/simple/price?ids={ids}&vs_currencies={vs_currencies}
For example, to get the exchange rate for Bitcoin (BTC) and Ethereum (ETH) in USD and EUR:
```shell
https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum&vs_currencies=usd,eur
```
2. CoinMarketCap API
CoinMarketCap is another popular choice. They offer a free tier with a limited number of API calls per month.

API Endpoint: https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest
Note: You need to sign up for an API key.

3. CryptoCompare API
CryptoCompare provides a free tier for fetching cryptocurrency data.

API Endpoint for Price: https://min-api.cryptocompare.com/data/price?fsym={from_symbol}&tsyms={to_symbols}
For example, to get the exchange rate for BTC in USD:
```bash
https://min-api.cryptocompare.com/data/price?fsym=BTC&tsyms=USD
```
