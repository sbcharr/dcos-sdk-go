// Example of a call to different APIs

package main

import (
	"log"
	auth "github.com/PTC-GLOBAL/dcos-sdk-go/auth/api/v1"
	marathon "github.com/PTC-GLOBAL/dcos-sdk-go/marathon/api/v2/lib"
)

type Service struct {
	Name string `json:"name"`
}

type Redis struct {
	Cpu int `json:"cpus"`
	Mem int `json:"mem"`
}

type PkgOptions struct {
	Ser Service `json:"service"`
	Red Redis `json:"redis"`
}
//install main
/*
func main() {
	uid := ""
	pass := ""
	url := "https://"
	t, err := auth.GetDCOSAccessToken(uid, pass, url)
	if err != nil {
		log.Fatalln("error getting token: ", err)
	}
	//log.Print(t)

	s := Service{
		Name: "elastic",
	}
	r := Redis{
		Cpu: 1,
		Mem: 1024,
	}

	po = PkgOptions{
		Ser: s,
		Red: r,
	}

	//opt := []byte(`{"service": {"name": "redis2"},"redis": {"cpus": 1,"mem": 1024}}`)
	input := cosmos.InstallPackageInput{
		PackageName: "beta-elastic",
		PackageVersion: "1.0.16-5.5.1-beta",
		//Options: po,

	}
	output, err := cosmos.InstallPackage(t, url, input)
	if err != nil {
		log.Fatalln("error creating redis cluster: ", err)
	}
	log.Print(output.ResponseStatus)
	log.Print(output.AppId)
	log.Print(string(output.ResponseMessage))


}


//uninstall main

func main() {
	uid := ""
	pass := ""
	url := "https://"
	t, err := auth.GetDCOSAccessToken(uid, pass, url)
	if err != nil {
		log.Fatalln("error getting token: ", err)
	}

	input := cosmos.UninstallPackageInput{
		PackageName: "beta-elastic",
		AppId: "/elastic",

	}
	output, err := cosmos.UninstallPackage(t, url, input)
	if err != nil {
		log.Fatalln("error deleting redis cluster: ", err)
	}
	log.Print(output.ResponseStatus)
	log.Print(string(output.ResponseMessage))

}
*/
//uninstall status main

func main() {
	uid := ""
	pass := ""
	url := "http://"
	t, err := auth.GetDCOSAccessToken(uid, pass, url)
	if err != nil {
		log.Fatalln("error getting token: ", err)
	}
	appId := "/elastic"
	output, err := marathon.UninstallPackageStatus(appId, t, url)
	if err != nil {
		log.Fatalln("error deleting cluster: ", err)
	}
	log.Print(output)


}


//Install status main
/*
func main() {
	uid := ""
	pass := ""
	url := "https://"
	t, err := auth.GetDCOSAccessToken(uid, pass, url)
	if err != nil {
		log.Fatalln("error getting token: ", err)
	}
	appId := "/elastic"
	output, err := marathon.InstallPackageStatus(appId, t, url)
	if err != nil {
		log.Fatalln("error creating cluster: ", err)
	}
	log.Print(output)


}

*/
