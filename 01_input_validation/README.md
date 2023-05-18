# Input Validation

Every user-input is a potential security risk and should be dealt with.

This is a simple inventory-service where we can try input validation techniques.

1. in `product.service_test.go` add a test with an sql-injection
2. Then try to use "prepared statements" to remove that risk
3. Then do url-parameter validation
