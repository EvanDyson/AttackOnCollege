describe('template spec', () => {
  
    it('Will get visible homepage elements',()=>{
         cy.visit('/')
        cy.get('#loginwhite').should('be.visible');
        cy.get('#register').should('be.visible');
    })
    it('Will get invisible landing page elements',()=>{
         cy.visit('/')
        cy.get('#profile').should('be.hidden');
        cy.get('#addassignment').should('be.hidden')
    })
    it('Will get visible login page elements',()=>{
         cy.visit('/login')
        cy.get('#register').should('be.visible');
    })
    it('Will get invisible login page elements',()=>{
         cy.visit('/login')
        cy.get('#profile').should('be.hidden');
        cy.get('#login').should('be.hidden');
        cy.get('#addassignment').should('be.hidden')
    })
    it('Login Button Links to Correct Page',()=>{
        cy.visit('/')
        cy.get('#loginwhite').click();
        cy.location('pathname').should('eq', '/login')
    })
    it('Register Button Links to Correct Page',()=>{
        cy.visit('/')
        cy.get('#register').click();
        cy.location('pathname').should('eq', '/register')
    })
    it('Should deny login',()=>{
        cy.visit('/login')
        cy.get("#username").type("someusername");
        cy.get("#password").type("somePASSword123");
        cy.get("#loginclick").click();
        cy.location('pathname').should('eq', '/login')
    })
    it('Show Profile invisibility',()=>{
        cy.visit('/profile')
        cy.get('#login').should('be.hidden');
        cy.get('#register').should('be.hidden');
    })
    it('Show Profile visibility',()=>{
       cy.visit('/profile')
       cy.get("#profile").should('be.visible');
       cy.get('#addassignment').should('be.visible')
    })
    it('Get Home from Anywhere',()=>{
       cy.visit('/')
       cy.get("#loginwhite").click();
       cy.get('#homelink').click();
       cy.location('pathname').should('eq', '/')
       cy.get('#register').click();
       cy.get('#homelink').click();
       cy.location('pathname').should('eq', '/')
       cy.get('#homelink').click();
       cy.location('pathname').should('eq', '/')
    })
})