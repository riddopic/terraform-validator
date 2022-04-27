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

const ApigeeInstanceAssetType string = "apigee.googleapis.com/Instance"

func resourceConverterApigeeInstance() ResourceConverter {
	return ResourceConverter{
		AssetType: ApigeeInstanceAssetType,
		Convert:   GetApigeeInstanceCaiObject,
	}
}

func GetApigeeInstanceCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//apigee.googleapis.com/{{org_id}}/instances/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetApigeeInstanceApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: ApigeeInstanceAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/apigee/v1/rest",
				DiscoveryName:        "Instance",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetApigeeInstanceApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandApigeeInstanceName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	locationProp, err := expandApigeeInstanceLocation(d.Get("location"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("location"); !isEmptyValue(reflect.ValueOf(locationProp)) && (ok || !reflect.DeepEqual(v, locationProp)) {
		obj["location"] = locationProp
	}
	peeringCidrRangeProp, err := expandApigeeInstancePeeringCidrRange(d.Get("peering_cidr_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("peering_cidr_range"); !isEmptyValue(reflect.ValueOf(peeringCidrRangeProp)) && (ok || !reflect.DeepEqual(v, peeringCidrRangeProp)) {
		obj["peeringCidrRange"] = peeringCidrRangeProp
	}
	ipRangeProp, err := expandApigeeInstanceIpRange(d.Get("ip_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ip_range"); !isEmptyValue(reflect.ValueOf(ipRangeProp)) && (ok || !reflect.DeepEqual(v, ipRangeProp)) {
		obj["ipRange"] = ipRangeProp
	}
	descriptionProp, err := expandApigeeInstanceDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	displayNameProp, err := expandApigeeInstanceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	diskEncryptionKeyNameProp, err := expandApigeeInstanceDiskEncryptionKeyName(d.Get("disk_encryption_key_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("disk_encryption_key_name"); !isEmptyValue(reflect.ValueOf(diskEncryptionKeyNameProp)) && (ok || !reflect.DeepEqual(v, diskEncryptionKeyNameProp)) {
		obj["diskEncryptionKeyName"] = diskEncryptionKeyNameProp
	}
	consumerAcceptListProp, err := expandApigeeInstanceConsumerAcceptList(d.Get("consumer_accept_list"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("consumer_accept_list"); !isEmptyValue(reflect.ValueOf(consumerAcceptListProp)) && (ok || !reflect.DeepEqual(v, consumerAcceptListProp)) {
		obj["consumerAcceptList"] = consumerAcceptListProp
	}

	return obj, nil
}

func expandApigeeInstanceName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandApigeeInstanceLocation(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandApigeeInstancePeeringCidrRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandApigeeInstanceIpRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandApigeeInstanceDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandApigeeInstanceDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandApigeeInstanceDiskEncryptionKeyName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandApigeeInstanceConsumerAcceptList(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
