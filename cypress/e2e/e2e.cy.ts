describe('template spec', () => {
    it('Should accept login',()=>{
        cy.visit('/login')
        cy.get("#username").type("AOCAdmin");
        
    
        cy.get("#password").type("SuperSecretP4ssFor4dmin");
        cy.get("#loginclick").click();
        cy.location('pathname').should('eq', '/')
    
        cy.get("#profile").click()
        cy.location('pathname').should('eq', '/profile')
        cy.get("#addcourse").click()
        cy.get('#courseCode').type("CEN3031")
        cy.get('#courseName').type("Intro to Software Engineering")
    
    })
})