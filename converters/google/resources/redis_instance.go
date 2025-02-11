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

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Is the new redis version less than the old one?
func isRedisVersionDecreasing(_ context.Context, old, new, _ interface{}) bool {
	return isRedisVersionDecreasingFunc(old, new)
}

// separate function for unit testing
func isRedisVersionDecreasingFunc(old, new interface{}) bool {
	if old == nil || new == nil {
		return false
	}
	re := regexp.MustCompile(`REDIS_(\d+)_(\d+)`)
	oldParsed := re.FindSubmatch([]byte(old.(string)))
	newParsed := re.FindSubmatch([]byte(new.(string)))

	if oldParsed == nil || newParsed == nil {
		return false
	}

	oldVersion, err := strconv.ParseFloat(fmt.Sprintf("%s.%s", oldParsed[1], oldParsed[2]), 32)
	if err != nil {
		return false
	}
	newVersion, err := strconv.ParseFloat(fmt.Sprintf("%s.%s", newParsed[1], newParsed[2]), 32)
	if err != nil {
		return false
	}

	return newVersion < oldVersion
}

// returns true if old=new or old='auto'
func secondaryIpDiffSuppress(_, old, new string, _ *schema.ResourceData) bool {
	if (strings.ToLower(new) == "auto" && old != "") || old == new {
		return true
	}
	return false
}

const RedisInstanceAssetType string = "redis.googleapis.com/Instance"

func resourceConverterRedisInstance() ResourceConverter {
	return ResourceConverter{
		AssetType: RedisInstanceAssetType,
		Convert:   GetRedisInstanceCaiObject,
	}
}

func GetRedisInstanceCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//redis.googleapis.com/projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetRedisInstanceApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: RedisInstanceAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/redis/v1/rest",
				DiscoveryName:        "Instance",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetRedisInstanceApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	alternativeLocationIdProp, err := expandRedisInstanceAlternativeLocationId(d.Get("alternative_location_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("alternative_location_id"); !isEmptyValue(reflect.ValueOf(alternativeLocationIdProp)) && (ok || !reflect.DeepEqual(v, alternativeLocationIdProp)) {
		obj["alternativeLocationId"] = alternativeLocationIdProp
	}
	authEnabledProp, err := expandRedisInstanceAuthEnabled(d.Get("auth_enabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("auth_enabled"); !isEmptyValue(reflect.ValueOf(authEnabledProp)) && (ok || !reflect.DeepEqual(v, authEnabledProp)) {
		obj["authEnabled"] = authEnabledProp
	}
	authorizedNetworkProp, err := expandRedisInstanceAuthorizedNetwork(d.Get("authorized_network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("authorized_network"); !isEmptyValue(reflect.ValueOf(authorizedNetworkProp)) && (ok || !reflect.DeepEqual(v, authorizedNetworkProp)) {
		obj["authorizedNetwork"] = authorizedNetworkProp
	}
	connectModeProp, err := expandRedisInstanceConnectMode(d.Get("connect_mode"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("connect_mode"); !isEmptyValue(reflect.ValueOf(connectModeProp)) && (ok || !reflect.DeepEqual(v, connectModeProp)) {
		obj["connectMode"] = connectModeProp
	}
	displayNameProp, err := expandRedisInstanceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	labelsProp, err := expandRedisInstanceLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	redisConfigsProp, err := expandRedisInstanceRedisConfigs(d.Get("redis_configs"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("redis_configs"); !isEmptyValue(reflect.ValueOf(redisConfigsProp)) && (ok || !reflect.DeepEqual(v, redisConfigsProp)) {
		obj["redisConfigs"] = redisConfigsProp
	}
	locationIdProp, err := expandRedisInstanceLocationId(d.Get("location_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("location_id"); !isEmptyValue(reflect.ValueOf(locationIdProp)) && (ok || !reflect.DeepEqual(v, locationIdProp)) {
		obj["locationId"] = locationIdProp
	}
	nameProp, err := expandRedisInstanceName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	maintenancePolicyProp, err := expandRedisInstanceMaintenancePolicy(d.Get("maintenance_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("maintenance_policy"); !isEmptyValue(reflect.ValueOf(maintenancePolicyProp)) && (ok || !reflect.DeepEqual(v, maintenancePolicyProp)) {
		obj["maintenancePolicy"] = maintenancePolicyProp
	}
	maintenanceScheduleProp, err := expandRedisInstanceMaintenanceSchedule(d.Get("maintenance_schedule"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("maintenance_schedule"); !isEmptyValue(reflect.ValueOf(maintenanceScheduleProp)) && (ok || !reflect.DeepEqual(v, maintenanceScheduleProp)) {
		obj["maintenanceSchedule"] = maintenanceScheduleProp
	}
	memorySizeGbProp, err := expandRedisInstanceMemorySizeGb(d.Get("memory_size_gb"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("memory_size_gb"); !isEmptyValue(reflect.ValueOf(memorySizeGbProp)) && (ok || !reflect.DeepEqual(v, memorySizeGbProp)) {
		obj["memorySizeGb"] = memorySizeGbProp
	}
	redisVersionProp, err := expandRedisInstanceRedisVersion(d.Get("redis_version"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("redis_version"); !isEmptyValue(reflect.ValueOf(redisVersionProp)) && (ok || !reflect.DeepEqual(v, redisVersionProp)) {
		obj["redisVersion"] = redisVersionProp
	}
	reservedIpRangeProp, err := expandRedisInstanceReservedIpRange(d.Get("reserved_ip_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("reserved_ip_range"); !isEmptyValue(reflect.ValueOf(reservedIpRangeProp)) && (ok || !reflect.DeepEqual(v, reservedIpRangeProp)) {
		obj["reservedIpRange"] = reservedIpRangeProp
	}
	tierProp, err := expandRedisInstanceTier(d.Get("tier"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("tier"); !isEmptyValue(reflect.ValueOf(tierProp)) && (ok || !reflect.DeepEqual(v, tierProp)) {
		obj["tier"] = tierProp
	}
	transitEncryptionModeProp, err := expandRedisInstanceTransitEncryptionMode(d.Get("transit_encryption_mode"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("transit_encryption_mode"); !isEmptyValue(reflect.ValueOf(transitEncryptionModeProp)) && (ok || !reflect.DeepEqual(v, transitEncryptionModeProp)) {
		obj["transitEncryptionMode"] = transitEncryptionModeProp
	}
	replicaCountProp, err := expandRedisInstanceReplicaCount(d.Get("replica_count"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("replica_count"); !isEmptyValue(reflect.ValueOf(replicaCountProp)) && (ok || !reflect.DeepEqual(v, replicaCountProp)) {
		obj["replicaCount"] = replicaCountProp
	}
	readReplicasModeProp, err := expandRedisInstanceReadReplicasMode(d.Get("read_replicas_mode"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("read_replicas_mode"); !isEmptyValue(reflect.ValueOf(readReplicasModeProp)) && (ok || !reflect.DeepEqual(v, readReplicasModeProp)) {
		obj["readReplicasMode"] = readReplicasModeProp
	}
	secondaryIpRangeProp, err := expandRedisInstanceSecondaryIpRange(d.Get("secondary_ip_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("secondary_ip_range"); !isEmptyValue(reflect.ValueOf(secondaryIpRangeProp)) && (ok || !reflect.DeepEqual(v, secondaryIpRangeProp)) {
		obj["secondaryIpRange"] = secondaryIpRangeProp
	}
	customerManagedKeyProp, err := expandRedisInstanceCustomerManagedKey(d.Get("customer_managed_key"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("customer_managed_key"); !isEmptyValue(reflect.ValueOf(customerManagedKeyProp)) && (ok || !reflect.DeepEqual(v, customerManagedKeyProp)) {
		obj["customerManagedKey"] = customerManagedKeyProp
	}

	return resourceRedisInstanceEncoder(d, config, obj)
}

func resourceRedisInstanceEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*Config)
	region, err := getRegionFromSchema("region", "location_id", d, config)
	if err != nil {
		return nil, err
	}
	if err := d.Set("region", region); err != nil {
		return nil, fmt.Errorf("Error setting region: %s", err)
	}
	return obj, nil
}

func expandRedisInstanceAlternativeLocationId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceAuthEnabled(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceAuthorizedNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	fv, err := ParseNetworkFieldValue(v.(string), d, config)
	if err != nil {
		return nil, err
	}
	return fv.RelativeLink(), nil
}

func expandRedisInstanceConnectMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandRedisInstanceRedisConfigs(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandRedisInstanceLocationId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return replaceVars(d, config, "projects/{{project}}/locations/{{region}}/instances/{{name}}")
}

func expandRedisInstanceMaintenancePolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCreateTime, err := expandRedisInstanceMaintenancePolicyCreateTime(original["create_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCreateTime); val.IsValid() && !isEmptyValue(val) {
		transformed["createTime"] = transformedCreateTime
	}

	transformedUpdateTime, err := expandRedisInstanceMaintenancePolicyUpdateTime(original["update_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUpdateTime); val.IsValid() && !isEmptyValue(val) {
		transformed["updateTime"] = transformedUpdateTime
	}

	transformedDescription, err := expandRedisInstanceMaintenancePolicyDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !isEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedWeeklyMaintenanceWindow, err := expandRedisInstanceMaintenancePolicyWeeklyMaintenanceWindow(original["weekly_maintenance_window"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWeeklyMaintenanceWindow); val.IsValid() && !isEmptyValue(val) {
		transformed["weeklyMaintenanceWindow"] = transformedWeeklyMaintenanceWindow
	}

	return transformed, nil
}

func expandRedisInstanceMaintenancePolicyCreateTime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceMaintenancePolicyUpdateTime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceMaintenancePolicyDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceMaintenancePolicyWeeklyMaintenanceWindow(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDay, err := expandRedisInstanceMaintenancePolicyWeeklyMaintenanceWindowDay(original["day"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDay); val.IsValid() && !isEmptyValue(val) {
			transformed["day"] = transformedDay
		}

		transformedDuration, err := expandRedisInstanceMaintenancePolicyWeeklyMaintenanceWindowDuration(original["duration"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDuration); val.IsValid() && !isEmptyValue(val) {
			transformed["duration"] = transformedDuration
		}

		transformedStartTime, err := expandRedisInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(original["start_time"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["startTime"] = transformedStartTime
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandRedisInstanceMaintenancePolicyWeeklyMaintenanceWindowDay(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceMaintenancePolicyWeeklyMaintenanceWindowDuration(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	if l[0] == nil {
		transformed := make(map[string]interface{})
		return transformed, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHours, err := expandRedisInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeHours(original["hours"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHours); val.IsValid() && !isEmptyValue(val) {
		transformed["hours"] = transformedHours
	}

	transformedMinutes, err := expandRedisInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeMinutes(original["minutes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinutes); val.IsValid() && !isEmptyValue(val) {
		transformed["minutes"] = transformedMinutes
	}

	transformedSeconds, err := expandRedisInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeSeconds(original["seconds"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSeconds); val.IsValid() && !isEmptyValue(val) {
		transformed["seconds"] = transformedSeconds
	}

	transformedNanos, err := expandRedisInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeNanos(original["nanos"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNanos); val.IsValid() && !isEmptyValue(val) {
		transformed["nanos"] = transformedNanos
	}

	return transformed, nil
}

func expandRedisInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeHours(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeMinutes(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeSeconds(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeNanos(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceMaintenanceSchedule(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedStartTime, err := expandRedisInstanceMaintenanceScheduleStartTime(original["start_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStartTime); val.IsValid() && !isEmptyValue(val) {
		transformed["startTime"] = transformedStartTime
	}

	transformedEndTime, err := expandRedisInstanceMaintenanceScheduleEndTime(original["end_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEndTime); val.IsValid() && !isEmptyValue(val) {
		transformed["endTime"] = transformedEndTime
	}

	transformedScheduleDeadlineTime, err := expandRedisInstanceMaintenanceScheduleScheduleDeadlineTime(original["schedule_deadline_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScheduleDeadlineTime); val.IsValid() && !isEmptyValue(val) {
		transformed["scheduleDeadlineTime"] = transformedScheduleDeadlineTime
	}

	return transformed, nil
}

func expandRedisInstanceMaintenanceScheduleStartTime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceMaintenanceScheduleEndTime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceMaintenanceScheduleScheduleDeadlineTime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceMemorySizeGb(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceRedisVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceReservedIpRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceTier(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceTransitEncryptionMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceReplicaCount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceReadReplicasMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceSecondaryIpRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceCustomerManagedKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
