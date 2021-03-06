# Ucum-Golang Version 0.9.4.beta
Golang library providing UCUM Services

*(inspired by Grahame Grieve's UCUM Java library: https://github.com/FHIR/Ucum-java )*

The library provides a set of services around UCUM:

- validate a UCUM unit (and also against a particular base unit)
- decide whether one unit can be converted/compared to another
- translate a quantity from one unit to another 
- prepare a human readable display of a unit 
- multiply 2 quantities together

To use the library, download the definitionFile: ucum-essence.xml from http://unitsofmeasure.org, and then create a UCUMEssenceService:

ucumSvc = NewUcumEssenceService(definitionFile);

Please find the library-API in the file Ucum.go

- Status: all API tests run fine.
