Feature: User login

  Scenario: Successful login with valid credentials
    Given the user "admin@viska.io" exists with password "admin123"
    Then the response code should be 200