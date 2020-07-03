/// <reference types="cypress" />

context('Actions', () => {
  beforeEach(() => {
    cy.visit('http://localhost:3000')
  });

  // https://on.cypress.io/interacting-with-elements

  it("should hit the login page", () => {
    cy.get("#gameId")
        .type("12344").should("have.value", "12344");

    cy.get("#userName")
        .type("test").should("have.value", "test");

    cy.get("#newGame")
        .click()
  });
});
