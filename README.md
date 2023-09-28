# andotp-converter

Sadly andOTP is not maintained anymore. While there are 2FA apps like https://getaegis.app/ which allow
importing TOTP secrets from andOTP I wanted to migrate to [2FAS](https://2fas.com/). Since 2FAS can't 
import andOTP data and I didn't want to take a detour over Aegis I created a small toool which
can convert andOTP exports into 2FAS backup files.
For my case it worked sufficiently well.

## Usage

`andotp-converter <path-to-andotp-export> <where-the-2fas-file-should-be-created>` That's it.
Please note that you need to use an unecrypted andOTP export and this tool creates an unecrypted
2FAS backup file.
