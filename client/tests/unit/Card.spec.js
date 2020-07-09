import { shallowMount } from '@vue/test-utils'
import Card from '../../src/components/Card.vue'

// Test helpers
const cardNumbers = [...Array(10).keys()]
const cardColors = ['red', 'green', 'blue', 'yellow']
// TODO: Add the special cards as they are supported

// Compute the cartesian product of numbers and colors
let allCards = [].concat(...cardNumbers.map(num => (cardColors.map(hue => [].concat(num, hue)))))

// Define a test suite to handle testing the properties of a card.
describe('Card properties', () => {

  test("Card has computed props", () => {
    expect(typeof Card.computed.card_specifics).toBe('function')
  })

  test.each(allCards)(
    'Card correctly computes specifics for %i %i', (num, col) => {
    const sampleProps = { number: num, color: col }

    expect(Card.computed.card_specifics.call(sampleProps)).toBe('card num-' + num + ' ' + col)
  })

})
