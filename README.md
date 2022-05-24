# Configuration File Parsers

Collection of packages to parse configuration files 

## aws-credentials Configuration Parsers

```go
package aws

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	ini "github.com/rajasoun/go-config-parsers/aws_credentials"
)

func ConfigProfilesHandler(w http.ResponseWriter, r *http.Request) {
	if handler.multiple {
		sections, err := ini.OpenFile(external.DefaultSharedCredentialsFilename())
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't parse credentials file")
		}
		respondWithJSON(w, 200, sections.List())
	} else {
		respondWithJSON(w, 200, []string{})
	}

}

```
