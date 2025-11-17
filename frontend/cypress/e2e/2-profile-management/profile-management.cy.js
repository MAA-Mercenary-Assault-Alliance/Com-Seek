//UAT-R05-M01

context('Actions', () => {
  beforeEach(() => {
    cy.visit('/')
  })

  it('Edit Company Profile', function () {
    cy.fixture("exampleCompany").then((data) => {
      cy.contains("Log in").click()

      cy.get("#email").type(data.email)
      cy.get("#login-password").type(data.password)

      cy.contains("LOGIN").click()

      cy.contains("HR Dashboard").should("be.visible");

      cy.contains("Profile").click()
      cy.contains("Edit").click()

      cy.get("#edit-name").clear().type(data.name)
      cy.get("#edit-website").clear().type(data.website);
      cy.get("#edit-location").clear().type(data.location);
      cy.get("#edit-contact-email").clear().type(data.contactEmail);
      cy.get("#edit-contact-number").clear().type(data.contactNumber);
      cy.get("#edit-desc").clear().type(data.description);


      cy.contains("Save").click()
    });
  })

  it('Edit Company Pictures', function () {
    cy.fixture("exampleCompany").then((data) => {
      cy.contains("Log in").click()

      cy.get("#email").type(data.email)
      cy.get("#login-password").type(data.password)

      cy.contains("LOGIN").click()

      cy.contains("HR Dashboard").should("be.visible");

      cy.contains("Profile").click()
      cy.contains("Edit").click()

      cy.get("#edit-profile").clear().type(data.name)
      cy.get("#edit-cover").clear().type(data.website);

      cy.contains("Save").click()
    });
  })

})