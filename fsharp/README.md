## Run application
Enter the following command in a terminal to run application from the source code.

```bash
cd ./csharp
dotnet run --project ./catdance/catdance.fsproj
```


Or, build and run the .dll file.

```bash
cd ./csharp
dotnet build ./catdance/catdance.fsproj
dotnet ./catdance/bin/Debug/netcoreapp3.1/catdance.dll
```

Then, navigate to https://localhost:5001/.
