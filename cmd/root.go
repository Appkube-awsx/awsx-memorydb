package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Appkube-awsx/awsx-memorydb/authenticator"
	"github.com/Appkube-awsx/awsx-memorydb/client"
	"github.com/Appkube-awsx/awsx-memorydb/cmd/memorydbcmd"

	"github.com/aws/aws-sdk-go/service/memorydb"
	"github.com/spf13/cobra"
)

var AwsxMemorydbMetadataCmd = &cobra.Command{
	Use:   "getListMemorydbMetaDataDetails",
	Short: "getListMemorydbMetaDataDetails command gets resource counts",
	Long:  `getListMemorydbMetaDataDetails command gets resource counts details of an AWS account`,

	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Command memory db started")
		vaultUrl := cmd.PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticator.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)

		if authFlag {
			getListCluster(region, crossAccountRoleArn, acKey, secKey, externalId)
		}
	},
}

type Clusters struct{
	Name string `json:"Name"`
}

type Response struct{
	Clusters []Clusters
}


// json.Unmarshal
func getListCluster(region string, crossAccountRoleArn string, accessKey string, secretKey string, externalId string) (*memorydb.DescribeClustersOutput, error) {
	log.Println("getting memorydb metadata list summary")

	listClusterClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)

	listClusterRequest := &memorydb.DescribeClustersInput{}
	
	listClusterResponse, err := listClusterClient.DescribeClusters(listClusterRequest)
	if err != nil {
		log.Fatalln("Error:in getting  user list", err)
	}
    var responseObject Response

	jsonedResponse, _ := json.Marshal(listClusterResponse)
	json.Unmarshal([]byte(string(jsonedResponse)), &responseObject)
	for _, cluster := range responseObject.Clusters{ 	
		log.Println(cluster.Name)
	}
	// log.Println(listClusterResponse)
	return listClusterResponse, err
}

func Execute() {
	err := AwsxMemorydbMetadataCmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		os.Exit(1)
	}
}

func init() {
	AwsxMemorydbMetadataCmd.AddCommand(memorydbcmd.GetConfigDataCmd)

	AwsxMemorydbMetadataCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxMemorydbMetadataCmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxMemorydbMetadataCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxMemorydbMetadataCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxMemorydbMetadataCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxMemorydbMetadataCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws crossAccountRoleArn is required")
	AwsxMemorydbMetadataCmd.PersistentFlags().String("externalId", "", "aws external id auth")

}
