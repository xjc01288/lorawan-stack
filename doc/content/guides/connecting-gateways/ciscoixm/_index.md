---
title: "Cisco LoRaWAN Gateway"
description: ""
weight: 1
---

The Cisco LoRaWAN Gateway is a LoRaWAN gateway, whose technical specifications can be found in [the official documentation](https://www.cisco.com/c/en/us/products/routers/wireless-gateway-lorawan/). This page guides you to connect it to {{% tts %}}.

## Prerequisites

1. User account on {{% tts %}} with rights to create Gateways.
2. Cisco LoRaWAN gateway with latest firmware version `2.0.30`.

## Registration

Create a gateway by following the instructions for the [Console]({{< ref "/guides/getting-started/console#create-gateway" >}}) or the [CLI]({{< ref "/guides/getting-started/cli#create-gateway" >}}). the **EUI** is derivated from the **MAC_ADDRESS** that can be found on the back panel of the gateway. To get the **EUI** from the **MAC_ADDRESS** insert FFFE after the first 6 characters to make it a 64bit EUI.

Create an API Key with Gateway Info rights for this gateway using the same instructions. Copy the key and save it for later use.

## Configuration

Find the IP address the gateway. This can be done in various ways. You can connect your machine to the same local network as that of the gateway ethernet connection and scan for open SSH ports or assign a static IP to the gateway and use that. Once the gateway IP address is found, SSH into it.

```bash
$ ssh admin@<GatewayIP>
```
The default password for the **admin** user is typically **admin** or **cisco**.

### System setup

First you need to enable the privileged mode

```bash
$ enable
```

enter your secret (see Enable Authentification section), the default password is typically **admin** or **cisco**.

#### Network

To configure your Cisco Gateway to your network, type the following commands:

```bash
$ configure terminal # Enter global configuration mode
$ interface FastEthernet 0/1 # Enter Ethernet configuration
```

If your local network has a **DHCP server** attributing IPs:

```bash
$ ip address dhcp
```

Otherwise, if you know the **static IP address** of your gateway:

```bash
$ ip address <ip-address> <subnet-mask>
```

Next, type the following to save the network configuration of your gateway:

```bash
$ description Ethernet # Replace "Ethernet" by any description you want
$ exit # Exit Ethernet configuration
$ exit # Exit global configuration mode
$ copy running-config startup-config # Save the configuration
```

You can test your Internet configuration with the `ping` command, for example:

```bash
$ ping ip 8.8.8.8 # Ping Google's DNS server
```

To see more information about the gateway's IP and the network, you can use `show interfaces FastEthernet 0/1`, `show ip interfaces FastEthernet 0/1` or `show ip route`.

#### Date and time

To configure your system's date and time, you can use **ntp**:

```bash
$ configure terminal # Enter global configuration mode"
$ ntp server address <NTP server address> # Configure ntp on an ntp address
# OR
$ ntp server ip <NTP server IP> # Configure ntp on an IP

$ exit # Exit global configuration mode
```

If you don't have production-grade **ntp** servers available, you can use **[pool.ntp.org](http://www.pool.ntp.org/en/use.html)'s servers**.

#### FPGA

If you needed to update your gateway firmware previously, your FPGA will need ~20 minutes to update once the new firmware is installed. The packet forwarder will not work until then, so we recommend at this point waiting until the FPGA is upgraded.

To show the status of the FPGA, you can use the following command:

```bash
$ show inventory
```

When the **FPGAStatus** line indicates **Ready**, this means you can go forward with this guide.

#### GPS

If you have a GPS connected to your Cisco gateway, enable it with the following commands:

```bash
$ configure terminal # Enter global configuration mode
$ gps ubx enable # Enable GPS
$ exit # Exit global configuration mode
```

> 📜 This command may return the message `packet-forwarder firmware is not installed`, this message can be ignored.

#### Enable radio

As a final step before setting up the packet forwarder software, we're going to **enable the radio**. You can see radio information with the `show radio` command:

```bash
$ show radio 
LORA_SN: FOC21028R8S
LORA_PN: 95.1602T01
LORA_SKU: 915
LORA_CALC: <NA,NA,NA,50,31,106,97,88,80,71,63,53,44,34,25,16-NA,NA,NA,54,36,109,100,91,83,74,66,57,48,39,30,21>
CAL_TEMP_CELSIUS: 31
CAL_TEMP_CODE_AD9361: 87
RSSI_OFFSET: -204.00,-204.40
LORA_REVISION_NUM: C0
RSSI_OFFSET_AUS: -203.00,-204.00

radio status: 
on
```

If the radio is off, enable with it `no radio off`.

> 📜 The `show radio` command also shows you more information about the LoRa concentrator powering the gateway. For example, **LORA_SKU** indicates the base frequency of the concentrator.

#### Enable authentication

To prevent unauthorized access to the gateway, you'll want to set up user authentication. The Cisco gateway has a **secret** system, that requires users to enter a secret to access privileged commands.

To enable this secret system, you can use the following commands:

+ `configure terminal` to enter global configuration mode.
+ To set the secret, you can use different commands:
  + `enable secret <secret>` to enter in plaintext the secret you wish to set, instead of `<secret>`. *Note*: Special characters cannot be used in plain secrets.
  + `enable secret 5 <secret>` to enter the secret **md5-encrypted**, instead of `<secret>`.
  + `enable secret 8 <secret>` to enter the secret **SHA512-encrypted**, instead of `<secret>`.
+ `exit` to exit global configuration mode.
+ `copy running-config startup-config` to save the configuration.

#### Verifications

Before we install the packet forwarder, let's run verifications to ensure that the gateway is ready.

+ Type `show radio` to verify that the **radio is enabled**. The result should indicate **radio status: on**.
+ Type `show inventory` to verify that the **FPGAStatus is Ready**.
+ Type `show gps status` to verify that the **GPS is correctly connected**. You can get additional GPS metadata by typing `show gps info`.
+ Verify that the **network connection is working**. You can test this by pinging common ping servers with `ping ip <IP>`, if your local network does not block ping commands. For example, you can ping Google's servers with `ping ip 8.8.8.8`.

If some of those checks fail, go back to the appropriate section earlier in order to fix it.

Then save the configuration by executing

```bash
$ copy running-config startup-config
```

### Packet forwarder configuration

> ⚠️ Keep in mind that the pre-installed packet forwarder is not supported by Cisco for production purposes.

To run the packet forwarder, we'll make use of the **container** that is running on the gateway at all times.

```bash
$  request shell container-console
```
You will be requested to enter the System Password. By default this is **admin**.

Create the directory to store the Packet Forwarder configuration:

```bash
$  mkdir /etc/pktfwd
```
Copy the packet forwarder to **/etc/pktfwd**:

```bash
$  cp /tools/pkt_forwarder /etc/pktfwd/pkt_forwarder
```

Cisco provides a list of configuration templates. To list the available templates:

```bash
$  ls -l /tools/templates
  total 136
  -rwxr-xr-x    1 65534    65534        11323 Oct  8 13:30 config_loc_dual_antenna_8ch_full_diversity_EU868.json
  -rwxr-xr-x    1 65534    65534        11248 Oct  8 13:30 config_loc_dual_antenna_8ch_full_diversity_JP920.json
  -rwxr-xr-x    1 65534    65534        11323 Oct  8 13:30 config_loc_dual_antenna_8ch_partial_diversity_EU868.json
  -rwxr-xr-x    1 65534    65534         7993 Oct  8 13:30 config_loc_single_antenna_16ch_EU868.json
  -rwxr-xr-x    1 65534    65534         7080 Oct  8 13:30 config_loc_single_antenna_16ch_US915.json
  -rwxr-xr-x    1 65534    65534        13519 Oct  8 13:30 config_loc_single_antenna_64ch_US915.json
  -rwxr-xr-x    1 65534    65534        13635 Oct  8 13:30 config_loc_single_antenna_full_duplex_64ch_US915.json
  -rwxr-xr-x    1 65534    65534        17478 Oct  8 13:30 config_test_dual_antenna_56ch_partial_diversity_EU868.json
  -rwxr-xr-x    1 65534    65534        14148 Oct  8 13:30 config_test_single_antenna_64ch_64x1_EU868.json
  -rwxr-xr-x    1 65534    65534        14148 Oct  8 13:30 config_test_single_antenna_64ch_8x8_EU868.json
```

Copy the configuration template **config_loc_dual_antenna_8ch_full_diversity_EU868.json** as **config.json** to **/etc/pktfwd**:

```bash
$  cp /tools/templates/config_loc_dual_antenna_8ch_full_diversity_EU868.json /etc/pktfwd/config.json
```

Edit the configuration using a text editor, such as **vi**:

```bash
$  vi /etc/pktfwd/config.json
```

>Note: Press the `i` key on your keyboard to start insert mode. Once finished editing, press `ESC` and enter `:wq` to write the file and quit.

Edit the **gateway_conf** parameters.

1. **gateway_ID**: The **EUI** that you used for registration.
2. **server_address**: Address of the Gateway Server. If you followed the [Getting Started guide]({{< ref "/guides/getting-started" >}}) this is the same as what you use instead of `thethings.example.com`.
3. **serv_port_up**: UDP upstream port of the Gateway Server, typically 1700.
4. **serv_port_down**: UDP downstream port of the Gateway Server, typically 1700.

Save the configuration.

You can now test the packet forwarder by executing:

```bash
$  /etc/pktfwd/pkt_forwarder -c /etc/pktfwd/config.json -g /dev/ttyS1
```

Your gateway will connect to {{% tts %}} after a couple of seconds.


Now that we know the packet forwarder is running, let's make it run permanently.

Save the following script as **/etc/init.d/S60pkt_forwarder**:

> 📜 If you are using another network than `tti.eu1.cloud.thethings.industries` replace it with the name of your network after `nslookup`.

```bash
#!/bin/sh

SCRIPT_DIR=/etc/pktfwd
SCRIPT=$SCRIPT_DIR/pkt_forwarder
CONFIG=$SCRIPT_DIR/config.json

PIDFILE=/var/run/pkt_forwarder.pid
LOGFILE=/var/log/pkt_forwarder.log

export NETWORKIP=$(nslookup tti.eu1.cloud.thethings.industries | grep -E -o "([0-9]{1,3}[\.]){3}[0-9]{1,3}" | tail -1)
sed -i 's/[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}/'$NETWORKIP'/g' "$CONFIG"

start() {
  echo "Starting pkt_forwarder"
  cd $SCRIPT_DIR
  start-stop-daemon \
        --start \
        --make-pidfile \
        --pidfile "$PIFDILE" \
        --background \
        --startas /bin/bash -- -c "exec $SCRIPT -- -c $CONFIG -g /dev/ttyS1 >> $LOGFILE 2>&1"
  echo $?
}

stop() {
  echo "Stopping pkt_forwarder"
  start-stop-daemon \
        --stop \
        --oknodo \
        --quiet \
        --pidfile "$PIDFILE"
}

restart() {
  stop
  sleep 1
  start
}

case "$1" in
  start)
    start
    ;;
  stop)
    stop
    ;;
  restart|reload)
    restart
    ;;
  *)
    echo "Usage: $0 {start|stop|restart}"
    exit 1
esac

exit $?
```

Then make the init script executable:

```bash
$  chmod +x /etc/init.d/S60pkt_forwarder
```

To enable it immediately, execute 

```bash
$  /etc/init.d/S60pkt_forwarder start
```

Now exit the container: enter **ctrl A** then **q** (the characters won't be displayed on the screen).

Finally reboot the gateway: ```reload```