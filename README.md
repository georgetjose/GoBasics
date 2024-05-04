Learning Tests Driven Development with Go.

The Red-Green-Blue method is the working process for TDD. The three steps of the method are:

1. The Red phase: Based on provided requirements, we begin by writing a test with specified inputs and expected outputs. We run this test to ensure it is failing. This helps us avoid any false positives later on.

2. The Green phase: Swapping over to source code, we write enough code to satisfy the previously written test and implement the functionality required. We re-run this test to ensure that it is now passing.

3. The Blue phase: While the test and implementation are successfully fulfilling the requirement we have been focusing on, we can take the time to make any improvements to it in the refactor phase.

The process repeats until we have successfully implemented all functionality that the unit requires.