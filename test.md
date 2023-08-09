# Rev. Audit Install Process

**Required Software**
- IGT EZPay Admin
- IGT Patron Management
- IGT Table Manager
- Global Payments VIP LightSpeed
- Laserfiche 11.0

### IGT Software Configuration
Rev. audit doesn't need additional configuration for IGT appliations.

### VIP LightSpeed
- Copy the contents of `\\sss-cat-netapp1.cnb-ss.com\ITSupport$\Software\Global Payments` to the deskop.
- Run the installer for the most current version.
- The only checkbox that should be checked is the one for the software itself.
- Snag the appropriate INI file from the directory copied from the share.
- Finish the installation.
- Go back to the INI directory and copy the necessary site files to `C:\ProgramData\Global Payments\reports`.

### Post-Install Checklist
- Run IGT Patron Management as administrator using `%admin!!!` *run as administrator, not as other user*. 
- Run `gpupdate /force` after running Patron Management as `%admin!!!`
  

![Gopher](gopher.png) "Gopher"

#### Try it with PowerShell
**Get Programs**
```ps1
gwmi win32_product -computername WKS-SAL-BZV932Z | select name, vendor
```
