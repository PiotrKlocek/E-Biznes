const USER = 'standard_user'
const PASSWORD = 'secret_sauce'

function login(username = USER, password = PASSWORD) {
  cy.visit('/')
  cy.get('[data-test="username"]').clear().type(username)
  cy.get('[data-test="password"]').clear().type(password)
  cy.get('[data-test="login-button"]').click()
}

describe('Zadanie 6 - 20 testów funkcjonalnych SauceDemo', () => {
  it('TC01: wyświetla formularz logowania', () => {
    cy.visit('/')
    cy.get('[data-test="username"]').should('be.visible')
    cy.get('[data-test="password"]').should('be.visible')
    cy.get('[data-test="login-button"]').should('be.visible').and('have.value', 'Login')
  })

  it('TC02: loguje poprawnego użytkownika', () => {
    login()
    cy.url().should('include', '/inventory.html')
    cy.get('.title').should('have.text', 'Products')
    cy.get('.inventory_item').should('have.length', 6)
  })

  it('TC03: blokuje błędne hasło', () => {
    login(USER, 'wrong_password')
    cy.get('[data-test="error"]').should('be.visible')
    cy.get('[data-test="error"]').should('contain', 'Username and password do not match')
    cy.url().should('not.include', '/inventory.html')
  })

  it('TC04: blokuje pusty login i hasło', () => {
    cy.visit('/')
    cy.get('[data-test="login-button"]').click()
    cy.get('[data-test="error"]').should('be.visible')
    cy.get('[data-test="error"]').should('contain', 'Username is required')
    cy.get('[data-test="username"]').should('have.class', 'input_error')
  })

  it('TC05: blokuje pusty password', () => {
    cy.visit('/')
    cy.get('[data-test="username"]').type(USER)
    cy.get('[data-test="login-button"]').click()
    cy.get('[data-test="error"]').should('be.visible')
    cy.get('[data-test="error"]').should('contain', 'Password is required')
    cy.get('[data-test="password"]').should('have.class', 'input_error')
  })

  it('TC06: wylogowuje użytkownika', () => {
    login()
    cy.get('#react-burger-menu-btn').click()
    cy.get('[data-test="logout-sidebar-link"]').should('be.visible').click()
    cy.url().should('eq', 'https://www.saucedemo.com/')
    cy.get('[data-test="login-button"]').should('be.visible')
  })

  it('TC07: dodaje jeden produkt do koszyka', () => {
    login()
    cy.get('[data-test="add-to-cart-sauce-labs-backpack"]').click()
    cy.get('.shopping_cart_badge').should('have.text', '1')
    cy.get('[data-test="remove-sauce-labs-backpack"]').should('be.visible')
    cy.get('.shopping_cart_link').click()
    cy.get('.cart_item').should('have.length', 1)
    cy.get('.inventory_item_name').should('contain', 'Sauce Labs Backpack')
  })

  it('TC08: usuwa produkt z koszyka na liście produktów', () => {
    login()
    cy.get('[data-test="add-to-cart-sauce-labs-bike-light"]').click()
    cy.get('.shopping_cart_badge').should('have.text', '1')
    cy.get('[data-test="remove-sauce-labs-bike-light"]').click()
    cy.get('.shopping_cart_badge').should('not.exist')
    cy.get('[data-test="add-to-cart-sauce-labs-bike-light"]').should('be.visible')
  })

  it('TC09: dodaje kilka produktów do koszyka', () => {
    login()
    cy.get('[data-test="add-to-cart-sauce-labs-backpack"]').click()
    cy.get('[data-test="add-to-cart-sauce-labs-bike-light"]').click()
    cy.get('[data-test="add-to-cart-sauce-labs-bolt-t-shirt"]').click()
    cy.get('.shopping_cart_badge').should('have.text', '3')
    cy.get('.shopping_cart_link').click()
    cy.get('.cart_item').should('have.length', 3)
    cy.get('.cart_item').first().should('contain', 'Sauce Labs Backpack')
  })

  it('TC10: usuwa produkt z poziomu koszyka', () => {
    login()
    cy.get('[data-test="add-to-cart-sauce-labs-backpack"]').click()
    cy.get('.shopping_cart_link').click()
    cy.get('.cart_item').should('have.length', 1)
    cy.get('[data-test="remove-sauce-labs-backpack"]').click()
    cy.get('.cart_item').should('not.exist')
    cy.get('.shopping_cart_badge').should('not.exist')
  })

  it('TC11: sortuje produkty od A do Z', () => {
    login()
    cy.get('[data-test="product-sort-container"]').select('az')
    cy.get('.inventory_item_name').first().should('have.text', 'Sauce Labs Backpack')
    cy.get('.inventory_item_name').last().should('have.text', 'Test.allTheThings() T-Shirt (Red)')
    cy.get('[data-test="product-sort-container"]').should('have.value', 'az')
  })

  it('TC12: sprawdza link LinkedIn w stopce', () => {
    login()
    cy.get('.footer').should('be.visible')
    cy.get('.social_linkedin a').should('exist')
    cy.get('.social_linkedin a')
      .should('have.attr', 'href')
      .and('include', 'linkedin.com')
  })

  it('TC13: sortuje produkty według ceny rosnąco', () => {
    login()
    cy.get('[data-test="product-sort-container"]').select('lohi')
    cy.get('.inventory_item_price').first().should('have.text', '$7.99')
    cy.get('.inventory_item_price').last().should('have.text', '$49.99')
    cy.get('[data-test="product-sort-container"]').should('have.value', 'lohi')
  })

  it('TC14: sortuje produkty według ceny malejąco', () => {
    login()
    cy.get('[data-test="product-sort-container"]').select('hilo')
    cy.get('.inventory_item_price').first().should('have.text', '$49.99')
    cy.get('.inventory_item_price').last().should('have.text', '$7.99')
    cy.get('[data-test="product-sort-container"]').should('have.value', 'hilo')
  })

  it('TC15: otwiera szczegóły produktu', () => {
    login()
    cy.contains('.inventory_item_name', 'Sauce Labs Backpack').click()
    cy.url().should('include', '/inventory-item.html')
    cy.get('.inventory_details_name').should('have.text', 'Sauce Labs Backpack')
    cy.get('.inventory_details_price').should('have.text', '$29.99')
    cy.get('[data-test="back-to-products"]').should('be.visible')
  })

  it('TC16: wraca ze szczegółów produktu do listy', () => {
    login()
    cy.contains('.inventory_item_name', 'Sauce Labs Bike Light').click()
    cy.get('[data-test="back-to-products"]').click()
    cy.url().should('include', '/inventory.html')
    cy.get('.title').should('have.text', 'Products')
    cy.get('.inventory_item').should('have.length', 6)
  })

  it('TC17: przechodzi do checkout z koszyka', () => {
    login()
    cy.get('[data-test="add-to-cart-sauce-labs-backpack"]').click()
    cy.get('.shopping_cart_link').click()
    cy.get('[data-test="checkout"]').click()
    cy.url().should('include', '/checkout-step-one.html')
    cy.get('.title').should('have.text', 'Checkout: Your Information')
    cy.get('[data-test="firstName"]').should('be.visible')
  })

  it('TC18: waliduje wymagane pola checkout', () => {
    login()
    cy.get('[data-test="add-to-cart-sauce-labs-backpack"]').click()
    cy.get('.shopping_cart_link').click()
    cy.get('[data-test="checkout"]').click()
    cy.get('[data-test="continue"]').click()
    cy.get('[data-test="error"]').should('be.visible')
    cy.get('[data-test="error"]').should('contain', 'First Name is required')
    cy.get('[data-test="firstName"]').should('have.class', 'input_error')
  })

  it('TC19: przechodzi przez podsumowanie zamówienia', () => {
    login()
    cy.get('[data-test="add-to-cart-sauce-labs-backpack"]').click()
    cy.get('.shopping_cart_link').click()
    cy.get('[data-test="checkout"]').click()
    cy.get('[data-test="firstName"]').type('Jan')
    cy.get('[data-test="lastName"]').type('Kowalski')
    cy.get('[data-test="postalCode"]').type('00-001')
    cy.get('[data-test="continue"]').click()
    cy.url().should('include', '/checkout-step-two.html')
    cy.get('.summary_info').should('contain', 'Payment Information')
    cy.get('.summary_subtotal_label').should('contain', 'Item total: $29.99')
  })

  it('TC20: kończy zamówienie', () => {
    login()
    cy.get('[data-test="add-to-cart-sauce-labs-backpack"]').click()
    cy.get('.shopping_cart_link').click()
    cy.get('[data-test="checkout"]').click()
    cy.get('[data-test="firstName"]').type('Jan')
    cy.get('[data-test="lastName"]').type('Kowalski')
    cy.get('[data-test="postalCode"]').type('00-001')
    cy.get('[data-test="continue"]').click()
    cy.get('[data-test="finish"]').click()
    cy.url().should('include', '/checkout-complete.html')
    cy.get('.complete-header').should('have.text', 'Thank you for your order!')
    cy.get('[data-test="back-to-products"]').should('be.visible')
  })
})
