//UAT-R05-M01
import 'cypress-file-upload';

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

      cy.get("#edit-profile").attachFile('images/newProfile.png')
      cy.get("#edit-cover").attachFile('images/redBanner.webp')

      cy.contains("Save").click()

      cy.get('img[alt="company avatar"]').then(($img) => {
        const src1 = $img.attr('src');
        cy.get("#edit-profile").attachFile('images/newProfile2.jpg')
        cy.contains("Save").click()
        cy.get('img[alt="company avatar"]').then(($img) => {
          const src2 = $img.attr('src');
          expect(src1).to.not.equal(src2);
        })
      });

      cy.get('img[alt="company cover"]').then(($img) => {
        const src1 = $img.attr('src');
        cy.get("#edit-cover").attachFile('images/greenBanner.jpg')
        cy.contains("Save").click()
        cy.get('img[alt="company cover"]').then(($img) => {
          const src2 = $img.attr('src');
          expect(src1).to.not.equal(src2);
        })
      });


      cy.contains("Save").click()
    });
  })

})
