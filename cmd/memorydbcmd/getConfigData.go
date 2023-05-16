package memorydbcmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-memorydb/authenticator"
	"github.com/Appkube-awsx/awsx-memorydb/client"
	"github.com/aws/aws-sdk-go/service/memorydb"
	"github.com/spf13/cobra"
)

// getConfigDataCmd represents the getConfigData command
var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		vaultUrl := cmd.Parent().PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.Parent().PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.Parent().PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.Parent().PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.Parent().PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.Parent().PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.Parent().PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticator.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)
		// print(authFlag)
		// authFlag := true
		if authFlag {
			clusterName, _ := cmd.Flags().GetString("clusterName")
			if clusterName != "" {
				getClusterDetails(region, crossAccountRoleArn, acKey, secKey, clusterName, externalId)
			} else {
				log.Fatalln("clusterName not provided. Program exit")
			}
		}
	},
}


func getClusterDetails(region string, crossAccountRoleArn string, accessKey string, secretKey string, clusterName string, externalId string) *memorydb.DescribeClustersOutput {
	log.Println("Getting aws cluster data")

	listClusterClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)

	input := &memorydb.DescribeClustersInput{}
	
	clusterDetailsResponse, err := listClusterClient.DescribeClusters(input)

	
	if err != nil { 
		log.Fatalln("Error:", err)
	}

	for _, cluster := range clusterDetailsResponse.Clusters{
		
		if *cluster.Name == clusterName{
			log.Println(cluster)
		}else{
			log.Println("No cluster Name present in memory db")
		}
	}
	return clusterDetailsResponse

}
	
func init() {
	GetConfigDataCmd.Flags().StringP("clusterName", "t", "", "Cluster name")

	if err := GetConfigDataCmd.MarkFlagRequired("clusterName"); err != nil {
		fmt.Println(err)
	}
}
