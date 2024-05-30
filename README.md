# Canto audit details
- Total Prize Pool: $100,000 in USDC
  - HM awards: $79,700 in USDC
  - QA awards: $3,300 in USDC 
  - Judge awards: $10,000 in USDC
  - Validator awards: $6,500 in USDC
  - Scout awards: $500 in USDC
- Join [C4 Discord](https://discord.gg/code4rena) to register
- Submit findings [using the C4 form](https://code4rena.com/contests/2024-05-canto/submit)
- [Read our guidelines for more details](https://docs.code4rena.com/roles/wardens)
- Starts May 30, 2024 20:00 UTC
- Ends June 20, 2024 20:00 UTC

❗️Awarding Note for Wardens, Judges, and Lookouts: If you want to claim your awards in $ worth of CANTO, you must follow the steps outlined in the audit channel before the audit ends; otherwise you'll be paid out in USDC.

## Automated Findings / Publicly Known Issues

The 4naly3er report can be found [here](https://github.com/code-423n4/2024-05-canto/blob/main/4naly3er-report.md).



_Note for C4 wardens: Anything included in this `Automated Findings / Publicly Known Issues` section is considered a publicly known issue and is ineligible for awards._

# Overview

## Key dependency changes:
- Cosmos-SDK v0.45.9 to v0.50.6
- CometBFT (Tendermint) v0.34.25 to v0.38.6
- ibc-go v3.2.0 to v8.2.1
- go-ethereum v1.10.19 to v1.10.26

## Canto

### notable changed files
```
app/app.go
app/ante/ante.go
app/ante/handler_options.go
app/upgrades/v8/upgrades.go
x/csr/keeper/keeper.go
x/csr/keeper/msg_server.go
x/csr/module.go
x/erc20/keeper/keeper.go
x/erc20/keeper/msg_server.go
x/erc20/module.go
x/govshuttle/keeper/keeper.go
x/govshuttle/keeper/msg_server.go
x/govshuttle/module.go
x/epochs/keeper/keeper.go
x/epochs/keeper/abci.go
x/epochs/module.go
x/inflation/keeper/keeper.go
x/inflation/keeper/hooks.go
x/inflation/module.go
x/coinswap/keeper/keeper.go
x/coinswap/keeper/msg_server.go
x/coinswap/module.go
x/onboarding/ibc_middleware.go
x/onboarding/keeper/keeper.go
x/onboarding/keeper/msg_server.go
x/onboarding/keeper/ibc_callbacks.go
x/onboarding/module.go
```

### app
- app.go
	- align the overall code with the structure of the cosmos-sdk's [app.go](https://github.com/cosmos/cosmos-sdk/blob/v0.50.6/simapp/app.go), except for some parts related canto's modules and issue(ref. https://github.com/Canto-Network/Canto/issues/122)
- ante
	-  name change and few other structural changes related to authz and issue(ref. https://github.com/Canto-Network/Canto/issues/124)
- upgrades.go
	- add migration code due to upgrade of cosmos-sdk and ibc-go 

### modules
- cli
	- remove legacy REST(ref. https://github.com/cosmos/cosmos-sdk/pull/9594)
- msgs
	- move `ValidateBasic` logic to msg_server
	- remove funcs like Route, Type, GetSignBytes, GetSigners because they are supported by `protobuf`. But in case of `MsgConvertERC20`, GetSigners func is alive and implemented GetSignersFromMsgConvertERC20V2 func additionally.(ref. https://github.com/b-harvest/Canto/pull/53#discussion_r1574149712)
	- the module-specific proposal handler and param change was switched to `msg level` so `authority` has been added to the required module.(ref. https://github.com/cosmos/cosmos-sdk/issues/10952)
- keeper
	- KVStore has been changed to use `KVStoreService` as opposed to the prev way of accessing KVStore via store key.(ref. https://github.com/cosmos/cosmos-sdk/pull/16324, ...)

### NOTE
- `v8 package bump` will be done in a separate pr after this pr is merged to avoid amplification of file changes in that pr
- `x/onboarding` module is a feature-driven of the `ibc`. So, need to ensure that features don't behave differently compared to canto v7, as the version of the `ibc-go` bump up.
- as the versions of different modules in `go.mod` are upgraded, need to ensure that there are no security issues.

## ethermint

### notable changed files
```
app/app.go
app/ante/ante.go
app/ante/handler_options.go
app/upgrades.go
encoding/config.go
server/start.go
server/util.go
rpc/backend/sign_tx.go
rpc/backend/node_info.go
rpc/namespaces/ethereum/eth/filters/api.go
rpc/namespaces/ethereum/eth/filters/filter_system.go
x/feemarket/types/msg.go
x/feemarket/keeper/keeper.go
x/feemarket/keeper/msg_server.go
x/feemarket/autocli.go
x/feemarket/module.go
x/evm/types/msg.go
x/evm/keeper/keeper.go
x/evm/keeper/msg_server.go
x/evm/autocli.go
x/evm/module.go
go.mod
```

### app
- app.go
	- align the overall code with the structure of the cosmos-sdk's [app.go](https://github.com/cosmos/cosmos-sdk/blob/v0.50.6/simapp/app.go), except for some parts related canto's modules and issue(ref. https://github.com/Canto-Network/Canto/issues/122)
- ante
	-  name change and few other structural changes related to authz.
- upgrades.go
	- add migration code due to upgrade of cosmos-sdk and ibc-go

### rpc
- as go-ethereum versions go up, the implementation of funcs used internally may have changed

### modules
- cli
	- remove legacy REST(ref. https://github.com/cosmos/cosmos-sdk/pull/9594)
	- introduce `autocli`, so removes contents of tx.go, query.go
- msgs
	- move `ValidateBasic` logic to msg_server
	- remove funcs like Route, Type, GetSignBytes, GetSigners because they are supported by `protobuf`. But in case of `MsgEthereumTx`, GetSigners func is alive and implemented GetSignersFromMsgEthereumTxV2 func additionally.(ref. https://github.com/b-harvest/Canto/pull/53#discussion_r1574149712)
	- changed so that all proposals(like MsgUpdateParams, ...) are done via proposal in the gov module.
	- the module-specific proposal handler and param change was switched to `msg level` so `authority` has been added to the required module.(ref. https://github.com/cosmos/cosmos-sdk/issues/10952)
- keeper
	- KVStore has been changed to use `KVStoreService` as opposed to the prev way of accessing KVStore via store key.(ref. https://github.com/cosmos/cosmos-sdk/pull/16324, ...)

## NOTE
- as the versions of different modules in `go.mod` are upgraded, need to ensure that there are no security issues.


## Links

- **Previous audits:**  
* https://code4rena.com/reports/2022-07-canto
* https://code4rena.com/reports/2022-06-canto
* https://code4rena.com/reports/2022-06-canto-v2
* https://code4rena.com/reports/2022-09-canto
* https://code4rena.com/reports/2022-11-canto
* https://code4rena.com/reports/2023-06-canto
* https://code4rena.com/reports/2023-08-verwa
* https://code4rena.com/reports/2023-10-canto
* https://code4rena.com/reports/2023-11-canto
* https://code4rena.com/reports/2024-01-canto
* https://code4rena.com/reports/2024-03-neobase
* https://code4rena.com/reports/2024-03-canto
- **Documentation:** https://docs.canto.io/
- **Website:** https://canto.io/
- **X/Twitter:** https://twitter.com/CantoPublic
- **Discord:** https://discord.com/invite/63GmEXZsVf

---

# Scope

## Scoping Q &amp; A

### General questions



| Question                                | Answer                                                                                                                                                                           |
| --------------------------------------- |----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| ERC20 used by the protocol              | None                                                                                                                                                                             |
| Test coverage                           | Canto<br/>./app: 64.7% files, 36.9% statements<br/>./x: 75.6% files, 15% statements<br/> Ethermint<br/>./app: 100% files, 71.5% statements<br/>./x: 83.8% files, 19.9 statements |
| ERC721 used  by the protocol            | None                                                                                                                                                                             |
| ERC777 used by the protocol             | None                                                                                                                                                                             |
| ERC1155 used by the protocol            | None                                                                                                                                                                             |
| Chains the protocol will be deployed on | OtherCanto, mainnet                                                                                                                                                              |


### External integrations (e.g., Uniswap) behavior in scope:


| Question                                                  | Answer |
| --------------------------------------------------------- | ------ |
| Enabling/disabling fees (e.g. Blur disables/enables fees) | No   |
| Pausability (e.g. Uniswap pool gets paused)               |  No   |
| Upgradeability (e.g. Uniswap gets upgraded)               |   No  |


### EIP compliance checklist
N/A

# Additional context

## Main invariants
N/A

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

| File | SLOC | Test Coverage |
| --- | --- |--------------|
| app/app.go | 1480 | 69.2%        |
| app/ante/ante.go | 64 | 0%           |
| app/ante/handler_options.go | 149 | 0%           |
| app/upgrades/v8/upgrades.go | 69 | 82.4%        |
| x/csr/keeper/keeper.go | 65 | 100%         |
| x/csr/keeper/msg_server.go | 38 | 75%          |
| x/csr/module.go | 170 | 60%          |
| x/erc20/keeper/keeper.go | 65 | 100%         |
| x/erc20/keeper/msg_server.go | 627 | 95.3%        |
| x/erc20/module.go | 159 | 47.1%        |
| x/govshuttle/keeper/keeper.go | 84 | 0%           |
| x/govshuttle/keeper/msg_server.go | 54 | 0%           |
| x/govshuttle/module.go | 152 | 0%           |
| x/epochs/keeper/keeper.go | 43 | 83.3%        |
| x/epochs/keeper/abci.go | 66 | 100%         |
| x/epochs/module.go | 183 | 44.1%        |
| x/inflation/keeper/keeper.go | 77 | 66.7%        |
| x/inflation/keeper/hooks.go | 141 | 90.2%        |
| x/inflation/module.go | 187 | 45.7%        |
| x/coinswap/keeper/keeper.go | 375 | 81.2%        |
| x/coinswap/keeper/msg_server.go | 175 | 9.4%         |
| x/coinswap/module.go | 167 | 45.2%        |
| x/onboarding/ibc_middleware.go | 76 | 62.5% |        
| x/onboarding/keeper/keeper.go | 102 | 66.7% |     
| x/onboarding/keeper/msg_server.go | 38 | 75% |    
| x/onboarding/keeper/ibc_callbacks.go | 170 | 93.8% |
| x/onboarding/module.go | 152 | 43.8% |
| go.mod | 260 |
| Total  |   5388 |


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

| File | SLOC | Test Coverage |
| --- | --- | --------------|
| app/app.go | 1094 | 74.7% |
| app/ante/ante.go | 177 | 74.5% |
| app/ante/handler_options.go | 109 | 61.5% |
| app/upgrades.go | 143 | 56.1% |
| server/start.go | 768 | 0% |
| server/util.go | 169 | 0% |
| rpc/backend/sign_tx.go | 168 | 70.4% |
| rpc/backend/node_info.go | 359 | 50.7% |
| rpc/namespaces/ethereum/eth/filters/api.go | 653 | 0% |
| rpc/namespaces/ethereum/eth/filters/filter_system.go | 340 | 0% |
| x/feemarket/types/msg.go | 5 | 0% |
| x/feemarket/keeper/keeper.go | 137 | 87.1% |
| x/feemarket/keeper/msg_server.go | 35 |
| x/feemarket/autocli.go | 48 | 100% |
| x/feemarket/module.go | 174 | 56.2% |
| x/evm/types/msg.go | 487 | 65.1 % |
| x/evm/keeper/keeper.go | 398 | 91.2% |
| x/evm/keeper/msg_server.go | 170 | 88.9% |
| x/evm/autocli.go | 74 | 100% |
| x/evm/module.go | 193 | 48.6% |
| go.mod | 256 ||
| Total  | 5957   ||

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


## All trusted roles in the protocol
N/A

## Describe any novel or unique curve logic or mathematical models implemented in the contracts:

N/A


## Running tests

To run local testnet

First:
* Install go (version must be >= 1.21 to avoid [this bug](https://github.com/jeremylong/DependencyCheck/issues/6258))
* Install [jq](https://jqlang.github.io/jq/download/)
* Make sure go environment is set up correctly
```bash
export PATH=$PATH:$HOME/go/bin
```

```bash
git clone https://github.com/code-423n4/2024-05-canto
cd 2024-05-canto

# for Canto-main
cd canto-main
make install
./init.sh

# for Ethermint-main
cd ethermint-main
make install
./init.sh
```

To run unit tests
```bash
cd canto-main
make test
```

To run simulation (fuzz test)
```bash
# canto needs to be installed and initialized ahead of simulations
cd canto-main
make test-sim-nondeterminism
make test-sim-custom-genesis-fast
make test-sim-import-export
make test-sim-after-import
make test-sim-custom-genesis-multi-seed
make test-sim-multi-seed-short
make test-sim-benchmark-invariants
```

Test coverage
```bash
cd canto-main
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
open coverage.html
```

## Miscellaneous
Employees of Canto and employees' family members are ineligible to participate in this audit.

