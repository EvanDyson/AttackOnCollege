describe('template spec', () => {
  it('should type username and password', () => {

        cy.visit('/src/app/register/register.component.html')
        cy.get("#firstname").type("Hi")

        cy.get("#lastname").type("Hi")
    
        cy.get("#email").type("someone@example.com")

        cy.get("#username").type("someusername")

        cy.get("#password").type("somepassword")

        cy.get("#confirmpassword").type("somepassword")
    
      })
})