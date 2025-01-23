# Linux setup

Just add your user to the dialout group so you have appropriate permissions on the device and Fyne error message at program startup: "Error parsing user locale C". Finally, reboot host.

```bash
sudo usermod -a -G dialout $USER
ls -l /dev/ttyUSB0
crw-rw---T 1 root dialout 188, 0 Feb 12 12:01 /dev/ttyUSB0
sudo update-locale LANG=fr_FR.utf8
sudo reboot
```


