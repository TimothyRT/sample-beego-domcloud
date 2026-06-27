Start-Process "http://localhost:9067/"
while ($true) {
    if ((Split-Path -Leaf (Get-Location)) -eq "sample-beego-domcloud") {
        break
    }

    $parent = Split-Path (Get-Location) -Parent
    if ($parent -eq (Get-Location).Path) {
        throw "No parent directory named 'sample-beego-domcloud' found."
    }

    Set-Location ..
}  # go upward until current folder is named "sample-beego-domcloud"
$env:PORT = "9067"
bee run