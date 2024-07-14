# UUID
an alternative to the `github.com/google/uuid` package

## Installation
```bash
go get -u github.com/malivvan/uuid
```

## Usage
```go
package main

import (
    "fmt"
    "github.com/malivvan/uuid"
)

func main() {
    // Generate a new UUID
    id, err := uuid.New("TYPE", 8)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(id)
}
```

## Features
- Type Prefix
- Host Fingerprinting
- Unix Timestamp
- Data bytes with dynamic length
- CRC Checksum

## Structure
```
| TYPE-SYSTEMID-UNIXSECOND-XXXXXXXXXXXXXXXX-CHECKSUM |
|  4  |    4   |    4     |        N       |    4    | = N + 16 bytes

> TEST-C05FA96E-1720977041-F0ADF56596E2C4FB-188A8A03
```

## Credits
- Special thanks to [Denis Brodbeck for his machineID implementation](https://github.com/denisbrodbeck/machineid)