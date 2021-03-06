package efs

import (
	gofigCore "github.com/akutz/gofig"
	gofig "github.com/akutz/gofig/types"
)

const (
	// Name is the provider's name.
	Name = "efs"

	// InstanceIDFieldRegion is the key to retrieve the region value from the
	// InstanceID Field map.
	InstanceIDFieldRegion = "region"

	// InstanceIDFieldAvailabilityZone is the key to retrieve the availability
	// zone value from the InstanceID Field map.
	InstanceIDFieldAvailabilityZone = "availabilityZone"
)

func init() {
	r := gofigCore.NewRegistration("EFS")
	r.Key(gofig.String, "", "", "", "efs.accessKey")
	r.Key(gofig.String, "", "", "", "efs.secretKey")
	r.Key(gofig.String, "", "",
		"Comma separated security group ids", "efs.securityGroups")
	r.Key(gofig.String, "", "", "AWS region", "efs.region")
	r.Key(gofig.String, "", "", "Tag prefix for EFS naming", "efs.tag")
	gofigCore.Register(r)
}
