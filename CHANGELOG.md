# Changelog

---
## [0.12.0] - 2024-06-06

### Changed
- Now dispatch window will perceive prover amount
- The project & prover contract history data will store in local db

---
## [0.10.1] - 2024-04-26

### Changed
- Use contract multi call to query block accurate prover&project contract data
- Change block memeory snapshot logic, now scheduler data more determinate
- Support "RequiredProverAmountHash" project attribute, now project can define required provers, and change it at running
- All chain monitor logic summarized in one place, avoid data inconsistencies

### Added
- task metrics

---
## [0.10.0] - 2024-04-12

### Changed
- Support new project & prover manager contract
- Support cache project & prover snapshot
- When prover restart, will rebuild schedule status when start
- Now scheduler more robust

### Added
- Support load local project file

---
## [0.9.0] - 2024-04-07

### Added
- Support task window
- Support task timeout
- Support task retry
- Support retrive task within project

---
## [0.8.0] - 2024-03-25

### Added
- Support prover scheduling
- Support multi prover
- Support wasm vm

### Changed
- Rename "znode" to "prover"
- Rename "enode" to "coordinator"
- Rename "http_plugin" to "sequencer"

---
## [0.7.0] - 2024-03-15

### Added
- Support output to textile
- Add http plugin for message receive and pack
- Support postgres data source

### Changed
- Enode now support pull task from data source

---
## [0.6.4] - 2024-02-26

### Added
- Sandbox contract support
- More unit test

### Changed
- Output support snark proof pack

---
## [0.6.1] - 2024-02-19

### Changed
- When message have not packed to task, query message will return reveived state
- Znode will auto join project p2p topic

---
## [0.6.0] - 2024-02-15

### Added
- Support project define message aggregation strategy
- Support did auth token

---
## [0.5.0] - 2023-12-28

### Added
- Support project self-define contract abi and method
- Support message raw data as output param

### Changed
- Powerc20, add a router contract, and use contract address for challenge

---
## [0.4.1] - 2023-12-20

### Added
- Support task
- Support powerc20 miner

---
## [0.4.0] - 2023-12-07

### Added
- Support risc0 local verify proof
- Support halo2 local verify proof
- Support zkwasm local verify proof
- Support p2p network

### Changed
- Update readme, replace Zero-Node with W3bstream  

---
## [0.3.1] - 2023-11-30

### Added
- Use ioctl control W3bstream sprout
- Support Zkwasm
- Support Halo2 VM
- Support build circuit and circuit template
- Support project config loading
- Support ioctl ws query message state

### Changed
- Refactor readme

### Fixed

