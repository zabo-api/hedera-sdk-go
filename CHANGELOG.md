# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Unreleased

### Changed

 - `hedera.GenerateSecretKey` now returns both the secret key and a 24-word mnemonic that can be used to recover the secret key.

### Added

 - `hedera.GenerateSecretKeyWithPassword` functions exactly as `hedera.GenerateSecretKey` except the
   returned mnemonic now will require the password when used to recover the secret key

### Removed

 - `hedera.Query.Cost()`

## 0.2.0 - 2018-12-07

### Fixed

 - Memory alignment in Windows 10 ([#11](https://github.com/hashgraph/hedera-sdk-go/issues/11))

### Changed

 - The client is now defined using a fluent builder pattern to prepare for
   further support for the Hedera API.

    ```go
    // v0.1.0: client.GetAccountBalance(id)
    client.Account(id).Balance()
    // upcoming: client.Account(id).Info()
    // upcoming: client.Account(id).Records()

    // v0.1.0: client.GetTransactionReceipt(id)
    client.Transaction(id).Receipt()
    // upcoming: client.Transaction(id).Record()
    ```

 - `client.CryptoTransfer` is now `client.TransferCrypto`.

 - For queries, `.Answer()` is now `.Get()`.

 - For transactions, `.Execute()` now returns the ID directly instead of being wrapped
   in a 1-element struct.

## 0.1.0 - 2018-11-29

Initial stable release.
