https://github.com/QuestofIranon/hedera-sdk-go.git# Hedera SDK for Go

[![GoDoc](https://godoc.org/github.com/launchbadge/hedera-sdk-go?status.svg)](https://godoc.org/github.com/launchbadge/hedera-sdk-go)
[![](https://goreportcard.com/badge/github.com/launchbadge/hedera-sdk-go)](https://goreportcard.com/report/github.com/launchbadge/hedera-sdk-go)

This repo contains the Go SDK for interacting with the [Hedera](https://hedera.com) platform. Hedera is the _only_ **public** distributed ledger licensed to use the Hashgraph consensus algorithm for fast, fair and secure transactions. By using any of the Hedera SDKs, developers will be empowered to build an entirely new class of decentralized applications.

Following the instructions below should help you to reach the position where you can send transactions and queries to a Hedera testnet or the Hedera mainnet.

## Table of Contents

* **[Developer Rewards][01-dev-rewards]**
* **[Architectural Overview][02-arch-overview]**
* **[Prerequisites][03-prerequisites]**
  * **[Software][03-01-software]**
  * **[Hedera Account][03-02-hedera-acct]**
  * **[Hedera testnet access][03-03-hedera-testnet]**
* **[Installing the Hedera SDK for Go][04-installing]**
* **[Creating a public/private keypair for testnet use][05-create-keypair]**
* **[Associating your public key with you Hedera testnet account][06-assoc-key]**
* **[Your First Hedera Application][07-first-hedera]**
  * **[Checking your account balance][07-01-check-balance]**
  * **[Enhance your application to check a friend's account balance][07-02-friend-balance]**
  * **[Next step: Transferring hbars to a friend's account][07-03-transfer-hbars]**
* **[Other resources][08-other-res]**
* **[Getting in touch][09-get-in-touch]**
* **[Contributing to this project][10-contribute]**
* **[License information][11-license]**

--------

## Developer Rewards

Developers who create Hedera accounts and test their applications on our testnet will be offered the opportunity to earn real mainnet __ℏ__ ([__hbars__](#a-hbar)) based upon that testing activity. The SDKs are intended to facilitate application development and encourage such testing. See the [Hedera Account][03-02-hedera-acct] section for further details.

Developers who create Hedera accounts and test their applications on our testnet are offered the opportunity to earn real mainnet __ℏ__ ([__hbars__](#a-hbar)) based upon that testing activity. The SDKs are intended to facilitate application development and encourage such testing. See the [Hedera Account][03-02-hedera-acct] section for further details.

> <a id="a-hbar">**Hbars**:</a>
>
> _Hbars_ are the native cryptocurrency that is used to pay for transactions on the Hedera platform and to secure the network from certain types of cyberattacks. They are the native platform coin needed to interact with and exchange value on Hedera.
>
> The symbol for hbars is "**ℏ**" so `5 ℏ` means 5 hbars
>
> <a id="a-tinybar">**Tinybars**:</a>
>
> Tinybars are (not surprisingly) smaller than hbars. They are used to divide hbars into smaller amounts. One hbar is equivalent to one hundred million tinybars.
>
> The symbol for tinybars is "**tℏ**" so it is correct to say `1 ℏ = 100,000,000 tℏ`
>
> _**Important Note**: The values of all fees and transfers throughout the Hedera SDKs are represented in tinybars, though the term hbars may be used for the purposes of brevity._

## Architectural Overview

All Hedera SDKs are intended to provide a developer-friendly means to leverage the Hedera API, which is based on [Google Protobuf](https://developers.google.com/protocol-buffers/). Protobuf supports code generation for a growing number of languages and is highly optimised for efficient communications. For those interesting in viewing the underlying protobuf message definitions, see the [Hedera Protobuf Message Definitions](https://github.com/hashgraph/hedera-protobuf) repo.

Developers who wish to work in other languages are at liberty to do so, but should be aware that implementation of cryptographic key generation and manipulation is not a trivial undertaking. The source code for the cryptography libraries used by this project can be found in the [Hedera SDK for Rust](https://github.com/hashgraph/hedera-sdk-rust) repository although they are compiled to C for inclusion in this project. We would recommend use of these same libraries for developers interested in adding support for other languages.

## Prerequisites

### Software

* [Go](https://golang.org/) – version 1.11 or higher ([download here](https://golang.org/dl/)).
  * Installation instructions for Go can be found [here](https://golang.org/doc/install).
  * An introduction to the Go programming language can be found [here](https://golang.org/).
* [Git LFS](https://git-lfs.github.com) is used to assist dependency and version management, speeding up repo cloning. In the case of this project, libraries built from the [Hedera SDK for Rust](https://github.com/hashgraph/hedera-sdk-rust) project are included in [`./libs`](./libs). It is important that Git LFS be installed _before_ the running `go get...` command below.

  * Installing [Git LFS](https://git-lfs.github.com) on **Mac OS X** (using homebrew). Run the following commands from a terminal window:

    ```sh
    brew install git-lfs
    git lfs install
    ```

  * Installing [Git LFS](https://git-lfs.github.com) on **Ubuntu** or **Debian**.  Run the following commands from a terminal window:

    ```sh
    sudo apt install git-lfs
    git lfs install
    ```

  * Installing [Git LFS](https://git-lfs.github.com) on **Windows**

    Download and run [git-lfs.exe](https://github.com/git-lfs/git-lfs/releases/download/v2.6.0/git-lfs-windows-v2.6.0.exe). Open your git console – under `Git Bash` if you installed Git through [git-scm.com](https://git-scm.com/).

    ```sh
    git lfs install
    ```

### Hedera Account

The [Hedera Portal](https://go.hedera.com) allows people to create a **Hedera Account** facilitating access to the Hedera mainnet and Hedera testnets. A Hedera Account allows entry of a testnet access code, in order to add a number of testnet __ℏ__ ([hbars](#a-hbar)) to a testnet account created (as can be seen [below][03-03-hedera-testnet]) using Hedera SDKs. The public key associated with a testnet account must also be associated with your Hedera account. A detailed explanation of this whole process is contained within this document.

* In order to gain early access (before Open Access) to a Hedera testnet or the Hedera mainnet users must create a Hedera account, including full identity verification. You can do this using the [Hedera Portal](https://go.hedera.com).
* We want to allow devs to earn __ℏ__ ([hbars](#a-hbar)) by helping us to test our SDKs. Hedera is based in the USA, so for us to be allowed to do this under US law we need to verify your identity as a part of the account creation process.

A full explanation of the Portal, Hedera accounts, identity verification and many other topics can be found at [Hedera Help](https://help.hedera.com). New users should head for the [Getting Started](https://help.hedera.com/hc/en-us/categories/360000099938-Getting-Started) section.

### Hedera testnet access

A Hedera testnet provides a test environment for testing your code without having to spend "real" mainnet __ℏ__ ([hbars](#a-hbar)). Testnet hbars are akin to "monopoly money" and have no intrinsic value, but testing against testnets will help you **earn** real **mainnet ℏ ([hbars](#a-hbar))**. It is worth noting that the virtual infrastructure used to provide testnets is not intended for performance testing, as the specification of nodes is not in any way equivalent to that of mainnet nodes. Further information on this topic is included within the "Testnet Performance and Throttling" section further on in these instructions.

* Once you have your Hedera account set up, you can request access to Hedera test networks by filling out the form [here](https://learn.hedera.com/HederaTestnetAccess).
* Check for answers to your testnet-related questions in the *Testnet Activation* section on [this](https://help.hedera.com/hc/en-us/categories/360000099938-Getting-Started) page.
* Please note that there is a waiting list for testnet access; please be patient. If you haven't had a response withing 10 days, reach out to to the team on [discord](https://hedera.com/discord).
* Once you have received a testnet access code, enter it into the testnet access box and push the proceed/arrow button. You should see a message "_Access code confirmed. By activating this code you will switch to testnetXXX with a starting value of 1000 hbars._" Please note that these are **testnet** hbars and should not to be confused with mainnet hbars.
* Pushing the "Activate and switch network" button will credit your testnet account with those testnet hbars and switch the portal to the testnet. You can switch the portal between mainnet and testnet at any time by using the drop-down at the top of the page.

## Installing the Hedera SDK for Go

To clone the repo and all required dependencies, run the following command from a terminal window. Note that this command will create a files and folders from within your current folder.

Remember that the [software prerequisites][03-01-software] described above must be installed before proceeding.

```sh
go get github.com/launchbadge/hedera-sdk-go
```

## Creating a public/private keypair for testnet use

As a general principle, it is bad practice to use your mainnet keys on a testnet. The code below shows the content of the [generate_key example](/examples/generate_key/main.go) file. This shows how you can create new public and private keys using the Hedera SDK for Go:

```go
package main

import (
  "fmt"
  "github.com/launchbadge/hedera-sdk-go"
)

func main() {
  secret, mnemonic := hedera.GenerateSecretKey()
  fmt.Printf("secret   = %v\n", secret)
  fmt.Printf("mnemonic = %v\n", mnemonic)

  public := secret.Public()
  fmt.Printf("public   = %v\n", public)
}
```

* It can be executed from a terminal window by changing to the `examples/generate_key` folder and typing:

```sh
go run main.go
```

* Make careful note of the 24-word mnemonic and both of the keys that are generateed. For a testnet you can copy and paste this information into a text file. For security reasons you should **_never do this for mainnet_**.

> **Development Environment Note**
>
> All of the commands described in these instructions can also be executed using an IDE – such as [VSCode](https://code.visualstudio.com/) or [Goland](https://www.jetbrains.com/go/), or editors such as [Atom](https://atom.io/) – according to each developer's preferences and budget. _Hedera has no affiliation with the companies providing these or other equivalent tools._
>
> Throughout these instructions you'll find the phrase "Run the following command from a terminal window." Feel free to use your IDE whenever you see this – if that's how you prefer to work. Terminal is used in this document to avoid ambiguity.

## Associating your public key with your Hedera testnet account

Once you have generated a public/private keypair as described [above][05-create-keypair], you need to link the **public** key to your Hedera testnet account. To do this, return to the Hedera portal and ensure that you have the testnet selected in the drop-down at the top of the page.

You should see the box asking you to `Enter your Public Key`. Copy the long hex value of the **public** key and paste it into that textbox in the portal. Do make sure you that you select all of the characters. Click `Submit`.

You should briefly see an "Account Pending" message, which will be replaced with an Account ID box. You should make a note of your account ID - perhaps in the same text file where you have stored your public and private testnet keys.

All Hedera IDs consist of three numbers (`int64s`) separated by the "`.`" symbol. The three numbers represent `Shard number`, `Realm number` and `Account number` respectively. Shards and Realms are not yet in use so expect the first two numbers to be zeros.

You should also scroll down until you see a box labelled "Network" and make a note of the `Address`, which will be something like `testnet.hedera.com:50222`. You should also take note of the `Node`, which represents the account ID of a node on the testnet; it will be something like `0.0.3`.

## Your First Hedera Application

### Checking your account balance

A more complete example, which includes code fragments in this section can be found in the [get_account_balance example](/examples/get_account_balance/main.go) file located in the [examples](/examples/) folder of this repo. This simplified example is broken into bite-sized pieces here so that accompanying explanations can be seen in context alongside each fragment.

Create a new `main.go` file importing `fmt` (for `Printf`) and the Hedera SDK. The `time` package is also imported but *commented* here. This package will be required later in this example and can be un-commented when required by removing the preceding `//`.

It's also useful to create a variable `oneHbar` to represent the number of **_tinybars_** in one **_hbar_**:

```go
package main

import (
  "fmt"
  //"time"
  "github.com/launchbadge/hedera-sdk-go"
)

func main() {

  oneHbar := 100000000
```

Create and set a `myAccount` variable, replacing `1234` with your own `Account ID` from the portal. Note that the `Shard ID` and `Realm ID` are not required at this time so they are defaulted to zero unless explicitly included. If the `accountID` shown in the portal is `0.0.1234` only the last number `1234` need be used. This is the account for which we will retrieve the balance.

```go
  myAccount := hedera.AccountID{Account: 1234}
```

Next, you will want to establish a connection to the Hedera testnet using the `Address` you noted earlier from the network section shown on the Hedera portal. Make sure that the `testnet address` contained within the quotation marks matches the one you copied from the portal.

You should also handle any errors and defer disconnection from the tesnet to keep things clean and orderly.

```go
  client, err := hedera.Dial("testnet.hedera.com:50222")
  if err == nil {
    defer client.Close()
  } else {
    panic(err)
  }
```

Once a connection has been established, you need to decide which node to send your query to. Testnets provide each developer with a single node's account ID (labelled `node` in the portal) in order to simplify this and limit testnet infrastructure burden. On mainnet it will be the responsibility of the application to choose a node - usually at random.

> **Node Account Defaults**: The `client.SetNode()` function sets the default node account for all future transactions until changed and unless overridden. In these examples this is not strictly necessary as all of the transaction examples in this document include `Node(nodeAccount)`. When using testnets, it is also worth noting that the SDK uses node account `0.0.3` by default – even when neither of these mechanisms is used.

The node account ID should look something like `0.0.3`. Since `Shard ID` and `Realm ID` are not yet in use, they are defaulted to zero. It is the last number (in this example `3`) that you should use in the following code:

```go
  nodeAccount := hedera.AccountID{Account: 3}
  client.SetNode(nodeAccount)
```

It is also important to specify which account is initiating this query – known as the **operator** account – so that the account can be charged a small fee for the execution of this query. In order to authorise the payment of such a fee, the operator must sign the request using their private key.

> **Security Tip**: In the [get_account_balance example](/examples/get_account_balance/main.go) file located in the [examples](/examples/) folder, the `os.Getenv("OPERATOR_SECRET")` function is used. This retrieves the private key from an environment variable called OPERATOR_SECRET. It is good practice to use this technique, as it avoids accidental publication of private keys when storing code in public repos or sharing them accidentally via collaboration tools. Don't forget: If someone else knows your private key, they effectively own your account! Although the impact of this is low when using testnets, it could be a very expensive mistake on mainnet. For purposes of clarity, the example below has been simplified and does not use an environment variable

Replace `<my-private-key>` with the private key you generated near the beginning of these instructions.

```go
  operatorSecret, err := hedera.SecretKeyFromString("<my-private-key>")
  if err != nil {
    panic(err)
  }
  client.SetOperator(myAccount, func() hedera.SecretKey {
    return operatorSecret
  })
```

Again, note that the `client.SetOperator()` function is used to set the default operator account for subsequent transactions. This is not necessary if both `Operator()` and `Sign()` are used when creating the transaction, but is included here in order to explain both approaches.

At this point, you're ready to query your account balance. The `client.Account(myAccount).Balance()` constructs the request; adding `.Get()` executes that request. Once again, don't forget to handle possible failures.

You can the output the balance using `fmt.Printf` and end the program by closing the braces for `func main`.

For illustrative purposes, we're showing the balance in **_[tinybars](#a-tinybar)_** and **_[hbars](#a-hbar)_**. The Hedera SDKs represent all quantities for transfers and fees as integers using _tinybars_. There are one hundred million (100,000,000) tinybars in one hbar.

```go
  myBalance, err := client.Account(myAccount).Balance().Get()
  if err != nil {
    panic(err)
  }

  fmt.Printf("Account %v balance = %v tinybars\n", myAccount.Account, myBalance)
  fmt.Printf("Account %v balance = %.5f hbars\n", myAccount.Account, float64(myBalance)/float64(oneHbar))
}
```

You should now be able to **run your first Hedera program** by executing `go run main.go` from terminal.

If everything went according to plan you should see something like this:

```sh
Account 1234 balance = 100500005000 tinybars
Account 1234 balance = 1005.00005 hbars
```

#### Testnet performance and throttling

> For the present, our testnets have been throttled, allowing a limited number of Hedera transactions per account per second. We're using virtual infrastructure to support the huge demand we have had for testnet access, and prefer to foster innovative use of these resources and discourage folks from trying to generate metrics using underspecified hardware.
>
> If you see error messages like `transaction failed the pre-check: Busy`, it is likely that you are exceeding these throttling thresholds. To avoid such errors, short delays can be added. To add a one second delay, for example, use the following code between transactions or queries:
>
>```go
>  time.Sleep(1 * time.Second)
>```
>
> Don't forget to remove the `//` comment symbols before `"time"` (in the `import` section at the beginning of the program) to gain access to the functions in the go time package. Some IDEs will automatically add missing packages when detected, so don't be surprised if you find that a new `"time"` line has appeared alongside `//"time"`.

### Enhance your application to check a friend's account balance

If you know the account ID of another account on your testnet – perhaps a friend or colleague – you can also check their balance.

> **Creating an additional testnet account**
>
> If your friends won’t share their accounts, or if you don’t have any friends, see the [Create Account Example](/examples/create_account/main.go) included in this repo.
>
> If you do choose to create an account using that example, don't forget to do the following:
>
> 1. Create a local environment variable OPERATOR_SECRET that contains your private key.
> 2. Adjust the `nodeAccountID` variable to the ID you see in the portal.
> 3. Update the `testnet.hedera.com:...` testnet address to the correct one.
> 4. Change the `operatorAccountID` variable value to your own testnet account ID.
> 5. Change the `InitialBalance` to an acceptable quantity of testnet tinybars.

* For the purposes of this example, an Account ID of `0:0:1235` will be used for that second account. Don't forget to amend `1235` to the account number of your friend's account. If you forget to do this will you will probably see a `transaction failed the pre-check: InvalidAccount` message.

* To continue with this example, add the following code into your existing `func main` function, just before the closing braces:

* As mentioned above, we will add a small delay to ensure that we do not exceed testnet throttling limits. For brevity, this statement will be included _without further comment_ in all subsequent examples. If you get an error here, you need to remove the `//` before `"time"` in the `import` block near the beginning of the program.

* Before executing any transfers, you can initialise a second variable `friendAccount` representing the second account, query its balance and output the result.

```go
  time.Sleep(1 * time.Second)

  friendAccount := hedera.AccountID{Account: 1235}

  friendBalance, err := client.Account(friendAccount).Balance().Get()
  if err != nil {
    panic(err)
  }

  fmt.Printf("Account %v balance = %v tinybars\n", friendAccount.Account, friendBalance)
  fmt.Printf("Account %v balance = %.5f hbars\n", friendAccount.Account, float64(friendBalance)/float64(oneHbar))
```

* Run the program again by executing `go run main.go` from terminal.

* You should see your balance followed by your friend's balance.

> **Transaction and Query Fees**
>
> Note that your balance will decrease slightly each time you execute your code. This is due to the small fees associated with each query or transaction on the Hedera platform. On a testnet, this is not all that important, but it's worth keeping on mind when using mainnet.

### Next step: Transferring _[hbars](#a-hbar)_ to a friend's account

* A `transferAmount` variable can be used to make the next steps more readable. In this case, we'll transfer **10 ℏ** and output details of the intended transaction.

```go
  transferAmount := int64(10 * oneHbar)
  fmt.Printf("Starting transfer of %v tinybars from Account %v to Account %v\n", transferAmount, myAccount.Account, friendAccount.Account)
```

* It is worth re-stating that a __secret__ key (also known as _private_ key) is required in order to transfer _hbars_ from an account. Separate signatures are required for the **operator** account and each `Transfer` line with a negative amount – even if they're the same account.

* The next statement is a little more complex so each line is explained individually below the code.

```go
  response, err := client.CryptoTransfer().
    Transfer(myAccount, -transferAmount).
    Transfer(friendAccount, transferAmount).
    Operator(myAccount).
    Node(nodeAccount).
    Memo("My first transfer of hbars! w00t!").
    Sign(operatorSecret).
    Sign(operatorSecret).
    Execute()
    if err != nil {
      panic(err)
    }
```

#### Explanation of the above code block by line number

__1__. `response, err := client.CryptoTransfer().` creates a transaction to transfer **_hbars_** between accounts.

__2__. `Transfer(myAccount, -transferAmount).` sets up part of the transfer. In this case _from_ **your** account. Note that the "`-`" (minus sign) denotes that **_hbars_** will be **deducted** from this account.

__3__. `Transfer(friendAccount, transferAmount).` sets up the second part of this transfer. In this case _to_ your **friend's**  account. A positive number indicates that this account will be **incremented** by the specified amount.

__4__. `Operator(myAccount).` specifies the account initiating the transaction and to which transaction fees will be charged.

__5__. `Node(nodeAccount).` indicates the node to which this transaction and associated fee payments will be sent.

__6__. `Memo("My first transfer of hbars! w00t!").` assigns a label to the transaction of up to 100 bytes. Use of this field is at the developer's discretion and does not affect the behaviour of the plaform.

__7__. `Sign(operatorSecret).` adds a signature for the **operator** account. This is required as fees will be deducted from this account.

__8__. `Sign(operatorSecret).` adds a signature for the account from which **_hbars_** will de debited. In this case the **operator** account is the same as the account **sending** **_hbars_** so the same signature can be re-used.

__9__. `Execute()` executes the transaction.

 If the `client.SetOperator()` function is used to set the default operator as illustrated above, lines __4__ and __7__ can be omitted. If the `client.SetNode()` function is used to set the default node as illustrated above, line __5__ can also be omitted.

> #### Multi-party transfers
>
> It is possible to create a transfer transaction containing **multiple** _to_ and **multiple** _from_accounts within that same transaction. In a case where multiple accounts were to be debited, signatures would be required for each one, and addition `Sign(...).` lines would be required.
>
> __Important__: the _sum of all amounts_ in `Transfer(...)` lines contained within in a `CryptoTransfer` _**must** add up to **zero**_.

* The ID of the transaction itself can now be captured from the `response` object in the above statement. The `transactionID` is made up of the account ID and the transaction timestamp – right down to nanoseconds.

* It makes sense to wait a little longer (2 seconds) after sending the transaction, so that the Hedera network can reach consensus on the transaction.

```go
  transactionID := response.ID
  fmt.Printf("Transfer Sent. Transaction ID is %v\n", transactionID)

  time.Sleep(2 * time.Second)
```

* To confirm that the transaction succeeded a `receipt` can be requested. Although this is not a mandatory step, it does verify that this transaction successfully reached network consensus.

```go
  receipt, err := client.Transaction(*transactionID).Receipt().Get()
  if err != nil {
    panic(err)
  }

  if receipt.Status == hedera.StatusSuccess {
    fmt.Printf("Transaction Successful. Consensus confirmed.\n")
  } else {
    panic(fmt.Errorf("Transaction unsuccessful. Status: %v", receipt.Status.String()))
  }

  time.Sleep(1 * time.Second)
```

* Finally, the balances of both accounts can be requeried to verify that the **10 ℏ** was indeed transferred from your account to that of your friend.

```go
  myBalance, err = client.Account(myAccount).Balance().Get()
  if err != nil {
    panic(err)
  }

  friendBalance, err = client.Account(friendAccount).Balance().Get()
  if err != nil {
    panic(err)
  }

  fmt.Printf("Account %v balance = %v tinybars\n", myAccount.Account, myBalance)
  fmt.Printf("Account %v balance = %.5f hbars\n", myAccount.Account, float64(myBalance)/float64(oneHbar))
  fmt.Printf("Account %v balance = %v tinybars\n", friendAccount.Account, friendBalance)
  fmt.Printf("Account %v balance = %.5f hbars\n", friendAccount.Account, float64(friendBalance)/float64(oneHbar))
```

* Run the program again by executing `go run main.go` from terminal.

* You should now see both balances prior to the transfer followed by details of the transfer including success/failure. You should then be able to see the balances of both accounts after the transfer, demonstrating that **10 ℏ** has been transferred from your account to you friend's account. Hopefully it looks something like this:

```txt
Account 1234 balance = 96495305000 tinybars
Account 1234 balance = 964.95305 hbars
Account 1235 balance = 4000000000 tinybars
Account 1235 balance = 40.00000 hbars
Transfering 1000000000 tinybars from Account 1234 to Account 1235
Transfer Sent. Transaction ID is 0:0:1234@1548679850.429332000
Transaction Successful. Consensus confirmed.
Account 1234 balance = 95494805000 tinybars
Account 1234 balance = 954.94805 hbars
Account 1235 balance = 5000000000 tinybars
Account 1235 balance = 50.00000 hbars
```

## Other resources

* For an explanation of the underlying hashgraph algorithm, please consult our [whitepaper](https://www.hedera.com/hh-whitepaper-v1.4-181017.pdf) or Dr. Leemon Baird's 52-minute [Simple Explanation](https://www.youtube.com/watch?v=wgwYU1Zr9Tg) video.
* Links to all Hedera news and information can be found in at [Hedera Help](https://help.hedera.com) – including Coq validation of the hashgraph ABFT algorithm.
* 300+ [Hedera interviews and videos](https://www.youtube.com/watch?v=v2M0eo9PRxw&list=PLuVX2ncHNKCwe1BdF6GH6RnjrF7J7yTZ4) on YouTube. Thanks to Arvydas – a Hedera MVP – for curating this list.

## Deeper Development

### Requirements

* [rustup](https://rustup.rs/)

* `musl`

  * Mac OS X

    ```sh
    brew install FiloSottile/musl-cross/musl-cross
    ```

* `MinGW`

  * Mac OS X

    ```sh
    brew install mingw-w64
    ```

#### Build

```sh
./x.py
```

## Getting in touch

Please reach out to us on the Hedera [discord channels](https://hedera.com/discord). We're fortunate to have an active community of over 5000 like-minded devs, who are passionate about our tech. The Hedera Developer Advocacy team also participates actively.

## Contributing to this Project

We welcome participation from all developers! For instructions on how to contribute to this repo, please review the [Contributing Guide](CONTRIBUTING.md).

## License Information

Licensed under Apache License,
Version 2.0 – see [LICENSE](LICENSE) in this repo or [apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0)

[//]: # (Internal reference links)
[01-dev-rewards]: #developer-rewards
[02-arch-overview]: #architectural-overview
[03-prerequisites]: #prerequisites
[03-01-software]: #software
[03-02-hedera-acct]: #hedera-account
[03-03-hedera-testnet]: #hedera-testnet-access
[04-installing]: #installing-the-hedera-sdk-for-go
[05-create-keypair]: #creating-a-publicprivate-keypair-for-testnet-use
[06-assoc-key]: #associating-your-public-key-with-you-hedera-tesnet-account
[07-first-hedera]: #your-first-hedera-application
[07-01-check-balance]: #checking-your-account-balance
[07-02-friend-balance]: #enhance-your-application-to-check-a-friends-account-balance
[07-03-transfer-hbars]: #next-step-transferring-hbars-to-a-friends-account
[08-other-res]: #other-resources
[09-get-in-touch]: #getting-in-touch
[10-contribute]: #contributing-to-this-project
[11-license]: #license-information
