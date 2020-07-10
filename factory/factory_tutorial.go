package main

import (
	"fmt"
	"github.com/edgargmc/hydra/factory/appliances"
)

func test() {
	var myTpe int
	fmt.Scan(&myTpe)

	myAppliance, error := appliances.CreateApplianceType(myTpe)

	if error == nil {
		myAppliance.Start()
		fmt.Println(myAppliance.GetPurpose())
	} else {
		fmt.Println(error)
	}
}
