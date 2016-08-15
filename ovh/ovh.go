package ovh

import (
	"net"

	"../config"

	log "github.com/Sirupsen/logrus"
	"github.com/parnurzeal/gorequest"
	"github.com/urfave/cli"
)

const (
	ovhApiBaseUrl = "https://www.ovh.com/nic/update"
)

func UpdateRecord(c *cli.Context) error {

	if !c.Args().Present() {
		log.Fatal("The domain to update (e.g. host.example.org) should be provided as argument")
	}

	ip := getIpAddress(c)
	sendUpdateApi(config.Config.Username, config.Config.Password, c.Args().First(), ip)

	return nil
}

// OVH APIs
func sendUpdateApi(username string, password string, domain string, ip string) {

	req := gorequest.
		New().
		SetBasicAuth(username, password).
		Get(ovhApiBaseUrl).
		Param("system", "dyndns").
		Param("hostname", domain)

	if ip != "" {
		log.Infof("Passing the IP address %s for update of domain %s", ip, domain)
		req.Param("myip", ip)
	}

	log.Debugf("Request to the OVH API: %s", req)

	resp, _, errs := req.End()
	log.Debugf("Response from update record call to the OVH API: %s", resp)
	if errs != nil {
		log.Fatalf("Error while issuing the update record call to the OVH API: %s", errs)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("Error while issuing the update record call to the OVH API: %s", resp)
	}

	log.Infof("IP Address of domain %s updated!", domain)
}

// IP/Network Interface management
func getIpAddress(c *cli.Context) string {

	if config.Config.IpAddress != "" {
		log.Debugf("An IP address was provided (%s), bypassing auto-detection.", config.Config.IpAddress)
		return config.Config.IpAddress
	}

	if config.Config.NetworkInterface != "" {
		log.Debugf("A network interface was provided (%s), using it for IP detection", config.Config.NetworkInterface)
		iface, err := net.InterfaceByName(config.Config.NetworkInterface)
		if err != nil {
			log.Fatalf("Unable to find the network interface %s: %s", config.Config.NetworkInterface, err)
		}

		return getIpAddressOfInterface(iface)
	}

	log.Info("No interface or IP address specified: will use OVH autodetection")
	return ""
}

func getIpAddressOfInterface(iface *net.Interface) string {

	addrs, err := iface.Addrs()
	if err != nil {
		log.Fatalf("Unable to get addresses for the network interface %s: %s", iface.Name, err)
	}
	if len(addrs) == 0 {
		log.Fatalf("The network interface %s has no addresses", iface.Name)
	}

	ip, _, err := net.ParseCIDR(addrs[0].String())
	if len(addrs) == 0 {
		log.Fatalf("Unable to parse the IP (%s) from the network interface %s: %s", addrs[0].String(), iface.Name, err)
	}
	log.Debugf("Returning IP address %s", ip.String())

	return ip.String()
}
