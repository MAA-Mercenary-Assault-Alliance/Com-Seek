//UAT-R04-M01

context('Actions', () => {
  beforeEach(() => {
    cy.visit('/')
  })

  // https://on.cypress.io/interacting-with-elements

  // it('Register as a student', () => {
  //   cy.contains('Get Started').click()
  //
  //   cy.get('#email').type("student@gmail.com")
  //   cy.get('#reg-password').type("12345678")
  //   cy.get('#reg-confirm').type("12345678")
  //   cy.get("#firstName").type("John")
  //   cy.get("#lastName").type("Doe")
  //
  //   cy.get('#register-button').click()
  //
  //   cy.get('#terms-scroll-area').scrollTo('bottom')
  //   cy.get("#agree").click()
  //   cy.get("#accept-terms").click()
  // })

  it('Login as a student', function () {
    cy.fixture("exampleStudent").then((data) => {
      cy.contains("Log in").click()

      cy.get("#email").type(data.email)
      cy.get("#login-password").type(data.password)

      cy.contains("LOGIN").click()

      cy.contains("Search for your jobs now").should("be.visible");
    });
  })

  it('Login as a company', function () {
    cy.fixture("exampleCompany").then((data) => {
      cy.contains("Log in").click()

      cy.get("#email").type(data.email)
      cy.get("#login-password").type(data.password)

      cy.contains("LOGIN").click()

      cy.contains("HR Dashboard").should("be.visible");
    });
  })

  it('Login as an admin', function () {
    cy.fixture("exampleAdmin").then((data) => {
      cy.contains("Log in").click()

      cy.get("#email").type(data.email)
      cy.get("#login-password").type(data.password)

      cy.contains("LOGIN").click()

      cy.contains("Admin Dashboard").should("be.visible");
    });
  })

})