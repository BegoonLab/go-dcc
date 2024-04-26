Commands to upload firmware using RPi

```bash
$ stty -F /dev/ttyACM0 1200
$ avrdude -C /etc/avrdude.conf -v -patmega4809 -cjtag2updi -P/dev/ttyACM0 -b115200 -e -D -Uflash:w:firmware.hex
```
