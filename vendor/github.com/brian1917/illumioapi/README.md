# Illumio API Package

[![GoDoc](https://godoc.org/github.com/brian1917/illumioapi?status.svg)](https://godoc.org/github.com/brian1917/illumioapi)

## Description

Go package to interact with the Illumio PCE API.

## Deprecated Method Announcement - June 2, 2022
Several methods have been deprecated as part of a naming convention standardization and leveraging the new PCE crud methods in `crud.go`. Old methods
are maintained in `depreceated.go` to keep backwards compatibility. The deprecated functions will be removed in August 2022.

## Example Code
All interaction with the PCE are done via methods on the PCE type. For example, the code below prints all hostnames:
```
// Create PCE
pce := illumioapi.PCE{
   FQDN: "bep-lab.poc.segmentationpov.com",
   Port: 443,
   DisableTLSChecking: true}

// Login and ignore error checking for example
pce.Login("brian@email.com", "Password123")

// Get all workloads and ignore error checking for example
wklds, _, _ := pce.GetWklds(nil)

// Iterate through workloads and print hostname
for _, w := range wklds {
    fmt.Println(w.
}

// Get just managed workloads using query parameter
managedWklds, _, _ := pce.GetWklds(map[string]string{"managed":"true"})
```