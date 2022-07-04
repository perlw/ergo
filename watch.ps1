# DISCLAIMER:
# It was late, I was tired, I don't know know enough powershell, and I got a surge of inspiration.
# Looking back at it now it just feels like a bloated monstrosity.
# It does, however, work pretty much as expected. Simply sourcing the file will set up and
# start the filewatcher. And it does the job, albeit without any useful output from the service
# or build.
# But, I'm happy with this. I guess.

$knownFiles = @{}

Function Register-Watcher {
    param ($folder)
    $absPath = ($folder | Resolve-Path)
    $filter = "*.*"
    $watcher = New-Object IO.FileSystemWatcher $absPath, $filter -Property @{
      IncludeSubdirectories = $true
      EnableRaisingEvents = $true
    }

    $changeAction = [scriptblock]::Create({
      try {
        $path = $Event.SourceEventArgs.FullPath
        $changedAt = $Event.TimeGenerated

        if (Test-Path -Path $path -PathType Container) {
          return
        }
        if ($path -like "*~") {
          return
        }

        # Write-Host ($knownFiles | Format-Table | Out-String)

        if ($knownFiles.ContainsKey($path)) {
          $compareWith = $knownFiles[$path]
          $difference = New-TimeSpan -Start $compareWith -End $changedAt
          if ($difference.Seconds -lt 1) {
            return
          }
        }
        $knownFiles[$path] = $changedAt

        Write-Host "Updated: $path"
        Write-Host "Ergo stopped!"
        Stop-Process -Name ergo
        if ($path -like "*.go") {
          Write-Host "Rebuilding Ergo..."
          go build -o .\bin\ergo.exe .\cmd\ergo
        }
        Write-Host "Starting Ergo..."
        Start-Job -ScriptBlock{.\bin\ergo.exe -port 1337 -web-base-dir .\web}
      } catch {
        Write-Host "An error occurred:"
        Write-Host $_
      }
    })

    Register-ObjectEvent $Watcher -EventName "Changed" -Action $changeAction
}

# I know this is bad practice (unregistering _all_ events), but works for me.
Get-EventSubscriber | Unregister-Event
Get-Job | Remove-Job -Force

Register-Watcher .\cmd
Register-Watcher .\web
