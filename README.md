#YUP! (Yaml Unbound Parser)

Yaml Parser for Effortless DCA

##Overview

This is a development repository for a Yaml parser needed for Effortless DCA. It accepts customer Yaml file as input, parses and looks for data suitable for waiver and remediation, and correspondingly generates a waiver and a remediation Yaml file in the output folder.

![Alt text](./YUP_overview.svg)
<img src="./YUP_overview.svg">

##Parsing Rules

1. Customer input Yaml file should have waiver_actions construct to generate an output waiver yaml file. Similarly remediate_actions is needed to generate a remediation yaml file.

2. If the waiver field is set to false or missing, the control is not included in the waiver yaml file. 

3. The run field is optional. Its default value is false for the waiver.

4. The expiration_date is optional. 

5. 