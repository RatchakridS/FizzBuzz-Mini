import {TukTuk} from './taxi'

describe('Taxi', () => {
  it('No Waiting Time', () => {
    const result = TukTuk(10.5, 0)
    expect(result).toEqual(42)
  })

  // it('Waiting Time', () => {
  //   const result = TukTuk(11.9, 5.5)
  //   expect(result).toEqual(54)
  // })

  // it('Drive only 0.1 Km', () => {
  //   const result = TukTuk(0.1, 0)
  //   expect(result).toEqual(4)
  // })

  // it('Drive for 0.9 Km with 0.1 Waiting Time', () => {
  //   const result = TukTuk(0.9, 0.1)
  //   expect(result).toEqual(5)
  // })

  // it('Drive into 0 Km with 0 Waiting Time', () => {
  //   const result = TukTuk(0 ,0)
  //   expect(result).toEqual(0)
  // })

  // it('Drive for -1 Km. with -1 Waiting Time', () => {
  //   const result = TukTuk(-1, -1)
  //   expect(result).toEqual(0)
  // })

  // it('Drive for -0.1 Km. with -0.1 Waiting Time', () => {
  //   const result = TukTuk(-0.1, -0.1)
  //   expect(result).toEqual(0)
  // })

  // it('Drive for -0.5 Km. With -0.5 Waiting Time', () => {
  //   const result = TukTuk(-0.5, -0.5)
  //   expect(result).toEqual(0)
  // })

  // it('Drive for -0.9 Km. with -0.9 Waiting Time', () => {
  //   const result = TukTuk(-0.9, -0.9)
  //   expect(result).toEqual(0)
  // })

  // it('Drive for 100 Km. with 100 Waiting Time', () => {
  //   const result = TukTuk(100, 100)
  //   expect(result).toEqual(500)
  // })

  // it('Drive for 10000000000000 Km with 100000000 Waiting Time', () => {
  //   const result = TukTuk(10000000000000, 100000000)
  //   let expected = (4 * 10000000000000 ) + (1 * 100000000)
  //   expect(result).toEqual(expected)
  // })
})
