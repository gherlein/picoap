# Pico W Access Point in TinyGo

This is my fledgling attempt at something like this.  

This is using the excellent [cyw43439](https://github.com/soypat/cyw43439) work.

## Status

I can activate the board as an AP and connect to it.  But, I don't yet know how to link the IP stack to it.  I am sure it uses the [seq](https://github.com/soypat/seqs/tree/main) library but I don't understand how to do that yet.  

I have opened an [issue](https://github.com/soypat/cyw43439/issues/51) to get some help.

## References

* [Pico W access point example code](https://github.com/raspberrypi/pico-examples/blob/master/pico_w/wifi/access_point/picow_access_point.c)
* [Another driver, just for reference](https://github.com/georgerobotics/cyw43-driver/tree/8ef38a6d32c54f850bff8f189bdca19ded33792a/src)
* [Infineon Data Sheet/App Notes](https://github.com/gherlein/picoap/tree/main/docs)