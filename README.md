- [What is awsx-memorydb](#awsx-memorydb)
- [How to write plugin subcommand](#how-to-write-plugin-subcommand)
- [How to build / Test](#how-to-build--test)
- [what it does ](#what-it-does)
- [command input](#command-input)
- [command output](#command-output)
- [How to run ](#how-to-run)

# awsx-memorydb

This is a plugin subcommand for awsx cli ( https://github.com/Appkube-awsx/awsx#awsx ) cli.

For details about awsx commands and how its used in Appkube platform , please refer to the diagram below:

![alt text](https://raw.githubusercontent.com/AppkubeCloud/appkube-architectures/main/LayeredArchitecture-phase2.svg)

This plugin subcommand will implement the Apis' related to Memorydb services , primarily the following API's:

- getConfigData

This cli collect data from metric/logs/traces of the Memorydb services and produce the data in a form that Appkube Platform expects.

This CLI , interacts with other Appkube services like Appkube vault , Appkube cloud CMDB so that it can talk with cloud services as
well as filter and sort the information in terms of product/services, so that Appkube platform gets the data that it expects from the cli.

# How to write plugin subcommand

Please refer to the instruction -
https://github.com/Appkube-awsx/awsx#how-to-write-a-plugin-subcommand

It has detailed instruction on how to write a subcommand plugin , build/test/debug/publish and integrate into the main commmand.

# How to build / Test

            go run main.go
                - Program will print Calling aws-cloudelements on console

            Another way of testing is by running go install command
            go install
            - go install command creates an exe with the name of the module (e.g. awsx-memorydb) and save it in the GOPATH
            - Now we can execute this command on command prompt as below
           awsx-memorydb getConfigData --zone=us-east-1 --accessKey=xxxxxxxxxx --secretKey=xxxxxxxxxx --crossAccountRoleArn=xxxxxxxxxx  --externalId=xxxxxxxxxx

# what it does

This subcommand implement the following functionalities -
getConfigData - It will get the resource count summary for a given AWS account id and region.

# command input

1. --valutURL = specifies the URL of the AWS Key Management Service (KMS) customer master key (CMK) that you want to use to encrypt a lambda.
2. --acountId = specifies the AWS account ID that the lambda belongs to.
3. --zone = specifies the AWS region where the memorydb is located.
4. --accessKey = specifies the AWS access key to use for authentication.
5. --secretKey = specifies the AWS secret key to use for authentication.
6. --crossAccountRoleArn = specifies the Amazon Resource Name (ARN) of the role that allows access to a memorydb in another account.
7. --external Id = specifies the AWS External id to use for authentication.
8. --clusterName= Insert your clusterName from memorydb service in aws account.

# command output

{
ACLName: "open-access",
ARN: "arn:aws:memorydb:us-east-1:657907747545:cluster/my-simple-cluster",
AutoMinorVersionUpgrade: true,
ClusterEndpoint: {
Address: "my-simple-cluster.mxfqj4.clustercfg.memorydb.us-east-1.amazonaws.com",
Port: 6379
},
DataTiering: "false",
EnginePatchVersion: "6.2.6",
EngineVersion: "6.2",
MaintenanceWindow: "thu:07:30-thu:08:30",
Name: "my-simple-cluster",
NodeType: "db.t4g.small",
NumberOfShards: 1,
ParameterGroupName: "default.memorydb-redis6",
ParameterGroupStatus: "in-sync",
SecurityGroups: [{
SecurityGroupId: "sg-0b0d51b62e18f147b",
Status: "active"
}],
SnapshotRetentionLimit: 0,
SnapshotWindow: "04:00-05:00",
Status: "available",
SubnetGroupName: "cluster-subnet",
TLSEnabled: false
}

# How to run

From main awsx command , it is called as follows:

```bash
awsx-memorydb  --zone=us-east-1 --accessKey=<> --secretKey=<> --crossAccountRoleArn=<>  --externalId=<>
```

If you build it locally , you can simply run it as standalone command as:

```bash
go run main.go  --zone=us-east-1 --accessKey=<> --secretKey=<> --crossAccountRoleArn=<> --externalId=<>
```

# awsx-memorydb

memorydb extension

# AWSX Commands for AWSX-Memorydb Cli's :

1. CMD used to get list of memorydb instance's :

```bash
./awsx-memorydb --zone=us-east-1 --accessKey=<> --secretKey=<> --crossAccountRoleArn=<> --externalId=<>
```

2. CMD used to get Config data (metadata) of AWS Memorydb instances :

```bash
./awsx-memorydb --zone=us-east-1 --accessKey=<> --secretKey=<> --crossAccountRoleArn=<> --externalId=<> getConfigData --clusterName=<>
```
