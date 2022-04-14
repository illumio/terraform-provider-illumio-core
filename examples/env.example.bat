set ILLUMIO_PCE_HOST="https://pce.my-company.com:8443"
set ILLUMIO_PCE_ORG_ID="1"
set ILLUMIO_API_KEY_USERNAME="api_xxxxxx"
set ILLUMIO_API_KEY_SECRET="xxxxxxxxxxxx"

set ILLUMIO_AES_GCM_KEY="aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

:: Define the TF variables set up for the examples

set TF_VAR_pce_url="%ILLUMIO_PCE_HOST%"
set TF_VAR_pce_org_id="%ILLUMIO_PCE_ORG_ID%"
set TF_VAR_pce_api_key="%ILLUMIO_API_KEY_USERNAME%"
set TF_VAR_pce_api_secret="%ILLUMIO_API_KEY_SECRET%"
