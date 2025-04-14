Feature: Get books

  Scenario: Fetch all books successfully
    Given the data-service is running
    When I send a GET request to "/api/books" with a valid token
    Then the response code should be 200
    And the response should contain a list of books
