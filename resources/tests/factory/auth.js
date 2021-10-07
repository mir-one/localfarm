export function login() {
  cy.clearLocalStorage()

  cy.visit('/#/')
  // Label
  cy.get('label#label-username').should('contain', 'Username')
  cy.get('label#label-password').should('contain', 'Password')

  // Typing the form
  cy.get('input#username').type('localfarm').should('have.value', 'localfarm')
  cy.get('input#password').type('localfarm')
  cy.get('button.btn').click()
}

export function logout() {
  cy.get('#signout').should('contain', 'Sign Out')
  cy.get('#signout').click()
  cy.location().should( location => {
    expect(location.hash).to.eq('#/auth/login')
  })
}