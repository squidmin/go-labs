@echo off
FOR /R %%G in (*.go) DO (
    pushd %%~dpG
    IF NOT EXIST "vendor" (
        go build
    )
    popd
)
echo Build complete.
