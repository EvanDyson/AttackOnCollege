describe('template spec', () => {
  it('This test will fill in the register form.', () => {

        cy.visit('/src/app/register/register.component.html')
        cy.get("#firstname").type("John")

        cy.get("#lastname").type("Joe")
    
        cy.get("#email").type("someone@example.com")

        cy.get("#username").type("someusername")

        cy.get("#password").type("somePASSword123")

        cy.get("#confirmPassword").type("somePASSword123")
      
        cy.get("#next1-3").click()
      
        cy.get("#dob").type("01/01/1999")
      
        cy.get("#college").type("University of Florida")
      
        cy.get("#major").type("Computer Science")
      })
})