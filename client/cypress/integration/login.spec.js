/// <reference types="cypress" />

describe('Login', () => {
  beforeEach(() => {
    cy.log("Hitting browser");
    cy.visit('http://localhost:3000')
  });

  it("should join a game", () => {
    cy.log("creating new game");
    cy.get("#newGame").click();

    cy.log("Confirm the new game is created");
    cy.get("#status").contains("New game id is:");
    cy.get("#gameId").should("have.value", "12234");
    cy.get("#userName").type("test").should("have.value", "test");

    cy.log("Join the game");
    cy.get("#joinGame").click();
  });
});
