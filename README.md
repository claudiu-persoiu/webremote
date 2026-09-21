# webremote
Forget about your keyboard and mouse, just sit back and use this web remote

## About

This app provides a virtual keyboard and mouse. It runs on Linux, Windows and Mac.

To see all the available options use the command:

    webremote -help


By default `uinput` is used, if you are using it on windows run it with:

    webremote -exec xdotool

When the app is started it will display the ip/host and port of the computer that is controlling. Open this link into a browser in the remote device and enjoy the remote.

## Build

    go build ./

The output file it will be `webremote`.

In Linux and MacOS make sure the file has execution rights.

## Execution options

The project uses [uinput](https://github.com/bendahl/uinput) and [xdotool](https://github.com/jordansissel/xdotool).

To change between options use the flag `-exec`.

To use the app with `xdotool` the utility must be installed on the computer.

## Limitations

### uinput on Linux
`uinput` requires root to run, this makes it unsafe. However this option doesn't need an additional utility to be installed, like in the case of xdotool.

To use it without `sudo` or `root`, considering using:

    sudo groupadd uinput
    sudo usermod -a -G uinput $USER
    sudo udevadm control --reload-rules
    echo "SUBSYSTEM==\"misc\", KERNEL==\"uinput\", GROUP=\"uinput\", MODE=\"0660\"" | sudo tee /etc/udev/rules.d/uinput.rules
    echo uinput | sudo tee /etc/modules-load.d/uinput.conf

### IPs can't be accessed inside the network

If you can't open the remote on the second device. Make sure the device and remote device (mobile phone) are on the same network. Make sure there isn't a network level restriction that is preventing the devices inside the network from accessing each other.

## Related projects
https://github.com/bendahl/uinput

https://github.com/jordansissel/xdotool