module github.com/aws/aws-sdk-go-v2/internal/configsources/configtesting

go 1.20

require (
	github.com/aws/aws-sdk-go-v2/config v1.27.19
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.10
)

require (
	github.com/aws/aws-sdk-go-v2 v1.28.0 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.17.19 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.6 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.10 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.11.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.11.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.20.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.24.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.28.13 // indirect
	github.com/aws/smithy-go v1.20.2 // indirect
)

replace github.com/aws/aws-sdk-go-v2 => ../../../

replace github.com/aws/aws-sdk-go-v2/config => ../../../config/

replace github.com/aws/aws-sdk-go-v2/credentials => ../../../credentials/

replace github.com/aws/aws-sdk-go-v2/feature/ec2/imds => ../../../feature/ec2/imds/

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../../../internal/endpoints/v2/

replace github.com/aws/aws-sdk-go-v2/internal/ini => ../../../internal/ini/

replace github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding => ../../../service/internal/accept-encoding/

replace github.com/aws/aws-sdk-go-v2/service/internal/presigned-url => ../../../service/internal/presigned-url/

replace github.com/aws/aws-sdk-go-v2/service/sso => ../../../service/sso/

replace github.com/aws/aws-sdk-go-v2/service/ssooidc => ../../../service/ssooidc/

replace github.com/aws/aws-sdk-go-v2/service/sts => ../../../service/sts/
