# Linux setup

Just add your user to the dialout group so you have appropriate permissions on the device.

```bash
sudo usermod -a -G dialout $USER

ls -l /dev/ttyUSB0
```

in my case the output is:

```bash
crw-rw---T 1 root dialout 188, 0 Feb 12 12:01 /dev/ttyUSB0
sudo reboot
```

and reboot


