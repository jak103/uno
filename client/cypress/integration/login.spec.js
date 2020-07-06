/// <reference types="cypress" />

describe('Login', () => {
  beforeEach(() => {
    cy.log("Hitting browser");
    cy.visit('http://localhost:3000')
  });

  it("should join a game", () => {
    cy.log("creating new game");
    cy.get('[test-id="login-new-game"]').click();

    cy.log("Confirm the new game is created");
    cy.get('[test-id="login-status"]').contains("New game id is:");

    // TODO - compare the value provided in status matches the value supplied for game-id
    cy.get('[test-id="login-game-id"]').invoke('val').should("not.be.empty");
    cy.get('[test-id="login-user-name"]').type("test").should("have.value", "test");

    cy.log("Join the game");
    cy.get('[test-id="login-join-game"]').click();
  });
});
