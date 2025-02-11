// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import "reflect"

const DatastreamConnectionProfileAssetType string = "datastream.googleapis.com/ConnectionProfile"

func resourceConverterDatastreamConnectionProfile() ResourceConverter {
	return ResourceConverter{
		AssetType: DatastreamConnectionProfileAssetType,
		Convert:   GetDatastreamConnectionProfileCaiObject,
	}
}

func GetDatastreamConnectionProfileCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//datastream.googleapis.com/projects/{{project}}/locations/{{location}}/connectionProfiles/{{connection_profile_id}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetDatastreamConnectionProfileApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: DatastreamConnectionProfileAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/datastream/v1/rest",
				DiscoveryName:        "ConnectionProfile",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetDatastreamConnectionProfileApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	labelsProp, err := expandDatastreamConnectionProfileLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	displayNameProp, err := expandDatastreamConnectionProfileDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	oracleProfileProp, err := expandDatastreamConnectionProfileOracleProfile(d.Get("oracle_profile"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("oracle_profile"); !isEmptyValue(reflect.ValueOf(oracleProfileProp)) && (ok || !reflect.DeepEqual(v, oracleProfileProp)) {
		obj["oracleProfile"] = oracleProfileProp
	}
	gcsProfileProp, err := expandDatastreamConnectionProfileGcsProfile(d.Get("gcs_profile"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("gcs_profile"); !isEmptyValue(reflect.ValueOf(gcsProfileProp)) && (ok || !reflect.DeepEqual(v, gcsProfileProp)) {
		obj["gcsProfile"] = gcsProfileProp
	}
	mysqlProfileProp, err := expandDatastreamConnectionProfileMysqlProfile(d.Get("mysql_profile"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("mysql_profile"); !isEmptyValue(reflect.ValueOf(mysqlProfileProp)) && (ok || !reflect.DeepEqual(v, mysqlProfileProp)) {
		obj["mysqlProfile"] = mysqlProfileProp
	}
	postgresqlProfileProp, err := expandDatastreamConnectionProfilePostgresqlProfile(d.Get("postgresql_profile"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("postgresql_profile"); !isEmptyValue(reflect.ValueOf(postgresqlProfileProp)) && (ok || !reflect.DeepEqual(v, postgresqlProfileProp)) {
		obj["postgresqlProfile"] = postgresqlProfileProp
	}
	forwardSshConnectivityProp, err := expandDatastreamConnectionProfileForwardSshConnectivity(d.Get("forward_ssh_connectivity"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("forward_ssh_connectivity"); !isEmptyValue(reflect.ValueOf(forwardSshConnectivityProp)) && (ok || !reflect.DeepEqual(v, forwardSshConnectivityProp)) {
		obj["forwardSshConnectivity"] = forwardSshConnectivityProp
	}

	return obj, nil
}

func expandDatastreamConnectionProfileLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDatastreamConnectionProfileDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamConnectionProfileOracleProfile(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHostname, err := expandDatastreamConnectionProfileOracleProfileHostname(original["hostname"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHostname); val.IsValid() && !isEmptyValue(val) {
		transformed["hostname"] = transformedHostname
	}

	transformedPort, err := expandDatastreamConnectionProfileOracleProfilePort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !isEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedUsername, err := expandDatastreamConnectionProfileOracleProfileUsername(original["username"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUsername); val.IsValid() && !isEmptyValue(val) {
		transformed["username"] = transformedUsername
	}

	transformedPassword, err := expandDatastreamConnectionProfileOracleProfilePassword(original["password"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPassword); val.IsValid() && !isEmptyValue(val) {
		transformed["password"] = transformedPassword
	}

	transformedDatabaseService, err := expandDatastreamConnectionProfileOracleProfileDatabaseService(original["database_service"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDatabaseService); val.IsValid() && !isEmptyValue(val) {
		transformed["databaseService"] = transformedDatabaseService
	}

	transformedConnectionAttributes, err := expandDatastreamConnectionProfileOracleProfileConnectionAttributes(original["connection_attributes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConnectionAttributes); val.IsValid() && !isEmptyValue(val) {
		transformed["connectionAttributes"] = transformedConnectionAttributes
	}

	return transformed, nil
}

func expandDatastreamConnectionProfileOracleProfileHostname(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamConnectionProfileOracleProfilePort(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamConnectionProfileOracleProfileUsername(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamConnectionProfileOracleProfilePassword(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamConnectionProfileOracleProfileDatabaseService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamConnectionProfileOracleProfileConnectionAttributes(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDatastreamConnectionProfileGcsProfile(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBucket, err := expandDatastreamConnectionProfileGcsProfileBucket(original["bucket"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBucket); val.IsValid() && !isEmptyValue(val) {
		transformed["bucket"] = transformedBucket
	}

	transformedRootPath, err := expandDatastreamConnectionProfileGcsProfileRootPath(original["root_path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRootPath); val.IsValid() && !isEmptyValue(val) {
		transformed["rootPath"] = transformedRootPath
	}

	return transformed, nil
}

func expandDatastreamConnectionProfileGcsProfileBucket(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamConnectionProfileGcsProfileRootPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamConnectionProfileMysqlProfile(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHostname, err := expandDatastreamConnectionProfileMysqlProfileHostname(original["hostname"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHostname); val.IsValid() && !isEmptyValue(val) {
		transformed["hostname"] = transformedHostname
	}

	transformedPort, err := expandDatastreamConnectionProfileMysqlProfilePort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !isEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedUsername, err := expandDatastreamConnectionProfileMysqlProfileUsername(original["username"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUsername); val.IsValid() && !isEmptyValue(val) {
		transformed["username"] = transformedUsername
	}

	transformedPassword, err := expandDatastreamConnectionProfileMysqlProfilePassword(original["password"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPassword); val.IsValid() && !isEmptyValue(val) {
		transformed["password"] = transformedPassword
	}

	transformedSslConfig, err := expandDatastreamConnectionProfileMysqlProfileSslConfig(original["ssl_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSslConfig); val.IsValid() && !isEmptyValue(val) {
		transformed["sslConfig"] = transformedSslConfig
	}

	return transformed, nil
}

func expandDatastreamConnectionProfileMysqlProfileHostname(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamConnectionProfileMysqlProfilePort(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamConnectionProfileMysqlProfileUsername(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamConnectionProfileMysqlProfilePassword(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamConnectionProfileMysqlProfileSslConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedClientKey, err := expandDatastreamConnectionProfileMysqlProfileSslConfigClientKey(original["client_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClientKey); val.IsValid() && !isEmptyValue(val) {
		transformed["clientKey"] = transformedClientKey
	}

	transformedClientKeySet, err := expandDatastreamConnectionProfileMysqlProfileSslConfigClientKeySet(original["client_key_set"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClientKeySet); val.IsValid() && !isEmptyValue(val) {
		transformed["clientKeySet"] = transformedClientKeySet
	}

	transformedClientCertificate, err := expandDatastreamConnectionProfileMysqlProfileSslConfigClientCertificate(original["client_certificate"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClientCertificate); val.IsValid() && !isEmptyValue(val) {
		transformed["clientCertificate"] = transformedClientCertificate
	}

	transformedClientCertificateSet, err := expandDatastreamConnectionProfileMysqlProfileSslConfigClientCertificateSet(original["client_certificate_set"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClientCertificateSet); val.IsValid() && !isEmptyValue(val) {
		transformed["clientCertificateSet"] = transformedClientCertificateSet
	}

	transformedCaCertificate, err := expandDatastreamConnectionProfileMysqlProfileSslConfigCaCertificate(original["ca_certificate"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCaCertificate); val.IsValid() && !isEmptyValue(val) {
		transformed["caCertificate"] = transformedCaCertificate
	}

	transformedCaCertificateSet, err := expandDatastreamConnectionProfileMysqlProfileSslConfigCaCertificateSet(original["ca_certificate_set"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCaCertificateSet); val.IsValid() && !isEmptyValue(val) {
		transformed["caCertificateSet"] = transformedCaCertificateSet
	}

	return transformed, nil
}

func expandDatastreamConnectionProfileMysqlProfileSslConfigClientKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamConnectionProfileMysqlProfileSslConfigClientKeySet(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamConnectionProfileMysqlProfileSslConfigClientCertificate(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamConnectionProfileMysqlProfileSslConfigClientCertificateSet(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamConnectionProfileMysqlProfileSslConfigCaCertificate(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamConnectionProfileMysqlProfileSslConfigCaCertificateSet(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamConnectionProfilePostgresqlProfile(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHostname, err := expandDatastreamConnectionProfilePostgresqlProfileHostname(original["hostname"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHostname); val.IsValid() && !isEmptyValue(val) {
		transformed["hostname"] = transformedHostname
	}

	transformedPort, err := expandDatastreamConnectionProfilePostgresqlProfilePort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !isEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedUsername, err := expandDatastreamConnectionProfilePostgresqlProfileUsername(original["username"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUsername); val.IsValid() && !isEmptyValue(val) {
		transformed["username"] = transformedUsername
	}

	transformedPassword, err := expandDatastreamConnectionProfilePostgresqlProfilePassword(original["password"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPassword); val.IsValid() && !isEmptyValue(val) {
		transformed["password"] = transformedPassword
	}

	transformedDatabase, err := expandDatastreamConnectionProfilePostgresqlProfileDatabase(original["database"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDatabase); val.IsValid() && !isEmptyValue(val) {
		transformed["database"] = transformedDatabase
	}

	return transformed, nil
}

func expandDatastreamConnectionProfilePostgresqlProfileHostname(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamConnectionProfilePostgresqlProfilePort(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamConnectionProfilePostgresqlProfileUsername(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamConnectionProfilePostgresqlProfilePassword(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamConnectionProfilePostgresqlProfileDatabase(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamConnectionProfileForwardSshConnectivity(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHostname, err := expandDatastreamConnectionProfileForwardSshConnectivityHostname(original["hostname"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHostname); val.IsValid() && !isEmptyValue(val) {
		transformed["hostname"] = transformedHostname
	}

	transformedUsername, err := expandDatastreamConnectionProfileForwardSshConnectivityUsername(original["username"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUsername); val.IsValid() && !isEmptyValue(val) {
		transformed["username"] = transformedUsername
	}

	transformedPort, err := expandDatastreamConnectionProfileForwardSshConnectivityPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !isEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedPassword, err := expandDatastreamConnectionProfileForwardSshConnectivityPassword(original["password"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPassword); val.IsValid() && !isEmptyValue(val) {
		transformed["password"] = transformedPassword
	}

	transformedPrivateKey, err := expandDatastreamConnectionProfileForwardSshConnectivityPrivateKey(original["private_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPrivateKey); val.IsValid() && !isEmptyValue(val) {
		transformed["privateKey"] = transformedPrivateKey
	}

	return transformed, nil
}

func expandDatastreamConnectionProfileForwardSshConnectivityHostname(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamConnectionProfileForwardSshConnectivityUsername(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamConnectionProfileForwardSshConnectivityPort(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamConnectionProfileForwardSshConnectivityPassword(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamConnectionProfileForwardSshConnectivityPrivateKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
