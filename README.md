Envver
======
[![Build Status](https://semaphoreci.com/api/v1/projects/5bea24be-d4c8-4236-9505-cda2299e66fc/495981/badge.svg)](https://semaphoreci.com/perceptive/envver)      


Retrieve Product Environment configuration from OpenAperture.

# Usage
* Run `./envver --id=[your OAuth Client ID] --secret=[your OAuth Client Secret] --auth_url=[your OAuth access token URL] --url=[your OpenAperture Manager URL] stage|build --product=[product name] --environment=[environment name]`
* Each option defaults to an environment variable (if set):
  * OA_CLIENT_ID
  * OA_CLIENT_SECRET
  * OA_AUTH_TOKEN_URL
  * OA_URL
  * OA_PRODUCT_NAME
  * OA_PRODUCT_ENVIRONMENT_NAME
* Run `./envver help` for more details.
