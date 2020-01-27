# YUP! (Yaml Unbound Parser)

Yaml Parser for Effortless DCA

## Overview

This is a development repository for a Yaml parser needed for Effortless DCA. It accepts customer Yaml file as input, parses and looks for data suitable for waiver and remediation, correspondingly generates a waiver and a remediation Yaml file in the output folder.


## Build and Run

```
git clone https://github.com/mudash/YUP.git
cd YUP
go build -o YUP
./YUP -input=./sample/customer_updated.yaml -output=./output

```

## Samples

### Customer Input File (Customer.yaml)

```
---
provider: CIS
benchmark: CentOS Linux 7
provider_version: v.2.2.0
controls:
# This is a permenant waiver that has no start or end date for both Scan and Remediate
  - id: 1.1.1.1_Ensure_mounting_of_cramfs_filesystems_is_disabled
    scan: 
      run: false
    remediate:
      run: false
    justification: "This is a waiver because John Snow says so"
# This is an example that remediates the control 
  - id: 1.1.1.2_Ensure_mounting_of_freevxfs_filesystems_is_disabled
    remediate:
      run: true
# This is an example that changes the value of the remediation and disables scan
  - id: 1.1.1.3_Ensure_mounting_of_jffs2_filesystems_is_disabled
    scan:
      run: false
    remediate: 
      run: true
      overlay_command:
        - local: echo "hello, this is an overlay command"
    justification: "Company policy is to do something other than what CIS says"
    # This is a time specfic waiver for remediation that also runs the scan
  - id: 1.1.1.4_Ensure_mounting_of_hfs_filesystems_is_disabled
    scan:
      run: true
    remediate:
      run: true
      Waiver:
        start_date_utc: "--- 2019-10-17 08:25:57.571436000 Z\n"
        expiration_date_utc: "--- 2029-10-14 08:25:57.571522000 Z\n"
        identifier: ticket_12345
    justification: "This is a temp waiver for ticket_12345"
    
```

### Waiver File Output (Waiver.yaml)

```
1.1.1.1_Ensure_mounting_of_cramfs_filesystems_is_disabled:
  justification: This is a waiver because John Snow says so
  run: false
1.1.1.3_Ensure_mounting_of_jffs2_filesystems_is_disabled:
  justification: Company policy is to do something other than what CIS says
  run: false

```

### Remediation File Output (Remediate.yaml)

```
provider: CIS
benchmark: CentOS Linux 7
provider_version: v.2.2.0
controls:
- id: CIS_CentOS_Linux_7_1_1_1_1
  enabled: false
- id: CIS_CentOS_Linux_7_1_1_1_2
  enabled: true
- id: CIS_CentOS_Linux_7_1_1_1_3
  enabled: true
  overlay_command:
  - local: echo "hello, this is an overlay command"
- id: CIS_CentOS_Linux_7_1_1_1_4
  enabled: true
  waiver:
    expiration_date_utc: '--- 2029-10-14 08:25:57.571522000 Z'
    identifier: ticket_12345
    start_date_utc: '--- 2019-10-17 08:25:57.571436000 Z'
    justification: This is a temp waiver for ticket_12345

```