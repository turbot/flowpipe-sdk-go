generate:
	java -jar -Dio.swagger.v3.parser.util.RemoteUrl.trustAll=true /home/vscode/openapi-generator/openapi-generator-cli.jar generate \
	-i https://host.docker.internal:7103/api/v0/docs/openapi.json -g go \
	-t ./templates/go \
	--package-name flowpipeapi --git-repo-id flowpipe-sdk-go --git-user-id turbot --remove-operation-id-prefix \
	--additional-properties structPrefix=true --additional-properties generateInterfaces=false

