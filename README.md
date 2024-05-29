# ✨ So you want to run an audit

This `README.md` contains a set of checklists for our audit collaboration.

Your audit will use two repos: 
- **an _audit_ repo** (this one), which is used for scoping your audit and for providing information to wardens
- **a _findings_ repo**, where issues are submitted (shared with you after the audit) 

Ultimately, when we launch the audit, this repo will be made public and will contain the smart contracts to be reviewed and all the information needed for audit participants. The findings repo will be made public after the audit report is published and your team has mitigated the identified issues.

Some of the checklists in this doc are for **C4 (🐺)** and some of them are for **you as the audit sponsor (⭐️)**.

---

# Audit setup

## 🐺 C4: Set up repos
- [ ] Create a new private repo named `YYYY-MM-sponsorname` using this repo as a template.
- [ ] Rename this repo to reflect audit date (if applicable)
- [ ] Rename audit H1 below
- [ ] Update pot sizes
  - [ ] Remove the "Bot race findings opt out" section if there's no bot race.
- [ ] Fill in start and end times in audit bullets below
- [ ] Add link to submission form in audit details below
- [ ] Add the information from the scoping form to the "Scoping Details" section at the bottom of this readme.
- [ ] Add matching info to the Code4rena site
- [ ] Add sponsor to this private repo with 'maintain' level access.
- [ ] Send the sponsor contact the url for this repo to follow the instructions below and add contracts here. 
- [ ] Delete this checklist.

# Repo setup

## ⭐️ Sponsor: Add code to this repo

- [ ] Create a PR to this repo with the below changes:
- [ ] Confirm that this repo is a self-contained repository with working commands that will build (at least) all in-scope contracts, and commands that will run tests producing gas reports for the relevant contracts.
- [ ] Please have final versions of contracts and documentation added/updated in this repo **no less than 48 business hours prior to audit start time.**
- [ ] Be prepared for a 🚨code freeze🚨 for the duration of the audit — important because it establishes a level playing field. We want to ensure everyone's looking at the same code, no matter when they look during the audit. (Note: this includes your own repo, since a PR can leak alpha to our wardens!)

## ⭐️ Sponsor: Repo checklist

- [ ] Modify the [Overview](#overview) section of this `README.md` file. Describe how your code is supposed to work with links to any relevent documentation and any other criteria/details that the auditors should keep in mind when reviewing. (Here are two well-constructed examples: [Ajna Protocol](https://github.com/code-423n4/2023-05-ajna) and [Maia DAO Ecosystem](https://github.com/code-423n4/2023-05-maia))
- [ ] Review the Gas award pool amount, if applicable. This can be adjusted up or down, based on your preference - just flag it for Code4rena staff so we can update the pool totals across all comms channels.
- [ ] Optional: pre-record a high-level overview of your protocol (not just specific smart contract functions). This saves wardens a lot of time wading through documentation.
- [ ] [This checklist in Notion](https://code4rena.notion.site/Key-info-for-Code4rena-sponsors-f60764c4c4574bbf8e7a6dbd72cc49b4#0cafa01e6201462e9f78677a39e09746) provides some best practices for Code4rena audit repos.

## ⭐️ Sponsor: Final touches
- [ ] Review and confirm the pull request created by the Scout (technical reviewer) who was assigned to your contest. *Note: any files not listed as "in scope" will be considered out of scope for the purposes of judging, even if the file will be part of the deployed contracts.*
- [ ] Check that images and other files used in this README have been uploaded to the repo as a file and then linked in the README using absolute path (e.g. `https://github.com/code-423n4/yourrepo-url/filepath.png`)
- [ ] Ensure that *all* links and image/file paths in this README use absolute paths, not relative paths
- [ ] Check that all README information is in markdown format (HTML does not render on Code4rena.com)
- [ ] Delete this checklist and all text above the line below when you're ready.

---

# Canto audit details
- Total Prize Pool: $100000 in USDC
  - HM awards: $79700 in USDC
  - (remove this line if there is no Analysis pool) Analysis awards: XXX XXX USDC (Notion: Analysis pool)
  - QA awards: $3300 in USDC
  - (remove this line if there is no Bot race) Bot Race awards: XXX XXX USDC (Notion: Bot Race pool)
 
  - Judge awards: $10000 in USDC
  - Lookout awards: XXX XXX USDC (Notion: Sum of Pre-sort fee + Pre-sort early bonus)
  - Scout awards: $500 in USDC
  - (this line can be removed if there is no mitigation) Mitigation Review: XXX XXX USDC (*Opportunity goes to top 3 backstage wardens based on placement in this audit who RSVP.*)
- Join [C4 Discord](https://discord.gg/code4rena) to register
- Submit findings [using the C4 form](https://code4rena.com/contests/2024-05-canto/submit)
- [Read our guidelines for more details](https://docs.code4rena.com/roles/wardens)
- Starts May 30, 2024 20:00 UTC
- Ends June 20, 2024 20:00 UTC

## Automated Findings / Publicly Known Issues

The 4naly3er report can be found [here](https://github.com/code-423n4/2024-05-canto/blob/main/4naly3er-report.md).



_Note for C4 wardens: Anything included in this `Automated Findings / Publicly Known Issues` section is considered a publicly known issue and is ineligible for awards._
## 🐺 C4: Begin Gist paste here (and delete this line)

# Overview

[ ⭐️ SPONSORS: add info here ]

## Links

- **Previous audits:**  https://code4rena.com/reports/2022-07-canto
https://code4rena.com/reports/2022-06-canto
https://code4rena.com/reports/2022-06-canto-v2
https://code4rena.com/reports/2022-09-canto
https://code4rena.com/reports/2022-11-canto
https://code4rena.com/reports/2023-06-canto
https://code4rena.com/reports/2023-08-verwa
https://code4rena.com/reports/2023-10-canto
https://code4rena.com/reports/2023-11-canto
https://code4rena.com/reports/2024-01-canto
https://code4rena.com/reports/2024-03-neobase
https://code4rena.com/reports/2024-03-canto
  - ✅ SCOUTS: If there are multiple report links, please format them in a list.
- **Documentation:** https://docs.canto.io/
- **Website:** 🐺 CA: add a link to the sponsor's website
- **X/Twitter:** 🐺 CA: add a link to the sponsor's Twitter
- **Discord:** 🐺 CA: add a link to the sponsor's Discord

---

# Scope

[ ✅ SCOUTS: add scoping and technical details here ]

### Files in scope
- ✅ This should be completed using the `metrics.md` file
- ✅ Last row of the table should be Total: SLOC
- ✅ SCOUTS: Have the sponsor review and and confirm in text the details in the section titled "Scoping Q amp; A"

*For sponsors that don't use the scoping tool: list all files in scope in the table below (along with hyperlinks) -- and feel free to add notes to emphasize areas of focus.*

| Contract | SLOC | Purpose | Libraries used |  
| ----------- | ----------- | ----------- | ----------- |
| [contracts/folder/sample.sol](https://github.com/code-423n4/repo-name/blob/contracts/folder/sample.sol) | 123 | This contract does XYZ | [`@openzeppelin/*`](https://openzeppelin.com/contracts/) |

### Files out of scope
✅ SCOUTS: List files/directories out of scope

## Scoping Q &amp; A

### General questions
### Are there any ERC20's in scope?: No

✅ SCOUTS: If the answer above 👆 is "Yes", please add the tokens below 👇 to the table. Otherwise, update the column with "None".




### Are there any ERC777's in scope?: 

✅ SCOUTS: If the answer above 👆 is "Yes", please add the tokens below 👇 to the table. Otherwise, update the column with "None".



### Are there any ERC721's in scope?: No

✅ SCOUTS: If the answer above 👆 is "Yes", please add the tokens below 👇 to the table. Otherwise, update the column with "None".



### Are there any ERC1155's in scope?: No

✅ SCOUTS: If the answer above 👆 is "Yes", please add the tokens below 👇 to the table. Otherwise, update the column with "None".



✅ SCOUTS: Once done populating the table below, please remove all the Q/A data above.

| Question                                | Answer                       |
| --------------------------------------- | ---------------------------- |
| ERC20 used by the protocol              |       🖊️             |
| Test coverage                           | ✅ SCOUTS: Please populate this after running the test coverage command                          |
| ERC721 used  by the protocol            |            🖊️              |
| ERC777 used by the protocol             |           🖊️                |
| ERC1155 used by the protocol            |              🖊️            |
| Chains the protocol will be deployed on | OtherCanto mainnet  |

### ERC20 token behaviors in scope

| Question                                                                                                                                                   | Answer |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ------ |
| [Missing return values](https://github.com/d-xo/weird-erc20?tab=readme-ov-file#missing-return-values)                                                      |    |
| [Fee on transfer](https://github.com/d-xo/weird-erc20?tab=readme-ov-file#fee-on-transfer)                                                                  |   |
| [Balance changes outside of transfers](https://github.com/d-xo/weird-erc20?tab=readme-ov-file#balance-modifications-outside-of-transfers-rebasingairdrops) |    |
| [Upgradeability](https://github.com/d-xo/weird-erc20?tab=readme-ov-file#upgradable-tokens)                                                                 |    |
| [Flash minting](https://github.com/d-xo/weird-erc20?tab=readme-ov-file#flash-mintable-tokens)                                                              |    |
| [Pausability](https://github.com/d-xo/weird-erc20?tab=readme-ov-file#pausable-tokens)                                                                      |    |
| [Approval race protections](https://github.com/d-xo/weird-erc20?tab=readme-ov-file#approval-race-protections)                                              |    |
| [Revert on approval to zero address](https://github.com/d-xo/weird-erc20?tab=readme-ov-file#revert-on-approval-to-zero-address)                            |    |
| [Revert on zero value approvals](https://github.com/d-xo/weird-erc20?tab=readme-ov-file#revert-on-zero-value-approvals)                                    |    |
| [Revert on zero value transfers](https://github.com/d-xo/weird-erc20?tab=readme-ov-file#revert-on-zero-value-transfers)                                    |    |
| [Revert on transfer to the zero address](https://github.com/d-xo/weird-erc20?tab=readme-ov-file#revert-on-transfer-to-the-zero-address)                    |    |
| [Revert on large approvals and/or transfers](https://github.com/d-xo/weird-erc20?tab=readme-ov-file#revert-on-large-approvals--transfers)                  |    |
| [Doesn't revert on failure](https://github.com/d-xo/weird-erc20?tab=readme-ov-file#no-revert-on-failure)                                                   |    |
| [Multiple token addresses](https://github.com/d-xo/weird-erc20?tab=readme-ov-file#revert-on-zero-value-transfers)                                          |    |
| [Low decimals ( < 6)](https://github.com/d-xo/weird-erc20?tab=readme-ov-file#low-decimals)                                                                 |    |
| [High decimals ( > 18)](https://github.com/d-xo/weird-erc20?tab=readme-ov-file#high-decimals)                                                              |    |
| [Blocklists](https://github.com/d-xo/weird-erc20?tab=readme-ov-file#tokens-with-blocklists)                                                                |    |

### External integrations (e.g., Uniswap) behavior in scope:


| Question                                                  | Answer |
| --------------------------------------------------------- | ------ |
| Enabling/disabling fees (e.g. Blur disables/enables fees) | No   |
| Pausability (e.g. Uniswap pool gets paused)               |  No   |
| Upgradeability (e.g. Uniswap gets upgraded)               |   No  |


### EIP compliance checklist
N/A

✅ SCOUTS: Please format the response above 👆 using the template below👇

| Question                                | Answer                       |
| --------------------------------------- | ---------------------------- |
| src/Token.sol                           | ERC20, ERC721                |
| src/NFT.sol                             | ERC721                       |


# Additional context

## Main invariants

 

✅ SCOUTS: Please format the response above 👆 so its not a wall of text and its readable.

## Attack ideas (where to focus for bugs)
### Areas of Concern

- **SDK Bumps and Interface Changes**
  - Ensure compatibility with Cosmos-SDK v0.50.x and identify any breaking changes from v0.45.x.
  - Check for updated or deprecated APIs and interfaces in the Cosmos-SDK that may affect current functionality.

- **IBC-go v8 Updates**
  - Verify that inter-blockchain communication (IBC) functionalities remain intact and secure with the new version.
  - Identify any new features or changes in ibc-go v8 that require modifications in the existing implementation.

- **CometBFT v0.38.x Updates**
  - Assess the impact of CometBFT updates on consensus mechanisms and blockchain performance.
  - Evaluate the changes in configuration settings and ensure they are correctly applied during the upgrade.

- **Migration Settings**
  - Ensure that migration scripts and processes are correctly implemented to transition from the older versions to the new versions without data loss or corruption.
  - Validate that state transitions and data migrations do not introduce inconsistencies or vulnerabilities.

- **Dependencies and Module Updates**
  - Check for any new dependencies introduced with the updates and ensure they are correctly integrated.
  - Review the Ethermint module within Canto for EVM compatibility and assess the changes made to accommodate the new versions of dependencies.
  - Ensure that all external libraries and modules are updated to their compatible versions to prevent any conflicts or security issues.

- **Functionality Testing**
  - Conduct thorough testing of all previously functioning features to confirm they work as expected after the upgrade.
  - Identify any features that malfunction or do not work post-upgrade and determine the root cause.

- **Security Assessment**
  - Perform a security audit to uncover any new vulnerabilities introduced by the upgrades.
  - Ensure that security patches and updates in the new versions are correctly applied to maintain the integrity of the blockchain.

- **Documentation and Support**
  - Update all relevant documentation to reflect the changes and new configurations required for the upgraded versions.
  - Ensure that support teams are informed about the changes and are prepared to handle any issues that may arise from the upgrade.

By focusing on these areas, the audit aims to ensure a smooth transition to the updated versions while maintaining the functionality, performance, and security of the Canto and Ethermint blockchain applications.


### **Background Information**

The current parameters and state of Canto v7 can be reviewed on [Polkachu](https://www.polkachu.com/networks/canto). For details on new and changed parameters, as well as module settings in v8, refer to the v8 upgrade handler code. The primary references for the upgrade include:

- Major changes from [v0.45 to v0.47](https://github.com/cosmos/cosmos-sdk/blob/v0.47.11/UPGRADING.md) and [v0.47 to v0.50](https://github.com/cosmos/cosmos-sdk/blob/v0.50.6/UPGRADING.md).
- The actual differences as seen in the [Canto pull request](https://github.com/Canto-Network/Canto/pull/126/).
- The actual differences as seen in the [Ethermint pull request](https://github.com/Canto-Network/ethermint/pull/2).

# Canto App

- target branch : [https://github.com/b-harvest/Canto/tree/dudong2/feat/canto-main-cosmos-sdk%40v0.50.x](https://github.com/b-harvest/Canto/tree/dudong2/feat/canto-main-cosmos-sdk%40v0.50.x)

## Scope

| File | SLOC |
| --- | --- |
| app/app.go | 1480 |
| app/ante/ante.go | 64 |
| app/ante/handler_options.go | 149 |
| app/upgrades/v8/upgrades.go | 69 |
| x/csr/keeper/keeper.go | 65 |
| x/csr/keeper/msg_server.go | 38 |
| x/csr/module.go | 170 |
| x/erc20/keeper/keeper.go | 65 |
| x/erc20/keeper/msg_server.go | 627 |
| x/erc20/module.go | 159 |
| x/govshuttle/keeper/keeper.go | 84 |
| x/govshuttle/keeper/msg_server.go | 54 |
| x/govshuttle/module.go | 152 |
| x/epochs/keeper/keeper.go | 43 |
| x/epochs/keeper/abci.go | 66 |
| x/epochs/module.go | 183 |
| x/inflation/keeper/keeper.go | 77 |
| x/inflation/keeper/hooks.go | 141 |
| x/inflation/module.go | 187 |
| x/coinswap/keeper/keeper.go | 375 |
| x/coinswap/keeper/msg_server.go | 175 |
| x/coinswap/module.go | 167 |
| x/onboarding/ibc_middleware.go | 76 |
| x/onboarding/keeper/keeper.go | 102 |
| x/onboarding/keeper/msg_server.go | 38 |
| x/onboarding/keeper/ibc_callbacks.go | 170 |
| x/onboarding/module.go | 152 |
| go.mod | 260 |

### app

- app.go
    - align the overall code with the structure of the cosmos-sdk's [app.go](https://github.com/cosmos/cosmos-sdk/blob/v0.50.6/simapp/app.go), except for some parts related canto's modules and issue(ref. [[Discussion] Policy Decision on BlockedAddrs in SDK v0.50 bump #122](https://github.com/Canto-Network/Canto/issues/122))
- ante
    - name change and few other structural changes related to authz and issue(ref. [[Discussion] Use min_commission_rate Instead of ValidatorCommissionDecorator #124](https://github.com/Canto-Network/Canto/issues/124))
- upgrades.go
    - add migration code due to upgrade of cosmos-sdk and ibc-go

### modules

- cli
    - remove legacy REST(ref. [feat!: remove legacy REST cosmos/cosmos-sdk#9594](https://github.com/cosmos/cosmos-sdk/pull/9594))
- msgs
    - move `ValidateBasic` logic to msg_server
    - remove funcs like Route, Type, GetSignBytes, GetSigners because they are supported by `protobuf`. But in case of `MsgConvertERC20`, GetSigners func is alive and implemented GetSignersFromMsgConvertERC20V2 func additionally.(ref. [feat: bump up cosmos-sdk v0.50.x, cometbft v0.38.x b-harvest/Canto#53 (comment)](https://github.com/b-harvest/Canto/pull/53#discussion_r1574149712))
    - the module-specific proposal handler and param change was switched to `msg level` so `authority` has been added to the required module.(ref. [https://github.com/cosmos/cosmos-sdk/issues/10952](https://github.com/cosmos/cosmos-sdk/issues/10952))
- keeper
    - KVStore has been changed to use `KVStoreService` as opposed to the prev way of accessing KVStore via store key.(ref. [refactor(x/staking)!: KVStoreService, return errors and use context.Context cosmos/cosmos-sdk#16324](https://github.com/cosmos/cosmos-sdk/pull/16324), ...)

## NOTE

- `x/onboarding` module is a feature-driven of the `ibc`. So, need to ensure that features don't behave differently compared to canto v7, as the version of the `ibc-go` bump up.
- as the versions of different modules in `go.mod` are upgraded, need to ensure that there are no security issues.

# Ethermint

- target branch : [https://github.com/b-harvest/ethermint/tree/dudong2/feat/cosmos-sdk%40v0.50.x-cometbft%40v0.38.0-2](https://github.com/b-harvest/ethermint/tree/dudong2/feat/cosmos-sdk%40v0.50.x-cometbft%40v0.38.0-2)

## Scope

| File | SLOC |
| --- | --- |
| app/app.go | 1094 |
| app/ante/ante.go | 177 |
| app/ante/handler_options.go | 109 |
| app/upgrades.go | 143 |
| server/start.go | 768 |
| server/util.go | 169 |
| rpc/backend/sign_tx.go | 168 |
| rpc/backend/node_info.go | 359 |
| rpc/namespaces/ethereum/eth/filters/api.go | 653 |
| rpc/namespaces/ethereum/eth/filters/filter_system.go | 340 |
| x/feemarket/types/msg.go | 5 |
| x/feemarket/keeper/keeper.go | 137 |
| x/feemarket/keeper/msg_server.go | 35 |
| x/feemarket/autocli.go | 48 |
| x/feemarket/module.go | 174 |
| x/evm/types/msg.go | 487 |
| x/evm/keeper/keeper.go | 398 |
| x/evm/keeper/msg_server.go | 170 |
| x/evm/autocli.go | 74 |
| x/evm/module.go | 193 |
| go.mod | 256 |

### app

- app.go
    - align the overall code with the structure of the cosmos-sdk's [app.go](https://github.com/cosmos/cosmos-sdk/blob/v0.50.6/simapp/app.go), except for some parts related canto's modules and issue(ref. [[Discussion] Policy Decision on BlockedAddrs in SDK v0.50 bump Canto#122](https://github.com/Canto-Network/Canto/issues/122))
- ante
    - name change and few other structural changes related to authz.
- upgrades.go
    - add migration code due to upgrade of cosmos-sdk and ibc-go

### rpc

- as go-ethereum versions go up, the implementation of funcs used internally may have changed

### modules

- cli
    - remove legacy REST(ref. [feat!: remove legacy REST cosmos/cosmos-sdk#9594](https://github.com/cosmos/cosmos-sdk/pull/9594))
    - introduce `autocli`, so removes contents of tx.go, query.go
- msgs
    - move `ValidateBasic` logic to msg_server
    - remove funcs like Route, Type, GetSignBytes, GetSigners because they are supported by `protobuf`. But in case of `MsgEthereumTx`, GetSigners func is alive and implemented GetSignersFromMsgEthereumTxV2 func additionally.(ref. [feat: bump up cosmos-sdk v0.50.x, cometbft v0.38.x b-harvest/Canto#53 (comment)](https://github.com/b-harvest/Canto/pull/53#discussion_r1574149712))
    - the module-specific proposal handler and param change was switched to `msg level` so `authority` has been added to the required module.(ref. [https://github.com/cosmos/cosmos-sdk/issues/10952](https://github.com/cosmos/cosmos-sdk/issues/10952))
- keeper
    - KVStore has been changed to use `KVStoreService` as opposed to the prev way of accessing KVStore via store key.(ref. [refactor(x/staking)!: KVStoreService, return errors and use context.Context cosmos/cosmos-sdk#16324](https://github.com/cosmos/cosmos-sdk/pull/16324), ...)

✅ SCOUTS: Please format the response above 👆 so its not a wall of text and its readable.

## All trusted roles in the protocol

N/A

✅ SCOUTS: Please format the response above 👆 using the template below👇

| Role                                | Description                       |
| --------------------------------------- | ---------------------------- |
| Owner                          | Has superpowers                |
| Administrator                             | Can change fees                       |

## Describe any novel or unique curve logic or mathematical models implemented in the contracts:

N/A

✅ SCOUTS: Please format the response above 👆 so its not a wall of text and its readable.

## Running tests

To run local testnet

```bash
git clone https://github.com/code-423n4/2024-05-canto
make install
./init.sh
```

To run unit tests
```bash
make test
```

To run simulation (fuzz test)
```bash
make test-sim-nondeterminism
make test-sim-custom-genesis-fast
make test-sim-import-export
make test-sim-after-import
make test-sim-custom-genesis-multi-seed
make test-sim-multi-seed-short
make test-sim-benchmark-invariants
```



✅ SCOUTS: Please format the response above 👆 using the template below👇

```bash
git clone https://github.com/code-423n4/2023-08-arbitrum
git submodule update --init --recursive
cd governance
foundryup
make install
make build
make sc-election-test
```
To run code coverage
```bash
make coverage
```
To run gas benchmarks
```bash
make gas
```

✅ SCOUTS: Add a screenshot of your terminal showing the gas report
✅ SCOUTS: Add a screenshot of your terminal showing the test coverage

## Miscellaneous
Employees of [SPONSOR NAME] and employees' family members are ineligible to participate in this audit.

