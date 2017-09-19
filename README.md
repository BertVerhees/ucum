# Ucum-Golang Version 0.1
Golang library providing UCUM Services

(inspired by Graham Grieve's UCUM Java library: https://github.com/FHIR/Ucum-java )

The library provides a set of services around UCUM:

- validate a UCUM unit (and also against a particular base unit)
- decide whether one unit can be converted/compared to another
- translate a quantity from one unit to another 
- prepare a human readable display of a unit 
- multiply 2 quantities together

To use the library, download ucum-essence from http://unitsofmeasure.org, and then create a UCUMEssenceService:

ucumSvc = NewUcumEssenceService(definitionFile);

Please find the library-API in the file Ucum.go

- Status: all tests run fine
