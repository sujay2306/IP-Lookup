package cmd

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"github.com/spf13/cobra"
)

var traceCmd = &cobra.Command{
		Use:   "trace",
		Short: "Trace the IP",
		Long: `Trace the IP"`,

		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				for _,ip := range args{
					fmt.Println(ip)
				}
			}else{
				fmt.Println("Please provide IP Address to trace")
			}
		},
	}

func init(){
	rootCmd.AddCommand(traceCmd)
}

func showData(){
	url := "http://ipinfo.io/1.1.1.1/geo"
	getData(url)
}

func getData(url string) [] byte{
	response, err := http.Get(url)
	if err != nil{
		log.Println("unable to get the response")
	}

	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil{
		log.Panicln("Unable to read response")
	}
	return responseByte
}
