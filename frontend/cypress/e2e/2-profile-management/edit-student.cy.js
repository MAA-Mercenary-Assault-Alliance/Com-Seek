//UAT-R06-M01
import 'cypress-file-upload';

context('Actions', () => {
  beforeEach(() => {
    cy.visit('/')
  })

  it('Edit Student Profile', function () {
    cy.fixture("exampleStudent").then((data) => {
      cy.contains("Log in").click()

      cy.get("#email").type(data.email)
      cy.get("#login-password").type(data.password)

      cy.contains("LOGIN").click()

      cy.contains("Profile").click()
      cy.get("#edit-profile").click()

      cy.get("#edit-first-name").clear().type(data.firstName)
      cy.get("#edit-last-name").clear().type(data.lastName);
      cy.get("#edit-desc").clear().type(data.description);
      cy.get("#edit-facebook").clear().type(data.facebook);
      cy.get("#edit-twitter").clear().type(data.twitter);
      cy.get("#edit-instagram").clear().type(data.instagram);
      cy.get("#edit-github").clear().type(data.github);

      cy.get("#save").click()
    });
  })

  it('Edit Student Pictures', function () {
    cy.fixture("exampleStudent").then((data) => {
      cy.contains("Log in").click()

      cy.get("#email").type(data.email)
      cy.get("#login-password").type(data.password)

      cy.contains("LOGIN").click()

      cy.contains("Profile").click()
      cy.get("#edit-profile").click()

      cy.get("#edit-profile").attachFile('images/newProfile.png')
      cy.get("#edit-cover").attachFile('images/redBanner.webp')

      cy.get("#save").click()

      cy.get("#edit-profile").click()

      cy.get('img[alt="avatar"]').then(($img) => {
        const src1 = $img.attr('src');
        cy.get("#edit-profile").attachFile('images/newProfile2.jpg')
        cy.get("#save").click()
        cy.get('img[alt="avatar"]').then(($img) => {
          const src2 = $img.attr('src');
          expect(src1).to.not.equal(src2);
        })
      });

      cy.get("#edit-profile").click()

      cy.get('img[alt="cover"]').then(($img) => {
        const src1 = $img.attr('src');
        cy.get("#edit-cover").attachFile('images/greenBanner.jpg')
        cy.get("#save").click()
        cy.get('img[alt="cover"]').then(($img) => {
          const src2 = $img.attr('src');
          expect(src1).to.not.equal(src2);
        })
      });


      cy.get("#save").click()
    });
  })

})
