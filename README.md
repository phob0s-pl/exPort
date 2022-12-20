# About

The project has 3 services:
 - api gateway which exposes REST API for scanning JSON file for ports
 - port-scanner which scans requested file for ports and publishes them as event
 - port-database which stores ports from events into memory database

```
         ┌─────────────────┐
         │                 │
         │   api gateway   │
         │                 │
         └────────┬────────┘
                ▲ │
                │ │
                │ ▼
┌───────────────┴──────────────────┐
│                                  │
│        event broker (nsq)        │
│                                  │
└───────┬──────────────┬───────────┘
      ▲ │              │  ▲
      │ │              │  │
      │ ▼              ▼  │
┌─────┴─────────┐ ┌───────┴────────┐
│               │ │                │
│ port scanner  │ │ port database  │
│               │ │                │
└───────────────┘ └────────────────┘

```

# API
The API gateway has following endpoint exposed
```
/port/process-file/[filename]
```
Where filename is name of the file to process.

# How to build
Required tools:
 - go
 - make
 - golangci-lint

To build and test binaries type:
```bash
make
```
To build docker images type:
```bash
make docker-images
```

# Running within docker compose

```bash
cd iac
docker compose up
```

Create request to api-gateway to process specified file:
```bash
curl http://localhost:8080/port/process-file/ports.json
```

# Concerns

### Security
 - api-gateway is unsecured, introduce https
 - event broker is unsecured, introduce certificates
 - size of messages is not checked, possible DOS
 - scanned file should be put into dedicated directory, where scanner would accept file argument relative to the directory
 - API gateway can be behind firewall or have other security features preventing from DDOS, ( eg. OVH hardware firewall)
 - add authorization and authentication system to determine who can invoke file scanning
 - introduce log and event collection for further analysis
 - scan code with dependabot for potentially malicious dependencies
 - scan docker images for known vulnerabilities


### Data integrity
- design allows paraller file processing, if whole system requires processing delivery JSON files one-by-one, proper logic must be introduced
- database can go out-of memory
- in case of malformed JSON file scanning process is interrupted, it would be better to continue scanning

### Optimization:
- compress data in database
- send scanned ports in batch to database
- create service for monitoring
- create service for handling non-optimistic scenarios, eg: notify when file processing fails
