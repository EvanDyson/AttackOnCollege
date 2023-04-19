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
    it('Will get visible landing page elements',()=>{
        cy.visit('/login')
        cy.get("#username").type("AOCAdmin");
        
    
        cy.get("#password").type("SuperSecretP4ssFor4dmin");
        cy.get("#loginclick").click();
        cy.location('pathname').should('eq', '/')
        
        cy.get('#profile').should('be.visibile');
        cy.get('#addassignment').should('be.be.visible')
        cy.get('#admin').should('be.visible')
    })
    it('Will get invisible homepage elements',()=>{
        cy.visit('/login')
        cy.get("#username").type("AOCAdmin");
        
    
        cy.get("#password").type("SuperSecretP4ssFor4dmin");
        cy.get("#loginclick").click();
        cy.location('pathname').should('eq', '/')
        cy.get('#loginwhite').should('be.hidden');
        cy.get('#register').should('be.hidden');
    })
    it('Navigate to Profile Page',()=>{
        cy.visit('/login')
        cy.get("#username").type("AOCAdmin");
        
    
        cy.get("#password").type("SuperSecretP4ssFor4dmin");
        cy.get("#loginclick").click();
        cy.location('pathname').should('eq', '/')
        cy.get('#profile').click();
        cy.location('pathname').should('eq', '/profile')
    })
    it('Should Log Out',()=>{
        cy.visit('/login')
        cy.get("#username").type("AOCAdmin");
        
    
        cy.get("#password").type("SuperSecretP4ssFor4dmin");
        cy.get("#loginclick").click();
        cy.location('pathname').should('eq', '/')
        cy.get('#logout').click();
        cy.location('pathname').should('eq', '/')
        cy.get('#profile').should('be.hidden');
    })
    it('Navigate to Achievement Page',()=>{
        cy.visit('/login')
        cy.get("#username").type("AOCAdmin");
        
    
        cy.get("#password").type("SuperSecretP4ssFor4dmin");
        cy.get("#loginclick").click();
        cy.location('pathname').should('eq', '/')
        cy.get('#achievementButton').click();
        cy.location('pathname').should('eq', '/achievement')
    })
    it('Navigate to Edit User Page',()=>{
        cy.visit('/login')
        cy.get("#username").type("AOCAdmin");
        
    
        cy.get("#password").type("SuperSecretP4ssFor4dmin");
        cy.get("#loginclick").click();
        cy.location('pathname').should('eq', '/')
        cy.get('#profile').click();
        cy.location('pathname').should('eq', '/profile')
        cy.get('#edit').click();
        cy.location('pathname').should('eq', '/edit')
    })
    it('Navigate to Add Assignment Page',()=>{
        cy.visit('/login')
        cy.get("#username").type("AOCAdmin");
        
    
        cy.get("#password").type("SuperSecretP4ssFor4dmin");
        cy.get("#loginclick").click();
        cy.location('pathname').should('eq', '/')
        cy.get('#profile').click();
        cy.location('pathname').should('eq', '/profile')
        cy.get('#addassignment').click();
        cy.location('pathname').should('eq', '/add-assignment')
    })
})

