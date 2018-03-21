# Fare Estimator
Given a set of coordinates that represent the segmented road paths among taxi rides, 
the estimator calculates the total price of each ride. Moreover, the estimator detects
erroneous paths based on a speed limit before attempting to evaluate the ride price.

### Overview
The script parses records from paths.csv, then filters, calculates 
and saves each ride estimation to output.csv.
```
a. 3 models: Path, Ride, Segment
b. 1 contract: FareService
c. 1 implementation with parameters: Taxi, Parameters
d. 1 helper for calculating the distance between points: Haversine
e. 1 estimation implementation: StartEstimator
f. End-to-end test suite
```
