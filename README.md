# About


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

# How to build

# Running within docker compose

```bash
ls
```

Create request to api-gateway to process specified file:
```bash

```

# Concerns

### Security
 - api-gateway is unsecured, introduce https
 - event broker is unsecured, introduce certificates
 - 

### Data integrity
- design allows paraller file processing, if whole system requires processing delivery JSON files one-by-one, proper logic must be introduced
- database can go out-of memory
- in case of malformed JSON file scanning process is interrupted, it would be better to continue scanning

### Optimization:
- compress data in database
- send scanned ports in batch to database

