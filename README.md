# 🤌 Portician 🤌

Periodically forward a port using [Universal Plug and Play (UPnP)](https://en.wikipedia.org/wiki/Universal_Plug_and_Play)

## 🤌 Example use case
Running Kubernetes in the cloud can get pretty expensive. For development purposes it can be beneficial to run the cluster locally instead. One way to achieve this is by using [Kind](https://kind.sigs.k8s.io/).

**However**, how do you access this cluster from outside of your home network? If you don't have a static IP, you can using Dynamic DNS (DDNS) to automatically manage your DNS records to always point at your home IP address. This is really simple to do if you use the [cloudflare-ddns project](https://github.com/timothymiller/cloudflare-ddns).

**However**, this still doesn't allow you to communicate with your cluster. The computer you're running your cluster on doesn't have an external IP. It will have an internal IP like 192.168.0.9 (because it's hidden behind a [NAT](https://en.wikipedia.org/wiki/Network_address_translation)). You now need to forward a port from your home router to the computer running the cluster. You can manually go into your router settings, assign a static IP to the computer you run your cluster on, and setup port forwarding that way.

**However**, this is kind of annoying, and all router software seems to be unbelievably terrible. In my case, my ISPs router is so trash that it won't obey static IP assignment. So instead, I made this project to just automatically forward the port(s) for me using the UPnP protocol. It figures out the internal IP that it's running on, and it tells the router to automatically forward some ports.

Niche use case? Probably.

## 🤌 How to use it
First things first, don't run it in a Linux based container on Windows. Your Docker Desktop for Windows installation uses WSL2 behind the scenes. WSL2 uses a virtual network switch (which is not configurable), so the container won't be able to communicate directly with your router, so it wont be able to use UPnP to forward the port.

Seems to work fine on Linux, or in a full Linux VM (if you set the virtual network switch to use 'Host' networking).

<ins>***First, configure it:***</ins>

Pretty simple to configure. Make sure there is a json file called `config.json` in the same directory as the executable. It should be in the following form:
```
{
    "externalport": 30778,
    "internalport": 30778,
    "internalip": "192.168.0.9",
    "portforwardduration": 3600,
    "updateInterval": 300,
    "protocol": "TCP",
    "description": "Some description"
}
```
Note that everything but `externalport` and `internalport` is optional. By default it will:
 * Use the internal IP of whatever computer the process is running on.
 * It will forward the port for an hour
 * Update the port forwarding every 5 mins
 * Use TCP protocol
 * Use some generic description 

<ins>**Then, run it:**</ins>

If you don't fall into the unfortunate category described above (ie: Docker on Windows), you can just run it in a docker container. Edit the `docker-compose.yml` file and change the volume path to point at wherever you config file sits. Then run:

```
docker compose up -d
```

Otherwise, don't run it in Docker. Go into the root directory, run
```
go build -o portician
```
followed by (assuming your config file is called `config.json` and is in the same directory):
```
./portician config.json
```

## 🤌 How does it work?
It makes significant use of [this amazing UPnP project](https://github.com/huin/goupnp). You should check it out.
