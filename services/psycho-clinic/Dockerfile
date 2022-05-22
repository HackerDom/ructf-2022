FROM mcr.microsoft.com/dotnet/sdk:6.0 AS builder
LABEL stage=builder

COPY ./src /app

WORKDIR /app

RUN dotnet restore && dotnet publish -c Release -o publish

FROM mcr.microsoft.com/dotnet/aspnet:6.0

WORKDIR /app

COPY --from=0 /app ./

RUN adduser --system --home /app --shell /usr/sbin/nologin --no-create-home --disabled-password deployer

USER deployer

CMD ["dotnet", "publish/psycho-clinic.dll"]