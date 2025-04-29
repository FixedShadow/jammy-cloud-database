package dto

type BasicRes struct {
	RequestId string `json:"requestId"`
}

type DBInstanceSpec struct {
	InstanceName        string `json:"instanceName"`
	Engine              string `json:"engine"`            //mysql sqlserver postgresql
	EngineVersion       string `json:"engineVersion"`     //mysql 5.7 8.0
	InstanceClass       string `json:"instanceClass"`     // db.mysql.s1.micro (1u 1g)   db.mysql.s1.small(1u 2g)   db.mysql.s1.medium(1u 4g) db.mysql.s1.large(2u 8g)
	InstanceStorageGB   int64  `json:"instanceStorageGB"` //GB
	ParameterGroup      string `json:"parameterGroup"`
	InstanceStorageType string `json:"instanceStorageType"` //LOCAL_SSD or NFS
	InstancePort        string `json:"instancePort"`        //the port of database, mysql default port is 3306
	StorageEncrypted    bool   `json:"storageEncrypted"`    //only valid when instanceStorageType is NFS
	InstanceType        string `json:"instanceType"`        //single or cluster
}

type CreateDBInstanceResult struct {
	InstanceId string `json:"instanceId"`
}

type CreateDBInstanceRes struct {
	BasicRes
	Result CreateDBInstanceResult `json:"result"`
}
